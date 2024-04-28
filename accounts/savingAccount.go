package accounts

import "github.com/BANK-POO-GO/bank-poo-go/clients"

type SavingAccount struct {
	Titular         clients.Client
	AgencyNumber    int
	AccountNumber   int
	OperationNumber int
	balance         float64
}

func (c *SavingAccount) DrawMoney(value float64) string {
	canDraw := value > 0 && value <= c.balance

	if canDraw {
		c.balance -= value
		return "Draw money done sucessfully!"
	} else {
		return "Insufficient balance :("
	}
}

func (c *SavingAccount) DepositMoney(value float64) (string, float64) {
	canDeposit := value > 0

	if canDeposit {
		c.balance += value
		return "Deposit done sucessfully!", c.balance
	} else {
		return "Inconsistent value :(", c.balance
	}

}

func (c *SavingAccount) TransferMoney(value float64, destinationAcc *SavingAccount) bool {
	if value > 0 && value <= c.balance {
		c.balance -= value
		destinationAcc.DepositMoney(value)
		return true
	} else {
		return false
	}
}

func (c *SavingAccount) GetBalance() float64 {
	return c.balance
}
