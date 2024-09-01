package customer

import (
	"bufio"
	"challenge-godb/entity"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"
)

func InputCreateCustomer(db *sql.DB) {
	var customer entity.Customer
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Input Customer ID: ")
	scanner.Scan()
	customer.Customer_id, _ = strconv.Atoi(scanner.Text())

	fmt.Print("Input Name: ")
	scanner.Scan()
	customer.Name = scanner.Text()

	fmt.Print("Input Phone: ")
	scanner.Scan()
	customer.Phone, _ = strconv.Atoi(scanner.Text())

	fmt.Print("Input Address: ")
	scanner.Scan()
	customer.Address = scanner.Text()

	customer.Created_at = time.Now()
	customer.Updated_at = time.Now()

	CreateCustomer(db, customer)
}

func InputViewListCustomer(db *sql.DB) {
	customers := ViewOfListCustomer(db)
	for _, customer := range customers {
		fmt.Printf("%+v\n", customer)
	}
}

func InputViewCustomerDetailsByID(db *sql.DB) {
	var id int
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Input Customer ID: ")
	scanner.Scan()
	id, _ = strconv.Atoi(scanner.Text())

	customer := ViewDetailsCustomerById(db, id)
	fmt.Printf("%+v\n", customer)
}

func InputUpdateCustomer(db *sql.DB) {
	var customer entity.Customer
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Input Customer ID: ")
	scanner.Scan()
	customer.Customer_id, _ = strconv.Atoi(scanner.Text())

	fmt.Print("Input New Name: ")
	scanner.Scan()
	customer.Name = scanner.Text()

	fmt.Print("Input New Phone: ")
	scanner.Scan()
	customer.Phone, _ = strconv.Atoi(scanner.Text())

	fmt.Print("Input New Address: ")
	scanner.Scan()
	customer.Address = scanner.Text()

	customer.Created_at = time.Now()
	customer.Updated_at = time.Now()

	UpdateCustomer(db, customer)
}

func InputDeleteCustomer(db *sql.DB) {
	var id int
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Input Customer ID: ")
	scanner.Scan()
	id, _ = strconv.Atoi(scanner.Text())

	DeleteCustomer(db, id)
}

func CreateCustomer(db *sql.DB, customer entity.Customer) {
	var err error

	exists, err := ValidasiCreateByCustomerId(db, customer.Customer_id)
	if err != nil {
		fmt.Printf("Error checking if customer exists: %v\n", err)
		return
	}
	if exists {
		fmt.Println("Customer ID already exists. Please enter a different ID.")
		return
	}

	Create := "INSERT INTO customer (customer_id, name, phone, address, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6);"

	_, err = db.Exec(Create, customer.Customer_id, customer.Name, customer.Phone, customer.Address, customer.Created_at, customer.Updated_at)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Succes create customer")
	}
}

func ValidasiCreateByCustomerId(db *sql.DB, id int) (bool, error) {
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM customer WHERE customer_id=$1);", id).Scan(&exists)
	return exists, err
}

func ViewOfListCustomer(db *sql.DB) []entity.Customer {

	sqlStatment := "SELECT * FROM customer;"

	rows, err := db.Query(sqlStatment)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	customers := ScanCustomer(rows)
	return customers
}

func ScanCustomer(rows *sql.Rows) []entity.Customer {
	customers := []entity.Customer{}
	var err error

	for rows.Next() {
		customer := entity.Customer{}
		err := rows.Scan(&customer.Customer_id, &customer.Name, &customer.Phone, &customer.Address, &customer.Created_at, &customer.Updated_at)
		if err != nil {
			panic(err)
		}

		customers = append(customers, customer)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return customers
}

func ViewDetailsCustomerById(db *sql.DB, customer_id int) entity.Customer {
	var err error

	sqlStatment := "SELECT * FROM customer WHERE customer_id = $1;"

	customer := entity.Customer{}
	err = db.QueryRow(sqlStatment, customer_id).Scan(&customer.Customer_id, &customer.Name, &customer.Phone, &customer.Address, &customer.Created_at, &customer.Updated_at)

	if err == sql.ErrNoRows {
		fmt.Println("Customer not found.")
	} else if err != nil {
		panic(err)
	}
	return customer
}

func UpdateCustomer(db *sql.DB, customer entity.Customer) {
	var err error

	var exists bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM customer WHERE customer_id = $1);", customer.Customer_id).Scan(&exists)
	if err != nil {
		fmt.Println("Error checking customer existence:", err)
		return
	}

	if !exists {
		fmt.Println("Customer not found.")
		return
	}

	sqlStatment := "UPDATE customer SET name = $2, phone = $3, address = $4, Created_at = $5, Updated_at = $6 WHERE customer_id = $1;"

	_, err = db.Exec(sqlStatment, customer.Customer_id, customer.Name, customer.Phone, customer.Address, customer.Created_at, customer.Updated_at)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Succesfully Update Data")
	}
}

func DeleteCustomer(db *sql.DB, id int) {
	var err error

	var exists bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM customer WHERE customer_id = $1);", id).Scan(&exists)
	if err != nil {
		fmt.Println("Error checking customer existence:", err)
		return
	}

	if !exists {
		fmt.Println("Customer ID not found. Please enter a different ID.")
		return
	}

	var used bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM orders WHERE customer_id = $1);", id).Scan(&used)
	if err != nil {
		fmt.Println("Error checking order usage:", err)
		return
	}

	if used {
		fmt.Println("Customer ID is being used in orders. Please delete the order first.")
		return
	}

	sqlStatment := "DELETE FROM customer WHERE customer_id =$1;"

	_, err = db.Exec(sqlStatment, id)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Succes Delete Data")
	}
}
