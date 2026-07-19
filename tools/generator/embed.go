package main

import "embed"

// templatesFS embeds the generator's templates so loadTemplate doesn't
// depend on the process's working directory. Previously it resolved
// tools/generator/templates/<name>.tmpl relative to the CWD, which broke
// the documented `cd tools/generator && go run . -endpoint X` workflow
// (CWD is tools/generator, so the path became
// tools/generator/tools/generator/templates/endpoint.tmpl).
//
//go:embed templates/*.tmpl
var templatesFS embed.FS

// fieldTypesFS embeds fieldtypes.json, the canonical field-name -> Go-type
// dictionary. See loadFieldTypes in generator.go for how it's used.
//
//go:embed fieldtypes.json
var fieldTypesFS embed.FS
