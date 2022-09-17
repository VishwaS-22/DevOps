/*
Instructions
In this exercise you'll be working with savings accounts. Each year, the balance of your savings account is updated based on its interest rate. 
The interest rate your bank gives you depends on the amount of money in your account (its balance):

->3.213% for a balance less than 0 dollars (balance gets more negative).
->0.5% for a balance greater than or equal to 0 dollars, and less than 1000 dollars.
->1.621% for a balance greater than or equal to 1000 dollars, and less than 5000 dollars.
->2.475% for a balance greater than or equal to 5000 dollars.
You have four tasks, each of which will deal your balance and its interest rate.
package interest

1.Calculate the interest rate.
2.Calculate the interest.
3.Calculate the annual balance update.
	Implement the AnnualBalanceUpdate() function to calculate the annual balance update, taking into account the interest rate
4.Calculate the years before reaching the desired balance.
	Implement the YearsBeforeDesiredBalance() function to calculate the minimum number of years required to reach the desired balance,
taking into account that each year, interest is added to the balance. This means that the balance after one year is: start balance + interest for start balance. 
The balance after the second year is: balance after one year + interest for balance after one year. And so on, until the current year's balance is greater than 
or equal to the target balance.
*/

package interest;

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	switch {
        case balance<0:
    	return 3.213;
        case balance<1000:
    	return 0.5;
        case balance<5000:
    	return 1.621;
        default:
    	return 2.475;
    }
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
    balance*=float64(InterestRate(balance))
    return balance/100;
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance+Interest(balance);
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int 
{
        ctc:=0;
        for balance<targetBalance{
        balance=AnnualBalanceUpdate(balance);
        ctc++;
        }
        return ctc;
}
