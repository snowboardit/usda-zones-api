package data

import (
	"encoding/csv"
	"log"
	"os"
)

type Row struct {
	Zipcode string
	Zone    string
	Trange  string
	Title   string
}

var Data []Row

func CheckError(err error, message string) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func Load() []Row {
	usData := ReadCSV("./data/phm_us_zipcode.csv")[1:]
	prData := ReadCSV("./data/phm_pr_zipcode.csv")[1:]
	hiData := ReadCSV("./data/phm_hi_zipcode.csv")[1:]
	akData := ReadCSV("./data/phm_ak_zipcode.csv")[1:]
	data := AggregateRows(usData, prData, hiData, akData)
	return data
}

func AggregateRows(rows ...[]Row) []Row {
	var result []Row
	for _, rowSlice := range rows {
		result = append(result, rowSlice...)
	}
	return result
}

func ReadCSV(path string) []Row {
	file, err := os.Open(path)
	CheckError(err, "Error while reading the file")
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	CheckError(err, "Error reading records")
	var returner []Row
	for _, eachrecord := range records {
		returner = append(returner, Row{
			Zipcode: eachrecord[0],
			Zone:    eachrecord[1],
			Trange:  eachrecord[2],
			Title:   eachrecord[3],
		})
	}
	return returner
}
