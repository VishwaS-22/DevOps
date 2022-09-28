/*
Instructions
A friend of yours has an old wholesale store called Gross Store. The name comes from the quantity of the item that the store sell: it's all in gross unit. 
Your friend asked you to implement a point of sale (POS) system for his store. First, you want to build a prototype for it. 
In your prototype, your system will only record the quantity. Your friend gave you a list of measurements to help you
1. Store the unit of measurement in your program
	In order to use the measurement, you need to store the measurement in your program.
2. Create a new customer bill	
	You need to implement a function that create a new (empty) bill for the customer.
3. Add an item to the customer bill
	->Return false if the given unit is not in the units map.
	->Otherwise add the item to the customer bill, indexed by the item name, then return true.
	->f the item is already present in the bill, increase its quantity by the amount that belongs to the provided unit.
4. Remove an item from the customer bill
	->Return false if the given item is not in the bill
	->Return false if the given unit is not in the units map.
	->Return false if the new quantity would be less than 0.
	->If the new quantity is 0, completely remove the item from the bill then return true.
	->Otherwise, reduce the quantity of the item and return true.
*/

package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	u := map[string]int {
        "quarter_of_a_dozen": 3,
        "half_of_a_dozen": 6,
        "dozen": 12,
        "small_gross": 120,
        "gross": 144,
        "great_gross": 1728,
    }
    return u
}

// NewBill creates a new bill.
func NewBill() map[string]int {
    newBill:=make(map[string]int)
    return newBill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item string, unit string) bool {
    v,c:=units[unit]
    if !c{
        return false
    }
    bill[item]+=v
    return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool{
    v,c:=bill[item]
    if !c{
        return false
    }
	v1,c1:=units[unit]
    if !c1{
        return false
    }
	newV:=v-v1
    if newV<0{
        return false
    }else if newV==0{
    	delete(bill,item)
       
    }else{
    	bill[item]-=units[unit]
    }
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
    v,c:=bill[item]
    return v,c
}
