package orders

type Usecase interface {
	NewOrder(order Order)
}
