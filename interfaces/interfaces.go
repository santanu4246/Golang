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

	// razorpayPaymentGw := razorpay{}
	// razorpayPaymentGw.pay(amount)

	p.gateway.pay(amount)

}

func (p payment) getRefund(amount float32, account string) {

	p.gateway.refund(amount, account)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {

	fmt.Println("making payment using razorpay", amount)
}

func (s razorpay) refund(amount float32, account string) {

	fmt.Println("refunding using stripe", amount, account)

}

type stripe struct{}

func (s stripe) pay(amount float32) {

	fmt.Println("makeing payment using stripe", amount)
}

func (s stripe) refund(amount float32, account string) {

	fmt.Println("refunding using stripe", amount, account)

}

func main() {

	// razorpayPaymentGw := razorpay{}
	stripeGw := stripe{}

	newPayment := payment{
		gateway: stripeGw,
	}

	newPayment.makePayment(100)
	newPayment.getRefund(50, "account123")
}
