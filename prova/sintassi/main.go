package main

import (
	"fmt"
	"strings"
)

func main() {
	msg := " { \"ID_device\": \"0002\", \"Time_device\": 1727428981, \"temperatura\": 10, \"pressione\": 0.63, \"umidita\": 0.37, \"dig1\": 0 }"

	p := strings.Split(msg, " ")
	fmt.Println(strings.Join(p[:6], " ") + " \"Data_value\": { " + strings.Join(p[6:], " ") + " }")
}

/*
package util

import (
	"context"
  "encoding/json"
	"fmt"
  "log"
  "strings"
  "os"
	"time"
	"cloud.google.com/go/bigquery"
)

type PubSubMessage struct {
	Data []byte `json:"data"`
}
type DataRow struct {
	ID_device  string `bigquery:"ID_device"`
	TimeDevice string `bigquery:"Time_device"`
	DataValue  string `bigquery:"Data_value"`
}

// Funzione che prende la stringa nel topic e poi la formatta correttamente per aggiungerla su BigQuery
func PubSubToBigQuery(ctx context.Context, m PubSubMessage) error {
	var rawData map[string]interface{}
	sms := []byte(messageFormatter(string(m.Data)))
	if err := json.Unmarshal(sms, &rawData); err != nil {
		return fmt.Errorf("errore nella decodifica del messaggio: %v", err)
	}
	data := &DataRow{
		ID_device:  rawData["ID_device"].(string),
		TimeDevice: time.Unix(int64(rawData["Time_device"].(float64)), 0).Format("2006-01-02 15:04:05"),
		DataValue:  jsonFormatter(rawData),
	}
	if err := insertToBigQuery(ctx, *data); err != nil {
		return fmt.Errorf("errore nell'inserimento in BigQuery: %v", err)
	}
	return nil
}

// Funzione per inserire i dati in BigQuery
func insertToBigQuery(ctx context.Context, data DataRow) error {
	projectID := os.Getenv("PROJECT_ID")
	datasetName := os.Getenv("DATASET_NAME")
	tableName := os.Getenv("TABLE_NAME")

	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("bigquery.NewClient: %v", err)
	}
	defer client.Close()
    log.Printf("creating inserter with dataset %s - table name %s:\n",datasetName,tableName)
	inserter := client.Dataset(datasetName).Table(tableName).Inserter()
    log.Printf("inserting: %v\n",data)
	if err := inserter.Put(ctx, data); err != nil {
		return fmt.Errorf("inserter.Put: %v", err)
	}
	return nil
}

// Formattazione del messaggio per convertirlo poi in dato inseribile in bigquery
func messageFormatter(msg string) string {
	p := strings.Split(msg, " ")
	return strings.Join(p[:6], " ") + " \"Data_value\": { " + strings.Join(p[6:], " ") + " }"
}

// Funzione che formatta il contenuto di Data_value restituendo la stringa json
func jsonFormatter(m map[string]interface{}) (json string) {
	i := 0
	for k, v := range m["Data_value"].(map[string]interface{}) {
		if i != len(m) {
			json += fmt.Sprintf("\"%v\": %v, ", k, fmt.Sprint(v))
		} else {
			json += fmt.Sprintf("\"%v\": %v", k, fmt.Sprint(v))
		}
		i++
	}
	fmt.Println("JSON formatter: { ", json, " }")
	return "{ " + json + " }"
}
*/
