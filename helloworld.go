package main

import (
	"fmt"
	"log"
)

func init() {
	log.SetPrefix("INFO: ")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	var test string = "HelloWorld\n"
	fmt.Print(test)
	fmt.Print("HelloWorld \n")

	log.Print("Hello World !")
	// Print + os.Exit = Fatal
	log.Fatal("Hello World !")
	// log.Fatalln("Hello World !")

}
