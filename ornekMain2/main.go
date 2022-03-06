package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Time struct {
	Updated    string
	UpdatedISO string
	Updateduk  string
}
type Cru struct {
	Code       string
	Symbol     string
	Rate       string
	Rate_Float float64
}
type Bpi struct {
	USD Cru
	EUR Cru
	GBP Cru
}

type ApiRes struct {
	Time      Time
	ChartName string
	Bpi       Bpi
}
type MQ struct {
	gorm.Model
	Data string
}

var giris string

func server() {

	resp, _ := http.Get("https://api.coindesk.com/v1/bpi/currentprice.json")
	data, _ := ioutil.ReadAll(resp.Body)
	var api ApiRes
	json.Unmarshal(data, &api)
	json.Marshal(api.Bpi)
	giris = api.Bpi.USD.Rate + api.Bpi.USD.Symbol

}
func database() {
	db, _ := gorm.Open(sqlite.Open("veri.db"), &gorm.Config{})
	db.AutoMigrate(&MQ{})
	db.Create(&MQ{Data: giris})
}

var handMessage mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())

	fmt.Printf("MSG: %s\n", msg.Payload())
	fmt.Println("---------------------------------------------------------------------")
}

func main() {
	server()
	database()

	opts := mqtt.NewClientOptions().AddBroker("18.197.171.34:1883")

	opts.SetKeepAlive(60 * time.Second)
	// Mesaj geri arama işleyicisini ayarlayın
	opts.SetDefaultPublishHandler(handMessage)
	opts.SetPingTimeout(1 * time.Second)

	c := mqtt.NewClient(opts)
	token := c.Connect()
	token.Wait()

	// mesajları almak için abone ol
	token = c.Subscribe("Cihaz", 0, nil)
	token.Wait()

	// mesajları yayınla
	token = c.Publish("API", 0, false, giris)
	token.Wait()

	time.Sleep(100 * time.Second)

	// Aboneliği iptal et
	token = c.Unsubscribe("Cihaz")
	token.Wait()

	// bağlantıyı kes
	c.Disconnect(250)
	time.Sleep(2 * time.Second)

}
