INSERT INTO customer (
	customer_id, name,
	phone, address, 
	created_at, 
	updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6);

SELECT EXISTS(SELECT 1 FROM customer WHERE customer_id=$1);

SELECT * FROM customer;

SELECT * FROM customer WHERE customer_id = $1;

SELECT EXISTS(SELECT 1 FROM customer WHERE customer_id = $1);

UPDATE customer SET 
	name = $2, 
	phone = $3, 
	address = $4, 
	Created_at = $5, 
	Updated_at = $6 
	WHERE customer_id = $1;

SELECT EXISTS(SELECT 1 FROM customer WHERE customer_id = $1);

SELECT EXISTS(SELECT 1 FROM orders WHERE customer_id = $1);

DELETE FROM customer WHERE customer_id =$1;

INSERT INTO service (
	service_id, 
	service_name, 
	unit, price, 
	created_at, 
	updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6);

SELECT EXISTS(SELECT 1 FROM service WHERE service_id=$1);

SELECT * FROM service;

SELECT * FROM service WHERE service_id = $1;

SELECT EXISTS(SELECT 1 FROM service WHERE service_id = $1);

UPDATE service SET 
	service_name = $2, 
	unit = $3, 
	price = $4, 
	Created_at = $5, 
	Updated_at = $6 
	WHERE service_id = $1;

SELECT EXISTS(SELECT 1 FROM service WHERE service_id = $1);

SELECT EXISTS(SELECT 1 FROM orders WHERE customer_id = $1);

DELETE FROM service WHERE service_id =$1;

SELECT EXISTS(SELECT 1 FROM customer WHERE customer_id=$1;

SELECT EXISTS(SELECT 1 FROM orders WHERE order_id=$1;

INSERT INTO orders (
	order_id, 
	customer_id, 
	order_date, 
	completion_date, 
	received_by, 
	created_at, 
	updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7);

SELECT EXISTS(SELECT 1 FROM orders WHERE order_id=$1);

UPDATE orders SET 
	completion_date=$1, 
	updated_at=NOW()
	WHERE order_id=$2;

SELECT * FROM orders;

SELECT * FROM orders WHERE order_id = $1;