package main

import (
	"fmt"
	"os"
	"strconv"
)

const DefaultGoroutines = 10

func main() {
    var urls []string
    // input ex : 1. /myhttp test.com
    //            2. /myhttp -parallel 3 test.com google.com facebook.com yahoo.com
    inputArgs := os.Args[1:]
  
    if (inputArgs[0] == "-parallel") {          
        func() {
            s, err := strconv.Atoi(inputArgs[1]);
            if err != nil {
				fmt.Println("failed to parse input commands %s", err)
			}

			fmt.Println(s)
        }()
    } 
}