package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func getJsonFromFile[T interface{}](filename string, model *T) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("Open file error: %v", err)
	}
	if err := json.Unmarshal(file, &model); err != nil {
		return fmt.Errorf("Json Unmarshal error: %v", err)
	}
	return nil
}

func saveToJsonFile[T interface{}](data []byte, model T, filename string) error {

	if err := json.Unmarshal(data, &model); err != nil {
		return fmt.Errorf("Json Unmarshal error: %v", err)
	}

	file, err := json.MarshalIndent(&model, "", " ")
	if err != nil {
		return fmt.Errorf("Json Marshal error: %v", err)
	}
	err = os.WriteFile(filename, file, 0644)
	if err != nil {
		return fmt.Errorf("Write error: %v", err)
	}
	return nil
}

func contains(s []BrawlerStat, str string) bool {
	for _, v := range s {
		if v.Name == str {
			return true
		}
	}
	return false
}

func parseTag(tag string) string {
	return "%23" + strings.TrimPrefix(strings.TrimPrefix(tag, "#"), "%23")
}

func savePlayer(tag string, c *Client) error {
	playerTag := parseTag(tag)

	fmt.Printf("\nGetting player: %v", playerTag)
	player, err := c.getPlayer(playerTag)
	if err != nil {
		fmt.Printf("\nError: %v", err)
		return err
	}
	var model Player
	saveToJsonFile(player, model, "data-Player.json")
	return nil

}

func getPlayer() (p Player) {
	if err := getJsonFromFile("data-Player.json", &p); err != nil {
		fmt.Print(err)
	}
	return
}

func saveBrawlers(c *Client) error {

	fmt.Printf("\nGetting List of Brawlers")
	brawlers, err := c.getBrawlers()
	if err != nil {
		return err
	}
	var model BrawlerList

	saveToJsonFile(brawlers, model, "data-Brawlers-List.json")
	return nil
}

func getBrawlerList() (bl BrawlerList) {
	if err := getJsonFromFile("data-Brawlers-List.json", &bl); err != nil {
		fmt.Print(err)
	}
	return
}

func hasAllBrawlers(p Player, bl BrawlerList, bs *BrawlerStatistics) error {

	var missingBrawlers []string
	for _, v := range bl.Brawlers {
		if !contains(p.Brawlers, v.Name) {
			missingBrawlers = append(missingBrawlers, v.Name)
		}
	}

	if len(missingBrawlers) != 0 {
		bn := ""
		for _, s := range missingBrawlers {
			if s == "BUZZ LIGHTYEAR" {
				continue
			}
			needs := addUntilMaxLevel(0, 0, 0)
			bs.addCoins(needs[0])
			bs.addPowerPoints(needs[1])
			bs.addGadgets(minGadgets)
			bs.addGears(minGears)
			bs.addStarPowers(minStarPowers)
			bn = bn + " " + s
		}
		if len(bn) > 0 {
			fmt.Printf("\n\nAinda faltam %v brawlers", len(bl.Brawlers)-(len(bl.Brawlers)-len(missingBrawlers)))
			fmt.Printf("\nSendo eles: %s", bn)
			return nil
		}
	}
	fmt.Printf("\n%s tem todos os brawlers\n", p.Name)

	return nil
}

func addUntilMaxLevel(l int, c int, pp int) []int {
	if l == maxLevel {
		return []int{c, pp}
	}

	c = c + coinsNeeded[l]
	pp = pp + ppNeeded[l]
	l++

	return addUntilMaxLevel(l, c, pp)
}

func printStatistics(bs *BrawlerStatistics) error {

	fmt.Println("\n--- Statistics ---")
	fmt.Printf("Coins to Level 11:          %v\n", bs.getCoins())
	fmt.Printf("Power Points to Level 11:   %v\n", bs.getPowerPoints())
	fmt.Printf("Gadgets:                    %v = %v Coins\n", bs.Gadgets, bs.getGadgetsCoins())
	fmt.Printf("Gears:                      %v = %v Coins\n", bs.Gears, bs.getGearsCoins())
	fmt.Printf("StarPowers:                 %v = %v Coins\n", bs.StarPowers, bs.getStarPowersCoins())
	fmt.Printf("Total Coins:                %v\n", bs.getTotalCoins())

	return nil
}
