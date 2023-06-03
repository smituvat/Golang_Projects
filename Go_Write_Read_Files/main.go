package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func CreateFile(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	_, err = file.WriteString("New file is created!")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}

func ReadFileName(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("File Name: %s \n", stat.Name())
}

func ReadFileContents(filename string) {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(contents))
}

func main() {
	CreateFile("test.txt")
	ReadFileName("test.txt")
	ReadFileContents("test.txt")
}
