package main

import (
	"fmt"
	"log"
	"os"
	"time"

	alert "github.com/dev-bittu/go-alert"
)

func main() {
	alert.Success("Programma in esecuzione...")
	fmt.Println("Istante di tempo in cui viene lanciato il sistema: ", time.Now().Unix())
	err := os.WriteFile("sum.log", []byte{}, 0644)
	if err != nil {
		log.Fatalf("Errore nell'azzerare il file: %v", err)
	}
	
	start_up()
}
