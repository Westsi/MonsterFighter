package main

import "fmt"

func monsterTester() {
	bmo := createMonster([]Monster{WormMonster, TrollMonster})
	fmt.Println(bmo.Name)
	fmt.Println(bmo.Rarity)
	fmt.Println(bmo.Health)
	fmt.Println(bmo.Generation)
	fmt.Println(bmo.Types)
	fmt.Println(bmo.Parents)

	bmt := createMonster([]Monster{DragonMonster, TrollMonster})
	fmt.Println(bmt.Name)
	fmt.Println(bmt.Rarity)
	fmt.Println(bmt.Health)
	fmt.Println(bmt.Generation)
	fmt.Println(bmt.Types)
	fmt.Println(bmt.Parents)

	bm := createMonster([]Monster{bmo, bmt})
	// fmt.Println(bm.Name)
	// fmt.Println(bm.Rarity)
	// fmt.Println(bm.Health)
	// fmt.Println(bm.Generation)
	fmt.Print("Types of final monster:")
	fmt.Println(bm.Types)

	bmtt := createMonster([]Monster{bm, bmt})
	// fmt.Println(bm.Name)
	// fmt.Println(bm.Rarity)
	// fmt.Println(bm.Health)
	fmt.Println(bmtt.Generation)
	fmt.Print("Types of final monster:")
	fmt.Println(bmtt.Types)
	// fmt.Println(bm.Parents)
}