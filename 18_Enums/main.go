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
	Confirmed OrderStatus = "confirmed"
	Prepared  OrderStatus = "prepared"
	Delivered OrderStatus = "dilevered"
)

func changingOrderStatus(status OrderStatus) {
	fmt.Println("Changing  order status to", status)
}

func main() {
	changingOrderStatus(Prepared)
}
