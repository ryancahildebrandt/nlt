// -*- coding: utf-8 -*-

// Created on Fri Jul 12 08:40:32 PM EDT 2024
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"fmt"
	"testing"
)

func reference_tables() (t1, t2, t3 TableData) {
	// t1 0, 0
	t1 = TableData{
		cells: []DataValue{
			{0, 0, []string{}, []string{}, "_"},
			{1, 0, []string{}, []string{}, "col1"},
			{2, 0, []string{}, []string{}, "col2"},
			{3, 0, []string{}, []string{}, "col3"},
			{0, 1, []string{}, []string{}, "row1"},
			{1, 1, []string{}, []string{}, "val11"},
			{2, 1, []string{}, []string{}, "val12"},
			{3, 1, []string{}, []string{}, "val13"},
			{0, 2, []string{}, []string{}, "row2"},
			{1, 2, []string{}, []string{}, "val21"},
			{2, 2, []string{}, []string{}, "val22"},
			{3, 2, []string{}, []string{}, "val23"},
			{0, 3, []string{}, []string{}, "row3"},
			{1, 3, []string{}, []string{}, "val31"},
			{2, 3, []string{}, []string{}, "val32"},
			{3, 3, []string{}, []string{}, "val33"},
			{0, 4, []string{}, []string{}, "row4"},
			{1, 4, []string{}, []string{}, "val41"},
			{2, 4, []string{}, []string{}, "val42"},
			{3, 4, []string{}, []string{}, "val43"},
		},
		rows: [][]string{
			{},
			{},
			{},
			{},
			{},
		},
		columns: [][]string{
			{},
			{},
			{},
			{},
			{},
		},
		x_dim: 4,
		y_dim: 4,
	}

	// t2 0, 1
	t2 = TableData{
		cells: []DataValue{
			{0, 0, []string{"_"}, []string{}, "_"},
			{1, 0, []string{"_"}, []string{}, "col1"},
			{2, 0, []string{"_"}, []string{}, "col2"},
			{3, 0, []string{"_"}, []string{}, "col3"},
			{0, 1, []string{"row1"}, []string{}, "row1"},
			{1, 1, []string{"row1"}, []string{}, "val11"},
			{2, 1, []string{"row1"}, []string{}, "val12"},
			{3, 1, []string{"row1"}, []string{}, "val13"},
			{0, 2, []string{"row2"}, []string{}, "row2"},
			{1, 2, []string{"row2"}, []string{}, "val21"},
			{2, 2, []string{"row2"}, []string{}, "val22"},
			{3, 2, []string{"row2"}, []string{}, "val23"},
			{0, 3, []string{"row3"}, []string{}, "row3"},
			{1, 3, []string{"row3"}, []string{}, "val31"},
			{2, 3, []string{"row3"}, []string{}, "val32"},
			{3, 3, []string{"row3"}, []string{}, "val33"},
			{0, 4, []string{"row4"}, []string{}, "row4"},
			{1, 4, []string{"row4"}, []string{}, "val41"},
			{2, 4, []string{"row4"}, []string{}, "val42"},
			{3, 4, []string{"row4"}, []string{}, "val43"},
		},
		rows: [][]string{
			{"_"},
			{"row1"},
			{"row2"},
			{"row3"},
			{"row4"},
		},
		columns: [][]string{
			{},
			{},
			{},
			{},
			{},
		},
		x_dim: 4,
		y_dim: 4,
	}

	// t3 5, 5
	t3 = TableData{
		cells: []DataValue{
			{0, 0, []string{"_"}, []string{"_"}, "_"},
			{1, 0, []string{"_", "col1"}, []string{"col1"}, "col1"},
			{2, 0, []string{"_", "col1", "col2"}, []string{"col2"}, "col2"},
			{3, 0, []string{"_", "col1", "col2", "col3"}, []string{"col3"}, "col3"},
			{0, 1, []string{"row1"}, []string{"_", "row1"}, "row1"},
			{1, 1, []string{"row1", "val11"}, []string{"col1", "val11"}, "val11"},
			{2, 1, []string{"row1", "val11", "val12"}, []string{"col2", "val12"}, "val12"},
			{3, 1, []string{"row1", "val11", "val12", "val13"}, []string{"col3", "val13"}, "val13"},
			{0, 2, []string{"row2"}, []string{"_", "row1", "row2"}, "row2"},
			{1, 2, []string{"row2", "val21"}, []string{"col1", "val11", "val21"}, "val21"},
			{2, 2, []string{"row2", "val21", "val22"}, []string{"col2", "val12", "val22"}, "val22"},
			{3, 2, []string{"row2", "val21", "val22", "val23"}, []string{"col3", "val13", "val23"}, "val23"},
			{0, 3, []string{"row3"}, []string{"_", "row1", "row2", "row3"}, "row3"},
			{1, 3, []string{"row3", "val31"}, []string{"col1", "val11", "val21", "val31"}, "val31"},
			{2, 3, []string{"row3", "val31", "val32"}, []string{"col2", "val12", "val22", "val32"}, "val32"},
			{3, 3, []string{"row3", "val31", "val32", "val33"}, []string{"col3", "val13", "val23", "val33"}, "val33"},
			{0, 4, []string{"row4"}, []string{"_", "row1", "row2", "row3", "row4"}, "row4"},
			{1, 4, []string{"row4", "val41"}, []string{"col1", "val11", "val21", "val31", "val41"}, "val41"},
			{2, 4, []string{"row4", "val41", "val42"}, []string{"col2", "val12", "val22", "val32", "val42"}, "val42"},
			{3, 4, []string{"row4", "val41", "val42", "val43"}, []string{"col3", "val13", "val23", "val33", "val43"}, "val43"},
		},
		rows: [][]string{
			{"_", "col1", "col2", "col3"},
			{"row1", "val11", "val12", "val13"},
			{"row2", "val21", "val22", "val23"},
			{"row3", "val31", "val32", "val33"},
			{"row4", "val41", "val42", "val43"},
		},
		columns: [][]string{
			{"_", "row1", "row2", "row3", "row4"},
			{"col1", "val11", "val21", "val31", "val41"},
			{"col2", "val12", "val22", "val32", "val42"},
			{"col3", "val13", "val23", "val33", "val43"},
			{},
		},
		x_dim: 4,
		y_dim: 4,
	}

	return t1, t2, t3
}

