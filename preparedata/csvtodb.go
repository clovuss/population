package preparedata

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func Preparecsv() [][]string {
	file, err := os.Open("./data/7.csv")
	if err != nil {
		fmt.Println(err)
	}
	myReader := csv.NewReader(file)
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
		line := make([]string, 3)
		line[0] = record[10]
		line[1] = record[0]
		line[2] = record[8]
		output = append(output, line)

	}
	return output
}
