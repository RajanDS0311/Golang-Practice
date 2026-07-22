// Interface bas check karta hai ki kisi type ke paas required methods hain ya nahi.
package main

import "fmt"

type paymenter interface {
	pay(amount float32)
	refund(amount float32, account string)
}

type payment struct {
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
	// razorpayGw := razorpay{}
	// razorpayGw.pay(amount)

	// stripeGw.pay(100)

	p.gateway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("making payments using razorpay", amount)
}

// type stripe struct{}

// func (s stripe) pay(amount float32) {
// 	fmt.Println("making payments using stripe", amount)
// }

type fakepayment struct{}

func (f fakepayment) pay(amount float32) {
	fmt.Println("Making payment using fake gateway for testing purpose")
}

type paypal struct{}

func (p paypal) pay(amount float32) {
	fmt.Println("Making payment using paypal", amount)
}

func (p paypal) refund(amount float32, account string) {
	fmt.Println("")
}

func main() {

	// stripeGw := stripe{}
	// razorpayGw := razorpay{}
	// fakeGw := fakepayment{}
	paypalGw := paypal{}
	newpayment := payment{
		gateway: paypalGw,
	}
	newpayment.makePayment(100)
}
