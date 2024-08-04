// -*- coding: utf-8 -*-

// Created on Sun Jul 21 04:00:15 PM EDT 2024
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"fmt"
	"testing"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func reference_dataframes() (df1, df2, df3 dataframe.DataFrame) {
	df1 = dataframe.LoadRecords(
		[][]string{
			{"_", "col1", "col2", "col3"},
			{"row1", "val11", "val12", "val13"},
			{"row2", "val21", "val22", "val23"},
			{"row3", "val31", "val32", "val33"},
			{"row4", "val41", "val42", "val43"},
		}, dataframe.DetectTypes(false), dataframe.DefaultType(series.String))

	df2 = dataframe.LoadRecords(
		[][]string{
			{"", "", "", ""},
			{"", "", "", ""},
			{"", "", "", ""},
			{"", "", "", ""},
			{"", "", "", ""},
		}, dataframe.DetectTypes(false), dataframe.DefaultType(series.String))

	df3 = dataframe.LoadRecords(
		[][]string{
			{"0", "0", "0", "0"},
			{"0", "0", "0", "0"},
			{"0", "0", "0", "0"},
			{"0", "0", "0", "0"},
			{"0", "0", "0", "0"},
		}, dataframe.DetectTypes(false), dataframe.DefaultType(series.String))

	return df1, df2, df3
}

func TestTSVParser(t *testing.T) {
	df1, df2, df3 := reference_dataframes()
	table := [3]struct {
		f   FileParser
		exp dataframe.DataFrame
	}{
		{&TSVParser{"data/test1.tsv"}, df1},
		{&TSVParser{"data/test2.tsv"}, df2},
		{&TSVParser{"data/test3.tsv"}, df3},
	}

	for _, test := range table {
		res := test.f.parse()
		if fmt.Sprint(res) != fmt.Sprint(test.exp) {
			t.Errorf("ReadConfig(%v) = %v, expected %v", test.f, res, test.exp)
		}
	}
}

func TestJSONLinesParser(t *testing.T) {
	df1, _, _ := reference_dataframes()
	table := []struct {
		f   FileParser
		exp dataframe.DataFrame
	}{
		{&JSONLinesParser{"data/test1.jsonl"}, df1},
	}

	for _, test := range table {
		res := test.f.parse()
		if fmt.Sprint(res) != fmt.Sprint(test.exp) {
			t.Errorf("ReadConfig(%v) = %v, expected %v", test.f, res, test.exp)
		}
	}
}

func TestJSONArrObjParser(t *testing.T) {
	df1, _, _ := reference_dataframes()
	table := []struct {
		f   FileParser
		exp dataframe.DataFrame
	}{
		{&JSONArrObjParser{"data/test_arr_obj1.json"}, df1},
	}

	for _, test := range table {
		res := test.f.parse()
		if fmt.Sprint(res) != fmt.Sprint(test.exp) {
			t.Errorf("ReadConfig(%v) = %v, expected %v", test.f, res, test.exp)
		}
	}
}

func TestJSONArrArrParser(t *testing.T) {
	df1, df2, _ := reference_dataframes()
	table := [2]struct {
		f   FileParser
		exp dataframe.DataFrame
	}{
		{&JSONArrArrParser{"data/test_arr_arr1.json"}, df1},
		{&JSONArrArrParser{"data/test_arr_arr2.json"}, df2},
	}

	for _, test := range table {
		res := test.f.parse()
		if fmt.Sprint(res) != fmt.Sprint(test.exp) {
			t.Errorf("ReadConfig(%v) = %v, expected %v", test.f, res, test.exp)
		}
	}
}

func TestMDParser(t *testing.T) {
	df1, df2, df3 := reference_dataframes()
	table := [3]struct {
		f   FileParser
		exp dataframe.DataFrame
	}{
		{&MDParser{"data/test1.md"}, df1},
		{&MDParser{"data/test2.md"}, df2},
		{&MDParser{"data/test3.md"}, df3},
	}

	for _, test := range table {
		res := test.f.parse()
		if fmt.Sprint(res) != fmt.Sprint(test.exp) {
			t.Errorf("ReadConfig(%v) = %v, expected %v", test.f, res, test.exp)
		}
	}
}

func TestHTMLParser(t *testing.T) {
	df1, df2, df3 := reference_dataframes()
	table := [3]struct {
		f   FileParser
		exp dataframe.DataFrame
	}{
		{&HTMLParser{"data/test1.html"}, df1},
		{&HTMLParser{"data/test2.html"}, df2},
		{&HTMLParser{"data/test3.html"}, df3},
	}

	for _, test := range table {
		res := test.f.parse()
		if fmt.Sprint(res) != fmt.Sprint(test.exp) {
			t.Errorf("ReadConfig(%v) = %v, expected %v", test.f, res, test.exp)
		}
	}
}
