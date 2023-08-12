package account

import "errors"

// Account represents a bank account with an owner and a balance.
type Account struct {
	owner   string // The owner of the account
	balance int    // The current balance of the account
}

var errInsufficientBalance = errors.New("error: insufficient balance")

// Note: we don't have to make the receiver a pointer since we are not modifying the value
func (a Account) Owner() string {
	return a.owner
}

func (a Account) Balance() int {
	return a.balance
}

// CreateAccount creates a new account with the given owner and an initial balance of 0.
// @return a pointer to the newly created account.
// Note: it is important to return the pointer so that we're not making duplicates of Account object created
func CreateAccount(owner string) *Account {
	return &Account{owner, 0}
}

// SetBalance sets the balance of the account to the given amount.
// Note: (a *Account) is called the pointer receiver, which:
// * binds SetBalance function with Account objects
// * modifies the values of the Account object who's the caller of this function
func (a *Account) SetBalance(amount int) {
	a.balance = amount
}

func (a *Account) Deposit(amount int) {
	a.balance += amount
}

func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errInsufficientBalance
	}
	a.balance -= amount
	return nil
}

func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}
