package services

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Peyton232/dual-go-reciever/database"
	model "github.com/Peyton232/dual-go-reciever/models"
)

func HandleWebhook(w http.ResponseWriter, r *http.Request) {

	var db = database.Connect()
	b, err := io.ReadAll(r.Body)
	// b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(b))

	t := model.Data{}

	_ = json.Unmarshal(b, &t)

	log.Println(t)

	mod := db.SendData(&model.NewSpaceData{
		Imei:             &t.Imei,
		Serial:           &t.Serial,
		Momsn:            &t.Momsn,
		TransmitTime:     &t.TransmitTime,
		IridiumLatitude:  &t.IridiumLatitude,
		IridiumLongitude: &t.IridiumLongitude,
		IridiumCep:       &t.IridiumCep,
		Data:             &t.Data,
	})

	if mod != nil {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

}
