// -*- coding: utf-8 -*-

// Created on Sun Jul 21 04:00:25 PM EDT 2024
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"fmt"
	"testing"
)

func reference_fields() (f1, f2 FormatFields) {
	f1 = FormatFields{"", "", "", "", "", "", ""}
	f2 = FormatFields{"delim", "link", "eq", "pre", "vallabel", "xlabel", "ylabel"}
	return f1, f2
}

func TestUnnamedCoordFormatter1(t *testing.T) {
	t1, _, _ := reference_tables()
	f1, f2 := reference_fields()
	table := [2]struct {
		formatter TableFormatter
		fields    FormatFields
		exp       []string
	}{
		{&UnnamedCoordFormatter1{t1}, f1, []string{
			"and  _",
			"and  col1",
			"and  col2",
			"and  col3",
			"and  row1",
			"and  val11",
			"and  val12",
			"and  val13",
			"and  row2",
			"and  val21",
			"and  val22",
			"and  val23",
			"and  row3",
			"and  val31",
			"and  val32",
			"and  val33",
			"and  row4",
			"and  val41",
			"and  val42",
			"and  val43",
		}},
		{&UnnamedCoordFormatter1{t1}, f2, []string{
			"vallabel link and eq _",
			"vallabel link and eq col1",
			"vallabel link and eq col2",
			"vallabel link and eq col3",
			"vallabel link and eq row1",
			"vallabel link and eq val11",
			"vallabel link and eq val12",
			"vallabel link and eq val13",
			"vallabel link and eq row2",
			"vallabel link and eq val21",
			"vallabel link and eq val22",
			"vallabel link and eq val23",
			"vallabel link and eq row3",
			"vallabel link and eq val31",
			"vallabel link and eq val32",
			"vallabel link and eq val33",
			"vallabel link and eq row4",
			"vallabel link and eq val41",
			"vallabel link and eq val42",
			"vallabel link and eq val43",
		}},
	}

	for _, test := range table {
		res := test.formatter.format(test.fields)
		if fmt.Sprint(res) != fmt.Sprint(test.exp) {
			t.Errorf("%v.format(%v) = %v, expected %v", test.formatter, test.fields, res, test.exp)
		}
	}
}

func TestUnnamedCoordFormatter2(t *testing.T) {
	t1, _, _ := reference_tables()
	f1, f2 := reference_fields()
	table := [2]struct {
		formatter TableFormatter
		fields    FormatFields
		exp       []string
	}{
		{&UnnamedCoordFormatter2{t1}, f1, []string{
			"and  _",
			"and  col1",
			"and  col2",
			"and  col3",
			"and  row1",
			"and  val11",
			"and  val12",
			"and  val13",
			"and  row2",
			"and  val21",
			"and  val22",
			"and  val23",
			"and  row3",
			"and  val31",
			"and  val32",
			"and  val33",
			"and  row4",
			"and  val41",
			"and  val42",
			"and  val43",
		}},
		{&UnnamedCoordFormatter2{t1}, f2, []string{
			"link and vallabel eq _",
			"link and vallabel eq col1",
			"link and vallabel eq col2",
			"link and vallabel eq col3",
			"link and vallabel eq row1",
			"link and vallabel eq val11",
			"link and vallabel eq val12",
			"link and vallabel eq val13",
			"link and vallabel eq row2",
			"link and vallabel eq val21",
			"link and vallabel eq val22",
			"link and vallabel eq val23",
			"link and vallabel eq row3",
			"link and vallabel eq val31",
			"link and vallabel eq val32",
			"link and vallabel eq val33",
			"link and vallabel eq row4",
			"link and vallabel eq val41",
			"link and vallabel eq val42",
			"link and vallabel eq val43",
		}},
	}

	for _, test := range table {
		res := test.formatter.format(test.fields)
		if fmt.Sprint(res) != fmt.Sprint(test.exp) {
			t.Errorf("%v.format(%v) = %v, expected %v", test.formatter, test.fields, res, test.exp)
		}
	}
}

