package delivery

import (
	"fileWatcher/orders"
	"log"
	"net/http"
)

type OrdersHandler struct {
	orderUcase orders.Usecase
}

func New(usecase orders.Usecase) *OrdersHandler {
	return &OrdersHandler{orderUcase: usecase}
}

func (oh *OrdersHandler) NewOrder(w http.ResponseWriter, r *http.Request) {
	o := orders.Order{}
	if err := o.FromJson(r.Body); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	oh.orderUcase.NewOrder(o)
}
