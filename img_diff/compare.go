package main

import (
	"fmt"
	"crypto/sha256"
	"io/ioutil"
	"log"
)

func main() {
	fichier1, err := ioutil.ReadFile("image_1.jpg")
	if err != nil {
		log.Fatal(err)
	}
	fichier2, err := ioutil.ReadFile("image_2.jpg")
	if err != nil {
		log.Fatal(err)
	}
	fichier3, err := ioutil.ReadFile("image_3.jpg")
	if err != nil {
		log.Fatal(err)
	}
	sum1 := sha256.Sum256(fichier1)
	sum2 := sha256.Sum256(fichier2)
	sum3 := sha256.Sum256(fichier3)
	
	if sum1==sum2{
		fmt.Printf("L'image differente est la 3\n")
	} else if sum1==sum3{
		fmt.Printf("L'image differente est la 2\n")
	} else {
		fmt.Printf("L'image differente est la 1\n")
	}
			
}
