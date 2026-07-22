package client

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"sync"
	"testing"
	"time"
)

type httpClientFunc func(*http.Request) (*http.Response, error)

func (f httpClientFunc) Do(req *http.Request) (*http.Response, error) {
	return f(req)
}

func mustNewClient(tb testing.TB, config Config) *Client {
	tb.Helper()
	c, err := NewClient(config)
	if err != nil {
		tb.Fatalf("NewClient() error = %v", err)
	}
	return c
}

func TestClient_Get(t *testing.T) {
	tests := []struct {
		name           string
		endpoint       string
		params         url.Values
		serverResponse string
		serverStatus   int
		wantErr        bool
	}{
		{
			name:           "successful request",
			endpoint:       "/test",
			params:         nil,
			serverResponse: `{"success": true}`,
			serverStatus:   http.StatusOK,
			wantErr:        false,
		},
		{
			name:           "404 error",
			endpoint:       "/notfound",
			params:         nil,
			serverResponse: `{"error": "not found"}`,
			serverStatus:   http.StatusNotFound,
			wantErr:        true,
		},
		{
			name:     "with query parameters",
			endpoint: "/test",
			params: url.Values{
				"param1": []string{"value1"},
				"param2": []string{"value2"},
			},
			serverResponse: `{"success": true}`,
			serverStatus:   http.StatusOK,
			wantErr:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.serverStatus)
				w.Write([]byte(tt.serverResponse))
			}))
			defer server.Close()

			client := mustNewClient(t, Config{
				BaseURL: server.URL,
			})

			resp, err := client.Get(context.Background(), tt.endpoint, tt.params)

			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && resp == nil {
				t.Errorf("Client.Get() returned nil response")
			}

			if !tt.wantErr && resp != nil {
				if resp.StatusCode != tt.serverStatus {
					t.Errorf("Client.Get() status = %d, want %d", resp.StatusCode, tt.serverStatus)
				}
				if string(resp.Body) != tt.serverResponse {
					t.Errorf("Client.Get() body = %s, want %s", string(resp.Body), tt.serverResponse)
				}
			}
		})
	}
}

