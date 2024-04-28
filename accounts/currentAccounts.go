package accounts

type CurrentAccount struct {
	Titular               string
	Agency                int
	AccountNumber         int
	Balance               float64
	privateExampleBalance float64 // lowercase fild name is like a private attribute
}

func (c *CurrentAccount) DrawMoney(value float64) string {
	canDraw := value > 0 && value <= c.Balance

	if canDraw {
		c.Balance -= value
		return "Draw money done sucessfully!"
	} else {
		return "Insufficient balance :("
	}
}

func (c *CurrentAccount) DepositMoney(value float64) (string, float64) {
	canDeposit := value > 0

	if canDeposit {
		c.Balance += value
		return "Deposit done sucessfully!", c.Balance
	} else {
		return "Inconsistent value :(", c.Balance
	}

}

func (c *CurrentAccount) TransferMoney(value float64, destinationAcc *CurrentAccount) bool {
	if value > 0 && value <= c.Balance {
		c.Balance -= value
		destinationAcc.DepositMoney(value)
		return true
	} else {
		return false
	}
}
