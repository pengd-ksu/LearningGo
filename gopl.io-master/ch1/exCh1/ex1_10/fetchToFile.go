// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 17.
//!+

// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

const total = 2

func main() {
	start := time.Now()
	ch := make(chan string)
	if len(os.Args) != 2 {
		fmt.Println("Require one argument.")
		return
	}
	for count := 0; count < total; count++ {
		url := os.Args[1]
		go fetch(url, ch, strconv.Itoa(count)) // start a goroutine
		fmt.Println(<-ch)                      // receive from channel ch
		fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	}
}

func fetch(url string, ch chan<- string, fileName string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	fileRoute := "./" + fileName + ".txt"
	file, err := os.Create(fileRoute)
	if err != nil {
		ch <- fmt.Sprintf("while creating file %s: %v", fileRoute, err)
	}
	nbytes, err := io.Copy(file, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

//!-
