package loops

import "fmt"

func PrimeNumbers() {
	var sayi int
	fmt.Print("Enter a number : ")
	fmt.Scan(&sayi)

	asalMi := true

	for i := 2; i < sayi; i++ {
		if sayi&i == 0 {
			asalMi = false
		}
	}

	if asalMi {
		fmt.Println("It's Prime")
	} else {
		fmt.Println("It's not Prime")
	}

}
