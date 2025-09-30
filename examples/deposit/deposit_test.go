package deposit_test

import (
	"context"
	"fmt"
	"testing"

	"godog-examples/examples/deposit"

	"github.com/cucumber/godog"
)

type (
	accountKey struct{}
	errorKey   struct{}
)

func iHaveANewAccount(ctx context.Context) context.Context {
	acct := deposit.NewSavingsAccount()

	return context.WithValue(ctx, accountKey{}, acct)
}

func getAccount(ctx context.Context) deposit.Account {
	return ctx.Value(accountKey{}).(deposit.Account)
}

func iHaveAnAccountWith(ctx context.Context, amount int) context.Context {
	acct := deposit.NewSavingsAccount(deposit.WithBalance(amount))
	return context.WithValue(ctx, accountKey{}, acct)
}

func iDeposit(ctx context.Context, amount int) (context.Context, error) {
	acct := getAccount(ctx)
	err := acct.Deposit(amount)
	return context.WithValue(ctx, accountKey{}, acct), err
}

func iWithdraw(ctx context.Context, amount int) (context.Context, error) {
	acct := getAccount(ctx)
	err := acct.Withdraw(amount)
	return context.WithValue(ctx, accountKey{}, acct), err
}

func iTryToWithdraw(ctx context.Context, amount int) context.Context {
	acct := getAccount(ctx)
	err := acct.Withdraw(amount)
	if err != nil {
		return context.WithValue(ctx, errorKey{}, err)
	}

	return context.WithValue(ctx, accountKey{}, acct)
}

func theAccountBalanceIs(ctx context.Context, amount int) error {
	acct := getAccount(ctx)
	if acct.Balance() != amount {
		return fmt.Errorf("expected the account balance to be %d by found %d", amount, acct.Balance())
	}
	return nil
}

func theTransactionShouldError(ctx context.Context) error {
	err := ctx.Value(errorKey{})
	if err == nil {
		return fmt.Errorf("the expected error was not found")
	}
	return nil
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: func(ctx *godog.ScenarioContext) {
			ctx.Step(`^I have a new account$`, iHaveANewAccount)
			ctx.Step(`^I have an account with (\d+)`, iHaveAnAccountWith)
			ctx.Step(`^I deposit (\d+)`, iDeposit)
			ctx.Step(`^I withdraw (\d+)`, iWithdraw)
			ctx.Step(`^I try to withdraw (\d+)`, iTryToWithdraw)
			ctx.Step(`^the account balance must be (\d+)`, theAccountBalanceIs)
			ctx.Step(`^the transaction should fail$`, theTransactionShouldError)
		},
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