func TestNamedCoordFormatter1(t *testing.T) {
	t1, _, _ := reference_tables()
	f1, f2 := reference_fields()
	table := [2]struct {
		formatter TableFormatter
		fields    FormatFields
		exp       []string
	}{
		{&NamedCoordFormatter1{t1}, f1, []string{
			"and   _",
			"and   col1",
			"and   col2",
			"and   col3",
			"and   row1",
			"and   val11",
			"and   val12",
			"and   val13",
			"and   row2",
			"and   val21",
			"and   val22",
			"and   val23",
			"and   row3",
			"and   val31",
			"and   val32",
			"and   val33",
			"and   row4",
			"and   val41",
			"and   val42",
			"and   val43",
		}},
		{&NamedCoordFormatter1{t1}, f2, []string{
			"vallabel link xlabel eq and ylabel eq eq _",
			"vallabel link xlabel eq and ylabel eq eq col1",
			"vallabel link xlabel eq and ylabel eq eq col2",
			"vallabel link xlabel eq and ylabel eq eq col3",
			"vallabel link xlabel eq and ylabel eq eq row1",
			"vallabel link xlabel eq and ylabel eq eq val11",
			"vallabel link xlabel eq and ylabel eq eq val12",
			"vallabel link xlabel eq and ylabel eq eq val13",
			"vallabel link xlabel eq and ylabel eq eq row2",
			"vallabel link xlabel eq and ylabel eq eq val21",
			"vallabel link xlabel eq and ylabel eq eq val22",
			"vallabel link xlabel eq and ylabel eq eq val23",
			"vallabel link xlabel eq and ylabel eq eq row3",
			"vallabel link xlabel eq and ylabel eq eq val31",
			"vallabel link xlabel eq and ylabel eq eq val32",
			"vallabel link xlabel eq and ylabel eq eq val33",
			"vallabel link xlabel eq and ylabel eq eq row4",
			"vallabel link xlabel eq and ylabel eq eq val41",
			"vallabel link xlabel eq and ylabel eq eq val42",
			"vallabel link xlabel eq and ylabel eq eq val43",
		}},
	}

	for _, test := range table {
		res := test.formatter.format(test.fields)
		if fmt.Sprint(res) != fmt.Sprint(test.exp) {
			t.Errorf("%v.format(%v) = %v, expected %v", test.formatter, test.fields, res, test.exp)
		}
	}
}

func TestNamedCoordFormatter2(t *testing.T) {
	t1, _, _ := reference_tables()
	f1, f2 := reference_fields()
	table := [2]struct {
		formatter TableFormatter
		fields    FormatFields
		exp       []string
	}{
		{&NamedCoordFormatter2{t1}, f1, []string{
			"and  ,  _",
			"and  ,  col1",
			"and  ,  col2",
			"and  ,  col3",
			"and  ,  row1",
			"and  ,  val11",
			"and  ,  val12",
			"and  ,  val13",
			"and  ,  row2",
			"and  ,  val21",
			"and  ,  val22",
			"and  ,  val23",
			"and  ,  row3",
			"and  ,  val31",
			"and  ,  val32",
			"and  ,  val33",
			"and  ,  row4",
			"and  ,  val41",
			"and  ,  val42",
			"and  ,  val43",
		}},
		{&NamedCoordFormatter2{t1}, f2, []string{
			"link xlabel eq and ylabel eq , vallabel eq _",
			"link xlabel eq and ylabel eq , vallabel eq col1",
			"link xlabel eq and ylabel eq , vallabel eq col2",
			"link xlabel eq and ylabel eq , vallabel eq col3",
			"link xlabel eq and ylabel eq , vallabel eq row1",
			"link xlabel eq and ylabel eq , vallabel eq val11",
			"link xlabel eq and ylabel eq , vallabel eq val12",
			"link xlabel eq and ylabel eq , vallabel eq val13",
			"link xlabel eq and ylabel eq , vallabel eq row2",
			"link xlabel eq and ylabel eq , vallabel eq val21",
			"link xlabel eq and ylabel eq , vallabel eq val22",
			"link xlabel eq and ylabel eq , vallabel eq val23",
			"link xlabel eq and ylabel eq , vallabel eq row3",
			"link xlabel eq and ylabel eq , vallabel eq val31",
			"link xlabel eq and ylabel eq , vallabel eq val32",
			"link xlabel eq and ylabel eq , vallabel eq val33",
			"link xlabel eq and ylabel eq , vallabel eq row4",
			"link xlabel eq and ylabel eq , vallabel eq val41",
			"link xlabel eq and ylabel eq , vallabel eq val42",
			"link xlabel eq and ylabel eq , vallabel eq val43",
		}},
	}

	for _, test := range table {
		res := test.formatter.format(test.fields)
		if fmt.Sprint(res) != fmt.Sprint(test.exp) {
			t.Errorf("%v.format(%v) = %v, expected %v", test.formatter, test.fields, res, test.exp)
		}
	}
}

