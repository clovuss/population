package preparedata

import (
	"encoding/xml"
	"fmt"
	"golang.org/x/text/encoding/charmap"
	"io"
	"os"
)

type xmlsource struct {
	Zglv    ZGLV      `xml:"ZGLV"`
	Prikrep []PRIKREP `xml:"PRIKREP"`
}

type ZGLV struct {
	Version  string `xml:"VERSION"`
	Data     string `xml:"DATA"`
	Year     string `xml:"YEAR"`
	Period   string `xml:"PERIOD"`
	Filename string `xml:"FILENAME"`
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

func dataprepare() *[]PRIKREP {
	xmlfile, err := os.Open("./data/211103.xml")
	if err != nil {
		return nil
	}
	defer xmlfile.Close()
	xmlDecoder := xml.NewDecoder(xmlfile)
	if err != nil {
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
	var all xmlsource
	err = xmlDecoder.Decode(&all)
	if err != nil {
		fmt.Println(err)
	}

	return &all.Prikrep
}
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

		switch typ := token.(type) {
		case xml.StartElement:
			var pacient PRIKREP
			if typ.Name.Local == "PRIKREP" {
				xmlDecoder.DecodeElement(&pacient, &typ)
				pacients[pacient.Pid] = pacient
				co++
			}

		}

	}

	return pacients
}
