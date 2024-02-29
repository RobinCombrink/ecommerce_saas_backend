// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package database

import (
	"database/sql"
)

type Customer struct {
	ID    int64
	Name  string
	Email string
}

type Order struct {
	ID         int64
	Paid       int64
	Customerid int64
	Total      float64
}

type OrderItem struct {
	ID        int64
	Orderid   int64
	Productid int64
}

type Product struct {
	ID          int64
	Name        string
	Description sql.NullString
	Price       float64
}
