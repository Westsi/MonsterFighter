package main

import (
	"fmt"

	"gorm.io/gorm"
)

type DBOriginalMonster struct {
	gorm.Model
	Name       string
	Health     int64
	Rarity     string
	Generation int64
	Type       string
}
type DBBredMonster struct {
	gorm.Model
	Name       string
	Health     int64
	Types      string
	Rarity     string
	Parents    string
	Generation int64
}

func (b BredMonster) prepareForDB() DBBredMonster {
	d := DBBredMonster{
		Name:       b.Name,
		Health:     b.Health,
		Types:      "TYPES GO HERE",
		Rarity:     string(b.Rarity),
		Parents:    "PARENTS GO HERE",
		Generation: b.Generation,
	}

	return d
}

func (o OriginalMonster) prepareForDB() DBOriginalMonster {
	d := DBOriginalMonster{
		Name:       o.Name,
		Health:     o.Health,
		Type:       fmt.Sprintf("%s", o.Type),
		Rarity:     string(o.Rarity),
		Generation: o.Generation,
	}

	return d
}
