// -*- coding: utf-8 -*-

// Created on Fri Jun 14 09:17:45 PM EDT 2024
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"fmt"
	"strings"
)

type TableFormatter interface {
	format(f FormatFields) []string
}

// Reformats DataValue structs into natural language for formatters that don't rely on arrays
func format_from_cells(t TableData, d string, f string, values ...any) []string {
	out := []string{}
	for _, cell := range t.cells {
		str := fmt.Sprintf(f, values...)
		x_head, y_head := cell.JoinHeaders(d)
		str = strings.Replace(str, "<x_head>", x_head, -1)
		str = strings.Replace(str, "<y_head>", y_head, -1)
		str = strings.Replace(str, "<cell_val>", cell.val, -1)
		str = strings.Replace(str, "  ", " ", -1)
		str = strings.TrimSpace(str)
		out = append(out, str)
	}
	return out
}

type CustomFormatter struct {
	TableData
	f_str  string
	values []any
}

// Formats DataValue data based on custom format string and specified values
func (f *CustomFormatter) format(ff FormatFields) []string {
	return format_from_cells(f.TableData, ff.Delim, f.f_str, f.values...)
}

type UnnamedCoordFormatter1 struct {
	TableData
}

// Format string: (val_label) (link) <x_head> and <y_head> (eq) <value>
// Example: price for extra pepperoni and no cheese is $12.00
func (f *UnnamedCoordFormatter1) format(ff FormatFields) []string {
	return format_from_cells(f.TableData, ff.Delim, "%s %s <x_head> and <y_head> %s <cell_val> \n", ff.ValLabel, ff.Link, ff.Eq)
}

type UnnamedCoordFormatter2 struct {
	TableData
}

// Format string: (link) <x_head> and (y_label), (val_label) (eq) <value>
// Example: For Extra pepperoni and no cheese, price will be $12.00
func (f *UnnamedCoordFormatter2) format(ff FormatFields) []string {
	return format_from_cells(f.TableData, ff.Delim, "%s <x_head> and <y_head> %s %s <cell_val> \n", ff.Link, ff.ValLabel, ff.Eq)
}

type NamedCoordFormatter1 struct {
	TableData
}

// Format string: (val_label) (link) (x_label) (eq) <x_head> and (y_label) (eq) <y_head> (eq) <value>
// Example: Price when size is medium and crust is thin is $15
func (f *NamedCoordFormatter1) format(ff FormatFields) []string {
	return format_from_cells(f.TableData, ff.Delim, "%s %s %s %s <x_head> and %s %s <y_head> %s <cell_val> \n", ff.ValLabel, ff.Link, ff.XLabel, ff.Eq, ff.YLabel, ff.Eq, ff.Eq)
}

type NamedCoordFormatter2 struct {
	TableData
}

// Format string: (link) (x_label) (eq) <x_head> and (y_label) (eq) <y_head>, (val_label) (eq) <value>
// Example: When size = medium and crust = thin, price = $15
func (f *NamedCoordFormatter2) format(ff FormatFields) []string {
	return format_from_cells(f.TableData, ff.Delim, "%s %s %s <x_head> and %s %s <y_head>, %s %s <cell_val> \n", ff.Link, ff.XLabel, ff.Eq, ff.YLabel, ff.Eq, ff.ValLabel, ff.Eq)
}

type NamedRowFormatter struct {
	TableData
}

// Format string: (link) (x_label) (eq) <x_head>, <y_head> (eq) <value>
// Example: If topping is meat, vegan is false
func (f *NamedRowFormatter) format(ff FormatFields) []string {
	return format_from_cells(f.TableData, ff.Delim, "%s %s %s <x_head>, <y_head> %s <cell_val> \n", ff.Link, ff.XLabel, ff.Eq, ff.Eq)
}

type NamedColFormatter struct {
	TableData
}

// Format string: (link) (y_label) (eq) <y_head>, <x_head> (eq) <value>
// Example: If crust is gluten free, price increases by $3
func (f *NamedColFormatter) format(ff FormatFields) []string {
	return format_from_cells(f.TableData, ff.Delim, "%s %s %s <y_head>, <x_head> %s <cell_val> \n", ff.Link, ff.YLabel, ff.Eq, ff.Eq)
}

type UnnamedRowKeyValFormatter struct {
	TableData
}

