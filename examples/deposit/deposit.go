package deposit

import (
	"fmt"
	"sync"
)

type Account interface {
	Balance() int
	Deposit(amount int) error
	Withdraw(amount int) error
}

type SavingsAccount struct {
	balance int
	sync.Mutex
}

type SavingsAccountOption func(*SavingsAccount)

func WithBalance(balance int) SavingsAccountOption {
	return func(s *SavingsAccount) {
		s.balance = balance
	}
}

func NewSavingsAccount(opts ...SavingsAccountOption) *SavingsAccount {
	acct := &SavingsAccount{
		balance: 0,
	}

	for _, opt := range opts {
		opt(acct)
	}

	return acct
}

func (s *SavingsAccount) Balance() int {
	return s.balance
}

func (s *SavingsAccount) Deposit(amount int) error {
	s.Lock()
	defer s.Unlock()

	newBalance := s.balance + amount
	s.balance = newBalance

	return nil
}

func (s *SavingsAccount) Withdraw(amount int) error {
	s.Lock()
	defer s.Unlock()

	newBalance := s.balance - amount

	if newBalance < 0 {
		return fmt.Errorf("insufficient balance")
	}

	s.balance = newBalance

	return nil
}
