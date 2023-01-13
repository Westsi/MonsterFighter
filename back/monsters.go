package main

import "fmt"

func initVars() {
	initOriginalMonsters()
	initMonsterTypes()
}

var WormType MonsterType
var WormMonster OriginalMonster
var DragonType MonsterType
var DragonMonster OriginalMonster
var TrollType MonsterType
var TrollMonster OriginalMonster
var AlienType MonsterType
var AlienMonster OriginalMonster


func createMonster(parents []Monster) BredMonster {
	return BredMonster{
		Name: nameMonster(parents),
	}
}

type Monster interface {
	getName() string
}

type Rarity string

const (
	Legendary Rarity = "Legendary"
	Epic      Rarity = "Epic"
	Rare      Rarity = "Rare"
	Uncommon  Rarity = "Uncommon"
	Common    Rarity = "Common"
)

type OriginalMonster struct {
	Name       string
	Health     int64
	Rarity     Rarity
	Generation int64
}

func (o OriginalMonster) getName() string {
	return o.Name
}

func initOriginalMonsters() {
	WormMonster = OriginalMonster{"Worm", 20, Common, 0}
	DragonMonster = OriginalMonster{"Dragon", 100, Rare, 0}
	TrollMonster = OriginalMonster{"Troll", 30, Uncommon, 0}
	AlienMonster = OriginalMonster{"Alien", 40, Uncommon, 0}
	
}

func initMonsterTypes() {
	WormType = MonsterType{"Worm", WormMonster, 100}
	DragonType = MonsterType{"Dragon", DragonMonster, 100}
	TrollType = MonsterType{"Troll", TrollMonster, 100}
	AlienType = MonsterType{"Alien", AlienMonster, 100}
}

type MonsterType struct {
	Name                string
	FromOriginalMonster OriginalMonster
	Percentage          int8
}

func (m MonsterType) printPercentage() {
	fmt.Printf("%s, making up %d%s of the monster\n", m.Name, m.Percentage, "%")
}

type BredMonster struct {
	Name       string
	Health     int64
	Types      []MonsterType
	Rarity     Rarity
	Parents    [2]Monster
	Generation int64
}

func (b BredMonster) getName() string {
	return b.Name
}

func (b BredMonster) determineRarity() Rarity {
	// something to do with generation, rarities of parents, and damage

	return Rare
}

func nameMonster(parents []Monster) string {
	name := ""
	name = name + parents[0].getName()[0:int(len(parents[0].getName())/2)]
	name = name + parents[1].getName()[int(len(parents[1].getName())/2):len(parents[1].getName())]

	return name
}
