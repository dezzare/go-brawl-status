package main

const maxLevel int = 11
const minGears int = 2
const minGadgets int = 1
const minStarPowers int = 1
const gadgetCost = 1000
const gearCost = 1000
const starPowerCost = 2000
const hyperchargeCost = 5000

var ppNeeded = []int{0, 20, 30, 50, 80, 130, 210, 340, 550, 890, 1440}
var coinsNeeded = []int{0, 20, 35, 75, 140, 290, 480, 800, 1250, 1875, 2800}

func newBrawlerStatistics() BrawlerStatistics {
	var bs BrawlerStatistics
	return bs
}

func (b *BrawlerStatistics) getPowerPoints() int {
	return b.PowerPoints
}

func (b *BrawlerStatistics) getCoins() int {
	return b.Coins
}

func (b *BrawlerStatistics) getGadgetsCoins() int {
	return b.Gadgets * gadgetCost
}

func (b *BrawlerStatistics) getStarPowersCoins() int {
	return b.StarPowers * starPowerCost
}

func (b *BrawlerStatistics) getGearsCoins() int {
	return b.Gears * gearCost
}

func (b *BrawlerStatistics) getTotalCoins() int {
	total := b.getCoins() + b.getGearsCoins() + b.getGadgetsCoins() + b.getStarPowersCoins()
	return total
}

func (b *BrawlerStatistics) addPowerPoints(pp int) {
	b.PowerPoints = b.PowerPoints + pp
}

func (b *BrawlerStatistics) addCoins(c int) {
	b.Coins = b.Coins + c
}

func (b *BrawlerStatistics) addGadgets(g int) {
	b.Gadgets = b.Gadgets + g
}

func (b *BrawlerStatistics) addStarPowers(sp int) {
	b.StarPowers = b.StarPowers + sp
}

func (b *BrawlerStatistics) addGears(g int) {
	b.Gears = b.Gears + g
}

func (b *BrawlerStatistics) setStatus(bs []BrawlerStat) error {
	for _, v := range bs {
		if len(v.Gadgets) < minGadgets {
			b.addGadgets(minGadgets - len(v.Gadgets))
		}

		if len(v.Gears) < minGears {
			b.addGears(minGears - len(v.Gears))
		}

		if len(v.StarPowers) < minStarPowers {
			b.addStarPowers(minStarPowers - len(v.StarPowers))

		}
		if v.Power < int(maxLevel) {
			needs := addUntilMaxLevel(v.Power, 0, 0)
			b.addCoins(needs[0])
			b.addPowerPoints(needs[1])
		}

	}

	return nil
}
