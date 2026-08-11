package main

import "fmt"

type paymenter interface{

	pay(amount float32)
}

type payment struct{

	gateway paymenter
}

func (p payment) makePayment(amount float32){

	// razorpayPaymentGw := razorpay{}
	// razorpayPaymentGw.pay(amount)

	p.gateway.pay(amount)
	
}

type razorpay struct{}

func (r razorpay) pay(amount float32){

	fmt.Println("making payment using razorpay", amount)
}

type stripe struct{}

func (s stripe) pay(amount float32){

	fmt.Println("makeing payment using stripe", amount)
}

func main(){

	// razorpayPaymentGw := razorpay{}
	stripeGw:= stripe{}

	newPayment := payment{
		gateway: stripeGw,
	}

	newPayment.makePayment(100)

}

