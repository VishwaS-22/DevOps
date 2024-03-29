//Instructions
/*
In this exercise you'll be writing code to analyze the production in a car factory.

1. Calculate the number of working cars produced per hour 
2. Calculate the number of working cars produced per minute
3. Calculate the cost of production
Each car normally costs $10,000 to produce individually, regardless of whether it is successful or not. 
But with a bit of planning,10 cars can be produced together for $95,000.
For example, 37 cars can be produced in the following way: 37 = 3 x groups of ten + 7 individual cars
So the cost for 37 cars is: 3*95,000+7*10,000=355,000
*/

package cars

// CalculateWorkingCarsPerHour calculates how many working cars are produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    v:=successRate/100.0
    s:=float64(productionRate)
    return s*v
}

// CalculateWorkingCarsPerMinute calculates how many working cars are produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    f:=CalculateWorkingCarsPerHour(productionRate,successRate)
    return int(f)/60
}

const x=10000
const y=95000

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
    single:=uint(x*(carsCount%10))
    tencar:=uint(y*(carsCount/10))
    return single+tencar
    }
