package main

import "fmt"

/*type sender struct{
	ratelimit int
	user //Embeded Struct
}

type user struct{
	name string
	number int
}

func test(s sender){
	fmt.Println(s.name)
	fmt.Println(s.number)
	fmt.Println(s.ratelimit)
}
func main() {
	test(sender{
		ratelimit: 100,
		user: user{
			name: "vishwa",
			number: 9600,

		},
	})
	//var smslimit int
	//var cost float64
	//var perm bool
	//var user string
	//smslimit, cost := 2.0, 5
	//smslimit := 2.0
	//const newsmslimit = 3.0
	//newsmslimit := int(smslimit)
	//msj := fmt.Sprintf(
		//"%v %f %v %q", smslimit, cost, perm, user,
		//"cost of %.2f is %v", smslimit, cost,
	//	"the updated limit is %.1f & the old one is %.1f ", newsmslimit, smslimit,
	//)

	//if smslimit<=newsmslimit{
	//	fmt.Println("lesser")
	//} else {
	//	fmt.Println("greater")
	//}

	/*smslimit=newadd(smslimit,newsmslimit)
	fmt.Printf("updated value is %.2f",smslimit)


	//fmt.Println(msj)
}
func newadd(sms1,sms2 float64)float64{
	sms1 = sms1+sms2
	return sms1

}
*/

type auth struct {
	name  string
	passd string
}

func (ai auth) getauth() string { // struct method
	return fmt.Sprintf("Auth: %s:%s", ai.name, ai.passd)
}

func test(Ai auth) {
	fmt.Println(Ai.getauth())
}

func main() {
	test(auth{
		name:  "Google",
		passd: "324",
	})
	test(auth{
		name:  "aws",
		passd: "567",
	})
}
