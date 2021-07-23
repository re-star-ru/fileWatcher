package usecase

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"golang.org/x/text/encoding/charmap"

	"fileWatcher/orders"
)

const OrdersDirName = "Orders"

type OrderUsecase struct {
	path string
}

func New(path string) orders.Usecase {
	go watcher(path)

	return &OrderUsecase{path: path}
}

func (ou OrderUsecase) NewOrder(o orders.Order) {
	// create file
	// todo: add TTL 1 day

	standPath := filepath.Join(ou.path, OrdersDirName, o.Stand)

	// preapare directories
	if err := os.MkdirAll(standPath, 0777); err != nil {
		log.Println(err)
		return
	}

	////write file todo: with encoding win-1251
	encoder := charmap.Windows1251.NewEncoder()
	encFileData, err := encoder.String(fmt.Sprintf("%s\n%s", o.Car, o.Unit))
	if err != nil {
		log.Println(err)
		return
	}

	// ASG_113.local name of stand, todo get nameStand from 1c
	if err := os.WriteFile(filepath.Join(standPath, "ASG_113.local", o.Name+".txt"), []byte(encFileData), 0777); err != nil {
		log.Println(err)
		return
	}

	log.Println("new order writed to filesystem successfull", o.Name)
}

func newJobHandler(w http.ResponseWriter, r *http.Request) {

}
