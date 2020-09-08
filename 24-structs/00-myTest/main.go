package main

import "fmt"

type car struct {
	model string
	year  int
	color string
}

type vehicle struct {
	model string
	year  int
	color string
}

func main() {
	tesla := car{
		model: "Tesla",
		year:  2020,
		color: "silver",
	}

	honda := car{
		"Honda",
		2007,
		"metallic",
	}

	motor := vehicle{
		model: "mootor",
		color: "blue",
	}

	ferrari := car(motor)

	fmt.Printf("Car1: %+v\n", tesla)
	fmt.Printf("Car2: %+v\n", honda)
	fmt.Printf("Veh: %+v\n", motor)
	fmt.Printf("Ferr: %+v\n", ferrari)
	fmt.Printf("Ferr: %T\n", ferrari)

	{
		fmt.Println()
		avengers := struct {
			title, genre string
			rating       float64
			released     bool
		}{
			"avengers: end game",
			"sci-fi",
			8.9,
			true,
		}

		fmt.Printf("%T\n", avengers)
	}
	{
		type Candidate struct {
			Name string   `name`
			Prio []string `priorities`
		}
		type VocalCandidate struct {
			Name string   `name`
			Prio []string `stated_priorities`
		}

		x := Candidate{}
		y := VocalCandidate(x)

		fmt.Printf("x: %+v\ny: %+v\n", x, y)
	}
}
