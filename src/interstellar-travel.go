package main

import "fmt"

func fuelGauge(fuel int) {
  fmt.Printf("You have %v gallons of fuel left.", fuel)
}

func calculateFuel(planet string) int {
  var fuel int

  switch planet {
  case "Venus":
   fuel = 300000
  case "Mercury":
   fuel = 500000
  case "Mars":
   fuel = 700000
  default:
   fuel = 0
  }
  return fuel
}

func greetPlanet(planet string) {
  greeting := fmt.Sprintf("Welcome to %v, alien!", planet)
	fmt.Println(greeting)
}

func cantFly() {
  fmt.Println("We do not have the available fuel to fly there.")
}

func flyToPlanet(planet string, fuel int) int {
  var fuelRemaining int
  var fuelCost int
  fuelRemaining = fuel
  fuelCost = calculateFuel(planet)
  if fuelRemaining >= fuelCost {
    greetPlanet(planet)
    fuelRemaining -= fuelCost
  } else {
    cantFly()
  }
  return fuelRemaining
}

func main() {
  
  var fuel int
  fuel = 1000000
  var planetChoice string
  planetChoice = "Venus"

  fuel = flyToPlanet(planetChoice, fuel)

  fuelGauge(fuel)
  
}