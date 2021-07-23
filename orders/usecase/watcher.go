package usecase

import (
	"bytes"
	"encoding/json"
	"github.com/rjeczalik/notify"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const ReportsDir = "Reports"

func getFileType(typeW string) string {
	switch typeW {
	case "I":
		return "До ремонта"
	case "W":
		return "В ремонте"
	case "O":
		return "После ремонта"
	default:
		return "statusNotFound"
	}
}

func split(r rune) bool {
	return r == '_' || r == '.'
}

// watch for file changes
func watcher(workDir string) {
	c := make(chan notify.EventInfo, 1)
	path := filepath.Join(workDir, ReportsDir)

	// /... is recursive
	if err := notify.Watch(path+"/...", c, notify.Write); err != nil {
		log.Fatal(err)
	}
	defer notify.Stop(c)

	for {
		ei := <-c
		log.Println("Got event:", ei.Event(), ei.Path())

		fPath := ei.Path()
		fileName := filepath.Base(fPath)
		a := strings.FieldsFunc(fileName, split)

		log.Println("splitted fileName:", a)
		//splitted fileName: [УТ113429 O 39 jpeg]
		if len(a) == 4 {
			order := a[0]
			fileType := getFileType(a[1])
			go addTo1c(order, fileType, fileName, fPath)
		}
	}
}

const host = "https://1c.re-star.ru/sm1/hs" // todo: extract host from there!
func addTo1c(order, fileType, fileName, fPath string) {
	time.Sleep(time.Millisecond * 200) // sleep for wat hmm?

	f, err := os.ReadFile(fPath) // todo: retry?
	if err != nil {
		log.Println(err)
		return
	}

	addFileReq := struct {
		Order string `json:"order"`
		Type  string `json:"type"`
		Photo []byte `json:"photo"`
	}{
		order,
		fileType,
		f,
	}

	b, err := json.Marshal(addFileReq)
	if err != nil {
		log.Println(err)
		return
	}

	reqBody := bytes.NewBuffer(b)
	req, err := http.NewRequest("POST", host+"/stand/1234/add", reqBody)
	if err != nil {
		log.Println(err)
		return
	}

	req.SetBasicAuth("API", "6O7EHDWS0Sk$yZ%i80p5") // todo: extract passsword from there!!!
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	// если оффер не найден то ошибка
	// если оффер записан не правильно то другая ошибка
	// если сервер не доступен то еще одна ошибка другая

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	if err := resp.Body.Close(); err != nil {
		log.Println(err)
		return
	}

	log.Println("order", order, "status code", resp.StatusCode, "body", string(body))
}
