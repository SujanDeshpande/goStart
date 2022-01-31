package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// CarrierPayload Carrier Payload
type FileInfo struct {
	name  string
	size  int64
	isDir bool
}

func main() {
	fmt.Println("Starting")
	argsWithoutProg := os.Args[1:]
	fmt.Println(argsWithoutProg)

	fileInfoSlice := make([]FileInfo, 1)
	fileNameSlice := []string{}
	filepath.Walk("/Users/sujandeshpande/test/",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			//fmt.Println(path, info.Size())
			fileInfo := FileInfo{info.Name(), info.Size(), info.IsDir()}
			fileInfoSlice = append(fileInfoSlice, fileInfo)
			fileNameSlice = append(fileNameSlice, info.Name())
			testPointer(fileInfo)
			fileInfoSlice = append(fileInfoSlice, fileInfo)

			return nil
		})

	//	fmt.Println(fileNameSlice)
	//list := removeDuplicateStr(fileNameSlice)
	testPrintPointer(fileInfoSlice)
	fmt.Println(len(fileNameSlice))
	//	fmt.Println(len(list))

	//	fmt.Println(list)
}

func testPointer(fileInfo FileInfo) {
	//fmt.Println(fileInfo)
	fileInfo.name = "Sujan"

}

func testPrintPointer(fileInfoSlice []FileInfo) {
	for _, item := range fileInfoSlice {
		fmt.Println(item)
	}
}

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		_, value := allKeys[item]
		if !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	fmt.Println(allKeys)
	return list
}
