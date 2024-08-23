package order

import (
	"challenge-godb/connection"
	"challenge-godb/entity"
	"database/sql"
	"fmt"
	"time"
)

func InputCreateOrder() {
	var order entity.Order
	fmt.Print("Input Order ID: ")
	fmt.Scan(&order.Order_id)
	fmt.Print("Input Customer ID: ")
	fmt.Scan(&order.Customer_id)
	fmt.Print("Input Receiver: ")
	fmt.Scan(&order.Received_by)
	fmt.Print("Input Order Date (YYYY-MM-DD): ")
	fmt.Scan(&order.Order_date)
	order.Created_at = time.Now()
	order.Updated_at = time.Now()

	CreateOrder(order)
}

func InputCompleteOrder() {
	var id int
	fmt.Print("Input Order ID: ")
	fmt.Scan(&id)
	var completionDate time.Time
	fmt.Print("Input Completion Date (YYYY-MM-DD): ")
	fmt.Scan(&completionDate)

	CompleteOrder(id, completionDate)
}

func InputViewListOrder() {
	orders := ViewOfListOrder()
	for _, order := range orders {
		fmt.Printf("%+v\n", order)
	}
}

func InputViewOrderDetailsByID() {
	var id int
	fmt.Print("Input Order ID: ")
	fmt.Scan(&id)
	order := ViewDetailsOrderById(id)
	fmt.Printf("%+v\n", order)
}

func CreateOrder(order entity.Order) {
	db := connection.ConnectDb()
	defer db.Close()
	var err error

	var customerExists bool
	customerQuery := "SELECT EXISTS(SELECT 1 FROM customer WHERE customer_id=$1;)"
	err = db.QueryRow(customerQuery, order.Customer_id).Scan(&customerExists)
	if err != nil {
		fmt.Printf("Failed to check if customer exists: %v\n", err)
		return
	}
	if !customerExists {
		fmt.Println("Customer not found.")
		return
	}

	var orderExists bool
	orderQuery := "SELECT EXISTS(SELECT 1 FROM orders WHERE order_id=$1;)"
	err = db.QueryRow(orderQuery, order.Order_id).Scan(&orderExists)
	if err != nil {
		fmt.Printf("Failed to check if order exists: %v\n", err)
		return
	}
	if orderExists {
		fmt.Println("Order ID already exists. Please enter a different ID.")
		return
	}

	Create := "INSERT INTO orders (order_id, customer_id, order_date, completion_date, received_by, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7);"

	_, err = db.Exec(Create, order.Order_id, order.Customer_id, order.Order_date, order.Completion_date, order.Received_by, order.Created_at, order.Updated_at)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Succes create order")
	}
}

func CompleteOrder(orderID int, completionDate time.Time) {
	db := connection.ConnectDb()
	defer db.Close()
	var err error

	var orderExists bool
	orderQuery := "SELECT EXISTS(SELECT 1 FROM orders WHERE order_id=$1);"
	err = db.QueryRow(orderQuery, orderID).Scan(&orderExists)
	if err != nil {
		fmt.Printf("Failed to check if order exists: %v\n", err)
		return
	}
	if !orderExists {
		fmt.Println("Order not found.")
		return
	}

	updateOrder := "UPDATE orders SET completion_date=$1, updated_at=NOW() WHERE order_id=$2;"
	_, err = db.Exec(updateOrder, completionDate, orderID)
	if err != nil {
		fmt.Printf("Failed to complete order: %v\n", err)
		return
	}

	fmt.Println("Order successfully completed.")
}

func ViewOfListOrder() []entity.Order {
	db := connection.ConnectDb()
	defer db.Close()

	sqlStatment := "SELECT * FROM orders;"

	rows, err := db.Query(sqlStatment)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	orders := ScanOrder(rows)
	return orders
}

func ScanOrder(rows *sql.Rows) []entity.Order {
	orders := []entity.Order{}
	var err error

	for rows.Next() {
		order := entity.Order{}
		err := rows.Scan(&order.Order_id, &order.Customer_id, &order.Order_date, &order.Completion_date, &order.Received_by, &order.Created_at, &order.Updated_at)

		if err != nil {
			panic(err)
		}

		orders = append(orders, order)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return orders
}

func ViewDetailsOrderById(order_id int) entity.Order {
	db := connection.ConnectDb()
	defer db.Close()
	var err error

	sqlStatment := "SELECT * FROM orders WHERE order_id = $1;"

	order := entity.Order{}
	err = db.QueryRow(sqlStatment, order_id).Scan(&order.Order_id, &order.Customer_id, &order.Order_date, &order.Completion_date, &order.Received_by, &order.Created_at, &order.Updated_at)

	if err == sql.ErrNoRows {
		fmt.Println("Order not found.")
	} else if err != nil {
		panic(err)
	}
	return order
}
