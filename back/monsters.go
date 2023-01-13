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
		Name:       nameMonster(parents),
		Health:     0,      // PLACEHOLDER
		Rarity:     Common, //PLACEHOLDER
		Generation: getGeneration(parents),
		Parents:    parents,
		Types:      getTypes(parents),
	}
}

type Monster interface {
	getName() string
	getGeneration() int64
	getTypes() []MonsterType
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
	Type       MonsterType
}

func (o OriginalMonster) getName() string {
	return o.Name
}

func (o OriginalMonster) getGeneration() int64 {
	return o.Generation
}

func (o OriginalMonster) getTypes() []MonsterType {
	return []MonsterType{o.Type}
}

func initOriginalMonsters() {
	WormMonster = OriginalMonster{"Worm", 20, Common, 0, WormType}
	DragonMonster = OriginalMonster{"Dragon", 100, Rare, 0, DragonType}
	TrollMonster = OriginalMonster{"Troll", 30, Uncommon, 0, TrollType}
	AlienMonster = OriginalMonster{"Alien", 40, Uncommon, 0, AlienType}

}

func initMonsterTypes() {
	WormType = MonsterType{"Worm", 100}
	DragonType = MonsterType{"Dragon", 100}
	TrollType = MonsterType{"Troll", 100}
	AlienType = MonsterType{"Alien", 100}
}

type MonsterType struct {
	Name       string
	Percentage int8
}

func (m MonsterType) printPercentage() {
	fmt.Printf("%s, making up %d%s of the monster\n", m.Name, m.Percentage, "%")
}

type BredMonster struct {
	Name       string
	Health     int64
	Types      []MonsterType
	Rarity     Rarity
	Parents    []Monster
	Generation int64
}

func (b BredMonster) getName() string {
	return b.Name
}
func (b BredMonster) getGeneration() int64 {
	return b.Generation
}

func (b BredMonster) getTypes() []MonsterType {
	var types []MonsterType
	for _, p := range b.Parents {
		types = append(types, p.getTypes()...)
	}
	return types
}

func (b BredMonster) determineRarity() Rarity {
	// something to do with generation, rarities of parents, and damage

	return Rare
}

func nameMonster(parents []Monster) string {
	name := ""
	fmt.Println(parents[0].getName())
	fmt.Println(parents)
	name = name + parents[0].getName()[0:int(len(parents[0].getName())/2)]
	name = name + parents[1].getName()[int(len(parents[1].getName())/2):len(parents[1].getName())]

	return name
}

func getGeneration(parents []Monster) int64 {
	if parents[0].getGeneration() > parents[1].getGeneration() {
		return parents[0].getGeneration()
	} else {
		return parents[1].getGeneration()
	}
}

func getTypes(parents []Monster) []MonsterType {
	var types []MonsterType
	for _, p := range parents {
		types = append(types, p.getTypes()...)
	}
	return types
}
