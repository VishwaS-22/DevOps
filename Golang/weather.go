/*
Instructions
Goblinocus is a country that takes its weather forecast very seriously. Since you are a renowned responsible and proficient developer, 
they asked you to write a program that can forecast the current weather condition of various cities in Goblinocus.
You were busy at the time and asked one of your friends to do the job instead. After a while,
the president of Goblinocus contacted you and said they do not understand your friend's code. 
When you check the code, you discover that your friend did not act as a responsible programmer and there are no comments in the code. 
You feel obligated to clarify the program so goblins can understand them as well.
1. Document package weather
2. Document the CurrentCondition and CurrentLocation variables
3. Document the Forecast() function
*/

//Package weather forcasts about the weather condition.
package weather

//CurrentCondition will said about current weather.
var CurrentCondition string

//CurrentLocation will said about current location.
var CurrentLocation string

//Forecast will said about current weather at current at current location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
