// Copyright (c) 2016, Socca Systems -- All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/soccasys/builder"
	"flag"
	"fmt"
	"log"
	"net/http"
)

var project *builder.Project

func main() {
	file := flag.String("load", "", "Load the project definition from a file")
	//root := flag.String("root", "/tmp/builder", "Directory where the builds are run")
	port := flag.Int("port", 8080, "HTTP port to use for the server")
	flag.Parse()
	if !flag.Parsed() {
		log.Fatal("Incorrect command line parameters")
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
