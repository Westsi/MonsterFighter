package main

import (
	"crypto/sha256"
	"fmt"
	"math"
	"time"
)

func initVars() {
	initMonsterTypes()
	initOriginalMonsters()
}

var WormType MonsterType
var WormMonster OriginalMonster
var DragonType MonsterType
var DragonMonster OriginalMonster
var TrollType MonsterType
var TrollMonster OriginalMonster
var AlienType MonsterType
var AlienMonster OriginalMonster

func createMonster(parentss []string) BredMonster {
	var parents []Monster
	for _, id := range parentss {
		parents = append(parents, getParentObjectFromID(id))
	}
	b := BredMonster{
		Name:       nameMonster(parents),
		Health:     determineHealth(parents),
		Rarity:     determineRarity(parents),
		Generation: getGeneration(parents),
		Parents:    parentss,
		Types:      getTypes(parents),
		ID:         genNewID(false, nameMonster(parents), parentss[0]),
	}
	b.writeToFile(b.ID)

	return b
}

type Monster interface {
	getName() string
	getGeneration() int64
	getTypes() []MonsterType
	getHealth() int64
	getID() string
	// ADD OWNERS
	// PARENTS SHOULD BE STORED AS IDS NOT OBJECTS
}

type Rarity string

const (
	Legendary Rarity = "Legendary"
	Epic      Rarity = "Epic"
	Rare      Rarity = "Rare"
	Uncommon  Rarity = "Uncommon"
	Common    Rarity = "Common"
)

// func (r Rarity)toString() string {
// 	select {
// 		case r == "Legendary"
// 	}
// }

type OriginalMonster struct {
	Name       string
	Health     int64
	Rarity     Rarity
	Generation int64
	Type       MonsterType
	ID         string
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

func (o OriginalMonster) getHealth() int64 {
	return o.Health
}

func (o OriginalMonster) getID() string {
	return o.ID
}

func initOriginalMonsters() {
	WormMonster = OriginalMonster{"Worm", 20, Common, 0, WormType, genNewID(true, "Worm", "")}
	DragonMonster = OriginalMonster{"Dragon", 100, Rare, 0, DragonType, genNewID(true, "Dragon", "")}
	TrollMonster = OriginalMonster{"Troll", 30, Uncommon, 0, TrollType, genNewID(true, "Troll", "")}
	AlienMonster = OriginalMonster{"Alien", 40, Uncommon, 0, AlienType, genNewID(true, "Alien", "")}

	WormMonster.writeToFile(WormMonster.getID())
	DragonMonster.writeToFile(DragonMonster.getID())
	TrollMonster.writeToFile(TrollMonster.getID())
	AlienMonster.writeToFile(AlienMonster.getID())

}

func initMonsterTypes() {
	WormType = MonsterType{"Worm", 100}
	DragonType = MonsterType{"Dragon", 100}
	TrollType = MonsterType{"Troll", 100}
	AlienType = MonsterType{"Alien", 100}
}

func genNewID(isOriginalMonster bool, name string, parentOneID string) string {
	if isOriginalMonster {
		sum := sha256.Sum256([]byte(name + fmt.Sprint(time.Now().UnixMicro())))
		return fmt.Sprintf("%x", sum)
	} else {
		sum := sha256.Sum256([]byte(name + parentOneID + fmt.Sprint(time.Now().UnixMicro())))
		return fmt.Sprintf("%x", sum)
	}

}

type MonsterType struct {
	Name       string
	Percentage float64
}

func (m MonsterType) printPercentage() {
	fmt.Printf("%s, making up %d%s of the monster\n", m.Name, m.Percentage, "%")
}

type BredMonster struct {
	Name       string
	Health     int64
	Types      []MonsterType
	Rarity     Rarity
	Parents    []string
	Generation int64
	ID         string
}

func (b BredMonster) getName() string {
	return b.Name
}
func (b BredMonster) getGeneration() int64 {
	return b.Generation
}

func (b BredMonster) getHealth() int64 {
	return b.Health
}

func (b BredMonster) getID() string {
	return b.ID
}

func getParentObjectFromID(id string) Monster {
	o := readOriginalMonsterFromFile(id)
	b := readBredMonsterFromFile(id)
	// HOW TO CHECK WHICH ONE TO USE?!
	fmt.Println(o)
	fmt.Println(b)

	return BredMonster{}
}

func (b BredMonster) getTypes() []MonsterType {
	var types []MonsterType
	for _, p := range b.Parents {
		types = append(types, getParentObjectFromID(p).getTypes()...)
	}
	return types
}

func determineRarity(parents []Monster) Rarity {
	// something to do with generation, rarities of parents, and damage

	return Rare
}

func nameMonster(parents []Monster) string {
	name := ""
	name = name + parents[0].getName()[0:int(len(parents[0].getName())/2)]
	name = name + parents[1].getName()[int(len(parents[1].getName())/2):len(parents[1].getName())]
	return name
}

func getGeneration(parents []Monster) int64 {
	if parents[0].getGeneration() > parents[1].getGeneration() {
		return parents[0].getGeneration() + 1
	} else {
		return parents[1].getGeneration() + 1
	}
}

func getTypes(parents []Monster) []MonsterType {
	var types []MonsterType
	for _, p := range parents {
		types = append(types, p.getTypes()...)
	}
	return workOutTypesPercentages(types)
}

func workOutTypesPercentages(types []MonsterType) []MonsterType {
	typecounts := make(map[MonsterType]int)
	for _, t := range types {
		typecounts[t] += 1
	}
	var ttr []MonsterType
	sum := 0

	for _, val := range typecounts {
		// sum vals
		sum += val
	}

	for t, val := range typecounts {
		// work out percentages
		tt := MonsterType{
			Name:       t.Name,
			Percentage: math.Round((float64(val)/float64(sum))*10000) / 100,
		}
		ttr = append(ttr, tt)
	}

	return ttr
}

func determineHealth(parents []Monster) int64 {
	sum := 0
	for _, p := range parents {
		sum += int(p.getHealth())
	}

	return int64(sum / len(parents))
}
