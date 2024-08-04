// -*- coding: utf-8 -*-

// Created on Thu Jul 11 08:28:52 PM EDT 2024
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"encoding/json"
	"fmt"
)

// Handles user inputs passed to TableFormatters
type FormatFields struct {
	// Delimter between headers for each row/column, for when there are multiple cells constituting the header
	Delim string `json:"delim,omitempty"`
	// Clause linking x/y/val labels to the rest of the sentence or another label. Can be used simialrly to preamble in some cases
	Link string `json:"link,omitempty"`
	// Statement of equality between labels and values
	Eq string `json:"eq,omitempty"`
	// Preamble clause setting the context for the relationship between labels and values. Can be used simialrly to link in some cases
	Pre string `json:"pre,omitempty"`
	// Semantic/category label for cell values
	ValLabel string `json:"val_label,omitempty"`
	// Semantic/category label for a given row
	XLabel string `json:"x_label,omitempty"`
	// Semantic/category label for a given column
	YLabel string `json:"y_label,omitempty"`
}

// Handles user input file paths and table parsing behavior settings
type ConfigFields struct {
	// The number of columns that should be counted as the header for each row
	NRowHeaders int `json:"row_headers"`
	// The number of rows that should be counted as the header for each column
	NColHeaders int `json:"col_headers"`
	// File containing tabular data to read in
	InFile string `json:"infile"`
	// text file to save reformatted data
	OutFile string `json:"outfile"`
	// TableFormatter to use when reformatting tabular data
	Formatter string `json:"formatter"`
	// FileParser to use when reading from InFile, corresponding to file format and structure
	Parser string `json:"parser"`
}

// Reads config.json at specified path into ConfigFields struct
func ReadConfig(p string) (ConfigFields, error) {
	var config ConfigFields
	dec := json.NewDecoder(ReaderFromFile(p))
	err := dec.Decode(&config)
	return config, err
}

// Reads config.json at specified path into FormatFields struct
func ReadFields(p string) (FormatFields, error) {
	var fields FormatFields
	dec := json.NewDecoder(ReaderFromFile(p))
	err := dec.Decode(&fields)
	return fields, err
}

// Returns a populated a TableFormatter based on the provided formatter name and TableData struct
func SetFormatter(t TableData, f string) TableFormatter {
	switch f {
	case "UnnamedCoordFormatter1":
		return &UnnamedCoordFormatter1{t}
	case "UnnamedCoordFormatter2":
		return &UnnamedCoordFormatter2{t}
	case "NamedCoordFormatter1":
		return &NamedCoordFormatter1{t}
	case "NamedCoordFormatter2":
		return &NamedCoordFormatter2{t}
	case "NamedRowFormatter":
		return &NamedRowFormatter{t}
	case "NamedColFormatter":
		return &NamedColFormatter{t}
	case "UnnamedRowKeyValFormatter":
		return &UnnamedRowKeyValFormatter{t}
	case "UnnamedColKeyValFormatter":
		return &UnnamedColKeyValFormatter{t}
	case "NamedRowKeyValFormatter":
		return &NamedRowKeyValFormatter{t}
	case "NamedColKeyValFormatter":
		return &NamedColKeyValFormatter{t}
	case "RowValFormatter":
		return &RowValFormatter{t}
	case "ColValFormatter":
		return &ColValFormatter{t}
	default:
		fmt.Println("Invalid formatter provided, defaulting to UnnamedCoordFormatter1")
		return &UnnamedCoordFormatter1{t}
	}
}

// Returns a populated FileParser based on the provided parser name and file path
func SetParser(p string, f string) FileParser {
	switch f {
	case "CSV":
		return &CSVParser{p}
	case "TSV":
		return &TSVParser{p}
	case "JSONLines":
		return &JSONLinesParser{p}
	case "JSONArrObj":
		return &JSONArrObjParser{p}
	case "JSONArrArr":
		return &JSONArrArrParser{p}
	case "MD":
		return &MDParser{p}
	case "HTML":
		return &HTMLParser{p}
	default:
		fmt.Println("Invalid parser provided, defaulting to CSVParser")
		return &CSVParser{p}
	}
}
