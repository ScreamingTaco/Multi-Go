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

// Project TODOS
// TODO tone down comments & make them more meaningful
// TODO improve 'Scrape'
// TODO finish email task
// TODO finish decompress (and review compress)
// TODO add 'tshark -r [file path]' task
// TODO add network scanner
// TODO add wifi password cracker

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/akamensky/argparse"
	ct "github.com/daviddengcn/go-colortext"

	"github.com/TheRedSpy15/Multi-Go/tasks"
	"github.com/TheRedSpy15/Multi-Go/utils"
)

func main() {
	dialogMode := false

	parser := argparse.NewParser("Multi-Go", "Runs multiple security orientated tasks")

	// Create flags
	t := parser.String("t", "Task", &argparse.Options{Required: false, Help: "Task to run"})
	r := parser.String("r", "Target", &argparse.Options{Required: false, Help: "Target to run task on"})

	err := parser.Parse(os.Args) // parse arguments
	utils.CheckErr(err)

	reader := bufio.NewReader(os.Stdin) // make reader object

	if *t == "" { // enter dialog mode
		dialogMode = true
		utils.PrintBanner()
		tasks.List()
	} else {
		ct.Foreground(ct.Yellow, false)
	}

	//Only continue execution in dialog mode
	for contExec := true; contExec; contExec = dialogMode {
		if dialogMode {
			fmt.Print("\nEnter task to run: ")
			choice, _ := reader.ReadString('\n')     // get choice
			choice = strings.TrimRight(choice, "\n") // trim choice so it can be check against properly

			if strings.Contains(choice, "-r") { // check for optional target
				inputs := strings.Split(choice, " -r ") // separate task & target
				*t = inputs[0]
				*r = inputs[1]
			} else { // no optional target
				*t = choice
			}
		}

		// Determine task to run
		switch *t {
		case "Hash":
			fmt.Println("\nRunning task:", *t, "\nTarget:", *r)
			tasks.HashFile(*r)
		case "pwnAccount":
			fmt.Println("\nRunning task:", *t, "\nTarget:", *r)
			tasks.PwnAccount(*r)
		case "encryptFile":
			fmt.Println("\nRunning task:", *t, "\nTarget:", *r)
			tasks.EncryptFile(*r)
		case "decryptFile":
			fmt.Println("\nRunning task:", *t, "\nTarget:", *r)
			tasks.DecryptFile(*r)
		case "Scrape":
			fmt.Println("\nRunning task:", *t, "\nTarget:", *r)
			tasks.Scrape(*r)
		case "DOS":
			fmt.Println("\nRunning task:", *t, "\nTarget:", *r)
			tasks.Dos(*r, nil)
		case "compress":
			fmt.Println("\nRunning task:", *t, "\nTarget:", *r)
			tasks.Compress(*r)
		case "decompress":
			fmt.Println("\nRunning task:", *t, "\nTarget:", *r)
			tasks.Decompress(*r)
		case "Firewall":
			fmt.Println("\nRunning task:", *t, "\nTarget:", *r)
			tasks.ToggleFirewall(*r)
		case "generatePassword":
			fmt.Println("\nRunning task:", *t, "\nTarget:", *r)
			tasks.GeneratePassword(*r)
		case "Install":
			fmt.Println("\nRunning task:", *t, "\nTarget:", *r)
			tasks.Install(*r)
		case "Bleach":
			fmt.Println("\nRunning task:", *t, "\nTarget:", *r)
			tasks.Bleach(*r)
		case "cyberNews":
			tasks.News()
		case "systemInfo":
			tasks.SystemInfo()
		case "Clean":
			tasks.Clean()
		case "Email":
			tasks.Email()
		case "Audit":
			tasks.Audit()
		case "About":
			tasks.About()
		case "List":
			tasks.List()
		case "Exit":
			os.Exit(0)
		default: // invalid
			ct.Foreground(ct.Red, true)
			fmt.Println("Invalid task -", *t)
			ct.Foreground(ct.Yellow, false)
			fmt.Println("Use '--help' or '-t List'")
		}
	}
}
