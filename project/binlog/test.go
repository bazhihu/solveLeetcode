package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func main() {
	fs, err := os.Open("data.csv")
	if err != nil {
		log.Fatalf("err :%+v", err)
		return
	}
	defer fs.Close()

	r := csv.NewReader(fs)
	for {
		row, err := r.Read()
		if err != nil && err != io.EOF {
			log.Fatalf("err: %+v", err)
		}
		if err == io.EOF {
			return
		}
		log.Println(row)
	}
}