func TestNamedRowFormatter(t *testing.T) {
	t1, _, _ := reference_tables()
	f1, f2 := reference_fields()
	table := [2]struct {
		formatter TableFormatter
		fields    FormatFields
		exp       []string
	}{
		{&NamedRowFormatter{t1}, f1, []string{
			",  _",
			",  col1",
			",  col2",
			",  col3",
			",  row1",
			",  val11",
			",  val12",
			",  val13",
			",  row2",
			",  val21",
			",  val22",
			",  val23",
			",  row3",
			",  val31",
			",  val32",
			",  val33",
			",  row4",
			",  val41",
			",  val42",
			",  val43",
		}},
		{&NamedRowFormatter{t1}, f2, []string{
			"link xlabel eq , eq _",
			"link xlabel eq , eq col1",
			"link xlabel eq , eq col2",
			"link xlabel eq , eq col3",
			"link xlabel eq , eq row1",
			"link xlabel eq , eq val11",
			"link xlabel eq , eq val12",
			"link xlabel eq , eq val13",
			"link xlabel eq , eq row2",
			"link xlabel eq , eq val21",
			"link xlabel eq , eq val22",
			"link xlabel eq , eq val23",
			"link xlabel eq , eq row3",
			"link xlabel eq , eq val31",
			"link xlabel eq , eq val32",
			"link xlabel eq , eq val33",
			"link xlabel eq , eq row4",
			"link xlabel eq , eq val41",
			"link xlabel eq , eq val42",
			"link xlabel eq , eq val43",
		}},
	}

	for _, test := range table {
		res := test.formatter.format(test.fields)
		if fmt.Sprint(res) != fmt.Sprint(test.exp) {
			t.Errorf("%v.format(%v) = %v, expected %v", test.formatter, test.fields, res, test.exp)
		}
	}
}

