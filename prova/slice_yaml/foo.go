package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"

	"gopkg.in/yaml.v3"
)

type Device struct {
	ID string
}
type Sensor struct {
	Name     string `yaml:"name"`
	Type     string `yaml:"type"`
	MaxValue *any   `yaml:"max_value,omitempty"`
	MinValue *any   `yaml:"min_value,omitempty"`
}
type Config struct {
	NumDevices    int      `yaml:"num_devices"`
	SendEverySecs int      `yaml:"send_every_secs"`
	Sensors       []Sensor `yaml:"sensors"`
}
type Field_Data struct {
	Name  string
	Value any
}
type Field struct {
	Dev  Device
	Time int64
	Data []Field_Data
}

// leggo file di configurazione e lo restituisco
func readConf() (config Config) {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("Errore nella lettura del file: %v", err)
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Errore nel parsing YAML: %v", err)
	}
	return config
}

// generazione random dei valori dei sensori per ogni device
func start(ch chan<- Field, d Device, shuffleSensors []Sensor) {
	config := readConf()
	wait := time.Duration(rand.Intn(config.SendEverySecs)) * time.Second

	campo := Field{
		Dev:  d,
		Data: []Field_Data{},
	}
	for _, sensor := range shuffleSensors {
		var value any
		switch sensor.Name {
		case "temperatura":
			value = rand.Intn(46) - 15
		case "umidità":
			value = rand.Float64()*0.833 + 0.05
		case "pressione":
			value = rand.Float64()*1.667 + 0.5
		case "dig1":
			value = rand.Intn(2)
		}
		campo.Data = append(campo.Data, Field_Data{Name: sensor.Name, Value: value})
	}
	time.Sleep(wait)
	campo.Time = time.Now().Unix()
	ch <- campo
}

// simulatore
func start_up() {
	config := readConf()
	s := config.Sensors
	dim := config.NumDevices

	devices := []Device{}
	for i := 0; i < dim; i++ {
		devices = append(devices, Device{
			ID: id_gen(i),
		})
	}
	dataChannel := make(chan Field)
	var wg sync.WaitGroup

	go func() {
		for field := range dataChannel {
			fmt.Print(stringifyData(field))
			wg.Done()
		}
	}()

	for {
		var shuffleSensors []Sensor
		for _, rand_I := range rand.Perm(len(s)) {
			shuffleSensors = append(shuffleSensors, s[rand_I])
		}
		for i, device := range devices {
			wg.Add(1)
			go start(dataChannel, device, shuffleSensors)
			if i < len(devices)-1 {
				time.Sleep(time.Duration(rand.Intn(6)) * time.Second)
			}
		}
		wg.Wait()
		fmt.Println()
	}
}

// generatore ID per ogni device
func id_gen(n int) string {
	return fmt.Sprintf("%04d", n)
}

// stampa a terminale dei dati estratti con formattazione simil JSON
func stringifyData(field Field) string {
	file, err := os.OpenFile("sum.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatalf("Errore nell'aprire il file: %v", err)
    }
	defer file.Close()
	p := fmt.Sprintf("{\"device\": \"%v\", \"timestamp\": %d,", field.Dev.ID, field.Time)
	for i, sensorData := range field.Data {
		if i < len(field.Data)-1 {
			switch v := sensorData.Value.(type) {
			case int:
				p += fmt.Sprintf("  \"%v\": %d,", sensorData.Name, v)
			case float64:
				p += fmt.Sprintf("  \"%v\": %.2f,", sensorData.Name, v)
			}
		} else {
			switch v := sensorData.Value.(type) {
			case int:
				p += fmt.Sprintf("  \"%v\": %d", sensorData.Name, v)
			case float64:
				p += fmt.Sprintf("  \"%v\": %.2f", sensorData.Name, v)
			}
		}
	}
	p += fmt.Sprintln("}")
	_,err=file.WriteString(p)
	if err!=nil{
		log.Fatalf("Errore nella scrittura del file, err: %v",err)
	}
	return p
}
