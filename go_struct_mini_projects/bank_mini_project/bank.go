package main
import  "fmt"
import "os"
func storeBalance(balance float64){
	balanceTxt := fmt.Sprint(balance)
	os.WriteFile("balance.txt" , []byte(balanceTxt) , 0664)

}

func main(){
	accountBalance :=1000.0
	var depositBal float64
	var withdrawBal float64
	var totalBalance float64
	fmt.Println("welcome to bank")
	for {
	
	fmt.Println("what do you want to do")
	fmt.Println("1.check the balance")
	fmt.Println("2.Deposit money")
	fmt.Println("3.Withdraw money")
	fmt.Println("4.Exit")
	var options int
	fmt.Print("chose your options: ")
	fmt.Scan(&options)
	// fmt.Println("your option is:",options)

	if options ==1 {
		// checkBalance()
		fmt.Println("Your Account balance is:",accountBalance)
		return 
	} else if options ==2{
		// depositBalance()
		fmt.Print("Please enter your Desired Amount to deposit:")
		fmt.Scan(&depositBal)
		if depositBal >0{
			totalBalance = depositBal + accountBalance
			storeBalance(totalBalance)
			fmt.Println("Your Total Account Balance is: ",totalBalance)
			return

		} else {
			fmt.Println("Invalid amount")
			continue
		}		

	} else if options ==3 {
		// withdrawBalance()
		fmt.Print("please Enter your desired amount to withdraw: ")
		fmt.Scan(&withdrawBal)
		if withdrawBal > 0 && withdrawBal <= accountBalance{
			totalBalance = accountBalance - withdrawBal
			storeBalance(totalBalance)
			fmt.Println("your total balance is: ",totalBalance)
			return

		} else{
			fmt.Println("Insufficient amount")
			continue
		}
	} else if options ==4{
		fmt.Println("Thank you for choosing backing sevice ")
		break
	}
}
}