func TestNamedColFormatter(t *testing.T) {
	t1, _, _ := reference_tables()
	f1, f2 := reference_fields()
	table := [2]struct {
		formatter TableFormatter
		fields    FormatFields
		exp       []string
	}{
		{&NamedColFormatter{t1}, f1, []string{
			",  _",
			",  col1",
			",  col2",
			",  col3",
			",  row1",
			",  val11",
			",  val12",
			",  val13",
			",  row2",
			",  val21",
			",  val22",
			",  val23",
			",  row3",
			",  val31",
			",  val32",
			",  val33",
			",  row4",
			",  val41",
			",  val42",
			",  val43",
		}},
		{&NamedColFormatter{t1}, f2, []string{
			"link ylabel eq , eq _",
			"link ylabel eq , eq col1",
			"link ylabel eq , eq col2",
			"link ylabel eq , eq col3",
			"link ylabel eq , eq row1",
			"link ylabel eq , eq val11",
			"link ylabel eq , eq val12",
			"link ylabel eq , eq val13",
			"link ylabel eq , eq row2",
			"link ylabel eq , eq val21",
			"link ylabel eq , eq val22",
			"link ylabel eq , eq val23",
			"link ylabel eq , eq row3",
			"link ylabel eq , eq val31",
			"link ylabel eq , eq val32",
			"link ylabel eq , eq val33",
			"link ylabel eq , eq row4",
			"link ylabel eq , eq val41",
			"link ylabel eq , eq val42",
			"link ylabel eq , eq val43",
		}},
	}

	for _, test := range table {
		res := test.formatter.format(test.fields)
		if fmt.Sprint(res) != fmt.Sprint(test.exp) {
			t.Errorf("%v.format(%v) = %v, expected %v", test.formatter, test.fields, res, test.exp)
		}
	}
}

func TestUnnamedRowKeyValFormatter(t *testing.T) {
	t1, _, _ := reference_tables()
	f1, f2 := reference_fields()
	table := [2]struct {
		formatter TableFormatter
		fields    FormatFields
		exp       []string
	}{
		{&UnnamedRowKeyValFormatter{t1}, f1, []string{
			" ,   _,   col1,   col2,   col3,   row1,   val11,   val12,   val13,   row2,   val21,   val22,   val23,   row3,   val31,   val32,   val33,   row4,   val41,   val42,   val43",
		}},
		{&UnnamedRowKeyValFormatter{t1}, f2, []string{
			"link ,  eq _,  eq col1,  eq col2,  eq col3,  eq row1,  eq val11,  eq val12,  eq val13,  eq row2,  eq val21,  eq val22,  eq val23,  eq row3,  eq val31,  eq val32,  eq val33,  eq row4,  eq val41,  eq val42,  eq val43",
		}},
	}

	for _, test := range table {
		res := test.formatter.format(test.fields)
		if fmt.Sprint(res) != fmt.Sprint(test.exp) {
			t.Errorf("%v.format(%v) = %v, expected %v", test.formatter, test.fields, res, test.exp)
		}
	}
}

func TestUnnamedColKeyValFormatter(t *testing.T) {
	t1, _, _ := reference_tables()
	f1, f2 := reference_fields()
	table := [2]struct {
		formatter TableFormatter
		fields    FormatFields
		exp       []string
	}{
		{&UnnamedColKeyValFormatter{t1}, f1, []string{
			" ,   _,   col1,   col2,   col3,   row1,   val11,   val12,   val13,   row2,   val21,   val22,   val23,   row3,   val31,   val32,   val33,   row4,   val41,   val42,   val43",
		}},
		{&UnnamedColKeyValFormatter{t1}, f2, []string{
			"link ,  eq _,  eq col1,  eq col2,  eq col3,  eq row1,  eq val11,  eq val12,  eq val13,  eq row2,  eq val21,  eq val22,  eq val23,  eq row3,  eq val31,  eq val32,  eq val33,  eq row4,  eq val41,  eq val42,  eq val43",
		}},
	}

	for _, test := range table {
		res := test.formatter.format(test.fields)
		if fmt.Sprint(res) != fmt.Sprint(test.exp) {
			t.Errorf("%v.format(%v) = %v, expected %v", test.formatter, test.fields, res, test.exp)
		}
	}
}

