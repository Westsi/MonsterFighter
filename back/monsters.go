package main

import (
	"crypto/sha256"
	"fmt"
	"math"
	"math/rand"
	"time"
)

var rng *rand.Rand

func initVars() {
	initMonsterTypes()
	initOriginalMonsters()
	rng = rand.New(rand.NewSource(time.Now().UnixMilli()))

}

var WormType MonsterType
var WormMonster OriginalMonster
var DragonType MonsterType
var DragonMonster OriginalMonster
var TrollType MonsterType
var TrollMonster OriginalMonster
var AlienType MonsterType
var AlienMonster OriginalMonster
var UnicornType MonsterType
var UnicornMonster OriginalMonster
var PhoenixType MonsterType
var PhoenixMonster OriginalMonster
var WolfType MonsterType
var WolfMonster OriginalMonster
var BearType MonsterType
var BearMonster OriginalMonster
var GorgonType MonsterType
var GorgonMonster OriginalMonster
var RabbitType MonsterType
var RabbitMonster OriginalMonster

func createMonster(parentss []string) BredMonster {
	var parents []Monster
	for _, id := range parentss {
		parents = append(parents, getParentObjectFromID(id))
	}
	b := BredMonster{
		Name:       "",
		Health:     determineHealth(parents),
		Rarity:     determineRarity(parents),
		Generation: getGeneration(parents),
		Parents:    parentss,
		Types:      getTypes(parents),
		ID:         parentss[0],
		Strength:   determineStrength(parents),
		Speed:      determineSpeed(parents),
		Stamina:    determineStamina(parents),
	}
	b = b.geneticallyRandomize()
	b.writeToFile(b.ID)

	return b
}

