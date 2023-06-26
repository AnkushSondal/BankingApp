package account

import (
	"errors"
	"strconv"

	uuid "github.com/satori/go.uuid"
)

type Account struct {
	ID       uuid.UUID
	UserName string
	Balance  float64
	Passbook []string
}
func NewAccount(userName string, balance float64) (*Account,error) {
	if balance<1000{
		return nil,errors.New("minimun balance at the time of account creation must be 1000")
	}
	newAcc:= &Account{
		ID: uuid.NewV4(),
		UserName: userName,
		Balance:  balance,
	}
	newAcc.Passbook =append(newAcc.Passbook, "Account created with amount :"+ strconv.Itoa(int(balance)))
	return newAcc,nil
}

func (ac *Account) ReadAccount() *Account {
	return ac
}

func (ac *Account) DepositAmount(amount float64)(*Account){
	ac.Balance+=amount
	ac.Passbook=append(ac.Passbook, "\nAccount credited with amount :"+strconv.Itoa(int(amount)))
	return ac
}

func (ac *Account) WithdrawAmount(amount float64)(*Account,error){
	if ac.Balance<amount{
		return nil,errors.New("you can't withdraw amount more than your balance")
	}
	ac.Balance-=amount
	ac.Passbook=append(ac.Passbook, "\nAccount debited with amount :"+strconv.Itoa(int(amount)))
	return ac,nil
}

func (ac *Account) GetPassbook ()([]string){
	return ac.Passbook
}
