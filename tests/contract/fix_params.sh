#!/bin/bash

# Fix parameters that need pointer wrappers
sed -i '' '
  # PlayerProfileV2
  /PlayerProfileV2Request/,/^[[:space:]]*}/ {
    s/PerMode:[[:space:]]*parameters\.PerModePerGame,/PerMode:  perModePtr(string(parameters.PerModePerGame)),/
    s/LeagueID:[[:space:]]*parameters\.LeagueIDNBA,/LeagueID: leagueIDPtr(string(parameters.LeagueIDNBA)),/
  }
  
  # TeamInfoCommonRequest - remove Season field and fix LeagueID and SeasonType
  /TeamInfoCommonRequest/,/^[[:space:]]*}/ {
    /Season:[[:space:]]*parameters\.Season/d
    s/LeagueID:[[:space:]]*parameters\.LeagueIDNBA,/LeagueID:   leagueIDPtr(string(parameters.LeagueIDNBA)),/
    s/SeasonType:[[:space:]]*parameters\.SeasonTypeRegular,/SeasonType: seasonTypePtr(string(parameters.SeasonTypeRegular)),/
  }
  
  # CommonTeamRosterRequest
  /CommonTeamRosterRequest/,/^[[:space:]]*}/ {
    s/Season:[[:space:]]*parameters\.Season/Season:   (*parameters.Season)(stringPtr(testSeason)),/
    s/LeagueID:[[:space:]]*parameters\.LeagueIDNBA,/LeagueID: leagueIDPtr(string(parameters.LeagueIDNBA)),/
  }
  
  # CommonAllPlayersRequest - fix IsOnlyCurrentSeason
  /CommonAllPlayersRequest/,/^[[:space:]]*}/ {
    s/IsOnlyCurrentSeason:[[:space:]]*intPtr(1),/IsOnlyCurrentSeason: intPtr(1),/
  }
  
  # ShotChartDetail - fix TeamID
  /ShotChartDetailRequest/,/^[[:space:]]*}/ {
    s/TeamID:[[:space:]]*intPtr(0),/TeamID:         stringPtr("0"),/
  }
' endpoints_test.go

echo "Fixed parameter types"
