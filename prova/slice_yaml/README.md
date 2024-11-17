# Obiettivo generale
Il programma simula un sistema dove dei dispositivi mandano dati presei da vari sensori a intervalli randomici

# Configurazione (config.yaml)
Il file contiene:
- il numero di dispositivi da simulare
- un intervallo temporale (in secondi) per l'invio dei dati per ogni device
- una lista di sensori, ognuno con un nome, tipo e, opzionalmente, valore minimo e valore massimo

# Funzioni
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