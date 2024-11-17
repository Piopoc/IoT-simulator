package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"IoT.io/slice_env/util"
)

func main() {
	fmt.Println("Tempo in cui viene lanciato il sistema: ", time.Now().Unix())
	err := os.WriteFile(util.LogFileName, []byte{}, 0644)
	if err != nil {
		log.Fatalf("Errore nell'azzerare il file: %v", err)
	}
	util.Start_up()
}
