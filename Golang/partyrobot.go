//String Formatting & Packages
/*
Instructions
Once there was an eccentric programmer living in a strange house with barred windows. One day he accepted a job from an online job board to build a party robot. 
The robot is supposed to greet people and help them to their seats. The first edition was very technical and showed the programmer's lack of human interaction. 
Some of which also made it into the next edition.
1. Welcome a new guest to the party
2. Welcome a new guest to the party whose birthday is today
3. Give directions
Implement the AssignTable function to give directions. It should accept 5 parameters.

->The name of the new guest to greet (string)
->The table number (int)
->The name of the seatmate (string)
->The direction where to find the table (string)
->The distance to the table (float64)
The exact result format can be seen in the example below.

The robot provides the table number in a 3 digits format. If the number is less than 3 digits it gets extra leading zeroes to become 3 digits 
(e.g. 3 becomes 003).The robot also mentions the distance of the table. Fortunately only with a precision that's limited to 1 digit.
*/

package partyrobot
import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
    a:="Welcome to my party, %s!"
    return fmt.Sprintf(a,name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
    b:="Happy birthday %s! You are now %d years old!"
    return fmt.Sprintf(b,name,age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	    return fmt.Sprintf("Welcome to my party, %s!\nYou have been assigned to table %03d. Your table is %v, exactly %.1f meters from here.\nYou will be sitting next to %s.", name, table, direction, distance, neighbor)
}