func TestJoinHeaders(t *testing.T) {
	table := []struct {
		input DataValue
		delim string
		exp1  string
		exp2  string
	}{
		{DataValue{1, 2, []string{"abc", "def"}, []string{"123", "456"}, "value"}, ";", "abc;def", "123;456"},
		{DataValue{0, 0, []string{"", "def"}, []string{"123", ""}, "value"}, "__", "__def", "123__"},
		{DataValue{-2, -1, []string{"", ""}, []string{"", ""}, ""}, "", "", ""},
	}

	for _, test := range table {
		res1, res2 := test.input.JoinHeaders(test.delim)
		if res1 != test.exp1 {
			t.Errorf("JoinHeaders(%v) = %v, expected %v", test.input, res1, test.exp1)
		}
		if res2 != test.exp2 {
			t.Errorf("JoinHeaders(%v) = %v, expected %v", test.input, res2, test.exp2)
		}
	}
}

func TestNewTableData(t *testing.T) {
	df, _, _ := reference_dataframes()
	t1, t2, t3 := reference_tables()
	table := [3]struct {
		x   int
		y   int
		exp TableData
	}{
		{0, 0, t1},
		{0, 1, t2},
		{5, 5, t3},
	}
	for _, test := range table {
		res := NewTableData(df, test.x, test.y)
		if fmt.Sprint(res) != fmt.Sprint(test.exp) {
			t.Errorf("NewTableData(df, %v, %v) = %v, expected %v", test.x, test.y, res, test.exp)
		}
	}
}
