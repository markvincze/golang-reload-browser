package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	log.Println("Starting reload server.")

	startReloadServer()

	log.Println("Reload server started.")
	log.Println("Press Enter to reload the browser!")
	for {
		reader := bufio.NewReader(os.Stdin)
		reader.ReadString('\n')

		log.Println("Reloading browser.")
		sendReload()
	}
}
