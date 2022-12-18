//Strings & Strings package.

/*
Instructions
There is an appliance store called "Tech Palace" nearby.The owner of the store recently installed a big display to use for marketing messages and
to show a special greeting when customers scan their loyalty cards at the entrance.The display consists of lots of small LED lights and can show multiple lines of text.
The store owner needs your help with the code that is used to generate the text for the new display.
package techpalace
1. Create the welcome message
2. Add a fancy border
3. Clean up old marketing messages
*/

package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.

func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.

func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	a:= strings.Repeat("*", numStarsPerLine)
	return a+ "\n" + welcomeMsg + "\n" +a
}

// CleanupMessage cleans up an old marketing message.

func CleanupMessage(oldMsg string) string {
	return strings.TrimSpace(strings.ReplaceAll(oldMsg, "*", ""))
}
