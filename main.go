// Copyright (c) 2016, Socca Systems -- All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/soccasys/automat"
	"encoding/json"
	"syscall"
	"flag"
	"fmt"
	"log"
	"net/http"
)

var server *automat.Server

func printExampleProject() {
	p := automat.NewProject()
        p.Name = "build-automat"
	p.AddComponent("src/github.com/soccasys/automat", "https://github.com/soccasys/automat.git", "master")
	p.AddComponent("src/github.com/soccasys/automat-server", "https://github.com/soccasys/automat-server.git", "master")
	p.AddBuildStep("Build All", ".", []string{"go", "install", "github.com/soccasys/automat-server"})
        text, _ := json.MarshalIndent(p, "", "    ")
	fmt.Printf("%s\n", text)
}

func main() {
	example := flag.Bool("example",  false, "Print an example of project definition on the standatd output, and exit")
	root := flag.String("serve", "", "Directory where the Automat server is stored")
	port := flag.Int("port", 8080, "HTTP port to use for the server")
	flag.Parse()
	if !flag.Parsed() {
		log.Fatal("Incorrect command line parameters")
	}
	if *example {
		printExampleProject()
		syscall.Exit(0)
        }
	if *root != "" {
		server = automat.NewServer(*root)
		http.Handle("/", server)
		err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
		if err != nil {
			log.Fatalf("ListenAndServe failed %s\n", err)
		}
	}
}
