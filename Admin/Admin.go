package admin

import (
	account "bankingapp/Account"
	bank "bankingapp/Bank"
	// user "bankingapp/User"
	"errors"

	uuid "github.com/satori/go.uuid"
)

type Admin struct {
	ID               uuid.UUID
	FirstName        string
	LastName         string
	userName         string
	IsAdmin          bool
	// usersCreatedByMe []*Admin
	// BanksCreatedByMe []*bank.Bank
}

var BanksCreatedByMe []*bank.Bank
var usersCreatedByMe []*Admin

func NewAdmin(firstName, lastName, username string) *Admin {
	return &Admin{
		ID:        uuid.NewV4(),
		FirstName: firstName,
		LastName:  lastName,
		userName:  username,
		IsAdmin:   true,
	}
}

func (ad *Admin) CreateBank(bankName string) (*bank.Bank, error) {
	if !ad.IsAdmin {
		return nil, errors.New("you are not an admin")
	}

	_, isBankExist := findBank(BanksCreatedByMe, bankName)
	if isBankExist {
		return nil, errors.New("bank already exists")
	}

	NewBank := bank.NewBank(bankName)
	BanksCreatedByMe = append(BanksCreatedByMe, NewBank)
	return NewBank, nil
}

func (ad *Admin) NewUser(firstName, lastName, username string) (*Admin, error) {
	if !ad.IsAdmin {
		return nil, errors.New("you are not an admin")
	}

	_, isUserExist := findUser(usersCreatedByMe, username)
	if isUserExist {
		return nil, errors.New("user already exists")
	}

	NewUser := &Admin{
		ID: uuid.NewV4(),
		FirstName: firstName,
		LastName: lastName,
		userName: username,
	}

	usersCreatedByMe = append(usersCreatedByMe, NewUser)
	return NewUser, nil
}

func findBank(banks []*bank.Bank, bankName string) (*bank.Bank, bool) {
	for i := 0; i < len(banks); i++ {
		if banks[i].BankName == bankName {
			return banks[i], true
		}
	}
	return nil, false

}

func findUser(users []*Admin, username string) (*Admin, bool) {
	// fmt.Println("...........................",users, username)
	for i := 0; i < len(users); i++ {
		// fmt.Println(">>>>>>>>>>>>>>>>>>>>>>", users[i].userName, username)
		if users[i].userName == username {
			return users[i], true
		}
	}
	return nil, false
}

func (ad *Admin) ReadBank() []*bank.Bank{
	return BanksCreatedByMe
}

func (ad *Admin) ReadUser() []*Admin {
	return usersCreatedByMe
}

func (ad *Admin) UpdateUser(username, fieldToUpdate, value string) error {
	if !ad.IsAdmin {
		return errors.New("you are not an admin")
	}

	userUpdate, isUserExist := findUser(usersCreatedByMe, username)
	if !isUserExist {
		return errors.New("user does not exists")
	}
	
	switch fieldToUpdate {
	case "userName":
		userUpdate.userName = value
	case "FirstName":
		userUpdate.FirstName = value
	case "LastName":
		userUpdate.LastName = value
	default:
		return errors.New("not valid field")
	}
	return nil
}

func (ad *Admin) DeleteUser(username string) error {
	if !ad.IsAdmin {
		return errors.New("you are not an admin")
	}
	for i := 0; i < len(usersCreatedByMe); i++ {
		if usersCreatedByMe[i].userName == username {
			usersCreatedByMe = append(usersCreatedByMe[:i], usersCreatedByMe[i+1:]...)
		}
	}
	return nil
}

func (ad *Admin) CreateUserAccount(username, bankName string, balance float64) (*account.Account, error) {
	// if ad.IsAdmin {
	// 	return nil, errors.New("you are an admin!! only user can create account")
	// }

	_, isUserExist := findUser(usersCreatedByMe, username)
	if !isUserExist {
		return nil, errors.New("user does not exists!! you can not create account")
	}

	bankObj, isBankExist := findBank(BanksCreatedByMe, bankName)
	if !isBankExist {
		return nil, errors.New("bank does not exists")
	}

	NewAccount, _ := bankObj.CreateAccount(username, balance)
	return NewAccount, nil
}

