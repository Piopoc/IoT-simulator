# Obiettivo generale
Il programma simula un sistema in cui vengono inviati dei dati da dei device composti da vari sensori al cloud (BigQuery) tramite l'appoggio dei servizi di Pub/Sub e Cloud Functions per la raccolta dei messaggi e la loro manipolazione. Lo scopo è quello di gestire e archiviare i dati generati ottimizzando le prestazioni e riducendo il più possibile i costi

## File di Configurazione (.env)
Il file contiene:
- il numero di dispositivi da simulare
- un intervallo temporale (in secondi) per l'invio dei dati per ogni device
- 3 sensori ognugno con id, name, tipo e, opzionalmente, valori massimi e minimi


## Funzioni 
- readConf()
    - legge il file di configurazione
    - il contenuto del file di configurazione e viene creato un oggetto 
    - restituisce l'oggetto Config
- start(ch chan<- Field, d Device, shuffleSensors []Sensor)
    - la funzione simula l'invio dei dati relativi ai vari sensori per un singolo dispositivo
    - viene utilizzato un canale per l'invio dei dati 
    - i dati relativi ad ogni sensore vengono generati randomicamente e si basano sul tipo e sui eventuali valori massimi e minimi
    - si attende un valore random (entro il valore descritto nel file di configurazione)
    - viene mandato nel canale l'oggetto Field che contiene l'ID del device, il timestamp e i dati generati dai sensori
- start_up()
    - la funzione gestisce l'intera simulazione
    - legge la configurazione e crea un canale per ricevere i Field creati nella funzione precedente
    - viene eseguita una goroutine per ricevere continuamente dati dal canale
    - tra il lancio di una goroutine e l'altra viene aggiunto un ritardo randomico
    - si aspetta che tutti i dispositivi finiscano di inviare i dati prima di passare all'iterazione successiva
- id_gen(n int)
    - la funzione genera una stringa di identificazione di 4 cifre per un singolo device
- stringify(field Field)
    - i dati vengono stampati a terminale utilizzando una stringa formattata in stile JSON
- publishData(projectID, topicID, msg string)
    - la funzione si occupa di convertire il messaggio sottoforma di stringa in un array di byte
    - pubblica il messaggio serialiazzato in un topic e prima di proseguire mi assicuro che la pubblicazione del dato non dia errori
    - riempio un file di log con il riassunto dei dati che sono stati pubblicati correttamente nel topic

# Da chiedere
- è problema togliere accento alle parole