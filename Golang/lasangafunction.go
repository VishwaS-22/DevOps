/*
Instructions
In this exercise you are going to write some more code related to preparing and cooking your brilliant lasagna from your favorite cookbook.
You have four tasks. The first one is related to the cooking itself, the other three are about the perfect preparation
package lasagna

1. Estimate the preparation time
Implement a function PreparationTime that accepts a slice of layers as a []string and the average preparation time per layer in minutes as an int. 
The function should return the estimate for the total preparation time based on the number of layers as an int. Go has no default values for functions. 
If the average preparation time is passed as 0 (the default initial value for an int), then the default value of 2 should be used.

2. Compute the amounts of noodles and sauce needed
For each noodle layer in your lasagna, you will need 50 grams of noodles. For each sauce layer in your lasagna, you will need 0.2 liters of sauce.
Define the function Quantities that takes a slice of layers as parameter as a []string. 
The function will then determine the quantity of noodles and sauce needed to make your meal. 
The result should be returned as two values of noodles as an int and sauce as a float64.	

3. Add the secret ingredient
AddSecretIngredient does not return anything - you should modify the list of your ingredients directly. 
The list with your friend's ingredients should not be modified. Also, since slice is passed into a function as pointers, 
changes to the two []string arguments passed into AddSecretIngredient will be modified directly.

4. Scale the recipe
Implement a function ScaleRecipe that takes two parameters.
->A slice of float64 amounts needed for 2 portions.
->The number of portions you want to cook.
The function should return a slice of float64 of the amounts needed for the desired number of portions. You want to keep the original recipe though. 
This means the quantities argument should not be modified in this function.
*/

package lasagna

// define the 'PreparationTime()' function
func PreparationTime(s []string,l int)int{
    if l==0{
         return len(s)*2;
    }
    return len(s)*l;
}

// define the 'Quantities()' function
func Quantities(s []string)(int,float64){
    var n int;
    var se float64;
    for i:=0;i<len(s);i++{
        if(s[i]=="noodles"){
            n++;
        }
		if(s[i]=="sauce"){
            se++;
        }
    }
	return n*50 , se*0.2;
}

// define the 'AddSecretIngredient()' function
func AddSecretIngredient(f []string,o []string){
   o[len(o)-1] = f[len(f)-1]
}

//define the 'ScaleRecipe()' function
func ScaleRecipe(s []float64,n int)[]float64{
	v:=make([]float64,len(s))
    for i:=0;i<len(s);i++ {
        v[i]=s[i]*float64(n)/2;
    }
	return v;
}
