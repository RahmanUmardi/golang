package order

import (
	"bufio"
	"challenge-godb/entity"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"
)

func InputCreateOrder(db *sql.DB) {
	var order entity.Order
	var orderDetail entity.OrderDetail
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Input Order ID: ")
	scanner.Scan()
	order.Order_id, _ = strconv.Atoi(scanner.Text())
	orderDetail.Order_id, _ = strconv.Atoi(scanner.Text())

	fmt.Print("Input Order Detail ID: ")
	scanner.Scan()
	orderDetail.Order_detail_id, _ = strconv.Atoi(scanner.Text())

	fmt.Print("Input Customer ID: ")
	scanner.Scan()
	order.Customer_id, _ = strconv.Atoi(scanner.Text())

	fmt.Print("Input Service ID: ")
	scanner.Scan()
	orderDetail.Service_id, _ = strconv.Atoi(scanner.Text())

	fmt.Print("Input Qty: ")
	scanner.Scan()
	orderDetail.Qty, _ = strconv.Atoi(scanner.Text())

	fmt.Print("Input Order Date (YYYY-MM-DD): ")
	scanner.Scan()
	orderDate, _ := time.Parse("2006-01-02", scanner.Text())
	order.Order_date = orderDate

	order.Completion_date = sql.NullTime{}

	fmt.Print("Input Receiver: ")
	scanner.Scan()
	order.Received_by = scanner.Text()

	order.Created_at = time.Now()
	order.Updated_at = time.Now()

	CreateOrder(db, order)
	CreateOrderDetail(db, orderDetail)
}

func InputCompleteOrder(db *sql.DB) {
	var id int
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Input Order ID: ")
	scanner.Scan()
	id, _ = strconv.Atoi(scanner.Text())

	fmt.Print("Input Completion Date (YYYY-MM-DD): ")
	scanner.Scan()
	completionDate, _ := time.Parse("2006-01-02", scanner.Text())

	CompleteOrder(db, id, completionDate)
}

func InputViewListOrder(db *sql.DB) {
	orders := ViewOfListOrder(db)
	for _, order := range orders {
		fmt.Printf("%+v\n", order)
	}
}

func InputViewOrderDetailsByID(db *sql.DB) {
	var id int
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Input Order ID: ")
	scanner.Scan()
	id, _ = strconv.Atoi(scanner.Text())

	order := ViewDetailsOrderById(db, id)
	fmt.Printf("%+v\n", order)
}

func CreateOrder(db *sql.DB, order entity.Order) {
	var err error

	var customerExists bool
	customerQuery := "SELECT EXISTS(SELECT 1 FROM customer WHERE customer_id=$1)"
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
	orderQuery := "SELECT EXISTS(SELECT 1 FROM orders WHERE order_id=$1)"
	err = db.QueryRow(orderQuery, order.Order_id).Scan(&orderExists)
	if err != nil {
		fmt.Printf("Failed to check if order exists: %v\n", err)
		return
	}
	if orderExists {
		fmt.Println("Order ID already exists. Please enter a different ID.")
		return
	}

	Create := "INSERT INTO orders (order_id, customer_id, order_date, completion_date, received_by, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)"

	_, err = db.Exec(Create, order.Order_id, order.Customer_id, order.Order_date, order.Completion_date, order.Received_by, order.Created_at, order.Updated_at)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Succes create order")
	}
}

func CreateOrderDetail(db *sql.DB, order_detail entity.OrderDetail) {
	var err error

	var orderExists bool
	orderQuery := "SELECT EXISTS(SELECT 1 FROM orders WHERE order_id=$1)"
	err = db.QueryRow(orderQuery, order_detail.Order_id).Scan(&orderExists)
	if err != nil {
		fmt.Printf("Failed to check if order exists: %v\n", err)
		return
	}
	if !orderExists {
		fmt.Println("order not found.")
		return
	}

	Create := "INSERT INTO order_detail (order_detail_id, order_id, service_id, qty) VALUES ($1, $2, $3, $4, )"

	_, err = db.Exec(Create, order_detail.Order_detail_id, order_detail.Order_id, order_detail.Service_id, order_detail.Qty)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Succes create order detail")
	}
}

func CompleteOrder(db *sql.DB, orderID int, completionDate time.Time) {
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

func ViewOfListOrder(db *sql.DB) []entity.Order {

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

func ViewDetailsOrderById(db *sql.DB, order_id int) entity.Order {
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
