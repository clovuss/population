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
		//line[0] = record[0] //карднум
		//line[1] = record[1] //фамилия
		//line[2] = record[2]
		//line[3] = record[3]//отчество
		//line[4] = record[4]//др
		//line[5] = record[5]//прописка
		//line[6] = record[6] //мж
		//line[8] = record[8]
		line = record[:7]
		line = append(line, record[8])
		output = append(output, line)

	}
	return output
}
