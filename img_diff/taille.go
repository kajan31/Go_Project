package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	fichier, err := ioutil.ReadFile("image_1.jpg")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Taille du fichier: %d \n", len(fichier),)

}
