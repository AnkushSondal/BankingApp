package main

import (
	admin "bankingapp/Admin"
	"fmt"
)

func main() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println(r)
		}
	}()

	admin1 := admin.NewAdmin("Ankush", "Sondal", "AS")
	fmt.Println(admin1)

	bank1, _ := admin1.CreateBank("SBI")
	fmt.Println("BANK CREATED :", bank1)

	user1, _ := admin1.NewUser("sanjeev", "yadav", "SY")
	fmt.Println(user1)

	_, err := user1.CreateUserAccount("SY", "SBI", 1000.0)
	if err != nil {
		panic(err)
	}

	accCreated, err2 := user1.ReadUserAccount("SY", "SBI")
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(accCreated.ReadAccount())

	userCreated := admin1.ReadUser()

	fmt.Print("All Users>>>>")
	for i := 0; i < len(userCreated); i++ {
		fmt.Print(userCreated[i].FirstName, " ")
	}
	admin1.UpdateUser("SY", "FirstName", "SANJIIV")
	userCreated = admin1.ReadUser()

	fmt.Print("\nAfter Updating user\nAll Users>>>>>  ")
	for i := 0; i < len(userCreated); i++ {
		fmt.Print(userCreated[i].FirstName, " ")
	}

	// admin1.DeleteUser("SY")
	// fmt.Println("\nUSERS after DELETION>>>>>>>>>>>>>",admin1.ReadUser())

	user1.DipositAmount("SY", "SBI", 100000)
	fmt.Println("\nBalance >>>>>>>>>> : ", accCreated.ReadAccount().Balance)

	// user1.WithdrawAmount("SY", "SBI", 2000)
	// fmt.Println("Balance >>>>>>>>>> : ", accCreated.ReadAccount().Balance)

	passbook, _ := user1.GetPassbook("SY", "SBI")
	fmt.Println("Passbook for", user1.FirstName, " :\n", passbook)

	bank2, _ := admin1.CreateBank("HDFC")
	fmt.Println("BANK CREATED :", bank2)

	fmt.Print("All banks>>>>>>>>>>> : ")
	banks := admin1.ReadBank()
	for i := 0; i < len(banks); i++ {
		fmt.Print(banks[i].BankName, ", ")
	}

	admin1.TranferToBanks("SBI", "HDFC", 1055, 10)

	fmt.Println("\nTransaction details of", bank1.BankName, "from 5th to 10th")
	bank1Transactions := bank1.BankTransactions(5, 10)
	fmt.Println(bank1Transactions)
	bank1finalBankAmount := bank1.GetTotalAmount()
	fmt.Println("final funds in",bank1.BankName,"account :", bank1finalBankAmount)

	// fmt.Print("\ndeleting HDFC bank\nAll banks>>>>>>>>>>>  ")
	// err = admin1.DeleteBank("HDFC")
	// if err != nil {
	// 	panic(err)
	// }
	// banks = admin1.ReadBank()
	// for i := 0; i < len(banks); i++ {
	// 	fmt.Print(banks[i].BankName, ", ")
	// }
	// fmt.Println("\ndeleting SBI bank")

	// err = admin1.DeleteBank("SBI")
	// if err != nil {
	// 	panic(err)
	// }
}
