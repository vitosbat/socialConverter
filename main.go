package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

func main() {
	t0 := time.Now()

	fileList, err := ioutil.ReadDir("input")
	if err != nil {
		fmt.Println(err)
	}

	if len(fileList) != 1 || filepath.Ext(fileList[0].Name()) != ".xml" {

		fmt.Println("ОШИБКА: В папке input должен находиться один XML-файл!")

	} else {

		Converter(fileList[0].Name())

	}

	fmt.Printf("Процесс завершен. Время исполнения: %v", time.Since(t0))
}

//Function read, convert and write new XML-file
func Converter(fileName string) {

	xmlFile, err := os.Open("input/" + fileName)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Входящий файл %s успешно прочитан.\nИдет процесс конвертации...\n", fileName)

	defer xmlFile.Close()

	//read opened xml file as a byte array
	byteValue, _ := ioutil.ReadAll(xmlFile)

	var request Request
	xml.Unmarshal(byteValue, &request)

	uoid := request.RequestInfo.Upid
	uoname := string(request.RequestInfo.Upname)

	countRecords := 0

	for x := range request.Records.Records {
		request.Records.Records[x].Uoid = uoid
		request.Records.Records[x].Uoname = uoname
		countRecords++
	}

	fmt.Printf("Общее количество обработанных записей: %d\n", countRecords)
	fmt.Println("Идет процесс записи файла output.xml...")

	request.RequestInfo.Upid = ""
	request.RequestInfo.Upname = ""
	request.RequestInfo.Filetype = "0"

	outData, err := xml.MarshalIndent(request, "", "  ")
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println(string(outData))
	outFile, err := os.Create("out/output.xml")
	if err != nil {
		fmt.Println("Can't create file", err)
	}

	fmt.Fprint(outFile, xml.Header+string(outData))
}
