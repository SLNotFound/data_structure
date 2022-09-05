package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var f *os.File

func getMsgFilePath(msgDir string) []string {
	var msgFilePathList []string
	dirs, err := ioutil.ReadDir(msgDir)
	if err != nil {
		fmt.Printf("read dir error: %v\n", err)
	}
	for _, dir := range dirs {
		if dir.IsDir() {
			msgFileDir := msgDir + "\\" + dir.Name()
			msgFile, _ := ioutil.ReadDir(msgFileDir)
			for _, msgFilePath := range msgFile {
				msgFilePathList = append(msgFilePathList, msgFileDir+"\\"+msgFilePath.Name())
			}
		}
	}
	return msgFilePathList
}

func SplitData(filePath string) (dataMap map[string]string) {
	var err error
	f, err = os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var dataList []string
	for scanner.Scan() {
		dataList = append(dataList, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
	InfoList := dataList[:9]
	var m = make(map[string]string)
	for _, v := range InfoList {
		split := strings.Split(v, ":")
		if len(split) > 1 {
			fmt.Sprintf("%#v", split)
			if strings.ToLower(split[0]) == "senderid" || strings.ToLower(split[0]) == "sendername" {
				m[strings.ToLower(split[0])] = strings.Trim(split[1], " ")
			}
		}
	}
	return m
}

func WriterToFile(dataMap map[string]string, fileName string) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open file failed, err: %v\n", err)
	}
	writer := bufio.NewWriter(file)
	writer.WriteString(dataMap["senderid"] + "," + dataMap["sendername"] + "\n")
	writer.Flush()
}

func main() {
	msgDir := os.Args[1]
	fileName := "user.txt"
	filePathList := getMsgFilePath(msgDir)
	fmt.Println(len(filePathList))
	for _, filePath := range filePathList {
		m := SplitData(filePath)
		WriterToFile(m, fileName)
	}
}
