package preparedata

import (
	"fmt"
	"os"
	"strings"
)

// SearchFiles возвращает список файлов
func SearchFiles() ([]os.FileInfo, error) {
	dir, err := os.ReadDir("./data")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	xmlInfos := make([]os.FileInfo, 0, 2)
	for _, entry := range dir {
		if strings.Contains(entry.Name(), "PRIKREP_M300025_") {
			info, err := entry.Info()
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			xmlInfos = append(xmlInfos, info)
		}
	}

	if len(xmlInfos) == 1 {
		return xmlInfos, nil
	}
	err = fmt.Errorf("проверьте файлы в папке")
	if len(xmlInfos) != 2 {
		return nil, err
	}
	if xmlInfos[0].ModTime().After(xmlInfos[1].ModTime()) {
		return xmlInfos, nil
	} else {
		xmlInfos[0], xmlInfos[1] = xmlInfos[1], xmlInfos[0]
		return xmlInfos, nil
	}

}

//CompareFiles возвращает список прибывшых и убывших
func CompareFiles(xml []os.FileInfo) (in, out map[string]PRIKREP, err error) {
	out = map[string]PRIKREP{}
	in = map[string]PRIKREP{}
	a := XmlToMapPrikrep("./data/" + xml[1].Name())
	b := XmlToMapPrikrep("./data/" + xml[0].Name())

	for k := range a { // перебор прошлой мапы //убывающие
		_, ok := b[k] //берем ключ от мапы а и по этому ключу смотрим в б. если нет, в новой мапе нет человека, т.е. он убыл
		if !ok {
			//fmt.Println(k) //делаем новый список людей убывабщих
			out[k] = a[k]

		}
	}

	a, b = b, a
	for k := range a { // перебор текущей мапы //прибывающие
		_, ok := b[k] //смотрим есть ли в новой мапе этот человек
		if !ok {
			//fmt.Println(k) //делаем новый список людей убывабщих
			in[k] = a[k]

		}
	}

	return in, out, nil

}
func Dbservice() (in, out map[string]PRIKREP, err error) {
	files, err := SearchFiles()
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}

	in, out, err = CompareFiles(files)
	return in, out, nil

}
