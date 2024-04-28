/**
* A linguagem Go não possui o conceito de classe.
 */

package accounts

import "github.com/BANK-POO-GO/bank-poo-go/clients"

type CurrentAccount struct {
	Titular       clients.Client
	Agency        int
	AccountNumber int
	balance       float64 // lowercase fild name is like a private attribute
}

// Métodos são definidos de maneira parecida com funções, mas de uma maneira diferente
// Existe um (c *CurrentAccount) que se refere a um ponteiro para a instância criada da estrutura
func (c *CurrentAccount) DrawMoney(value float64) string {
	canDraw := value > 0 && value <= c.balance

	if canDraw {
		c.balance -= value
		return "Draw money done sucessfully!"
	} else {
		return "Insufficient balance :("
	}
}

func (c *CurrentAccount) DepositMoney(value float64) (string, float64) {
	canDeposit := value > 0

	if canDeposit {
		c.balance += value
		return "Deposit done sucessfully!", c.balance
	} else {
		return "Inconsistent value :(", c.balance
	}

}

func (c *CurrentAccount) TransferMoney(value float64, destinationAcc *CurrentAccount) bool {
	if value > 0 && value <= c.balance {
		c.balance -= value
		destinationAcc.DepositMoney(value)
		return true
	} else {
		return false
	}
}

func (c *CurrentAccount) GetBalance() float64 {
	return c.balance
}
