package service

import (
	"bufio"
	"challenge-godb/entity"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"
)

func InputCreateService(db *sql.DB) {
	scanner := bufio.NewScanner(os.Stdin)
	var service entity.Service

	fmt.Print("Input Service ID: ")
	scanner.Scan()
	service.Service_id, _ = strconv.Atoi(scanner.Text())

	fmt.Print("Input Service Name: ")
	scanner.Scan()
	service.Service_name = scanner.Text()

	fmt.Print("Input Unit: ")
	scanner.Scan()
	service.Unit = scanner.Text()

	fmt.Print("Input Price: ")
	scanner.Scan()
	service.Price, _ = strconv.Atoi(scanner.Text())

	service.Created_at = time.Now()
	service.Updated_at = time.Now()

	CreateService(db, service)
}

func InputViewListService(db *sql.DB) {
	services := ViewOfListService(db)
	for _, service := range services {
		fmt.Printf("%+v\n", service)
	}
}

func InputViewServiceDetailsByID(db *sql.DB) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Input Service ID: ")
	scanner.Scan()
	id, _ := strconv.Atoi(scanner.Text())

	service := ViewDetailsServiceById(db, id)
	fmt.Printf("%+v\n", service)
}

func InputUpdateService(db *sql.DB) {
	scanner := bufio.NewScanner(os.Stdin)
	var service entity.Service

	fmt.Print("Input Service ID: ")
	scanner.Scan()
	service.Service_id, _ = strconv.Atoi(scanner.Text())

	fmt.Print("Input New Service Name: ")
	scanner.Scan()
	service.Service_name = scanner.Text()

	fmt.Print("Input New Unit: ")
	scanner.Scan()
	service.Unit = scanner.Text()

	fmt.Print("Input New Price: ")
	scanner.Scan()
	service.Price, _ = strconv.Atoi(scanner.Text())

	service.Created_at = time.Now()
	service.Updated_at = time.Now()

	UpdateService(db, service)
}

func InputDeleteService(db *sql.DB) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Input Service ID: ")
	scanner.Scan()
	id, _ := strconv.Atoi(scanner.Text())

	DeleteService(db, id)
}

func CreateService(db *sql.DB, service entity.Service) {
	var err error

	exists, err := ValidasiCreateByServiceId(db, service.Service_id)
	if err != nil {
		fmt.Printf("Error checking if customer exists: %v\n", err)
		return
	}
	if exists {
		fmt.Println("Service ID already exists. Please enter a different ID.")
		return
	}

	Create := "INSERT INTO service (service_id, service_name, unit, price, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6);"

	_, err = db.Exec(Create, service.Service_id, service.Service_name, service.Unit, service.Price, service.Created_at, service.Updated_at)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Succes create service")
	}
}

func ValidasiCreateByServiceId(db *sql.DB, id int) (bool, error) {
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM service WHERE service_id=$1);", id).Scan(&exists)
	return exists, err
}

func ViewOfListService(db *sql.DB) []entity.Service {

	sqlStatment := "SELECT * FROM service;"

	rows, err := db.Query(sqlStatment)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	service := ScanService(rows)
	return service
}

func ScanService(rows *sql.Rows) []entity.Service {
	services := []entity.Service{}
	var err error

	for rows.Next() {
		service := entity.Service{}
		err := rows.Scan(&service.Service_id, &service.Service_name, &service.Unit, &service.Price, &service.Created_at, &service.Updated_at)

		if err != nil {
			panic(err)
		}
		services = append(services, service)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return services
}

func ViewDetailsServiceById(db *sql.DB, service_id int) entity.Service {
	var err error

	sqlStatment := "SELECT * FROM service WHERE service_id = $1;"

	service := entity.Service{}
	err = db.QueryRow(sqlStatment, service_id).Scan(&service.Service_id, &service.Service_name, &service.Unit, &service.Price, &service.Created_at, &service.Updated_at)

	if err == sql.ErrNoRows {
		fmt.Println("Service not found.")
	} else if err != nil {
		panic(err)
	}
	return service
}

func UpdateService(db *sql.DB, service entity.Service) {
	var err error

	var exists bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM service WHERE service_id = $1);", service.Service_id).Scan(&exists)
	if err != nil {
		fmt.Println("Error checking service existence:", err)
		return
	}

	if !exists {
		fmt.Println("service not found.")
		return
	}

	sqlStatment := "UPDATE service SET service_name = $2, unit = $3, price = $4, Created_at = $5, Updated_at = $6 WHERE service_id = $1;"

	_, err = db.Exec(sqlStatment, service.Service_id, service.Service_name, service.Unit, service.Price, service.Created_at, service.Updated_at)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Succesfully Update Data")
	}
}

func DeleteService(db *sql.DB, id int) {
	var err error

	var exists bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM service WHERE service_id = $1);", id).Scan(&exists)
	if err != nil {
		fmt.Println("Error checking service existence:", err)
		return
	}

	if !exists {
		fmt.Println("service ID not found. Please enter a different ID.")
		return
	}

	var used bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM orders WHERE customer_id = $1);", id).Scan(&used)
	if err != nil {
		fmt.Println("Error checking order usage:", err)
		return
	}

	if used {
		fmt.Println("Service ID is being used in orders. Please delete the order first.")
		return
	}

	sqlStatment := "DELETE FROM service WHERE service_id =$1;"

	_, err = db.Exec(sqlStatment, id)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Succes Delete Data")
	}
}
