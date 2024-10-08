package main

import "fmt"

type BaseProcessor struct{}

func (bp BaseProcessor) SharedMethod() {
	fmt.Println("This is a shared method")
}

type PaymentProcessor interface {
	ProcessPayment(amount float64) error
	GetPaymentStatus(paymentID string) (interface{}, error)
}

// Billogram implementation
type Billogram struct {
	BaseProcessor
}

func (b Billogram) ProcessPayment(amount float64) error {
	b.SharedMethod()

	fmt.Printf("Processing Billogram payment of %.2f\n", amount)
	return nil
}

func (b Billogram) GetPaymentStatus(paymentID string) (interface{}, error) {
	fmt.Printf("Getting Billogram payment status for %s\n", paymentID)
	return map[string]string{"status": "Completed"}, nil
}

// Stripe implementation
type Stripe struct {
	BaseProcessor
}

func (s Stripe) ProcessPayment(amount float64) error {
	s.SharedMethod()
	fmt.Printf("Processing Stripe payment of %.2f\n", amount)
	return nil
}

func (s Stripe) GetPaymentStatus(paymentID string) (interface{}, error) {
	fmt.Printf("Getting Stripe payment status for %s\n", paymentID)
	return "Completed", nil
}

func main() {
	var processor PaymentProcessor

	// Use Billogram
	processor = Billogram{}
	processor.ProcessPayment(100.0)
	status, _ := processor.GetPaymentStatus("12345")
	fmt.Println("Billogram Payment Status:", status)

	// Use Stripe
	processor = Stripe{}
	processor.ProcessPayment(200.0)
	status, _ = processor.GetPaymentStatus("67890")
	fmt.Println("Stripe Payment Status:", status)
}
