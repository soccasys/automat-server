// Copyright (c) 2016, Socca Systems -- All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/soccasys/builder"
	"encoding/json"
	"syscall"
	"flag"
	"fmt"
	"log"
	"net/http"
)

var project *builder.Project

func printExampleProject() {
	p := builder.NewProject()
        p.Name = "build-automat"
	p.AddComponent("src/github.com/builder", "https://github.com/soccasys/build-automat.git", "master")
	p.AddComponent("src/github.com/build-automat", "https://github.com/soccasys/builder.git", "master")
	p.AddBuildStep("Build All", ".", []string{"go", "install", "github.com/soccasys/build-automat"})
        text, _ := json.MarshalIndent(p, "", "    ")
	fmt.Printf("%s\n", text)
}

func main() {
	example := flag.Bool("example",  false, "Print an example of project definition on the standatd output, and exit")
	file := flag.String("load", "", "Load the project definition from a file")
	//root := flag.String("root", "/tmp/builder", "Directory where the builds are run")
	port := flag.Int("port", 8080, "HTTP port to use for the server")
	flag.Parse()
	if !flag.Parsed() {
		log.Fatal("Incorrect command line parameters")
	}
	if *example {
		printExampleProject()
		syscall.Exit(0)
        }
	project = builder.NewProject()
	if *file != "" {
		project.Load(*file)
	}
	http.Handle("/project/", project)
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		log.Fatalf("ListenAndServe failed %s\n", err)
	}
}
