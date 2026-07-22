package main

import "fmt"

//enumerated types

// type OrderStatus int

// const (
// 	Recieved OrderStatus = iota
// 	Confirmed
// 	Prepared
// 	Delivered
// )

type OrderStatus string

const (
	Recieved  OrderStatus = "recieved"
	Confirmed             = "confirmed"
	Prepared              = "prepared"
	Delivered             = "dilevered"
)

func changingOrderStatus(status OrderStatus) {
	fmt.Println("Changing  order status to", status)
}

func main() {
	changingOrderStatus(Prepared)
}
