package main

/*
Copyright 2018 TheRedSpy15

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
limitations under the License.
*/

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"syscall"

	"github.com/daviddengcn/go-colortext"
	"github.com/gocolly/colly"
	"golang.org/x/crypto/ssh/terminal"
)

// BUG: still runs scrape task
// check if target is empty, panic if is
func checkTarget(target string) {
	if target == "" {
		ct.Foreground(ct.Red, true)
		panic("target cannot be empty when performing this task!")
	}
}

// Util function - used for getting []byte of file
func readFileIntoByte(filename string) []byte {
	var data []byte                // specify type
	file, err := os.Open(filename) // make file object
	defer file.Close()
	if err != nil {
		panic(err.Error())
	} else {
		data, err = ioutil.ReadAll(file) // read all
		if err != nil {
			panic(err.Error())
		}
	}
	return data
}

// Util function - securely get password from user
func getPassword() string {
	print("Enter Password: ")
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin)) // run password command, make var with result
	password := string(bytePassword)                             // cast to string var

	return password
}

// Util function - displays banner text
func printBanner() {
	ct.Foreground(ct.Red, true)
	file, err := os.Open("banner.txt") // make file object
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file) // make scanner object to read file object
	for scanner.Scan() {              // internally, it advances token based on sperator
		fmt.Println(scanner.Text()) // token in unicode-char
	}

	ct.Foreground(ct.Yellow, false)
}

// TODO: scrape list of files from text file
// Util function - scrapes a website link
func scrape(site string) {
	c := colly.NewCollector() // make colly object
	c.IgnoreRobotsTxt = true

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) { // print error message on error
		ct.Foreground(ct.Red, true)
		log.Println("Something went wrong:", err)
		ct.ResetColor()
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
		println("Response:", r.StatusCode)
	})

	c.OnScraped(func(r *colly.Response) { // finished with site
		fmt.Println("Finished", r.Request.URL)

		err := r.Save(r.FileName()) // saving data

		if err != nil {
			ct.Foreground(ct.Red, true)
			panic("Error saving")
		} else {
			ct.Foreground(ct.Green, true)
			println("Saved - ", r.FileName())
			ct.ResetColor()
		}
	})

	c.Visit(site)
}
