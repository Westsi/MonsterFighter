package main

// currency is called GakZunn

// initial $ -> GZ ratio
// $1 -> 100,000 GZ
// $0.00001 -> 1 GZ

// amount of GZ in circulation
var gz_amount int64


func initCurrency() {
	// set starting amount of GZ to 100 billion, or 1 million dollars
	gz_amount = 100000000000
}

func getMonsterWorth(m Monster){
	worth := 0
	switch m.getRarity() {
		case Common:
			worth += 1000
		case Uncommon:
			worth += 5000
		case Rare:
			worth += 10000
		case Epic:
			worth += 50000
		case Legendary:
			worth += 100000
	}

	worth += int(m.getGeneration()) * 10
}