package main

import (
	"fmt"
)

func main() {
	fulanoAccount := accountAlias.CurrentAccount{titular: "Fulano", agency: 847,
		accountNumber: 123456, balance: 1200.52}

	var cicranoAccount *CurrentAccount
	cicranoAccount = new(CurrentAccount)
	cicranoAccount.titular = "Cicrano"

	fmt.Println(*cicranoAccount)
	fmt.Println(fulanoAccount)

	fmt.Println(fulanoAccount.DrawMoney(200))
	fmt.Println(fulanoAccount)

	status, newBalance := cicranoAccount.DepositMoney(100)
	fmt.Println(status, newBalance)
	fmt.Println(*cicranoAccount)

	//-------------------------------------------
	beltranoAccount := CurrentAccount{titular: "Beltrano", balance: 100}
	zutanoAccount := CurrentAccount{titular: "Zutano", balance: 300}

	fmt.Println()
	fmt.Println(beltranoAccount)
	fmt.Println(zutanoAccount)

	zutanoAccount.TransferMoney(200, &beltranoAccount)

	fmt.Println(beltranoAccount)
	fmt.Println(zutanoAccount)
}
