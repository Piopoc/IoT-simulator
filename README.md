
# IoT Data Viewer & Simulator

![Go Version](https://img.shields.io/badge/Go-1.20%2B-blue)
![License](https://img.shields.io/badge/license-MIT-green)
![Status](https://img.shields.io/badge/status-active-brightgreen)

Questo progetto è una **web app sviluppata in Go** che permette di:
- visualizzare i dati storici dei dispositivi IoT memorizzati su **Google BigQuery**
- servire un’interfaccia HTML statica per la visualizzazione dei grafici
- eseguire query dinamiche sui dati in base a un `deviceID` e un intervallo temporale

> Il progetto è pensato per supportare scenari di simulazione e visualizzazione di dati IoT in ambienti Cloud.

---

## ⚙️ Architettura del progetto

La struttura del progetto è la seguente:

```
progetto/
├── dataViewer/
│   ├── main.go          # Codice Go che gestisce le API REST e serve i file statici
│   ├── go.mod           # Modulo Go
│   ├── go.sum
│   └── static/
│       └── index.html   # Pagina HTML con grafici e logica frontend
├── README.md
└── .git/                # Repository Git
```

---

## 🛰 Schema architetturale

```
┌──────────────┐      HTTP GET /data/{ID_device}
│  Browser     │ ────────────────────────────────▶ ┌────────────┐
│ (index.html) │                                    │  Server Go │
└──────┬───────┘      JSON response ◀──────────────┤ (main.go)  │
       │                                            └────┬───────┘
       │ Serve static HTML                               │
       │───────────────────────────────────────────────▶│
                                                         │
                                                  Query BigQuery
                                                         │
                                                         ▼
                                                ┌─────────────────────────┐
                                                │ Google BigQuery         │
                                                │ test_fcorradi.DataTable │
                                                └─────────────────────────┘
```

---

## 💻 Tecnologie utilizzate

- **Go** (linguaggio backend)
- **Google BigQuery** (database per la storicizzazione dei dati)
- **HTML / JavaScript** (frontend statico)
- **Cloud SDK (client BigQuery)** per l'interazione con BigQuery
- **HTTP server integrato** (senza framework esterni)

---

## 🚀 Come eseguire il progetto

### Prerequisiti
- Go >= 1.18
- Credenziali Google Cloud configurate (eseguire `gcloud auth application-default login`)

### Avvio
```bash
git clone https://github.com/Piopoc/IoT-simulator.git
cd IoT-simulator/dataViewer
go run main.go
```

Il server partirà sulla porta `8080`. Accedi a:
```
http://localhost:8080/0000
http://localhost:8080/0001
http://localhost:8080/0002
http://localhost:8080/0003
http://localhost:8080/0004
```

---

## 🌐 API e routing

### Interfacce frontend
```
/0000, /0001, /0002, /0003, /0004
```
Servono la pagina `index.html`.

### API REST
```
/data/{ID_device}
/data/{ID_device}/update
```
Restituiscono un JSON con i dati BigQuery relativi al dispositivo.

---

## 🔍 Dettagli tecnici

- Parsing dell'URL per estrarre `ID_device`
- Query parametrica su BigQuery:
  ```sql
  SELECT FORMAT_DATETIME('%Y-%m-%dT%H:%M:%S', Time_creation) AS Time_creation, ID_device, Data_value 
  FROM test_fcorradi.DataTable
  WHERE ID_device = @deviceID AND Time_creation >= @initDate AND Time_creation <= @finalDate
  ORDER BY Time_creation ASC
  LIMIT 100
  ```
- `initDate` e `finalDate` hardcoded nel codice:
  ```
  2024-10-02T09:27:21.301391 - 2024-10-25T19:42:32.894262
  ```
- Restituisce un array JSON:
  ```json
  [
    {
      "ID_device": "0000",
      "Time_creation": "2024-10-03T10:00:00",
      "Data_value": "{...}"
    },
    ...
  ]
  ```

---

## 📈 Possibili estensioni

- Consentire filtro dinamico via query string (es. `/data/0000?init=...&final=...`)
- Autenticazione / autorizzazione alle API
- Logging avanzato e metriche (es. Prometheus, Grafana)
- Frontend dinamico (React, Vue)
- Deployment su **Cloud Run** o simili

---

## 👨‍💻 Autore

Progetto sviluppato da **Filippo Corradi**  
🎯 Contesto: simulazione e visualizzazione dati IoT in ambiente Cloud per studio e sperimentazione.

---

## 📌 Note

⚠️ **Attenzione:**  
Il progetto richiede accesso a BigQuery e configurazione delle credenziali.  
La query è limitata a **100 record** per ottimizzazione.

---

## 📝 License

MIT License
