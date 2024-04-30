package main

import (
	"fmt"
)

func main() {

	var bigIntA BigInt
	var bigIntB BigInt
	var bigIntC BigInt

	bigIntA.setHex("36f028580bb02cc8272a9a020f4200e346e276ae664e45ee80745574e2f5ab80")
	bigIntB.setHex("70983d692f648185febe6d6fa607630ae68649f7e6fc45b94680096c06e4fadb")

	fmt.Println("A is 36f028580bb02cc8272a9a020f4200e346e276ae664e45ee80745574e2f5ab80")
	fmt.Println("B is 70983d692f648185febe6d6fa607630ae68649f7e6fc45b94680096c06e4fadb")

	bigIntC = bigIntA.ADD(bigIntB)
	fmt.Println("ADD = ", bigIntC.GetHex())

	fmt.Printf("A > B is %v\n", bigIntA.moreThan(bigIntB))
	fmt.Printf("A >= B is %v\n", bigIntA.moreOrEqualThan(bigIntB))
	fmt.Printf("A < B is %v\n", bigIntA.lessThan(bigIntB))
	fmt.Printf("A <= B is %v\n", bigIntA.lessOrEqualThan(bigIntB))
	fmt.Printf("A = B is %v\n\n", bigIntA.equal(bigIntB))

	bigIntA.setHex("FFFFFFFFFFFFFFFF")
	bigIntB.setHex("1")

	fmt.Println("A is FFFFFFFFFFFFFFFF")
	fmt.Println("B is 1")

	bigIntC = bigIntA.ADD(bigIntB)
	fmt.Println("ADD = ", bigIntC.GetHex())

	fmt.Printf("A > B is %v\n", bigIntA.moreThan(bigIntB))
	fmt.Printf("A >= B is %v\n", bigIntA.moreOrEqualThan(bigIntB))
	fmt.Printf("A < B is %v\n", bigIntA.lessThan(bigIntB))
	fmt.Printf("A <= B is %v\n", bigIntA.lessOrEqualThan(bigIntB))
	fmt.Printf("A = B is %v\n\n", bigIntA.equal(bigIntB))

}
