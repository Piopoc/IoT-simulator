package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

type ChartData struct {
	ID_device     string `json:"ID_device"`
	Time_creation string `json:"Time_creation"`
	Data_value    string `json:"Data_value"`
}

type DatetimeValues struct {
	InitDatetime  string `json:"initDatetime"`
	FinalDatetime string `json:"finalDatetime"`
}

var startDate string = "2024-10-02T09:27:21.301391"
var endDate string = "2024-10-25T19:42:32.894262"

func getDataFromBigQuery(w http.ResponseWriter, r *http.Request) {
	getDataFromBigQuerySupport(w, r, startDate, endDate)
}

func getDataFromBigQuerySupport(w http.ResponseWriter, r *http.Request, initDate, finalDate string) {
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, "m31-academy")
	if err != nil {
		log.Printf("Errore nel connettersi a BigQuery: %v", err)
		http.Error(w, "Errore nel connettersi a BigQuery", http.StatusInternalServerError)
		return
	}
	defer client.Close()
	var deviceID string
	log.Println(r.URL.Path)

	if strings.Contains(r.URL.Path, "/update") {
		deviceID = strings.TrimSuffix(r.URL.Path[len("/data/"):], "/update")
	} else {
		deviceID = r.URL.Path[len("/data/"):]
	}

	query := client.Query(`
		SELECT FORMAT_DATETIME('%Y-%m-%dT%H:%M:%S', Time_creation) AS Time_creation, ID_device, Data_value 
		FROM test_fcorradi.DataTable
		WHERE ID_device = @deviceID AND Time_creation >= @initDate AND Time_creation <= @finalDate
		ORDER BY Time_creation ASC LIMIT 100
	`)

	query.Parameters = []bigquery.QueryParameter{
		{
			Name:  "deviceID",
			Value: deviceID,
		},
		{
			Name:  "initDate",
			Value: initDate,
		},
		{
			Name:  "finalDate",
			Value: finalDate,
		},
	}

	log.Printf("Esecuzione query con deviceID: %s, initDate: %s, finalDate: %s", deviceID, initDate, finalDate)
	it, err := query.Read(ctx)
	if err != nil {
		log.Printf("Errore nell'esecuzione della query: %v", err)
		http.Error(w, "Errore nell'esecuzione della query", http.StatusInternalServerError)
		return
	}

	var results []ChartData
	for {
		var row ChartData
		err := it.Next(&row)
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("Errore durante l'iterazione dei risultati: %v", err)
			http.Error(w, "Errore durante l'iterazione dei risultati", http.StatusInternalServerError)
			return
		}
		results = append(results, row)
	}

	log.Printf("Errore nell'esecuzione della query: %v; Parametri: initDate=%s, finalDate=%s", err, initDate, finalDate)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func main() {
	http.Handle("/0000", http.StripPrefix("/0000", http.FileServer(http.Dir("./static"))))
	http.Handle("/0001", http.StripPrefix("/0001", http.FileServer(http.Dir("./static"))))
	http.Handle("/0002", http.StripPrefix("/0002", http.FileServer(http.Dir("./static"))))
	http.Handle("/0003", http.StripPrefix("/0003", http.FileServer(http.Dir("./static"))))
	http.Handle("/0004", http.StripPrefix("/0004", http.FileServer(http.Dir("./static"))))

	http.HandleFunc("/data/", getDataFromBigQuery)

	port := "8080"
	fmt.Printf("Server is running on http://localhost:%s/0000\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
