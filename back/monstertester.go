package main

import "fmt"

func monsterTester() {
	testUser := User{Name: "Test", Password: "test", Email: "test", GakZunnCount: int(new_user_currency)}
	bmo := createMonster(&testUser, []string{WormMonster.getID(), TrollMonster.getID()})
	fmt.Println(bmo.Name)
	fmt.Println(bmo.Rarity)
	fmt.Println(bmo.Health)
	fmt.Println(bmo.Generation)
	fmt.Println(bmo.Types)
	fmt.Println(bmo.Parents)

	bmt := createMonster(&testUser, []string{DragonMonster.getID(), TrollMonster.getID()})
	fmt.Println(bmt.Name)
	fmt.Println(bmt.Rarity)
	fmt.Println(bmt.Health)
	fmt.Println(bmt.Generation)
	fmt.Println(bmt.Types)
	fmt.Println(bmt.Parents)

	bm := createMonster(&testUser, []string{bmo.getID(), bmt.getID()})
	fmt.Println(bm.Name)
	fmt.Println(bm.Rarity)
	fmt.Println(bm.Health)
	fmt.Println(bm.Generation)
	fmt.Print("Types of final monster:")
	fmt.Println(bm.Types)

	// bmtt := createMonster(&testUser, []string{bm.getID(), bmt.getID()})
	// fmt.Println(bmtt.Name)
	// fmt.Println(bm.Rarity)
	// fmt.Println(bm.Health)
	// fmt.Println(bmtt.Generation)
	// fmt.Print("Types of final monster:")
	// fmt.Println(bmtt.Types)
	// fmt.Println(bm.Parents)

}
