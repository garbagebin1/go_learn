package wallet


import (
	"fmt"
)
type Stringer interface {
	String() string
}

type Bitcoin int


func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Println("address of balance in Deposit is", &w.balance)
	w.balance += amount
}


func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}