func (ad *Admin) ReadUserAccount(username, bankName string) (*account.Account, error) {
	if ad.IsAdmin {
		return nil, errors.New("you are an admin!! only user can read account")
	}

	_, isUserExist := findUser(usersCreatedByMe, username)
	if !isUserExist {
		return nil, errors.New("user does not exists")
	}

	bankObj, isBankExist := findBank(BanksCreatedByMe, bankName)
	if !isBankExist {
		return nil, errors.New("bank does not exists")
	}

	accobj, _ := bankObj.ReadAccount(username)
	return accobj, nil
}

func (ad *Admin) DipositAmount(username, bankName string, amount float64)(*account.Account,error){
	if ad.IsAdmin{
		return nil, errors.New("you are an admin!! only user can deposit account")
	}
	_, isUserExist := findUser(usersCreatedByMe, username)
	if !isUserExist {
		return nil, errors.New("user does not exists")
	}

	bankObj, isBankExist := findBank(BanksCreatedByMe, bankName)
	if !isBankExist {
		return nil, errors.New("bank does not exists")
	}

	accountObj,_:= bankObj.DepositAmount(username, amount)
	return accountObj, nil
}

func (ad *Admin) WithdrawAmount(username, bankName string, amount float64)(*account.Account,error){
	if ad.IsAdmin{
		return nil, errors.New("you are an admin!! only user can deposit account")
	}
	_, isUserExist := findUser(usersCreatedByMe, username)
	if !isUserExist {
		return nil, errors.New("user does not exists")
	}

	bankObj, isBankExist := findBank(BanksCreatedByMe, bankName)
	if !isBankExist {
		return nil, errors.New("bank does not exists")
	}

	accountObj,_:= bankObj.WithdrawAmount(username, amount)
	return accountObj, nil
}

func (ad *Admin) GetPassbook (username, bankName string)([]string,error){
	if ad.IsAdmin{
		return nil, errors.New("you are an admin!! only user can access passbook")
	}
	_, isUserExist := findUser(usersCreatedByMe, username)
	if !isUserExist {
		return nil, errors.New("user does not exists")
	}

	bankObj, isBankExist := findBank(BanksCreatedByMe, bankName)
	if !isBankExist {
		return nil, errors.New("bank does not exists")
	}
	passbook,_ := bankObj.GetPassbook(username)
	return passbook,nil
}

func (ad *Admin) DeleteBank(bankname string)(error){
	if !ad.IsAdmin{
		return errors.New("you are not an admin!! only admin can delete bank")
	}

	bankObj, isBankExist := findBank(BanksCreatedByMe, bankname)
	if !isBankExist {
		return errors.New("bank does not exists")
	}

	if len(bankObj.AccountsCreatedByMe)!=0{
		return errors.New("bank contains active account, you cannot delete the bank")
	}

	for i := 0; i < len(BanksCreatedByMe); i++ {
		if BanksCreatedByMe[i].BankName == bankname {
			BanksCreatedByMe = append(BanksCreatedByMe[:i], BanksCreatedByMe[i+1:]...)
		}
	}
	return nil
}

func (ad* Admin) TranferToBanks(sourceBankName, destBankName string, amount float64, date int)(error){
	if !ad.IsAdmin{
		return errors.New("you are not an admin")
	}

	sourceBankObj,isSourceBankExist := findBank(BanksCreatedByMe, sourceBankName)
	if !isSourceBankExist {
		return errors.New(" source bank does not exists")
	}

	destBankObj,isDestBankExist := findBank(BanksCreatedByMe, destBankName)
	if !isDestBankExist {
		return errors.New("destination bank does not exists")
	}
	
	sourceBankObj.TranferToBanks(amount,destBankObj,date)
	return nil
}

// func (ad* Admin) TranferToUserAccount(sourceAccountUsername, destAccountUserName string, amount float64)(error){
// 	if !ad.IsAdmin{
// 		return errors.New("you are not an admin")
// 	}

// 	sourceUserObj,isSourceUserExist := findUser(usersCreatedByMe, sourceAccountUsername)
// 	if !isSourceUserExist {
// 		return errors.New(" source bank does not exists")
// 	}

// 	destUserObj,isDestBankExist := findUser(usersCreatedByMe, destAccountUserName)
// 	if !isDestBankExist {
// 		return errors.New("destination bank does not exists")
// 	}

// 	sourceUserObj.TranferToUserAccount(amount,destUserObj)
// 	return nil
// }
