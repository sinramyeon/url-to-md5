package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

type Results struct {
	Url string
	MD5 [md5.Size]byte
}

func MdHasher(url string) (*Results, error) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("failed to get url : ", err)
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("failed to read url : ", err)
		return nil, err
	}

	md5Result := md5.Sum(body)

	return &Results{url, md5Result}, nil
}

func MdHasherHandler(urls []string, numOfGoroutines int, done chan struct{}) (chan *Results, chan string) {
	length := len(urls)

	wg := sync.WaitGroup{}
	wg.Add(length)

	processed := make(chan string, length)
	failed := make(chan string)
	result := make(chan *Results, length)

	for g := 0; g < numOfGoroutines; g++ {
		go func() {
			for {
				select {
				case url := <-processed:
					res, err := MdHasher(url)
					if err != nil {
						failed <- err.Error()
					} else {
						result <- res
					}
					wg.Done()
				case <-done:
					return
				}
			}
		}()
	}

	for _, url := range urls {
		processed <- url
	}

	go func() {
		wg.Wait()
		close(done)
	}()

	return result, failed
}
