package main

type Player struct {
	Club                 PlayerClub    `json:"club"`
	IsPro                bool          `json:"isQualifiedFromChampionshipChallenge"`
	Win3x3               int           `json:"3vs3Victories"`
	Icon                 PlayerIcon    `json:"icon"`
	Tag                  string        `json:"tag"`
	Name                 string        `json:"name"`
	Trophies             int           `json:"trophies"`
	ExpLevel             int           `json:"expLevel"`
	ExpPoints            int           `json:"expPoints"`
	HighestTrophies      int           `json:"highestTrophies"`
	SoloVictories        int           `json:"soloVictories"`
	DuoVictories         int           `json:"duoVictories"`
	BestRoboRumbleTime   int           `json:"bestRoboRumbleTime"`
	BestTimeAsBigBrawler int           `json:"bestTimeAsBigBrawler"`
	Brawlers             []BrawlerStat `json:"brawlers"`
	NameColor            string        `json:"nameColor"`
}

type PlayerClub struct {
	Tag  string `json:"tag"`
	Name string `json:"name"`
}

type PlayerIcon struct {
	Id int `json:"id"`
}

type Brawler struct {
	Name       string      `json:"name"`
	Gadgets    []Accessory `json:"gadgets"`
	StarPowers []StarPower `json:"starPowers"`
	Id         int         `json:"id"`
}

type BrawlerList struct {
	Brawlers []Brawler `json:"items"`
}

type BrawlerStat struct {
	Name            string      `json:"name"`
	Gadgets         []Accessory `json:"gadgets"`
	StarPowers      []StarPower `json:"starPowers"`
	Id              int         `json:"id"`
	Rank            int         `json:"rank"`
	Trophies        int         `json:"trophies"`
	HighestTrophies int         `json:"highestTrophies"`
	Power           int         `json:"power"`
	Gears           []GearInfo  `json:"gears"`
}

type Accessory struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type GearInfo struct {
	Name  string `json:"name"`
	Id    int    `json:"id"`
	Level int    `json:"level"`
}

type StarPower struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type BrawlerStatistics struct {
	PowerPoints int
	Coins       int
	Gadgets     int
	StarPowers  int
	Gears       int
}
