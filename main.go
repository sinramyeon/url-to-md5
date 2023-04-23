package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const DefaultGoroutines = 10

func main() {
	var urls []string
	var startIndex int
	numOfGoRoutines := DefaultGoroutines

	// input ex : 1. /md5hasher test.com
	//            2. /md5hasher -parallel 3 test.com google.com facebook.com yahoo.com
	inputArgs := os.Args[1:]

	if len(inputArgs) == 0 {
		fmt.Println("not enough arguments to call md5 hasher")
	}

	if inputArgs[0] == "-parallel" {
		num, err := strconv.Atoi(inputArgs[1])
		if err != nil {
			fmt.Sprintln("failed to parse input commands %w", err)
		}
		numOfGoRoutines = num
		startIndex = 2
	}

	for i := startIndex; i < len(inputArgs); i++ {
		if !(strings.HasPrefix(inputArgs[i], "http://") || strings.HasPrefix(inputArgs[i], "https://")) {
			inputArgs[i] = "http://" + inputArgs[i]
		}
		urls = append(urls, inputArgs[i])
	}

	done := make(chan struct{})
	result, fail := MdHasherHandler(urls, numOfGoRoutines, done)
	defer close(result)

	for {
		select {
		//example : http://google.com 8ff1c478ccca08cca025b028f68b352f
		case resp := <-result:
			fmt.Println(resp.Url, fmt.Sprintf("%x", resp.MD5))
		case failedResponse := <-fail:
			fmt.Println("md5 parsing failed : %w", failedResponse)
		case <-done:
			return
		}
	}
}
