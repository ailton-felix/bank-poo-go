package main

import (
	"fmt"

	"github.com/BANK-POO-GO/bank-poo-go/accounts"
	"github.com/BANK-POO-GO/bank-poo-go/clients"
)

func main() {
	client1 := clients.Client{Nome: "Fulano", Cpf: "123.123.123-12",
		Occupation: "Programer", Age: 18}
	account1 := accounts.CurrentAccount{Titular: client1, Agency: 847,
		AccountNumber: 123456}

	account1.DepositMoney(1200.52)

	fmt.Print("Saldo da conta do " + account1.Titular.Nome + ": ")
	fmt.Println(account1.GetBalance())

	// -------------------------------------
	// interfaces
	fmt.Println("\n\n# Testes com uso de interface\n")

	accountPaulo := accounts.CurrentAccount{}
	accountPaulo.DepositMoney(100)
	fmt.Println("Conta Paulo")
	fmt.Println("Saldo:", accountPaulo.GetBalance())
	accounts.PayTicket(&accountPaulo, 50)
	fmt.Println("Saldo pós pagamento:", accountPaulo.GetBalance())
	println()

	accountMaya := accounts.SavingAccount{}
	accountMaya.DepositMoney(500)
	fmt.Println("Conta Maya")
	fmt.Println("Saldo:", accountMaya.GetBalance())
	accounts.PayTicket(&accountMaya, 400)
	fmt.Println("Saldo pós pagamento:", accountMaya.GetBalance())
	println()

	/* Using "CurrentAccount.Titular" as a string
	fulanoAccount := acc.CurrentAccount{Titular: "Fulano", Agency: 847,
		AccountNumber: 123456, Balance: 1200.52}

	var cicranoAccount *acc.CurrentAccount
	cicranoAccount = new(acc.CurrentAccount)
	cicranoAccount.Titular = "Cicrano"

	fmt.Println(*cicranoAccount)
	fmt.Println(fulanoAccount)

	fmt.Println(fulanoAccount.DrawMoney(200))
	fmt.Println(fulanoAccount)

	status, newBalance := cicranoAccount.DepositMoney(100)
	fmt.Println(status, newBalance)
	fmt.Println(*cicranoAccount)

	//-------------------------------------------
	beltranoAccount := acc.CurrentAccount{Titular: "Beltrano", Balance: 100}
	zutanoAccount := acc.CurrentAccount{Titular: "Zutano", Balance: 300}

	fmt.Println()
	fmt.Println(beltranoAccount)
	fmt.Println(zutanoAccount)

	zutanoAccount.TransferMoney(200, &beltranoAccount)

	fmt.Println(beltranoAccount)
	fmt.Println(zutanoAccount)*/
}
