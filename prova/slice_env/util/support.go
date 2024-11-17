package util

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// leggo file di configurazione
func readConf() (config Config) {
	err := godotenv.Load("./util/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	num, _ := strconv.Atoi(os.Getenv("NUM_DEVICES"))
	send, _ := strconv.Atoi(os.Getenv("SEND_EVERY_SECS"))

	var sensors []Sensor
	for i := 0; ; i++ {
		name := os.Getenv(fmt.Sprintf("SENSORS_%d_NAME", i))
		if name == "" {
			break
		}
		sensorType := os.Getenv(fmt.Sprintf("SENSORS_%d_TYPE", i))
		var maxValue, minValue any

		switch sensorType {
		case "integer":
			maxValue, _ = strconv.Atoi(os.Getenv(fmt.Sprintf("SENSORS_%d_MAX_VALUE", i)))
			minValue, _ = strconv.Atoi(os.Getenv(fmt.Sprintf("SENSORS_%d_MIN_VALUE", i)))
		case "float":
			maxValue, _ = strconv.ParseFloat(os.Getenv(fmt.Sprintf("SENSORS_%d_MAX_VALUE", i)), 64)
			minValue, _ = strconv.ParseFloat(os.Getenv(fmt.Sprintf("SENSORS_%d_MIN_VALUE", i)), 64)
		case "digital":
			maxValue = nil
			minValue = nil
		}
		sensors = append(sensors, Sensor{
			Name:     name,
			Type:     sensorType,
			MaxValue: &maxValue,
			MinValue: &minValue,
		})
	}
	config = Config{
		NumDevices:    num,
		SendEverySecs: send,
		Sensors:       sensors,
	}
	return config
}

// generatore ID per ogni device
func id_gen(n int) string {
	return fmt.Sprintf("%04d", n)
}

// stampa a terminale dei dati estratti con formattazione simil JSON
func stringifyData(field Field) (string, string) {
	var q string
	p := fmt.Sprintf("{ \"ID_device\": \"%v\", \"Time_device\": %d,", field.ID_device, field.Time_device)
	for i, sensorData := range field.Data_value {
		if i < len(field.Data_value)-1 {
			switch v := sensorData.Value.(type) {
			case int:
				q += fmt.Sprintf("  \"%v\": %d,", sensorData.Name, v)
			case float64:
				q += fmt.Sprintf("  \"%v\": %.2f,", sensorData.Name, v)
			}
		} else {
			switch v := sensorData.Value.(type) {
			case int:
				q += fmt.Sprintf("  \"%v\": %d", sensorData.Name, v)
			case float64:
				q += fmt.Sprintf("  \"%v\": %.2f", sensorData.Name, v)
			}
		}
	}
	p += q + fmt.Sprintln(" }")
	dataJson := "{" + q + "}"
	return p, dataJson
}