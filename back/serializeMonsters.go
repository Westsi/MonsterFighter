package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// path format
// monsters/_id_.json

type FileOriginalMonster struct {
	Name       string
	Health     int64
	Rarity     string
	Generation int64
	Type       string
	ID         string
	Strength   int64
	Speed      int64
	Stamina    int64
}

type FileBredMonster struct {
	Name       string
	Health     int64
	Types      []string
	Rarity     string
	Parents    []string
	Generation int64
	ID         string
	Strength   int64
	Speed      int64
	Stamina    int64
	Owner      string
}

const BASEPATH string = "monsters/"

func (b BredMonster) writeToFile(path string) {
	p := path

	if path[len(path)-5:len(path)-1] != ".json" {
		p = p + ".json"
	}

	var types []string
	for _, ty := range b.Types {
		types = append(types, ty.getString())
	}
	t := FileBredMonster{
		Name:       b.Name,
		Health:     b.Health,
		Types:      types,
		Rarity:     string(b.Rarity),
		Parents:    b.Parents,
		Generation: b.Generation,
		ID:         b.ID,
		Strength:   b.Strength,
		Speed:      b.Speed,
		Stamina:    b.Stamina,
		Owner:      b.Owner,
	}
	monster, _ := json.Marshal(t)
	// fmt.Println(string(monster))

	ioutil.WriteFile(BASEPATH+p, monster, 0666)

}

func (o OriginalMonster) writeToFile(path string) {

	p := path

	if path[len(path)-5:len(path)-1] != ".json" {
		p = p + ".json"
	}

	t := FileOriginalMonster{
		Name:       o.Name,
		Health:     o.Health,
		Rarity:     string(o.Rarity),
		Generation: o.Generation,
		Type:       o.Type.getString(),
		ID:         o.ID,
		Strength:   o.Strength,
		Speed:      o.Speed,
		Stamina:    o.Stamina,
	}

	monster, _ := json.Marshal(t)
	// fmt.Println(string(monster))

	ioutil.WriteFile(BASEPATH+p, monster, 0666)

}

func readBredMonsterFromFile(path string) BredMonster {

	p := path

	fmt.Println(path[len(path)-5:])

	if path[len(path)-5:] != ".json" {
		p = p + ".json"
	}
	dat, err := os.ReadFile(BASEPATH + p)
	if err != nil {
		fmt.Println(err)
	}
	var f FileBredMonster
	json.Unmarshal(dat, &f)

	var types []MonsterType

	for _, st := range f.Types {
		types = append(types, deserializeMonsterType(st))
	}

	b := BredMonster{
		Name:       f.Name,
		Health:     f.Health,
		Types:      types,
		Rarity:     Rarity(f.Rarity),
		Parents:    f.Parents,
		Generation: f.Generation,
		ID:         f.ID,
		Strength:   f.Strength,
		Speed:      f.Speed,
		Stamina:    f.Stamina,
		Owner:      f.Owner,
	}

	return b
}

func readOriginalMonsterFromFile(path string) OriginalMonster {

	p := path

	if path[len(path)-5:len(path)-1] != ".json" {
		p = p + ".json"
	}
	dat, err := os.ReadFile(BASEPATH + p)
	if err != nil {
		fmt.Println(err)
	}
	var f FileOriginalMonster
	json.Unmarshal(dat, &f)

	o := OriginalMonster{
		Name:       f.Name,
		Health:     f.Health,
		Type:       deserializeMonsterType(f.Type),
		Rarity:     Rarity(f.Rarity),
		Generation: f.Generation,
		ID:         f.ID,
		Strength:   f.Strength,
		Speed:      f.Speed,
		Stamina:    f.Stamina,
	}

	return o
}

func (mt MonsterType) getString() string {
	//name percentage
	s := ""
	s = s + mt.Name + "~" + fmt.Sprintf("%f", mt.Percentage) + "~" + mt.getDominantSyllable()
	return s
}

func deserializeMonsterType(s string) MonsterType {
	tempstrings := strings.Split(s, "~")
	name := tempstrings[0]
	percentage, err := strconv.ParseFloat(tempstrings[1], 64)
	if err != nil {
		fmt.Println(err)
	}

	return MonsterType{Name: name, Percentage: percentage, DominantSyllable: tempstrings[2]}
}
