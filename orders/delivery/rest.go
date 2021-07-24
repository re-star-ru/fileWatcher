package delivery

import (
	"fileWatcher/orders"
	"github.com/pkg/errors"
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
		err = errors.Wrap(err, "wrong post data")

		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := oh.orderUcase.NewOrder(o); err != nil {
		err = errors.Wrap(err, "err while make new order")

		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