func TestClient_buildURL(t *testing.T) {
	client := mustNewClient(t, Config{
		BaseURL: "https://api.example.com",
	})

	tests := []struct {
		name     string
		endpoint string
		params   url.Values
		want     string
	}{
		{
			name:     "simple endpoint",
			endpoint: "/test",
			params:   nil,
			want:     "https://api.example.com/test",
		},
		{
			name:     "with parameters",
			endpoint: "/test",
			params: url.Values{
				"a": []string{"1"},
				"b": []string{"2"},
			},
			want: "https://api.example.com/test?a=1&b=2",
		},
		{
			name:     "sorted parameters",
			endpoint: "/test",
			params: url.Values{
				"z": []string{"last"},
				"a": []string{"first"},
				"m": []string{"middle"},
			},
			want: "https://api.example.com/test?a=first&m=middle&z=last",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := client.buildURL(tt.endpoint, tt.params)
			if got != tt.want {
				t.Errorf("Client.buildURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewClientRejectsInvalidBaseURL(t *testing.T) {
	if _, err := NewClient(Config{BaseURL: "://invalid"}); err == nil {
		t.Fatal("NewClient() succeeded with an invalid base URL")
	}
}

// TestNewClientRejectsUnusableBaseURL covers BaseURL values url.Parse
// itself accepts without error - a relative reference, an opaque string
// with no host, the empty string, and a host with a port but no hostname
// - which the plain url.Parse check alone let through uncaught. Found by
// the 2026-07-22 (9eb3a9a) maintainability assessment: NewClient checked
// only url.Parse's error, never IsAbs()/Scheme/Host, so a mistyped
// BaseURL like "not-a-url" would construct a Client successfully and
// only fail confusingly on the first Get. "https://:443" is a later
// addition (2026-07-22, 0e400d1 assessment): url.URL.Host includes an
// optional port, so this input has a non-empty Host (":443") that passed
// the original Host == "" check even though Hostname() - the actual
// destination - is empty.
func TestNewClientRejectsUnusableBaseURL(t *testing.T) {
	for _, baseURL := range []string{
		"",
		"not-a-url",
		"example.com",       // no scheme - url.Parse treats this as a relative path, not a host
		"//example.com",     // protocol-relative - still no scheme
		"ftp://example.com", // parses fine, but not a scheme this client can use
		"https://:443",      // Host is ":443" (non-empty) but Hostname() is ""
	} {
		t.Run(baseURL, func(t *testing.T) {
			if _, err := NewClient(Config{BaseURL: baseURL}); err == nil {
				t.Fatalf("NewClient(BaseURL: %q) succeeded, want an error", baseURL)
			}
		})
	}
}

// TestNewClientAcceptsValidBaseURL is the positive-case counterpart to
// TestNewClientRejectsUnusableBaseURL/
// TestNewClientRejectsBaseURLWithUserinfoQueryOrFragment: every prior
// BaseURL test in this file only asserts rejection, so a future change
// that made validation stricter in some new way (e.g. rejecting a bare
// host with no path, or an IPv6 host) could ship without any test
// noticing a previously-valid configuration broke. Found by the
// 2026-07-22 (1b428f6) maintainability assessment. Also asserts buildURL
// preserves each base's host/port/path correctly when joining an
// endpoint, not just that construction succeeds.
func TestNewClientAcceptsValidBaseURL(t *testing.T) {
	tests := []struct {
		baseURL  string
		wantHost string
		wantURL  string // buildURL("test", nil)
	}{
		{"https://example.com", "example.com", "https://example.com/test"},
		{"https://example.com/stats", "example.com", "https://example.com/stats/test"},
		{"http://localhost:8080", "localhost:8080", "http://localhost:8080/test"},
		{"http://127.0.0.1:8080", "127.0.0.1:8080", "http://127.0.0.1:8080/test"},
		{"http://[::1]:8080", "[::1]:8080", "http://[::1]:8080/test"},
		{"https://sub.example.com:8443/base/path", "sub.example.com:8443", "https://sub.example.com:8443/base/path/test"},
	}

	for _, tt := range tests {
		t.Run(tt.baseURL, func(t *testing.T) {
			c, err := NewClient(Config{BaseURL: tt.baseURL})
			if err != nil {
				t.Fatalf("NewClient(BaseURL: %q) error = %v, want success", tt.baseURL, err)
			}
			if c.baseURL.Host != tt.wantHost {
				t.Errorf("baseURL.Host = %q, want %q", c.baseURL.Host, tt.wantHost)
			}
			if got := c.buildURL("test", nil); got != tt.wantURL {
				t.Errorf("buildURL(%q) = %q, want %q", tt.baseURL, got, tt.wantURL)
			}
		})
	}
}

// TestNewClientRejectsBaseURLWithUserinfoQueryOrFragment covers three
// BaseURL shapes that are syntactically valid absolute http(s) URLs -
// passing every check TestNewClientRejectsUnusableBaseURL exercises - but
// are never a sensible way to configure a BaseURL. Found by the
// 2026-07-22 (1b428f6) maintainability assessment, itself following a
// lead from an external review of v3.1.2.
func TestNewClientRejectsBaseURLWithUserinfoQueryOrFragment(t *testing.T) {
	for _, baseURL := range []string{
		"https://user:pass@example.com",
		"https://example.com?token=secret",
		"https://example.com#fragment",
	} {
		t.Run(baseURL, func(t *testing.T) {
			if _, err := NewClient(Config{BaseURL: baseURL}); err == nil {
				t.Fatalf("NewClient(BaseURL: %q) succeeded, want an error", baseURL)
			}
		})
	}
}

// TestNewClientRejectionErrorsDoNotLeakBaseURL covers every path in
// NewClient that can reject a BaseURL potentially carrying a credential
// or token: userinfo, a query string, an invalid scheme on an otherwise
// credential-bearing URL (the pre-existing instance of the same defect,
// present since v3.1.2's original scheme check, not introduced by
// v3.1.3's userinfo/query/fragment checks), and several shapes of
// url.Parse failure - an invalid percent-escape, an invalid port, and a
// malformed IPv6 host. The url.Parse-failure cases are the uncovered
// path this table originally missed (v3.1.4's fix wrapped url.Parse's
// own %w-formatted error, which embeds the complete input), and then,
// after v3.1.5 unwrapped to what its own comment incorrectly called an
// "input-free" reason, the third instance of the same defect class: the
// invalid-port and malformed-IPv6-host cases below leaked via that
// unwrapped reason too, since net/url builds those specific reasons
// directly from the input. Found by the 2026-07-22 (b3c605d, 0e400d1,
// and f4801ef) maintainability assessments across three consecutive
// cycles; fixed for good by returning a fixed, input-free message for
// every url.Parse failure rather than trying to find a "safe" layer of
// the parser's own error to render - see NewClient's comment. A caller
// passing "https://admin:secret@host" got back an error containing the
// literal string "admin:secret" in earlier versions - disclosing exactly
// the credential the check exists to keep out of use, in whatever logs
// or error trackers capture NewClient's error.
func TestNewClientRejectionErrorsDoNotLeakBaseURL(t *testing.T) {
	tests := []struct {
		name    string
		baseURL string
		secrets []string
	}{
		{"userinfo", "https://admin:hunter2@example.com", []string{"admin", "hunter2"}},
		{"query", "https://example.com?api_key=hunter2", []string{"api_key", "hunter2"}},
		{"fragment", "https://example.com#hunter2", []string{"hunter2"}},
		{"invalid scheme with userinfo", "ftp://admin:hunter2@example.com", []string{"admin", "hunter2"}},
		{"parse failure with userinfo", "https://admin:hunter2@example.com/%zz", []string{"admin", "hunter2"}},
		{"parse failure with query token", "https://example.com/%zz?token=hunter2", []string{"token", "hunter2"}},
		{"parse failure with secret in invalid port", "https://example.com:sk_live_123/path", []string{"sk_live_123"}},
		{"parse failure with secret in malformed IPv6 host", "https://[::1sk_live_123]:443/path", []string{"sk_live_123"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewClient(Config{BaseURL: tt.baseURL})
			if err == nil {
				t.Fatalf("NewClient(BaseURL: %q) succeeded, want an error", tt.baseURL)
			}
			for _, secret := range tt.secrets {
				if strings.Contains(err.Error(), secret) {
					t.Errorf("NewClient(BaseURL: %q) error = %q leaks secret %q", tt.baseURL, err.Error(), secret)
				}
			}
		})
	}
}

// FuzzNewClientErrorDoesNotEchoInput is the structural counterpart to
// TestNewClientRejectionErrorsDoNotLeakBaseURL: rather than relying on a
// hand-enumerated list of NewClient's rejection paths to be complete -
// exactly the technique that missed the url.Parse failure path until the
// 2026-07-22 (0e400d1) maintainability assessment found it via an
// external review - this asserts the underlying invariant directly, for
// whatever return path NewClient actually takes: if construction fails,
// a marker planted in the BaseURL never appears in the returned error.
//
// Caveat on what this test actually covers, added after the 2026-07-22
// (f4801ef) maintainability assessment found the invariant held for
// these five template positions but not for a marker placed in the port
// or in a malformed IPv6 host - two url.Parse failure shapes none of the
// original three templates exercised. NewClient's url.Parse-failure
// branch no longer depends on this test (or any enumeration) to hold the
// invariant: it now returns a fixed, input-free message regardless of
// input, so no future parser-error shape can reopen this defect class
// through that branch. The templates below (including the two added for
// port and host) exist as regression evidence for the specific inputs
// found across three cycles, not as the thing guaranteeing the
// invariant - that guarantee now comes from NewClient's implementation,
// not from this test enumerating every parser failure mode.
//
// Markers are restricted to at least 4 characters and containing a
// digit: NewClient's own error text is hand-written English prose with
// no digits anywhere in it (confirmed by reading every message in
// NewClient), so a marker without this constraint can - and, verified
// directly during development of this test, does - spuriously "leak" by
// coincidentally matching a substring of a real word (e.g. a fuzzed
// marker "esca" against the stdlib's "invalid URL escape" text is not a
// leak of anything). Requiring a digit is also a more realistic shape
// for a secret than arbitrary English-letter fuzzing: API keys and
// tokens are essentially never purely alphabetic.
func FuzzNewClientErrorDoesNotEchoInput(f *testing.F) {
	for _, seed := range []string{
		"hunter2",
		"sk_live_abcdef0123456789",
		"a1%zzbadescape2",
		"秘密トークン0123",
		"a marker 42 with spaces",
	} {
		f.Add(seed)
	}

	templates := []string{
		"https://MARKER@example.com",       // userinfo
		"https://example.com?token=MARKER", // query string
		"https://MARKER@example.com/%zz",   // forces the url.Parse failure path, with a credential present
		"https://example.com:MARKER/path",  // forces an invalid-port url.Parse failure, marker in the port
		"https://[::1MARKER]:443/path",     // forces a malformed-IPv6-host url.Parse failure, marker in the host
	}

	f.Fuzz(func(t *testing.T, marker string) {
		if len(marker) < 4 || !strings.ContainsAny(marker, "0123456789") {
			t.Skip()
		}
		for _, tmpl := range templates {
			baseURL := strings.ReplaceAll(tmpl, "MARKER", marker)
			_, err := NewClient(Config{BaseURL: baseURL})
			if err == nil {
				continue // a BaseURL this marker happens to make valid isn't this test's concern
			}
			if strings.Contains(err.Error(), marker) {
				t.Fatalf("NewClient(BaseURL: %q) error = %q leaks marker %q", baseURL, err.Error(), marker)
			}
		}
	})
}

func TestClientHeaderMutationsAreSafeDuringRequests(t *testing.T) {
	client := mustNewClient(t, Config{
		BaseURL: "https://api.example.com",
		HTTPClient: httpClientFunc(func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Header:     make(http.Header),
				Body:       io.NopCloser(http.NoBody),
			}, nil
		}),
	})

	const iterations = 200
	errs := make(chan error, iterations)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < iterations; i++ {
			client.SetHeader("X-Request-ID", "set")
			client.AddHeader("X-Request-ID", "add")
			client.SetHeaders(http.Header{"X-Request-ID": {"replace"}})
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < iterations; i++ {
			if _, err := client.Get(context.Background(), "health", nil); err != nil {
				errs <- err
			}
		}
	}()
	wg.Wait()
	close(errs)
	for err := range errs {
		t.Errorf("Get() returned error: %v", err)
	}
}

// TestClientTimeoutAppliesWithCustomHTTPClient guards the fix for Config.Timeout
// being silently ignored whenever a caller supplies their own HTTPClient (which
// the SDK can't stamp with an http.Client.Timeout). Get now imposes the timeout
// as a per-request context deadline, so a custom client that respects the
// request context is still bounded by it.
func TestClientTimeoutAppliesWithCustomHTTPClient(t *testing.T) {
	client := mustNewClient(t, Config{
		BaseURL: "https://api.example.com",
		Timeout: 20 * time.Millisecond,
		// A custom client with no timeout of its own that only returns once
		// the request context is done - i.e. it would hang forever if Get
		// didn't impose a deadline.
		HTTPClient: httpClientFunc(func(req *http.Request) (*http.Response, error) {
			<-req.Context().Done()
			return nil, req.Context().Err()
		}),
	})

	done := make(chan error, 1)
	go func() {
		_, err := client.Get(context.Background(), "health", nil)
		done <- err
	}()

	select {
	case err := <-done:
		if err == nil {
			t.Fatal("Get() returned nil error; expected a deadline-exceeded failure")
		}
		if !errors.Is(err, context.DeadlineExceeded) {
			t.Errorf("Get() error = %v; want it to wrap context.DeadlineExceeded", err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("Get() did not return within 2s; Config.Timeout was not enforced for the custom HTTPClient")
	}
}

// TestClientNegativeTimeoutDisablesContextDeadline guards the documented
// behavior of a negative Config.Timeout: unlike 0 (normalized to
// DefaultTimeout), a negative value is left as-is and Get's context-deadline
// branch (c.timeout > 0) is skipped entirely, so the request context carries
// no deadline from the client at all.
func TestClientNegativeTimeoutDisablesContextDeadline(t *testing.T) {
	var sawDeadline bool
	client := mustNewClient(t, Config{
		BaseURL: "https://api.example.com",
		Timeout: -1,
		HTTPClient: httpClientFunc(func(req *http.Request) (*http.Response, error) {
			_, sawDeadline = req.Context().Deadline()
			return &http.Response{StatusCode: 200, Body: http.NoBody, Header: make(http.Header)}, nil
		}),
	})

	if _, err := client.Get(context.Background(), "health", nil); err != nil {
		t.Fatalf("Get() error = %v", err)
	}
	if sawDeadline {
		t.Error("request context had a deadline; negative Config.Timeout should disable the context-deadline enforcement entirely")
	}
}

// TestClientNegativeTimeoutSDKBuiltClient guards the same behavior for the
// other half of the "uniform" timeout mechanism: when Config.HTTPClient is
// left nil, the SDK builds its own http.Client and stamps it with
// config.Timeout directly. A negative value must pass through unmodified
// (net/http treats <= 0 as "no timeout"), not get normalized like 0 does.
func TestClientNegativeTimeoutSDKBuiltClient(t *testing.T) {
	client := mustNewClient(t, Config{
		BaseURL: "https://api.example.com",
		Timeout: -1,
	})

	httpClient, ok := client.httpClient.(*http.Client)
	if !ok {
		t.Fatalf("client.httpClient is %T, want *http.Client (SDK-built path)", client.httpClient)
	}
	if httpClient.Timeout != -1 {
		t.Errorf("httpClient.Timeout = %v, want -1 (negative Config.Timeout must not be normalized)", httpClient.Timeout)
	}
}
