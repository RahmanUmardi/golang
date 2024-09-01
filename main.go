package main

import (
	"challenge-godb/customer"
	"challenge-godb/order"
	"challenge-godb/service"
	"database/sql"

	"bufio"
	"challenge-godb/connection"
	"fmt"
	"os"
	"strconv"
)

func main() {
	db := connection.ConnectDb()
	defer db.Close()

	for {
		switch MainMenu() {
		case 1:
			CustomerMenu(db)
		case 2:
			ServiceMenu(db)
		case 3:
			OrderMenu(db)
		case 4:
			fmt.Println("Exit")
			os.Exit(0)
		default:
			fmt.Println("Invalid. Please input a number between 1 and 4.")
		}
	}
}

func MainMenu() int {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enigma Laundry")
	fmt.Println("1. Customer")
	fmt.Println("2. Service")
	fmt.Println("3. Order")
	fmt.Println("4. Exit")
	fmt.Print("Input : ")

	scanner.Scan()
	input, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Invalid input")
		return 0
	}
	return input
}

func CustomerMenu(db *sql.DB) {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("\nCustomer Menu:")
		fmt.Println("1. Create Customer")
		fmt.Println("2. View Of List Customer")
		fmt.Println("3. View Details Customer by ID")
		fmt.Println("4. Update Customer")
		fmt.Println("5. Delete Customer")
		fmt.Println("6. Back to Main Menu")
		fmt.Print("Input : ")
		scanner.Scan()
		input, _ := strconv.Atoi(scanner.Text())

		switch input {
		case 1:
			customer.InputCreateCustomer(db)
		case 2:
			customer.InputViewListCustomer(db)
		case 3:
			customer.InputViewCustomerDetailsByID(db)
		case 4:
			customer.InputUpdateCustomer(db)
		case 5:
			customer.InputDeleteCustomer(db)
		case 6:
			return
		default:
			fmt.Println("Invalid. Please input a number between 1 and 6.")
		}
	}
}

func ServiceMenu(db *sql.DB) {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("\nService Menu:")
		fmt.Println("1. Create Service")
		fmt.Println("2. View Of List Services")
		fmt.Println("3. View Details Service by ID")
		fmt.Println("4. Update Service")
		fmt.Println("5. Delete Service")
		fmt.Println("6. Back to Main Menu")
		fmt.Print("Input : ")
		scanner.Scan()
		input, _ := strconv.Atoi(scanner.Text())

		switch input {
		case 1:
			service.InputCreateService(db)
		case 2:
			service.InputViewListService(db)
		case 3:
			service.InputViewServiceDetailsByID(db)
		case 4:
			service.InputUpdateService(db)
		case 5:
			service.InputDeleteService(db)
		case 6:
			return
		default:
			fmt.Println("Invalid. Please input a number between 1 and 6.")
		}
	}
}

func OrderMenu(db *sql.DB) {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("\nOrder Menu:")
		fmt.Println("1. Create Order")
		fmt.Println("2. Complete Order")
		fmt.Println("3. View of List Order")
		fmt.Println("4. View Order Details by ID")
		fmt.Println("5. Back to Main Menu")
		fmt.Print("Input : ")
		scanner.Scan()
		input, _ := strconv.Atoi(scanner.Text())

		switch input {
		case 1:
			order.InputCreateOrder(db)
		case 2:
			order.InputCompleteOrder(db)
		case 3:
			order.InputViewListOrder(db)
		case 4:
			order.InputViewOrderDetailsByID(db)
		case 5:
			return
		default:
			fmt.Println("Invalid. Please input a number between 1 and 5.")
		}
	}
}
