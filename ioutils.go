package main

import (
    "encoding/csv"
	"io"
    "io/ioutil"
	"strings"
    "os"
)

// WriteOutput writes a csv file (specified in filename) from a slice of strings (specified in words)
func WriteOutput(words []string, filename string) {
    // create file
    f, err := os.Create(filename)
    Check(err)
    
    //writer
    w := csv.NewWriter(f)

    err = w.Write(words)
    Check(err)    
	w.Flush()
    Check(err)
}

// CsvPuzzleSlice creates a "word puzzle" slice by loading the file specified in filename
func CsvPuzzleSlice(filename string) [][]string {
    f, err := ioutil.ReadFile(filename)
    Check(err)
	r := csv.NewReader(strings.NewReader(string(f)))
    slice := make([][]string, 0, 0)
    
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
        Check(err)

        slice = append(slice, record)
	}
    return slice
}

// CsvWords creates a slice of strings by loading the csv file specified in filename
func CsvWords(filename string) []string {
    f, err := ioutil.ReadFile(filename)
    Check(err)
	r := csv.NewReader(strings.NewReader(string(f)))

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
        Check(err)

        return record
	}
    return nil
}