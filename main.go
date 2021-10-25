package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func getRoster(teamLink string) (players []string) {
	base := "https://statsapi.web.nhl.com"
	URI := base + teamLink + "?expand=team.roster&season=20212022"
	resp, err := http.Get(URI)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var teams Teams
	err = json.Unmarshal(body, &teams)
	if err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	for _, team := range teams.Teams {
		for _, player := range team.Roster.Roster {
			players = append(players, player.Person.Link)
		}
	}
	return players
}

func getPlayerStats(URI string) (stats PlayerStats) {
	resp, err := http.Get(URI)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &stats)
	if err != nil {
		fmt.Println("Can not unmarshal Stat JSON")
		fmt.Println(err)
	}
	return stats
}

func getPlayer(playerLink string) (customPlayer Player) {
	base := "https://statsapi.web.nhl.com"
	URI := base + playerLink
	resp, err := http.Get(URI)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var people People
	err = json.Unmarshal(body, &people)
	if err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	for _, player := range people.People {
		customPlayer.Profile = player
		statsURI := URI + "/stats?stats=statsSingleSeason&season=20212022"
		stats := getPlayerStats(statsURI)
		for _, stat := range stats.Stats {
			for _, split := range stat.Splits {
				customPlayer.Stat = split.Stat
			}
		}
	}
	return customPlayer
}

func getCaptain(playerLink string) {
	base := "https://statsapi.web.nhl.com"
	URI := base + playerLink
	resp, err := http.Get(URI)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var people People
	err = json.Unmarshal(body, &people)
	if err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	var captain string
	for _, player := range people.People {
		if player.Captain {
			captain = "Captain"
		} else if player.AlternateCaptain {
			captain = "Alternate Captain"
		}
		if captain != "" {
			fmt.Printf("%s : %s, #%s, %s, Age: %d, Height %s, Weight %d\n", player.FullName, captain, player.PrimaryNumber, player.PrimaryPosition.Name, player.CurrentAge, player.Height, player.Weight)
		}
	}
}

func getTeams(teamName string) (teamLinks []string) {
	resp, err := http.Get("https://statsapi.web.nhl.com/api/v1/teams?season=20212022")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var teams Teams
	err = json.Unmarshal(body, &teams)
	if err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	for _, team := range teams.Teams {
		if strings.Contains(strings.ToLower(team.Name), strings.ToLower(teamName)) {
			teamLinks = append(teamLinks, team.Link)
		}
	}
	return teamLinks
}

func printRoster(teams []string) {
	for _, teamLink := range teams {
		var customPlayers []Player
		players := getRoster(teamLink)
		for _, p := range players {
			player := getPlayer(p)
			customPlayers = append(customPlayers, player)
		}
		fmt.Println("================ Forwards ================")
		fmt.Println("Name, Number, Position, Age, Height, Weight, Goals, Assists, +-, Shots, Time On Ice")
		for _, player := range customPlayers {
			if player.Profile.PrimaryPosition.Name != "Defenseman" {
				if player.Profile.PrimaryPosition.Name != "Goalie" {
					fmt.Printf("%s, #%s, %s, %d, %s, %d, %d, %d, %d, %d, %s\n", player.Profile.FullName, player.Profile.PrimaryNumber, player.Profile.PrimaryPosition.Name, player.Profile.CurrentAge,
						player.Profile.Height, player.Profile.Weight, player.Stat.Goals, player.Stat.Assists, player.Stat.PlusMinus, player.Stat.Shots, player.Stat.TimeOnIce)
				}
			}
		}
		fmt.Println("================ Defensemen ================")
		fmt.Println("Name, Number, Position, Age, Height, Weight, Goals, Assists, +-, Shots, Time On Ice")
		for _, player := range customPlayers {
			if player.Profile.PrimaryPosition.Name == "Defenseman" {
				fmt.Printf("%s, #%s, %s, %d, %s, %d, %d, %d, %d, %d, %s\n", player.Profile.FullName, player.Profile.PrimaryNumber, player.Profile.PrimaryPosition.Name, player.Profile.CurrentAge,
					player.Profile.Height, player.Profile.Weight, player.Stat.Goals, player.Stat.Assists, player.Stat.PlusMinus, player.Stat.Shots, player.Stat.TimeOnIce)
			}
		}
		fmt.Println("================ Goalies ================")
		fmt.Println("Name, Number, Position, Age, Height, Weight, Wins, Losses, SavePercentage, EvenStrengthSavePercentage, PowerPlaySavePercentage, ShortHandedSavePercentage, Time On Ice")
		for _, player := range customPlayers {
			if player.Profile.PrimaryPosition.Name == "Goalie" {
				fmt.Printf("%s, #%s, %s, %d, %s, %d, %d, %d, %f, %f, %f, %f, %s\n", player.Profile.FullName, player.Profile.PrimaryNumber, player.Profile.PrimaryPosition.Name, player.Profile.CurrentAge,
					player.Profile.Height, player.Profile.Weight, player.Stat.Wins, player.Stat.Losses, player.Stat.SavePercentage, player.Stat.EvenStrengthSavePercentage,
					player.Stat.PowerPlaySavePercentage, player.Stat.ShortHandedSavePercentage, player.Stat.TimeOnIce)
			}
		}
	}
}

func main() {
	teamPtr := flag.String("team", "Carolina Hurricanes", "Team to lookup")
	captainsPtr := flag.Bool("captains", false, "To print the captains of the team or not")
	rosterPtr := flag.Bool("roster", true, "To print the team roster or not")
	flag.Parse()

	teams := getTeams(*teamPtr)
	if *captainsPtr {
		for _, teamLink := range teams {
			players := getRoster(teamLink)
			for _, player := range players {
				getCaptain(player)
			}
		}
		os.Exit(0)
	}
	if *rosterPtr {
		printRoster(teams)
		os.Exit(0)
	}

}
