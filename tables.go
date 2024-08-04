// -*- coding: utf-8 -*-

// Created on Sat Jun  1 07:49:59 PM EDT 2024
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"strings"

	"github.com/go-gota/gota/dataframe"
)

// Represents a single value in a data table
type DataValue struct {
	// Position on the x axis, or which column it belongs to
	x int
	// Position on the y axis, or which row it belongs to
	y int
	// First values on the x axis, defining the row header
	x_head []string
	// First values on the y axis, defining the column header
	y_head []string
	// Value of the cell at [x, y]
	val string
}

// Combines the provided number of headers into a single string for a DataValue
func (v DataValue) JoinHeaders(d string) (string, string) {
	return strings.Join(v.x_head, d), strings.Join(v.y_head, d)
}

// Represents all data in a table
type TableData struct {
	// All DataValues in the table
	cells []DataValue
	// All row header values in the table
	rows [][]string
	// All column header values in the table
	columns [][]string
	// Size of the x axis, or the number of columns
	x_dim int
	// Size of the y axis, or the number of rows
	y_dim int
}

// Pulls dimensions of dataframe to DataTable
func (t *TableData) populateDims(df dataframe.DataFrame) {
	t.y_dim, t.x_dim = df.Dims()
}

// Breaks table into individual cells and stores in TableData.cells
func (t *TableData) populateCells(df dataframe.DataFrame) {
	for y, row := range df.Records() {
		for x, cell := range row {
			t.cells = append(t.cells, DataValue{x: x, y: y, val: cell})
		}
	}
}

// Extracts first n cells from each row and stores in TableData.rows
func (t *TableData) populateRows(n int) {
	heads := make(map[int][]string)
	t.rows = make([][]string, t.y_dim+1)
	for i, cell := range t.cells {
		if cell.x < n {
			heads[cell.y] = append(heads[cell.y], cell.val)
		}
		t.cells[i].x_head = heads[cell.y]
		t.rows[cell.y] = heads[cell.y]
	}
}

// Extracts first n cells from each column and stores in TableData.columns
func (t *TableData) populateColumns(n int) {
	heads := make(map[int][]string)
	t.columns = make([][]string, t.x_dim+1)
	for i, cell := range t.cells {
		if cell.y < n {
			heads[cell.x] = append(heads[cell.x], cell.val)
		}
		t.cells[i].y_head = heads[cell.x]
		t.columns[cell.x] = heads[cell.x]
	}
}

// Creates a new TableData struct from provided data frame
// Populates all TableData with specified number of row and col headers
func NewTableData(df dataframe.DataFrame, y, x int) TableData {
	out := TableData{}
	out.populateDims(df)
	out.populateCells(df)
	out.populateRows(x)
	out.populateColumns(y)
	return out
}
