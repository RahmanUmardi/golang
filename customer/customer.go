package customer

import (
	"challenge-godb/connection"
	"challenge-godb/entity"
	"database/sql"
	"fmt"
	"time"
)

func InputCreateCustomer() {
	var customer entity.Customer
	fmt.Print("Input Customer ID: ")
	fmt.Scan(&customer.Customer_id)
	fmt.Print("Input Name: ")
	fmt.Scan(&customer.Name)
	fmt.Print("Input Phone: ")
	fmt.Scan(&customer.Phone)
	fmt.Print("Input Address: ")
	fmt.Scan(&customer.Address)
	customer.Created_at = time.Now()
	customer.Updated_at = time.Now()

	CreateCustomer(customer)
}

func InputViewListCustomer() {
	customers := ViewOfListCustomer()
	for _, customer := range customers {
		fmt.Printf("%+v\n", customer)
	}
}

func InputViewCustomerDetailsByID() {
	var id int
	fmt.Print("Input Customer ID: ")
	fmt.Scan(&id)
	customer := ViewDetailsCustomerById(id)
	fmt.Printf("%+v\n", customer)
}

func InputUpdateCustomer() {
	var customer entity.Customer
	fmt.Print("Input Customer ID: ")
	fmt.Scan(&customer.Customer_id)
	fmt.Print("Input New Name: ")
	fmt.Scan(&customer.Name)
	fmt.Print("Input New Phone: ")
	fmt.Scan(&customer.Phone)
	fmt.Print("Input New Address: ")
	fmt.Scan(&customer.Address)
	customer.Created_at = time.Now()
	customer.Updated_at = time.Now()

	UpdateCustomer(customer)
}

func InputDeleteCustomer() {
	var id int
	fmt.Print("Input Customer ID: ")
	fmt.Scan(&id)
	DeleteCustomer(id)
}

func CreateCustomer(customer entity.Customer) {
	db := connection.ConnectDb()
	defer db.Close()
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

func ViewOfListCustomer() []entity.Customer {
	db := connection.ConnectDb()
	defer db.Close()

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

func ViewDetailsCustomerById(customer_id int) entity.Customer {
	db := connection.ConnectDb()
	defer db.Close()
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

func UpdateCustomer(customer entity.Customer) {
	db := connection.ConnectDb()
	defer db.Close()
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

func DeleteCustomer(id int) {
	db := connection.ConnectDb()
	defer db.Close()
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
