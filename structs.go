package main

// Teams : Teams as returned by the NHL API
type Teams struct {
	Copyright string `json:"copyright"`
	Teams     []struct {
		Abbreviation string `json:"abbreviation"`
		Active       bool   `json:"active"`
		Conference   struct {
			ID   int64  `json:"id"`
			Link string `json:"link"`
			Name string `json:"name"`
		} `json:"conference"`
		Division struct {
			Abbreviation string `json:"abbreviation"`
			ID           int64  `json:"id"`
			Link         string `json:"link"`
			Name         string `json:"name"`
			NameShort    string `json:"nameShort"`
		} `json:"division"`
		FirstYearOfPlay string `json:"firstYearOfPlay"`
		Franchise       struct {
			FranchiseID int64  `json:"franchiseId"`
			Link        string `json:"link"`
			TeamName    string `json:"teamName"`
		} `json:"franchise"`
		FranchiseID     int64  `json:"franchiseId"`
		ID              int64  `json:"id"`
		Link            string `json:"link"`
		LocationName    string `json:"locationName"`
		Name            string `json:"name"`
		OfficialSiteURL string `json:"officialSiteUrl"`
		Roster          struct {
			Link   string `json:"link"`
			Roster []struct {
				JerseyNumber string `json:"jerseyNumber"`
				Person       struct {
					FullName string `json:"fullName"`
					ID       int64  `json:"id"`
					Link     string `json:"link"`
				} `json:"person"`
				Position struct {
					Abbreviation string `json:"abbreviation"`
					Code         string `json:"code"`
					Name         string `json:"name"`
					Type         string `json:"type"`
				} `json:"position"`
			} `json:"roster"`
		} `json:"roster"`
		ShortName string `json:"shortName"`
		TeamName  string `json:"teamName"`
		Venue     struct {
			City     string `json:"city"`
			ID       int64  `json:"id"`
			Link     string `json:"link"`
			Name     string `json:"name"`
			TimeZone struct {
				ID     string `json:"id"`
				Offset int64  `json:"offset"`
				Tz     string `json:"tz"`
			} `json:"timeZone"`
		} `json:"venue"`
	} `json:"teams"`
}

// People : A list of playerProfiles
type People struct {
	Copyright string          `json:"copyright"`
	People    []PlayerProfile `json:"people"`
}

// PlayerProfile : A profile of a player
type PlayerProfile struct {
	Active           bool   `json:"active"`
	AlternateCaptain bool   `json:"alternateCaptain"`
	BirthCity        string `json:"birthCity"`
	BirthCountry     string `json:"birthCountry"`
	BirthDate        string `json:"birthDate"`
	Captain          bool   `json:"captain"`
	CurrentAge       int64  `json:"currentAge"`
	CurrentTeam      struct {
		ID   int64  `json:"id"`
		Link string `json:"link"`
		Name string `json:"name"`
	} `json:"currentTeam"`
	FirstName       string `json:"firstName"`
	FullName        string `json:"fullName"`
	Height          string `json:"height"`
	ID              int64  `json:"id"`
	LastName        string `json:"lastName"`
	Link            string `json:"link"`
	Nationality     string `json:"nationality"`
	PrimaryNumber   string `json:"primaryNumber"`
	PrimaryPosition struct {
		Abbreviation string `json:"abbreviation"`
		Code         string `json:"code"`
		Name         string `json:"name"`
		Type         string `json:"type"`
	} `json:"primaryPosition"`
	Rookie        bool   `json:"rookie"`
	RosterStatus  string `json:"rosterStatus"`
	ShootsCatches string `json:"shootsCatches"`
	Weight        int64  `json:"weight"`
}

// PlayerStats : List of Stats for an individual player
type PlayerStats struct {
	Copyright string `json:"copyright"`
	Stats     []struct {
		Splits []struct {
			Season string     `json:"season"`
			Stat   PlayerStat `json:"stat"`
		} `json:"splits"`
		Type struct {
			DisplayName string `json:"displayName"`
			GameType    struct {
				Description string `json:"description"`
				ID          string `json:"id"`
				Postseason  bool   `json:"postseason"`
			} `json:"gameType"`
		} `json:"type"`
	} `json:"stats"`
}

// PlayerStat : Stats for an individual player
type PlayerStat struct {
	Assists                     int64   `json:"assists"`
	Blocked                     int64   `json:"blocked"`
	EvenTimeOnIce               string  `json:"evenTimeOnIce"`
	EvenTimeOnIcePerGame        string  `json:"evenTimeOnIcePerGame"`
	FaceOffPct                  float64 `json:"faceOffPct"`
	GameWinningGoals            int64   `json:"gameWinningGoals"`
	Games                       int64   `json:"games"`
	Goals                       int64   `json:"goals"`
	Hits                        int64   `json:"hits"`
	OverTimeGoals               int64   `json:"overTimeGoals"`
	PenaltyMinutes              string  `json:"penaltyMinutes"`
	Pim                         int64   `json:"pim"`
	PlusMinus                   int64   `json:"plusMinus"`
	Points                      int64   `json:"points"`
	PowerPlayGoals              int64   `json:"powerPlayGoals"`
	PowerPlayPoints             int64   `json:"powerPlayPoints"`
	PowerPlayTimeOnIce          string  `json:"powerPlayTimeOnIce"`
	PowerPlayTimeOnIcePerGame   string  `json:"powerPlayTimeOnIcePerGame"`
	Shifts                      int64   `json:"shifts"`
	ShortHandedGoals            int64   `json:"shortHandedGoals"`
	ShortHandedPoints           int64   `json:"shortHandedPoints"`
	ShortHandedTimeOnIce        string  `json:"shortHandedTimeOnIce"`
	ShortHandedTimeOnIcePerGame string  `json:"shortHandedTimeOnIcePerGame"`
	ShotPct                     float64 `json:"shotPct"`
	Shots                       int64   `json:"shots"`
	TimeOnIce                   string  `json:"timeOnIce"`
	TimeOnIcePerGame            string  `json:"timeOnIcePerGame"`
	EvenSaves                   int64   `json:"evenSaves"`
	EvenShots                   int64   `json:"evenShots"`
	EvenStrengthSavePercentage  float64 `json:"evenStrengthSavePercentage"`
	GamesStarted                int64   `json:"gamesStarted"`
	GoalAgainstAverage          float64 `json:"goalAgainstAverage"`
	GoalsAgainst                int64   `json:"goalsAgainst"`
	Losses                      int64   `json:"losses"`
	Ot                          int64   `json:"ot"`
	PowerPlaySavePercentage     float64 `json:"powerPlaySavePercentage"`
	PowerPlaySaves              int64   `json:"powerPlaySaves"`
	PowerPlayShots              int64   `json:"powerPlayShots"`
	SavePercentage              float64 `json:"savePercentage"`
	Saves                       int64   `json:"saves"`
	ShortHandedSavePercentage   float64 `json:"shortHandedSavePercentage"`
	ShortHandedSaves            int64   `json:"shortHandedSaves"`
	ShortHandedShots            int64   `json:"shortHandedShots"`
	ShotsAgainst                int64   `json:"shotsAgainst"`
	Shutouts                    int64   `json:"shutouts"`
	Ties                        int64   `json:"ties"`
	Wins                        int64   `json:"wins"`
}

// Player :  A player has a profile and a stat
type Player struct {
	Profile PlayerProfile `json:"profile"`
	Stat    PlayerStat    `json:"stat"`
}
