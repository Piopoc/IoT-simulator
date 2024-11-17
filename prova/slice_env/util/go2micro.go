package util

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"cloud.google.com/go/pubsub"
)

// Funzione che pubblica il messaggio in un topic e scrive un resoconto dei dati inviati in un file di log
func PublishData(projectID, topicID, msg string) error {
	file, err := os.OpenFile(LogFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Errore nell'aprire il file: %v", err)
	}
	defer file.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	credPath := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	if credPath == "" {
		log.Fatalf("Le credenziali non sono impostate, assicurati di avere GOOGLE_APPLICATION_CREDENTIALS configurato")
	} else {
		log.Printf("Credenziali caricate dal percorso: %s\n", credPath)
	}

	cli, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		fmt.Printf("pubsub.NewClient: %v\n", err)
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}
	defer cli.Close()

	topic := cli.Topic(topicID)
	pub := topic.Publish(ctx, &pubsub.Message{Data: []byte(msg)})
	_, err = pub.Get(ctx)
	if err != nil {
		log.Printf("pub.Get: %v",err)
		return fmt.Errorf("pub.Get: %v", err)
	}
	log_string := strings.Split(msg, ",")
	_, err = file.WriteString("Dato pubblicato: " + strings.Join(log_string[:2], ",") + " }\n")
	if err != nil {
		log.Fatalf("Errore nella scrittura del file, err: %v", err)
	}
	return nil
}
