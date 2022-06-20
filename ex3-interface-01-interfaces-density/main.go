package main

import "fmt"

// create an interface with properties that is extensible across different types of matters
type Dense interface {
	Density() float64
}

// Gas and Metal have their respective properties and methods (Density)
// Gas and Metal share also IMPLEMENTS interface Dense, which all share the same method signature (Density)
// Therefore function isDenser is able to accept both Metal and Gas objects as Dense objects to process

type Metal struct {
	mass   float64
	volume float64
}

type Gas struct {
	pressure      float64
	temperature   float64
	molecularmass float64
}

func (m *Metal) Density() float64 {
	return m.mass / m.volume
}

func (g *Gas) Density() float64 {
	return float64((g.molecularmass * g.pressure) / (0.0821 * (g.temperature + 273)))
}

func isDenser(matter1, matter2 Dense) bool {
	return matter1.Density() > matter2.Density()
}

func main() {

	Gold := Metal{mass: 478, volume: 24}
	Silver := Metal{mass: 100, volume: 10}

	Oxygen := Gas{pressure: 5, temperature: 27, molecularmass: 32}
	Hydrogen := Gas{pressure: 1, temperature: 0, molecularmass: 2}

	fmt.Println("Gold is denser than silver:", isDenser(&Gold, &Silver))
	fmt.Println("Oxygen is denser than hydrogen:", isDenser(&Oxygen, &Hydrogen))

}
