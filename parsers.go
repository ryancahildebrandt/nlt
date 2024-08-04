// -*- coding: utf-8 -*-

// Created on Sat Jun  1 07:49:59 PM EDT 2024
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

// Reads in a file and returns a bytes reader, for flexibility in different file formats
func ReaderFromFile(p string) *bytes.Reader {
	content, err := os.ReadFile(p)
	if err != nil {
		log.Fatal(err)
	}
	return bytes.NewReader(content)
}

// Replaces target html tags with specified replacement
func swapHtmlTags(d *goquery.Document, old string, new string) {
	d.Find(old).Each(func(i int, s *goquery.Selection) {
		for _, node := range s.Nodes {
			node.Data = new
		}
	})
}

// Replaces each target html element with its children elements
func replaceWithChildren(d *goquery.Document, t string) {
	d.Find(t).Each(func(i int, s *goquery.Selection) {
		s.ReplaceWithNodes(s.Children().Nodes...)
	})
}

// Removes th, thead, and tbody elements from html tables for more predictable inputs
func StandardizeTables(doc *goquery.Document) {
	swapHtmlTags(doc, "th", "td")
	replaceWithChildren(doc, "thead")
	replaceWithChildren(doc, "tbody")
}

type FileParser interface {
	parse() dataframe.DataFrame
}

type CSVParser struct {
	path string
}

// Reads CSV file into dataframe
func (p *CSVParser) parse() dataframe.DataFrame {
	return dataframe.ReadCSV(ReaderFromFile(p.path), dataframe.DetectTypes(false), dataframe.DefaultType(series.String))
}

type TSVParser struct {
	path string
}

// Reads TSV file into dataframe
func (p *TSVParser) parse() dataframe.DataFrame {
	return dataframe.ReadCSV(ReaderFromFile(p.path), dataframe.WithDelimiter('\t'), dataframe.DetectTypes(false), dataframe.DefaultType(series.String))
}

type JSONLinesParser struct {
	path string
}

// Reads JSONL file into dataframe
func (p *JSONLinesParser) parse() dataframe.DataFrame {
	dec := json.NewDecoder(ReaderFromFile(p.path))
	jsonl := []map[string]interface{}{}
	for {
		res := map[string]interface{}{}
		err := dec.Decode(&res)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		jsonl = append(jsonl, res)
	}
	return dataframe.LoadMaps(jsonl, dataframe.DetectTypes(false), dataframe.DefaultType(series.String))
}

type JSONArrObjParser struct {
	path string
}

// Reads an array of objects JSON file into dataframe
func (p *JSONArrObjParser) parse() dataframe.DataFrame {
	return dataframe.ReadJSON(ReaderFromFile(p.path), dataframe.DetectTypes(false), dataframe.DefaultType(series.String))
}

type JSONArrArrParser struct {
	path string
}

// Reads an array of arrays JSON file into dataframe
func (p *JSONArrArrParser) parse() dataframe.DataFrame {
	dec := json.NewDecoder(ReaderFromFile(p.path))
	records := [][]string{}
	err := dec.Decode(&records)
	if err != nil {
		log.Fatal(err)
	}
	return dataframe.LoadRecords(records, dataframe.DetectTypes(false), dataframe.DefaultType(series.String))
}

type MDParser struct {
	path string
}

// Reads MD file into dataframe
func (p *MDParser) parse() dataframe.DataFrame {
	md, err := io.ReadAll(ReaderFromFile(p.path))
	if err != nil {
		log.Fatal(err)
	}
	html := markdown.Render(parser.New().Parse(md), html.NewRenderer(html.RendererOptions{}))
	reader := strings.NewReader(string(html))
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Fatal(err)
	}

	StandardizeTables(doc)
	out, err := doc.Html()
	if err != nil {
		log.Fatal(err)
	}
	return dataframe.ReadHTML(strings.NewReader(out), dataframe.DetectTypes(false), dataframe.DefaultType(series.String))[0]
}

type HTMLParser struct {
	path string
}

// Reads HTML file into dataframe
func (p *HTMLParser) parse() dataframe.DataFrame {
	doc, err := goquery.NewDocumentFromReader(ReaderFromFile(p.path))
	if err != nil {
		log.Fatal(err)
	}

	StandardizeTables(doc)
	out, err := doc.Html()
	if err != nil {
		log.Fatal(err)
	}
	return dataframe.ReadHTML(strings.NewReader(out), dataframe.DetectTypes(false), dataframe.DefaultType(series.String))[0]
}
