package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
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
func main() {
	server()
	db, _ := gorm.Open(sqlite.Open("veri.db"), &gorm.Config{})
	db.AutoMigrate(&MQ{})
	db.Create(&MQ{Data: giris})

	mqtt.DEBUG = log.New(os.Stdout, "", 0)
	mqtt.ERROR = log.New(os.Stdout, "", 0)
	opts := mqtt.NewClientOptions().AddBroker("3.67.253.250:1883").SetClientID("deneme1")

	opts.SetKeepAlive(60 * time.Second)
	// Mesaj geri arama işleyicisini ayarlayın
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(1 * time.Second)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// mesajları almak için abone ol
	if token := c.Subscribe("deneme1", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	// mesajları yayınla
	token := c.Publish("deneme1", 0, false, giris)
	token.Wait()

	time.Sleep(6 * time.Second)

	// Aboneliği iptal et
	if token := c.Unsubscribe("deneme1"); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	// bağlantıyı kes
	c.Disconnect(250)
	time.Sleep(1 * time.Second)

}
