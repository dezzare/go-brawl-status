package main

import (
	"fmt"
	"os"
)

func main() {
	c := newClient()
	tag := os.Args[1]

	if err := savePlayer(tag, c); err != nil {
		fmt.Println(err)
	}

	if err := saveBrawlers(c); err != nil {
		fmt.Println(err)
	}
	player := getPlayer()
	brawlers := getBrawlerList()
	brawlerStatistics := newBrawlerStatistics()

	if err := hasAllBrawlers(player, brawlers, &brawlerStatistics); err != nil {
		fmt.Println(err)
	}

	if err := brawlerStatistics.setStatus(player.Brawlers); err != nil {
		fmt.Println(err)
	}

	if err := printStatistics(&brawlerStatistics); err != nil {
		fmt.Println(err)
	}
}

func init() {
	loadConfig()
}
