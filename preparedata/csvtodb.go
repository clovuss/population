package preparedata

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// Prepare1csv saves data from csv to db
func Prepare1csv(uch string) map[string][]string {
	file, err := os.Open("./data/" + uch + ".csv")
	if err != nil {
		fmt.Println(err)
	}
	myReader := csv.NewReader(file)
	myReader.Comma = ';'
	line := make(map[string][]string, 7)
	c := 0
	mapaXML := XmlToMapPrikrep("13")

	for {

		record, err := myReader.Read()

		if err == io.EOF {
			break
		}
		c++
		if err != nil {
			fmt.Println(err)
			break
		}
		if record[3] == "" {
			fmt.Println("ошибка на строке №", c)
			break
		}

		for _, v := range record {
			v = strings.TrimSpace(v)
		}
		fio := record[3] + record[4] + record[5] + record[7]
		switch record[4] {
		case "АРТЕМ": //Если есть имя АРТЕМ
			if _, ok := mapaXML[fio]; !ok { //поиск в мапе по ключу +АРТЕМ+.

				fio = record[3] + "АРТЁМ" + record[5] + record[7]

				if _, osk := mapaXML[fio]; !osk {
					fio = record[3] + record[4] + record[5] + record[7]
				}
			}
		case "СЕМЕН":
			if _, ok := mapaXML[fio]; !ok { //поиск в мапе по ключу +АРТЕМ+.

				fio = record[3] + "СЕМЁН" + record[5] + record[7]

				if _, osk := mapaXML[fio]; !osk {
					fio = record[3] + record[4] + record[5] + record[7]
				}
			}

		}
		//fmt.Println(record[3], record[4], record[5], record[7], fio)

		if mapaXML[fio].ENP != "" {

			line[fio] = []string{record[3], record[4], record[5], record[7], mapaXML[fio].ENP, record[8]}

			line[fio] = append(line[fio], telValidator(record[24])...)
			telValidator(record[24])

		}

	}

	return line
}

func telValidator(str string) []string {
	StringSlice := strings.Split(str, ",")
	//fmt.Println(StringSlice)
	output := make([]string, 0, 2)

	for _, w := range StringSlice {
		str = ""
		for _, v := range w {
			if unicode.IsDigit(v) {
				d := fmt.Sprintf("%c", v)
				str += d
			}
		}
		str = strings.TrimLeft(str, "78")
		countD := len(str)
		if countD == 10 || countD == 6 || countD == 0 {
			output = append(output, str)
		}
	}
	if len(output) == 1 {
		output = append(output, "")
	}

	return output
}

func NewtelValidator(str string) (int, error) {
	newtelstring := ""
	for _, v := range str {
		if unicode.IsDigit(v) {
			d := fmt.Sprintf("%c", v)
			newtelstring += d
		}
	}
	if len(newtelstring) != 10 {
		err := fmt.Errorf("безобразный номер")
		return 0, err
	}
	phoneNumber, err := strconv.Atoi(newtelstring)
	if err != nil {
		return 0, err
	}

	return phoneNumber, nil
}