// Format string: (link) (x_head),  [(y_head) (eq) <value>]
// Example: For daily specials, [Monday is none, Tuesday is taco pizza, Wednesday is wing pizza]
func (f *UnnamedRowKeyValFormatter) format(ff FormatFields) []string {
	outmap := map[string][]string{}
	for _, cell := range f.cells {
		x_head, y_head := cell.JoinHeaders(ff.Delim)
		id := fmt.Sprintf("%s %s", ff.Link, x_head)
		str := fmt.Sprintf("%s %s %s", y_head, ff.Eq, cell.val)
		outmap[id] = append(outmap[id], str)
	}

	out := []string{}
	for id, arr := range outmap {
		str := fmt.Sprintf("%s, %s", id, strings.Join(arr, ", "))
		out = append(out, str)
	}
	return out
}

type UnnamedColKeyValFormatter struct {
	TableData
}

// Format string: (link) (y_head), [(x_head) (eq) <value>]
// Example: For sides, [Wings are $5, Mozz sticks are $7, Cheese curds are $6]
func (f *UnnamedColKeyValFormatter) format(ff FormatFields) []string {
	outmap := map[string][]string{}
	for _, cell := range f.cells {
		x_head, y_head := cell.JoinHeaders(ff.Delim)
		id := fmt.Sprintf("%s %s", ff.Link, y_head)
		str := fmt.Sprintf("%s %s %s", x_head, ff.Eq, cell.val)
		outmap[id] = append(outmap[id], str)
	}

	out := []string{}
	for id, arr := range outmap {
		str := fmt.Sprintf("%s, %s", id, strings.Join(arr, ", "))
		out = append(out, str)
	}
	return out
}

type NamedRowKeyValFormatter struct {
	TableData
}

// Format string: (link) (x_label) (eq) <x_head>, [(y_head) (eq) <value>]
// Example: When country = South Korea, [Dominos is #1, Pizza Alvolo is #2, PizzaHut is #3]
func (f *NamedRowKeyValFormatter) format(ff FormatFields) []string {
	outmap := map[string][]string{}
	for _, cell := range f.cells {
		x_head, y_head := cell.JoinHeaders(ff.Delim)
		id := fmt.Sprintf("%s %s %s %s", ff.Link, ff.XLabel, ff.Eq, x_head)
		str := fmt.Sprintf("%s %s %s", y_head, ff.Eq, cell.val)
		outmap[id] = append(outmap[id], str)
	}

	out := []string{}
	for id, arr := range outmap {
		str := fmt.Sprintf("%s, %s", id, strings.Join(arr, ", "))
		out = append(out, str)
	}
	return out
}

type NamedColKeyValFormatter struct {
	TableData
}

// Format string: (link) (y_label) (eq) <y_head>, [(x_head) (eq) <value>]
// Example: In the case that chain is Sbarro, [locations is 600, year founded is 1956, hq is Columbus, Ohio]
func (f *NamedColKeyValFormatter) format(ff FormatFields) []string {
	outmap := map[string][]string{}
	for _, cell := range f.cells {
		x_head, y_head := cell.JoinHeaders(ff.Delim)
		id := fmt.Sprintf("%s %s %s %s", ff.Link, ff.YLabel, ff.Eq, y_head)
		str := fmt.Sprintf("%s %s %s", x_head, ff.Eq, cell.val)
		outmap[id] = append(outmap[id], str)
	}

	out := []string{}
	for id, arr := range outmap {
		str := fmt.Sprintf("%s, %s", id, strings.Join(arr, ", "))
		out = append(out, str)
	}
	return out
}

type RowValFormatter struct {
	TableData
}

// Format string: (pre) <x_head> (link) [<value>]
// Example: All possible topping are [sausage, mushroom, olives]
func (f *RowValFormatter) format(ff FormatFields) []string {
	outmap := map[string][]string{}
	for _, cell := range f.cells {
		x_head, _ := cell.JoinHeaders(ff.Delim)
		id := fmt.Sprintf("%s %s %s", ff.Pre, x_head, ff.Link)
		str := cell.val
		outmap[id] = append(outmap[id], str)
	}

	out := []string{}
	for id, arr := range outmap {
		str := fmt.Sprintf("%s %s", id, strings.Join(arr, ", "))
		out = append(out, str)
	}
	return out
}

type ColValFormatter struct {
	TableData
}

// Format string: (pre) <y_head> (link) [<value>]
// Example: Size can be one of [small, medium, large]
func (f *ColValFormatter) format(ff FormatFields) []string {
	outmap := map[string][]string{}
	for _, cell := range f.cells {
		_, y_head := cell.JoinHeaders(ff.Delim)
		id := fmt.Sprintf("%s %s %s", ff.Pre, y_head, ff.Link)
		str := cell.val
		outmap[id] = append(outmap[id], str)
	}

	out := []string{}
	for id, arr := range outmap {
		str := fmt.Sprintf("%s %s", id, strings.Join(arr, ", "))
		out = append(out, str)
	}
	return out
}
