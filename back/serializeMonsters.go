package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// path format
// monsters/_timeHashed_.json

func (b BredMonster) writeToFile(path string) {

	monster, _ := json.Marshal(b)
	fmt.Println(string(monster))

	ioutil.WriteFile(path, monster, 0666)

}

func (o OriginalMonster) writeToFile(path string) {

	monster, _ := json.Marshal(o)
	fmt.Println(string(monster))

	ioutil.WriteFile(path, monster, 0666)

}

func readBredMonsterFromFile(path string) BredMonster {
	dat, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	var b BredMonster
	json.Unmarshal(dat, &b)
	fmt.Println(b)

	return b
}

func readOriginalMonsterFromFile(path string) OriginalMonster {
	dat, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	var o OriginalMonster
	json.Unmarshal(dat, &o)

	return o
}
