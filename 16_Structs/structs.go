package main

import (
	"fmt"
	"time"
)

// we can use struct for OOPs
// to reate struct
// data inside a struct is called fields
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time //nanosecond precision
}

// go do not have anything like Constructor, but if we want then we need to use this hack
// i.e... use a function

// func newOrder(id string, amount float32, status string) *order {
// 	myorder := order{
// 		id:     id,
// 		amount: amount,
// 		status: status,
// 	}

// 	return &myorder
// }

// attaching functions to struct
// reciever type
// only use * / pointer if we are modifying the struct value , and not when only returning or getting them
// func (o *order) chnageStatus(status string) {
// 	o.status = status
// }

// func (o order) getAmount() float32 {
// 	return o.amount
// }

func main() {

	//if struct is going to use for once only then we can do this inline way.
	language := struct {
		name   string
		isGood bool
	}{"Golang", true}

	fmt.Println(language)

	// myorder := newOrder("1", 100, "recieved") // constructor wala
	// fmt.Println(myorder.amount)

	// to instantiate struct without constructor
	// myorder := order{
	// 	id:     "1",
	// 	amount: 500.00,
	// 	status: "recieved",
	// }

	// if we dont set the value of any field then default value is ZERO value
	// int and Float -> 0, string -> " ", boolean -> false

	// myorder.chnageStatus("pending")
	// fmt.Println(myorder.getAmount())

	// myorder2 := order{
	// 	id:        "2",
	// 	amount:    200,
	// 	status:    "dilevered",
	// 	createdAt: time.Now(),
	// }

	// myorder.createdAt = time.Now()
	// fmt.Println(myorder)
	// fmt.Println(myorder.amount) // to get any field

	// myorder.status = "confirmed" // all instances are seperate from each other, no effect on anothe if something is modified in one.
	// fmt.Println(myorder)
	// fmt.Println(myorder2)

}
