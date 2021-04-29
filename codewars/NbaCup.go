package codewars

import (
	"fmt"
	"strconv"
	"strings"
)

type teaminfo struct {
	wins     int
	draws    int
	losses   int
	scored   int
	conceded int
	points   int
}

func NbaCup(resultSheet, toFind string) string {
	teamDataMap := map[string]teaminfo{}
	for _, info := range strings.Split(resultSheet, ",") {
		teamAName := ""
		teamBName := ""
		teamAScores := 0
		teamBScores := 0
		teamAStats := [4]int{}
		teamBStats := [4]int{}
		var doneWithTeamAName = false
		var doneWithTeamBName = false
		for _, chunk := range strings.Split(info, " ") {
			if score, ok := strconv.Atoi(string(chunk)); ok == nil {
				if !doneWithTeamAName {
					doneWithTeamAName = true
					teamAScores = score
					continue
				} else {
					doneWithTeamBName = true
					teamBScores = score
					continue
				}
			}
			if !doneWithTeamAName {
				teamAName += " " + chunk
			} else if !doneWithTeamBName {
				teamBName += " " + chunk
			}
		}

		if teamAScores == teamBScores {
			teamAStats[0] = 1
			teamAStats[3] = 0
			teamBStats[0] = 1
			teamBStats[3] = 0
		} else if teamAScores < teamBScores {
			teamAStats[2] = 1
			teamBStats[1] = 1
			teamAStats[3] = 0
			teamBStats[3] = 3
		} else {
			teamAStats[1] = 1
			teamBStats[2] = 1
			teamBStats[3] = 0
			teamAStats[3] = 3
		}

		updateData(teamDataMap, strings.Trim(teamAName, " "), teamAStats, teamAScores, teamBScores)
		updateData(teamDataMap, strings.Trim(teamBName, " "), teamBStats, teamBScores, teamAScores)
	}

	data, ok := teamDataMap[toFind]
	if !ok {
		return fmt.Sprintf("%v:This team didn't play!", toFind)
	} else {
		return fmt.Sprintf("%v:W=%v;D=%v;L=%v;Scored=%v;Conceded=%v;Points=%v", toFind, data.wins, data.draws, data.losses, data.scored, data.conceded, data.points)
	}
}

func updateData(teamData map[string]teaminfo, teamname string, stats [4]int, score int, conceded int) {
	if _, k := teamData[teamname]; !k {
		teamData[teamname] = teaminfo{
			draws:    stats[0],
			wins:     stats[1],
			losses:   stats[2],
			points:   stats[3],
			scored:   score,
			conceded: conceded,
		}
	} else {
		var v teaminfo = teamData[teamname]
		teamData[teamname] = teaminfo{
			draws:    v.draws + stats[0],
			wins:     v.wins + stats[1],
			losses:   v.losses + stats[2],
			points:   v.points + stats[3],
			scored:   v.scored + score,
			conceded: v.conceded + conceded,
		}
	}
}
