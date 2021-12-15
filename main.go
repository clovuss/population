package main

import (
	"encoding/xml"
	"fmt"
	"golang.org/x/text/encoding/charmap"
	"io"
	"os"
)

type xmlPopulation struct {
	Zglv    ZGLV
	Prikrep []PRIKREP
}
type ZGLV struct {
	Version  string `xml:"VERSION"`
	Data     string `xml:"DATA"`
	Year     string `xml:"YEAR"`
	Period   string `xml:"PERIOD"`
	Filename string `xml:"FILENAME"`
}

type Docs struct {
}

type PRIKREP struct {
	Pid          string `xml:"pid"`
	ENP          string `xml:"ENP"`
	FAM          string `xml:"FAM"`
	IM           string `xml:"IM"`
	OT           string `xml:"OT"`
	BIRTHDAY     string `xml:"DR"`
	GENDER       string `xml:"W"`
	SNILS        string `xml:"SS"`
	PLACEOFBIRTH string `xml:"MR"`
	//Q string `xml:"Q"`
	OPDOC   string `xml:"OPDOC"`
	SPOL    string `xml:"SPOL"`
	NPOL    string `xml:"NPOL"`
	DOCTP   string `xml:"DOCTP"`
	DOCS    string `xml:"DOCS"`
	DOCN    string `xml:"DOCN"`
	DOCDT   string `xml:"DOCDT"`
	DOCORG  string `xml:"DOCORG"`
	CN      string `xml:"CN"`
	SUBJ    string `xml:"SUBJ"`
	RN      string `xml:"RN"`
	INDX    string `xml:"INDX"`
	RNNAME  string `xml:"RNNAME"`
	CITY    string `xml:"CITY"`
	NP      string `xml:"NP"`
	UL      string `xml:"UL"`
	DOM     string `xml:"DOM"`
	KOR     string `xml:"KOR"`
	KV      string `xml:"KV"`
	LPU     string `xml:"LPU"`
	LPUAUTO string `xml:"LPUAUTO"`
	LPUDT   string `xml:"LPUDT"`
	LPUUCH  string `xml:"LPUUCH"`
	KODPODR string `xml:"KODPODR"`
	SSD     string `xml:"SSD"`
}

type xmlsource struct {
	Zglv    ZGLV      `xml:"ZGLV"`
	Prikrep []PRIKREP `xml:"PRIKREP"`
}

func main() {
	xmlfile, err := os.Open("./data/211103.xml")

	if err != nil {
		fmt.Println(err)
	}

	defer xmlfile.Close()
	xmlDecoder := xml.NewDecoder(xmlfile)
	if err != nil {
		fmt.Println(err)
	}

	xmlDecoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		switch charset {
		case "windows-1251":
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		default:
			return nil, fmt.Errorf("unknown charset: %s", charset)
		}
	}
	//var all xmlsource
	//err = xmlDecoder.Decode(&all)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println("r:", all.Zglv)
	population := make([]PRIKREP, 0)
	for {
		token, err := xmlDecoder.Token()
		if token == nil {
			break
		}
		if err != nil {
			fmt.Println(11, err)
		}

		switch tp := token.(type) {
		case xml.StartElement:
			if tp.Name.Local == "PRIKREP" {
				var p PRIKREP
				err := xmlDecoder.DecodeElement(&p, &tp)
				if err != nil {
					fmt.Println(err)
				}
				population = append(population, p)

			}

		}

	}

	fmt.Println(population[0])
}
