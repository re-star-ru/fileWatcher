package usecase

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/text/encoding/charmap"

	"fileWatcher/orders"
)

type OrderUsecase struct {
	path string
}

func New() orders.Usecase {
	const defaultReportsDir = "/Sync/Reports/"
	if err := os.MkdirAll(defaultReportsDir, 0777); err != nil {
		log.Fatal(err) // exit from proggramm
	}

	go watcher(defaultReportsDir)

	const defaultOrdersDir = "/Sync/Orders/"
	return &OrderUsecase{path: defaultOrdersDir}
}

func (ou OrderUsecase) NewOrder(o orders.Order) error {
	// create file
	// todo: add TTL 1 day

	standPath := filepath.Join(ou.path, o.Stand)
	log.Println("path", standPath)

	// preapare directories
	if err := os.MkdirAll(standPath, 0777); err != nil {
		return err
	}

	////write file todo: with encoding win-1251
	encoder := charmap.Windows1251.NewEncoder()
	encFileData, err := encoder.String(fmt.Sprintf("%s\n%s", o.Car, o.Unit))
	if err != nil {
		return err
	}

	// ASG_113.local name of stand, todo get nameStand from 1c
	if err := os.WriteFile(filepath.Join(standPath, o.Name+".txt"), []byte(encFileData), 0777); err != nil {
		return err
	}

	log.Println("new order writed to filesystem successfull", o.Name)
	return nil
}
