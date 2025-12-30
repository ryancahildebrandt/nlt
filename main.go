// -*- coding: utf-8 -*-

// Created on Sat Jun  1 07:49:59 PM EDT 2024
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

/*
nlt reformats tabular data to natural language
The basic executable can handle c/tsv, html, md, and json/l formats and outputs to plain text
Inputs are provided by either 1) config.json in the current directory, or 2) a user specified file given with -c flag

Usage:

	nlt [flags] [path]

Flags:

	-c
		Path to user specified config.json file
	-l
		If nlt should run using config from lastrun.json
*/
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli"
)

// Saves the reformatted output slice to the specified path
func writeOutput(s []string, p string) error {
	bytes := []byte(strings.Join(s, "\n"))
	err := os.WriteFile(p, bytes, 0o644)
	return err
}

// Saves a copy of the current config to lastrun.json
func writeLastrun(c ConfigFields, f FormatFields) error {
	fields := struct {
		ConfigFields
		FormatFields
	}{c, f}
	bytes, err := json.Marshal(fields)
	if err != nil {
		return err
	}
	err = os.WriteFile("lastrun.json", bytes, 0o644)
	return err
}

func main() {
	var configPath string
	var lastrun bool

	app := &cli.App{
		Name:  "NLT",
		Usage: "Reformat tabular data into natural language",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "c",
				Value:       "./config.json",
				Usage:       "Location of config.json file",
				Destination: &configPath,
			},
			&cli.BoolFlag{
				Name:        "l",
				Usage:       "Use lastrun.json for the current run",
				Destination: &lastrun,
			},
		},
		Action: func(*cli.Context) {
			if lastrun {
				fmt.Println("lastrun true")
				configPath = "./lastrun.json"
				fmt.Printf("lastrunpath %s, configpath %s", "./lastrun.json", configPath)
			}

			fmt.Printf("Reading from config file at %s \n", configPath)

			config, err := ReadConfig(configPath)
			if err != nil {
				log.Fatalf("Unable to load config fields\nError: %v", err)
			}
			fmt.Printf("Config fields read as:\n%#v\n", config)

			fields, err := ReadFields(configPath)
			if err != nil {
				log.Fatalf("Unable to load formatter fields\nError: %v", err)
			}
			fmt.Printf("Formatter fields read as:\n%#v\n", fields)

			parser := SetParser(config.InFile, config.Parser)
			df := parser.parse()
			table := NewTableData(df, 0, 0)
			fmt.Printf("Table read from %s \n", config.InFile)
			fmt.Printf("Table:\n%v\n", df)

			formatter := SetFormatter(table, config.Formatter)
			out := formatter.format(fields)
			fmt.Printf("Table reformatted to natural language using %v\n", config.Formatter)
			fmt.Printf("Output:\n%v\n", strings.Join(out, "\n"))

			err = writeOutput(out, config.OutFile)
			if err != nil {
				log.Fatalf("Unable to save output file\nError: %v", err)
			}
			fmt.Printf("Output written to %v\n", config.OutFile)

			err = writeLastrun(config, fields)
			if err != nil {
				log.Fatalf("Unable to save lastrun file\nError: %v", err)
			}
		},
	}
	app.Run(os.Args)
}
