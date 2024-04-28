package accounts

func PayTicket(account account, value float64) {
	account.DrawMoney(value)
}

type account interface {
	// Every method that implements this function and refers to a struct,
	// causes this struct to "inherit" from "account"
	DrawMoney(value float64) string
}
