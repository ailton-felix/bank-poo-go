package accounts

type CurrentAccount struct {
	titular       string
	agency        int
	accountNumber int
	balance       float64
}

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
