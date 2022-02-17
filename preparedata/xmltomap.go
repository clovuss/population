package preparedata

import (
	"encoding/xml"
	"fmt"
	"golang.org/x/text/encoding/charmap"
	"io"
	"os"
	"time"
)

type zl struct {
	Date string `xml:"ZGLV>DATA"`
}

//type xmlprikrep struct {
//	Prikrep []PRIKREP `xml:"PRIKREP"`
//}
//type xmlzglv struct {
//	Zglv ZGLV `xml:"ZGLV"`
//}

//type ZGLV struct {
//	Version  string `xml:"VERSION"`
//	Data     string `xml:"DATA"`
//	Year     string `xml:"YEAR"`
//	Period   string `xml:"PERIOD"`
//	Filename string `xml:"FILENAME"`
//}

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
	DOCTP        string `xml:"DOCTP"`
	DOCS         string `xml:"DOCS"`
	DOCN         string `xml:"DOCN"`
	DOCDT        string `xml:"DOCDT"`
	DOCORG       string `xml:"DOCORG"`
	OPDOC        string `xml:"OPDOC"`
	SPOL         string `xml:"SPOL"`
	NPOL         string `xml:"NPOL"`
	CN           string `xml:"CN"`
	SUBJ         string `xml:"SUBJ"`
	RN           string `xml:"RN"`
	INDX         string `xml:"INDX"`
	RNNAME       string `xml:"RNNAME"`
	CITY         string `xml:"CITY"`
	NP           string `xml:"NP"`
	UL           string `xml:"UL"`
	DOM          string `xml:"DOM"`
	KOR          string `xml:"KOR"`
	KV           string `xml:"KV"`
	LPU          string `xml:"LPU"`
	LPUAUTO      string `xml:"LPUAUTO"`
	LPUDT        string `xml:"LPUDT"`
	LPUUCH       string `xml:"LPUUCH"`
	KODPODR      string `xml:"KODPODR"`
	SSD          string `xml:"SSD"`
}

//func dataprepare() *[]PRIKREP {
//	xmlfile, err := os.Open("./data/211103.xml")
//	if err != nil {
//		return nil
//	}
//	defer xmlfile.Close()
//	xmlDecoder := xml.NewDecoder(xmlfile)
//	if err != nil {
//		return nil
//	}
//	xmlDecoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
//		switch charset {
//		case "windows-1251":
//			return charmap.Windows1251.NewDecoder().Reader(input), nil
//		default:
//			return nil, fmt.Errorf("unknown charset: %s", charset)
//		}
//	}
//	var all xmlprikrep
//	err = xmlDecoder.Decode(&all)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	return &all.Prikrep
//}

func ImprooveDataPrepare(f string) map[string]PRIKREP {
	fp := "./data/PRIKREP_M300025_2112" + f + ".xml"
	xmlfile, err := os.Open(fp)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer xmlfile.Close()
	xmlDecoder := xml.NewDecoder(xmlfile)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	xmlDecoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		switch charset {
		case "windows-1251":
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		default:
			return nil, fmt.Errorf("unknown charset: %s", charset)
		}
	}

	pacients := make(map[string]PRIKREP)
	for {
		token, err := xmlDecoder.Token()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)

		}
		if token == nil {
			break
		}

		switch typ := token.(type) {
		case xml.StartElement:
			var pacient PRIKREP
			if typ.Name.Local == "PRIKREP" {
				xmlDecoder.DecodeElement(&pacient, &typ)
				pacients[pacient.Pid] = pacient
			}

		}

	}

	return pacients
}

//XmlToMapPrikrep generates map os string with date format as dd.mm.yyyy

func XmlToMapPrikrep(file string) map[string]PRIKREP {
	xmlfile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer xmlfile.Close()
	xmlDecoder := xml.NewDecoder(xmlfile)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	xmlDecoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		switch charset {
		case "windows-1251":
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		default:
			return nil, fmt.Errorf("unknown charset: %s", charset)
		}
	}
	co := 0
	pacients := make(map[string]PRIKREP)
	for {
		token, err := xmlDecoder.Token()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		if token == nil {
			break
		}

		switch typ := token.(type) { //StartElement, EndElement, CharData, Comment, ProcInst, or Directive
		case xml.StartElement:
			var pacient PRIKREP
			if typ.Name.Local == "PRIKREP" {
				xmlDecoder.DecodeElement(&pacient, &typ)
				point := []byte(".")
				temps := []byte(pacient.BIRTHDAY)
				newslice := []byte{temps[8], temps[9], point[0], temps[5], temps[6], point[0],
					temps[0], temps[1], temps[2], temps[3],
				}

				pacient.BIRTHDAY = string((newslice))
				pacients[pacient.FAM+pacient.IM+pacient.OT+pacient.BIRTHDAY] = pacient
				co++
			}

		}

	}

	return pacients
}

func getDateFromXml(file string) time.Time {
	xmlfile, err := os.Open(file)
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
	var xml zl
	xmlDecoder.Decode(&xml)
	dateOfFile, err := time.Parse("2006-01-02", xml.Date)
	if err != nil {
		fmt.Println(err)
	}
	return dateOfFile
}