type Monster interface {
	getName() string
	getGeneration() int64
	getTypes() []MonsterType
	getHealth() int64
	getSpeed() int64
	getStrength() int64
	getStamina() int64
	getID() string
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
	ID         string
	Strength   int64
	Speed      int64
	Stamina    int64
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

func (o OriginalMonster) getSpeed() int64 {
	return o.Speed
}

func (o OriginalMonster) getStamina() int64 {
	return o.Stamina
}

func (o OriginalMonster) getStrength() int64 {
	return o.Strength
}

func (o OriginalMonster) getID() string {
	return o.ID
}

func initOriginalMonsters() {
	// need to update strength, speed, stamina for these (last 3 vals)
	WormMonster = OriginalMonster{"Worm", 200, Common, 0, WormType, genNewID(true, "Worm", ""), 100, 100, 100}
	DragonMonster = OriginalMonster{"Dragon", 1000, Epic, 0, DragonType, genNewID(true, "Dragon", ""), 100, 100, 100}
	TrollMonster = OriginalMonster{"Troll", 300, Common, 0, TrollType, genNewID(true, "Troll", ""), 100, 100, 100}
	AlienMonster = OriginalMonster{"Alien", 400, Uncommon, 0, AlienType, genNewID(true, "Alien", ""), 100, 100, 100}
	UnicornMonster = OriginalMonster{"Unicorn", 600, Rare, 0, UnicornType, genNewID(true, "Unicorn", ""), 100, 100, 100}
	PhoenixMonster = OriginalMonster{"Phoenix", 700, Rare, 0, PhoenixType, genNewID(true, "Phoenix", ""), 100, 100, 100}
	WolfMonster = OriginalMonster{"Wolf", 500, Uncommon, 0, WolfType, genNewID(true, "Wolf", ""), 100, 100, 100}
	BearMonster = OriginalMonster{"Bear", 400, Uncommon, 0, BearType, genNewID(true, "Bear", ""), 100, 100, 100}
	GorgonMonster = OriginalMonster{"Gorgon", 900, Epic, 0, GorgonType, genNewID(true, "Gorgon", ""), 100, 100, 100}
	RabbitMonster = OriginalMonster{"Rabbit", 1500, Legendary, 0, RabbitType, genNewID(true, "Rabbit", ""), 100, 100, 100}
	//
	WormMonster.writeToFile(WormMonster.getID())
	DragonMonster.writeToFile(DragonMonster.getID())
	TrollMonster.writeToFile(TrollMonster.getID())
	AlienMonster.writeToFile(AlienMonster.getID())
	UnicornMonster.writeToFile(UnicornMonster.getID())
	PhoenixMonster.writeToFile(PhoenixMonster.getID())
	WolfMonster.writeToFile(WolfMonster.getID())
	BearMonster.writeToFile(BearMonster.getID())
	GorgonMonster.writeToFile(GorgonMonster.getID())
	RabbitMonster.writeToFile(RabbitMonster.getID())
	fmt.Println("initialized monsters")
}

func initMonsterTypes() {
	WormType = MonsterType{"Worm", 100, "wrm"}
	DragonType = MonsterType{"Dragon", 100, "drag"}
	TrollType = MonsterType{"Troll", 100, "tro"}
	AlienType = MonsterType{"Alien", 100, "lien"}
	UnicornType = MonsterType{"Unicorn", 100, "corn"}
	PhoenixType = MonsterType{"Phoenix", 100, "nix"}
	WolfType = MonsterType{"Wolf", 100, "wol"}
	BearType = MonsterType{"Bear", 100, "bea"}
	GorgonType = MonsterType{"Gorgon", 100, "gon"}
	RabbitType = MonsterType{"Rabbit", 100, "bit"}

	fmt.Println("initialized monster types")
}

func genNewID(isOriginalMonster bool, name string, parentOneID string) string {
	if isOriginalMonster {
		sum := sha256.Sum256([]byte(name + fmt.Sprint(time.Now().UnixMicro())))
		return fmt.Sprintf("o%x", sum)
	} else {
		sum := sha256.Sum256([]byte(name + parentOneID + fmt.Sprint(time.Now().UnixMicro())))
		return fmt.Sprintf("b%x", sum)
	}

}

type MonsterType struct {
	Name             string
	Percentage       float64
	DominantSyllable string
}

func (m MonsterType) percentageMakeup() string {
	return fmt.Sprintf("%s, making up %f%s of the monster\n", m.Name, m.Percentage, "%")
}

func (m MonsterType) getDominantSyllable() string {
	return m.DominantSyllable
}

type BredMonster struct {
	Name       string
	Health     int64
	Types      []MonsterType
	Rarity     Rarity
	Parents    []string
	Generation int64
	ID         string
	Strength   int64
	Speed      int64
	Stamina    int64
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

func (b BredMonster) getSpeed() int64 {
	return b.Speed
}

func (b BredMonster) getStamina() int64 {
	return b.Stamina
}

func (b BredMonster) getStrength() int64 {
	return b.Strength
}

func (b BredMonster) geneticallyRandomize() BredMonster {

	// 10% chance for one of the stats to be boosted by 1-50 points, 7.5% chance for random base type to be added

	n := rng.Intn(100)

	if n <= 10 {
		fmt.Println("INCREASING STAT")
		increase := rng.Intn(49) + 1
		statChoice := rng.Intn(4)
		if statChoice == 0 {
			b.Health += int64(increase)
		} else if statChoice == 1 {
			b.Strength += int64(increase)
		} else if statChoice == 2 {
			b.Speed += int64(increase)
		} else if statChoice == 3 {
			b.Stamina += int64(increase)
		}
	}

	n = rng.Intn(1000)
	if n <= 80 {
		fmt.Println("ADDING RANDOM TYPE")
		if n <= 40 {
			// COMMON
			if n <= 20 {
				b.Types = workOutTypesPercentages(append(b.Types, WormType))
			} else {
				b.Types = workOutTypesPercentages(append(b.Types, TrollType))
			}
		} else if n <= 64 {
			// UNCOMMON
			if n <= 48 {
				b.Types = workOutTypesPercentages(append(b.Types, AlienType))
			} else if n <= 56 {
				b.Types = workOutTypesPercentages(append(b.Types, WolfType))
			} else {
				b.Types = workOutTypesPercentages(append(b.Types, BearType))
			}
		} else if n <= 72 {
			// RARE
			if n <= 68 {
				b.Types = workOutTypesPercentages(append(b.Types, UnicornType))
			} else {
				b.Types = workOutTypesPercentages(append(b.Types, PhoenixType))
			}
		} else if n <= 78 {
			// EPIC
			if n <= 75 {
				b.Types = workOutTypesPercentages(append(b.Types, DragonType))
			} else {
				b.Types = workOutTypesPercentages(append(b.Types, GorgonType))
			}
		} else {
			// LEGENDARY
			b.Types = workOutTypesPercentages(append(b.Types, RabbitType))
		}
	}

	b.Name = nameMonster(b.Types)
	b.ID = genNewID(false, b.Name, b.ID)
	return b
}

func getParentObjectFromID(id string) Monster {
	if string(id[0]) == "o" {
		return readOriginalMonsterFromFile(id)
	}
	return readBredMonsterFromFile(id)
}

func (b BredMonster) getTypes() []MonsterType {
	var types []MonsterType
	for _, p := range b.Parents {
		types = append(types, getParentObjectFromID(p).getTypes()...)
	}
	return types
}

// NEED TO REWRITE determineX functions and nameMonster

func determineRarity(parents []Monster) Rarity {
	// something to do with generation, rarities of parents, and damage

	// later generation - more powerful. should not directly affect, but will affect damage and health
	return Rare
}

func nameMonster(types []MonsterType) string {
	// name by comparing dominant syllables of types?
	name := ""

	for _, t := range types {
		name = name + t.getDominantSyllable()
	}

	// name = name + parents[0].getName()[0:int(len(parents[0].getName())/2)]
	// name = name + parents[1].getName()[int(len(parents[1].getName())/2):len(parents[1].getName())]
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
	typecounts := make(map[string]int)
	domsyls := make(map[string]string)
	for _, t := range types {
		typecounts[t.Name] += 1
		domsyls[t.Name] = t.getDominantSyllable()
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
			Name:             t,
			Percentage:       math.Round((float64(val)/float64(sum))*10000) / 100,
			DominantSyllable: domsyls[t],
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

func determineSpeed(parents []Monster) int64 {
	sum := 0
	for _, p := range parents {
		sum += int(p.getSpeed())
	}

	return int64(sum / len(parents))
}

func determineStrength(parents []Monster) int64 {
	sum := 0
	for _, p := range parents {
		sum += int(p.getStrength())
	}

	return int64(sum / len(parents))
}

func determineStamina(parents []Monster) int64 {
	sum := 0
	for _, p := range parents {
		sum += int(p.getStamina())
	}

	return int64(sum / len(parents))
}
