/*
Instructions
In this exercise you'll be organizing races between various types of remote controlled cars. Each car has its own speed and battery drain characteristics.
Cars start with full (100%) batteries. Each time you drive the car using the remote control, 
it covers the car's speed in meters and decreases the remaining battery percentage by its battery drain.
If a car's battery is below its battery drain percentage, you can't drive the car anymore.
Each race track has its own distance. Cars are tested by checking if they can finish the track without running out of battery.

1. Create a remote controlled car
Define a Car struct with the following int type fields:
1.battery
2.batteryDrain
3.speed
4.distance

2. Create a race track
Define another struct type called Track with the field distance of type integer. 
Allow creating a race track by defining a function NewTrack that takes the track's distance in meters as its sole parameter 

3. Drive the car
Implement the Drive function that updates the number of meters driven based on the car's speed, 
and reduces the battery according to the battery drainage. 
If there is not enough battery to drive one more time the car will not move.

4. Check if a remote controlled car can finish a race
To finish a race, a car has to be able to drive the race's distance. This means not draining its battery before having crossed the finish line.
Implement the CanFinish function that takes a Car and a Track instance as its parameter and returns true if the car can finish the race; otherwise, return false.

Assume that you are currently at the starting line of the race and start the engine of the car for the race. Take into account that the car's battery 
might not necessarily be fully charged when starting the race
*/

package speed

// TODO: define the 'Car' type struct
type Car struct{
    battery int
    batteryDrain int
    speed int
    distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
        battery: 100,
        batteryDrain: batteryDrain,
        speed: speed,
    }
}

// TODO: define the 'Track' type struct
type Track struct{
    distance int
} 

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{
        distance: distance,
    }
}

// Drive drives the car one time. If there is not enough battery to drive one more time, the car will not move.
func Drive(car Car) Car {
	if newbattery:=car.battery-car.batteryDrain;newbattery>=0{
        car.battery=newbattery
        car.distance+=car.speed
    }
	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	return car.battery/car.batteryDrain*car.speed>=track.distance
}