func TestNamedRowKeyValFormatter(t *testing.T) {
	t1, _, _ := reference_tables()
	f1, f2 := reference_fields()
	table := [2]struct {
		formatter TableFormatter
		fields    FormatFields
		exp       []string
	}{
		{&NamedRowKeyValFormatter{t1}, f1, []string{
			"   ,   _,   col1,   col2,   col3,   row1,   val11,   val12,   val13,   row2,   val21,   val22,   val23,   row3,   val31,   val32,   val33,   row4,   val41,   val42,   val43",
		}},
		{&NamedRowKeyValFormatter{t1}, f2, []string{
			"link xlabel eq ,  eq _,  eq col1,  eq col2,  eq col3,  eq row1,  eq val11,  eq val12,  eq val13,  eq row2,  eq val21,  eq val22,  eq val23,  eq row3,  eq val31,  eq val32,  eq val33,  eq row4,  eq val41,  eq val42,  eq val43",
		}},
	}

	for _, test := range table {
		res := test.formatter.format(test.fields)
		if fmt.Sprint(res) != fmt.Sprint(test.exp) {
			t.Errorf("%v.format(%v) = %v, expected %v", test.formatter, test.fields, res, test.exp)
		}
	}
}

func TestNamedColKeyValFormatter(t *testing.T) {
	t1, _, _ := reference_tables()
	f1, f2 := reference_fields()
	table := [2]struct {
		formatter TableFormatter
		fields    FormatFields
		exp       []string
	}{
		{&NamedColKeyValFormatter{t1}, f1, []string{
			"   ,   _,   col1,   col2,   col3,   row1,   val11,   val12,   val13,   row2,   val21,   val22,   val23,   row3,   val31,   val32,   val33,   row4,   val41,   val42,   val43",
		}},
		{&NamedColKeyValFormatter{t1}, f2, []string{
			"link ylabel eq ,  eq _,  eq col1,  eq col2,  eq col3,  eq row1,  eq val11,  eq val12,  eq val13,  eq row2,  eq val21,  eq val22,  eq val23,  eq row3,  eq val31,  eq val32,  eq val33,  eq row4,  eq val41,  eq val42,  eq val43",
		}},
	}

	for _, test := range table {
		res := test.formatter.format(test.fields)
		if fmt.Sprint(res) != fmt.Sprint(test.exp) {
			t.Errorf("%v.format(%v) = %v, expected %v", test.formatter, test.fields, res, test.exp)
		}
	}
}

func TestRowValFormatter(t *testing.T) {
	t1, _, _ := reference_tables()
	f1, f2 := reference_fields()
	table := [2]struct {
		formatter TableFormatter
		fields    FormatFields
		exp       []string
	}{
		{&RowValFormatter{t1}, f1, []string{
			"   _, col1, col2, col3, row1, val11, val12, val13, row2, val21, val22, val23, row3, val31, val32, val33, row4, val41, val42, val43",
		}},
		{&RowValFormatter{t1}, f2, []string{
			"pre  link _, col1, col2, col3, row1, val11, val12, val13, row2, val21, val22, val23, row3, val31, val32, val33, row4, val41, val42, val43",
		}},
	}

	for _, test := range table {
		res := test.formatter.format(test.fields)
		if fmt.Sprint(res) != fmt.Sprint(test.exp) {
			t.Errorf("%v.format(%v) = %v, expected %v", test.formatter, test.fields, res, test.exp)
		}
	}
}

func TestColValFormatter(t *testing.T) {
	t1, _, _ := reference_tables()
	f1, f2 := reference_fields()
	table := [2]struct {
		formatter TableFormatter
		fields    FormatFields
		exp       []string
	}{
		{&ColValFormatter{t1}, f1, []string{
			"   _, col1, col2, col3, row1, val11, val12, val13, row2, val21, val22, val23, row3, val31, val32, val33, row4, val41, val42, val43",
		}},
		{&ColValFormatter{t1}, f2, []string{
			"pre  link _, col1, col2, col3, row1, val11, val12, val13, row2, val21, val22, val23, row3, val31, val32, val33, row4, val41, val42, val43",
		}},
	}

	for _, test := range table {
		res := test.formatter.format(test.fields)
		if fmt.Sprint(res) != fmt.Sprint(test.exp) {
			t.Errorf("%v.format(%v) = %v, expected %v", test.formatter, test.fields, res, test.exp)
		}
	}
}
