package bank

import (
	account "bankingapp/Account"
	"errors"
	"strconv"

	uuid "github.com/satori/go.uuid"
)

var AllBanks []string

type Bank struct {
	ID uuid.UUID
	BankName string
	totalAmont float64
	AccountsCreatedByMe []*account.Account
	// BankPassbook []string
	BankPB map[int]string

}

func NewBank(bankName string) *Bank{
	AllBanks = append(AllBanks, bankName)
	return &Bank{
		BankName: bankName,
		totalAmont: 0.0,
		BankPB: make(map[int]string),
	}
}

func findAccount(accounts []*account.Account, username string) (*account.Account,bool) {
	for i := 0; i < len(accounts); i++ {
		if accounts[i].UserName==username{
			return accounts[i], true
		}
	}
	return nil, false
}

func (b *Bank) CreateAccount(username string, balance float64) (*account.Account,error){
	_,isAccountExist := findAccount(b.AccountsCreatedByMe,username)
	if isAccountExist{
		return nil, errors.New("account already exists")
	}
	NewAccount,_:= account.NewAccount(username, balance)
	b.AccountsCreatedByMe = append(b.AccountsCreatedByMe, NewAccount)
	b.totalAmont+=balance
	return NewAccount, nil
}

func (b *Bank) ReadAccount(username string) (*account.Account,error){
	accountObj,isAccountExist := findAccount(b.AccountsCreatedByMe,username)
	if !isAccountExist{
		return nil, errors.New("account does not exists")
	}
	return accountObj.ReadAccount(),nil
	
}

func (b *Bank) DepositAmount(username string, amount float64)(*account.Account,error){
	accountObj,isAccountExist := findAccount(b.AccountsCreatedByMe,username)
	if !isAccountExist{
		return nil, errors.New("account does not exists")
	}
	b.totalAmont+=amount

	return accountObj.DepositAmount(amount), nil
}

func (b *Bank) WithdrawAmount(username string, amount float64)(*account.Account,error){
	accountObj,isAccountExist := findAccount(b.AccountsCreatedByMe,username)
	if !isAccountExist{
		return nil, errors.New("account does not exists")
	}

	b.totalAmont-=amount
	newAccObj,_:= accountObj.WithdrawAmount(amount)
	return newAccObj, nil
}

func (b *Bank) GetPassbook(username string)([]string,error){
	accountObj,isAccountExist := findAccount(b.AccountsCreatedByMe,username)
	if !isAccountExist{
		return nil, errors.New("account does not exists")
	}

	Passbook := accountObj.GetPassbook()
	return Passbook, nil
}

func (b *Bank) TranferToBanks(amount float64, destBank *Bank, date int) error{
	if b.totalAmont<amount{
		return errors.New("insufficient funds in bank")
	}
	b.totalAmont-=amount
	destBank.totalAmont+=amount
	b.BankPB[date]=strconv.Itoa(int(amount))+" transferred to "+destBank.BankName
	destBank.BankPB[date]=strconv.Itoa(int(amount))+" received from "+b.BankName
	// b.BankPassbook=append(b.BankPassbook, strconv.Itoa(int(amount))+" transferred to "+destBank.BankName)
	// destBank.BankPassbook=append(destBank.BankPassbook, strconv.Itoa(int(amount))+" received from "+b.BankName)
	return nil
}

func (b *Bank) BankTransactions (fromDate, toDate int) map[int]string{
	transaction := make(map[int]string)
	for i := fromDate; i <= toDate; i++ {
		if b.BankPB[i]!=""{
			transaction[i] = b.BankPB[i]
		}else{
			transaction[i]= "No transaction"
		}
	}
	return transaction
}

func (b *Bank) GetTotalAmount ()(float64){
	return b.totalAmont
}


