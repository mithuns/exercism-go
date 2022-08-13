package tournament

import (
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type TeamNames string

const (
	AllegoricAlaskans      TeamNames = "Allegoric Alaskians"
	BlitheringBadgers      TeamNames = "Blithering Badgers"
	DevastatingDonkeys     TeamNames = "Devastating Donkeys"
	CourageousCalifornians TeamNames = "Courageous Californians"
)

type MatchOutcome string

const (
	Win  MatchOutcome = "win"
	Loss MatchOutcome = "loss"
	Draw MatchOutcome = "draw"
)

type TallyOutput struct {
	matchesPlayed int
	matchesWon    int
	matchesLost   int
	matchesDrawn  int
	totalPoints   int
}

//TODO type TallyResults map[TeamNames]TallyOutput

func Tally(reader io.Reader, writer io.Writer) error {
	b := make([]byte, 512)
	inputLines := make([]string, 0)
	for {
		n, err := reader.Read(b)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if n > 0 {
			inputLines = append(inputLines, string(b[0:n]))
		}
	}
	if output, err := computeTally(inputLines); err != nil {
		return err
	} else {
		sortedKeys := sortData(output)
		if _, err := writer.Write([]byte(prettyPrint(sortedKeys, output))); err != nil {
			return err
		}
	}

	return nil
}

func computeTally(inputLines []string) (map[TeamNames]TallyOutput, error) {
	tallyResults := make(map[TeamNames]TallyOutput)
	inputLines2 := strings.Split(inputLines[0], "\n")
	for i := range inputLines2 {
		line := strings.Trim(inputLines2[i], " ")
		if len(line) > 0 {
			tokens := strings.Split(inputLines2[i], ";")
			if strings.HasPrefix(line, "#") {
				continue
			}
			if len(tokens) != 3 {
				return nil, errors.New("ignore comments and newlines")
			}
			if !validateTeamName(TeamNames(tokens[0])) || !validateTeamName(TeamNames(tokens[1])) {
				return nil, errors.New("invalid team name")
			}
			if err := updateTallyForGivenTeam(TeamNames(tokens[0]), TeamNames(tokens[1]), tokens[2], tallyResults); err != nil {
				return nil, errors.New("invalid result")
			}

		}
	}
	return tallyResults, nil
}

func prettyPrint(sortedKeys []TeamNames, tallyResults map[TeamNames]TallyOutput) string {
	header := fmt.Sprintln("Team                           | MP |  W |  D |  L |  P")
	format := "%-31s|%3d |%3d |%3d |%3d |%3d"

	for key := range sortedKeys {
		info := fmt.Sprintf(format, sortedKeys[key], tallyResults[sortedKeys[key]].matchesPlayed, tallyResults[sortedKeys[key]].matchesWon, tallyResults[sortedKeys[key]].matchesDrawn, tallyResults[sortedKeys[key]].matchesLost, tallyResults[sortedKeys[key]].totalPoints)
		header = header + fmt.Sprintln(info)
	}
	return header
}
func updateTallyForGivenTeam(teamA TeamNames, teamB TeamNames, matchResult string, tallyResults map[TeamNames]TallyOutput) error {
	teamATally := tallyResults[teamA]
	teamBTally := tallyResults[teamB]

	teamATally.matchesPlayed++
	teamBTally.matchesPlayed++

	switch matchResult {
	case "win":
		teamATally.matchesWon += 1
		teamATally.totalPoints += 3
		teamBTally.matchesLost += 1
	case "loss":
		teamATally.matchesLost += 1
		teamBTally.matchesWon += 1
		teamBTally.totalPoints += 3
	case "draw":
		teamATally.matchesDrawn += 1
		teamBTally.matchesDrawn += 1

		teamATally.totalPoints += 1
		teamBTally.totalPoints += 1
	default:
		return errors.New("")
	}
	tallyResults[teamA] = teamATally
	tallyResults[teamB] = teamBTally
	return nil
}

// sortData sorts keys for value totalpoints, in case of tie, sorts by key alphabetically
func sortData(tallyResults map[TeamNames]TallyOutput) []TeamNames {
	keys := make([]TeamNames, 0)
	for k := range tallyResults {
		keys = append(keys, k)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		if tallyResults[keys[i]].totalPoints == tallyResults[keys[j]].totalPoints {
			return keys[i] < keys[j]
		}
		return tallyResults[keys[i]].totalPoints > tallyResults[keys[j]].totalPoints
	})
	return keys
}

func validateTeamName(teamname TeamNames) bool {
	if teamname == AllegoricAlaskans || teamname == BlitheringBadgers || teamname == CourageousCalifornians || teamname == DevastatingDonkeys {
		return true
	}
	return false
}
