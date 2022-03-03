package preparedata

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func PreparePromedCsv() [][]string {
	file, err := os.Open("./data/prom.csv")
	if err != nil {
		fmt.Println(err)
	}
	myReader := csv.NewReader(file)
	myReader.Comma = ';'
	output := make([][]string, 0)
	for {
		record, err := myReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		line := make([]string, 8)
		line = record[:7]
		line = append(line, record[8])
		output = append(output, line)

	}
	return output
}
