// Interface bas check karta hai ki kisi type ke paas required methods hain ya nahi.
package main

import "fmt"

type payment struct {
	gateway razorpay
}

func (p payment) makePayment(amount float32) {
	// razorpayGw := razorpay{}
	// razorpayGw.pay(amount)

	// stripeGw.pay(100)

	p.gateway.pay(100)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("making payments using razorpay", amount)
}

type stripe struct{}

func (s stripe) pay(amount float32) {
	fmt.Println("making payments using stripe", amount)
}
func main() {

	// stripeGw := stripe{}
	razorpayGw := razorpay{}
	newpayment := payment{
		gateway: razorpayGw,
	}
	newpayment.makePayment(100)
}
