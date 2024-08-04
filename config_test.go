// -*- coding: utf-8 -*-

// Created on Sun Jul 21 04:00:25 PM EDT 2024
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import "testing"

func TestReadConfig(t *testing.T) {
	table := [2]struct {
		path string
		exp  ConfigFields
	}{
		{"data/test_config1.json", ConfigFields{0, 0, "", "", "", ""}},
		{"data/test_config2.json", ConfigFields{10, 1000, "test.csv", "test.txt", "test", "test"}},
	}

	for _, test := range table {
		res, err := ReadConfig(test.path)
		if err != nil {
			t.Errorf("%v", err)
		}
		if res != test.exp {
			t.Errorf("ReadConfig(%v) = %v, expected %v", test.path, res, test.exp)
		}
	}
}

func TestReadFields(t *testing.T) {
	table := [2]struct {
		path string
		exp  FormatFields
	}{
		{"data/test_config1.json", FormatFields{"", "", "", "", "", "", ""}},
		{"data/test_config2.json", FormatFields{"test", "test", "test", "test", "test", "test", "test"}},
	}

	for _, test := range table {
		res, err := ReadFields(test.path)
		if err != nil {
			t.Errorf("%v", err)
		}
		if res != test.exp {
			t.Errorf("ReadFields(%v) = %v, expected %v", test.path, res, test.exp)
		}
	}
}
