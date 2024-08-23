package main

import (
	"challenge-godb/customer"
	"challenge-godb/order"
	"challenge-godb/service"
	"fmt"
	"os"
)

func main() {
	for {
		switch MainMenu() {
		case 1:
			CustomerMenu()
		case 2:
			ServiceMenu()
		case 3:
			OrderMenu()
		case 4:
			fmt.Println("Exit")
			os.Exit(0)
		default:
			fmt.Println("Invalid. Please input a number between 1 and 4.")
		}
	}
}

func MainMenu() int {
	fmt.Println("Enigma Laundry")
	fmt.Println("1. Customer")
	fmt.Println("2. Service")
	fmt.Println("3. Order")
	fmt.Println("4. Exit")
	fmt.Print("Input : ")
	var input int
	fmt.Scan(&input)
	return input
}

func CustomerMenu() {
	for {
		fmt.Println("\nCustomer Menu:")
		fmt.Println("1. Create Customer")
		fmt.Println("2. View Of List Customer")
		fmt.Println("3. View Details Customer by ID")
		fmt.Println("4. Update Customer")
		fmt.Println("5. Delete Customer")
		fmt.Println("6. Back to Main Menu")
		fmt.Print("Input : ")
		var input int
		fmt.Scan(&input)

		switch input {
		case 1:
			customer.InputCreateCustomer()
		case 2:
			customer.InputViewListCustomer()
		case 3:
			customer.InputViewCustomerDetailsByID()
		case 4:
			customer.InputUpdateCustomer()
		case 5:
			customer.InputDeleteCustomer()
		case 6:
			return
		default:
			fmt.Println("Invalid. Please input a number between 1 and 6.")
		}
	}
}

func ServiceMenu() {
	for {
		fmt.Println("\nService Menu:")
		fmt.Println("1. Create Service")
		fmt.Println("2. View Of List Services")
		fmt.Println("3. View Details Service by ID")
		fmt.Println("4. Update Service")
		fmt.Println("5. Delete Service")
		fmt.Println("6. Back to Main Menu")
		fmt.Print("Input : ")
		var input int
		fmt.Scan(&input)

		switch input {
		case 1:
			service.InputCreateService()
		case 2:
			service.InputViewListService()
		case 3:
			service.InputViewServiceDetailsByID()
		case 4:
			service.InputUpdateService()
		case 5:
			service.InputDeleteService()
		case 6:
			return
		default:
			fmt.Println("Invalid. Please input a number between 1 and 6.")
		}
	}
}

func OrderMenu() {
	for {
		fmt.Println("\nOrder Menu:")
		fmt.Println("1. Create Order")
		fmt.Println("2. Complete Order")
		fmt.Println("3. View of List Order")
		fmt.Println("4. View Order Details by ID")
		fmt.Println("5. Back to Main Menu")
		fmt.Print("Input : ")
		var input int
		fmt.Scan(&input)

		switch input {
		case 1:
			order.InputCreateOrder()
		case 2:
			order.InputCompleteOrder()
		case 3:
			order.InputViewListOrder()
		case 4:
			order.InputViewOrderDetailsByID()
		case 5:
			return
		default:
			fmt.Println("Invalid. Please input a number between 1 and 5.")
		}
	}
}
