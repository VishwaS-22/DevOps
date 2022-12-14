//Learn Go Functions: Interstellar Travel
//You’ll be writing Go functions to perform calculations and build out an interstellar travel agency!

package main
import "fmt"

// Create the function fuelGauge() here
func fuelGauge(fuel int){
  fmt.Println("You have",fuel,"left")
}

// Create the function calculateFuel() here
func calculateFuel(planet string) int{
  var fuel int
  switch planet {
  case "Venus":
  fuel=300000
  case "Mercury":
  fuel=500000
  case "Mars":
  fuel=700000
  default :
  fuel=0
  }
  return fuel
}

// Create the function greetPlanet() here
func greetPlanet(planet string){
  fmt.Println("Welcome to the planet",planet)
}

// Create the function cantFly() here
func cantFly(){
  fmt.Println("We do not have the available fuel to fly there.")
}

// Create the function flyToPlanet() here
func flyToPlanet(planet string,fuel int) int{
var fuelRemaining ,fuelCost int
fuelRemaining=fuel
return fuelRemaining
fuelCost=calculateFuel(planet)
if fuelRemaining >=fuelCost{
    greetPlanet(planet)
    fuelRemaining-=fuelCost
  }else{
    cantFly()
  }
  return fuelRemaining
  }
func main(){
  // Test your functions!
  var fuel int
  fuel=1000000
  // Create `planetChoice` and `fuel`
  planetChoice:="Venus"
  fuel=flyToPlanet(planetChoice,fuel)
  fuelGauge(fuel)
  // And then liftoff!
}
