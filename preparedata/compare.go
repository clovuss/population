package preparedata

import (
	"fmt"
	"os"
	"strings"
)

// CompareFiles returns 2 maps
func CompareFiles() (in, out map[string]PRIKREP) {
	dir, err := os.ReadDir("./data")
	if err != nil {
		fmt.Println(err)
	}
	xml := make([]string, 0, 2)
	for _, entry := range dir {
		if strings.Contains(entry.Name(), "PRIKREP_M300025_") {
			xml = append(xml, "./data/"+entry.Name())
		}
	}
	if len(xml) != 2 {
		fmt.Println("Недостаточно файлов хмл для сравнения")
		os.Exit(0)
	}
	var past, present string // имена файлов
	if getDateFromXml(xml[0]).Before(getDateFromXml(xml[1])) {
		past, present = xml[0], xml[1]
	} else {
		past, present = xml[1], xml[0]
	}
	out = map[string]PRIKREP{}
	in = map[string]PRIKREP{}
	a := XmlToMapPrikrep(past)
	b := XmlToMapPrikrep(present)
	for k, _ := range a { // перебор прошлой мапы //убывающие
		_, ok := b[k] //смотрим есть ли в новой мапе этот человек
		if !ok {
			//fmt.Println(k) //делаем новый список людей убывабщих
			out[k] = a[k]
		}
	}
	fmt.Println(out)
	a, b = b, a
	for k, _ := range a { // перебор прошлой мапы //прибывающие
		_, ok := b[k] //смотрим есть ли в новой мапе этот человек
		if !ok {
			//fmt.Println(k) //делаем новый список людей убывабщих
			in[k] = a[k]
		}
	}
	fmt.Println(in)
	if err := os.Remove(past); err != nil {
		fmt.Println(err)
	}
	return in, out
}
