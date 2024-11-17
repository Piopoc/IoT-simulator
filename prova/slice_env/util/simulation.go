package util

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

var (
	projectID   = "m31-academy"
	topicID     = "data"
	LogFileName = "./util/last_pub.log"
)

type Sensor struct {
	Name     string
	Type     string
	MaxValue *any
	MinValue *any
}
type Config struct {
	NumDevices    int
	SendEverySecs int
	Sensors       []Sensor
}
type Field_Data struct {
	Name  string
	Value any
}
type Field struct {
	ID_device   string
	Time_device int64
	Data_value  []Field_Data
}

// generazione random dei valori dei sensori per ogni device
func start(ch chan<- Field, d string, shuffleSensors []Sensor, wg *sync.WaitGroup) {
	defer wg.Done()
	wait := time.Duration(rand.Intn(readConf().SendEverySecs)) * time.Second
	campo := Field{
		ID_device:  d,
		Data_value: []Field_Data{},
	}
	for _, sensor := range shuffleSensors {
		var value any
		switch sensor.Type {
		case "integer":
			max := (*sensor.MaxValue).(int)
			min := (*sensor.MinValue).(int)
			value = rand.Intn(max-min+1) + min
		case "float":
			max := (*sensor.MaxValue).(float64)
			min := (*sensor.MinValue).(float64)
			value = rand.Float64()*(max-min) + min
		case "digital":
			value = rand.Intn(2)
		}
		campo.Data_value = append(campo.Data_value, Field_Data{Name: sensor.Name, Value: value})
	}
	time.Sleep(wait)
	campo.Time_device = time.Now().Unix()
	ch <- campo
}

// simulatore
func Start_up() {
	config := readConf()
	s := config.Sensors
	dim := config.NumDevices

	devices := []string{}
	for i := 0; i < dim; i++ {
		devices = append(devices, id_gen(i))
	}
	dataChannel := make(chan Field)
	var wg sync.WaitGroup
	go func() {
		for field := range dataChannel {
			p, _ := stringifyData(field)
			err := PublishData(projectID, topicID, p)
			if err != nil {
				log.Printf("Errore durante la pubblicazione del messaggio: %v", err)
			}
			fmt.Print(p)
		}
	}()
	for {
		var shuffleSensors []Sensor
		for _, rand_I := range rand.Perm(len(s)) {
			shuffleSensors = append(shuffleSensors, s[rand_I])
		}
		for i, device := range devices {
			wg.Add(1)
			go start(dataChannel, device, shuffleSensors, &wg)
			if i < len(devices)-1 {
				time.Sleep(time.Duration(rand.Intn(6)) * time.Second)
			}
		}
		wg.Wait()
		// fmt.Println()
	}
}
