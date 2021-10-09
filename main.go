package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/cconcannon/czml"
	"github.com/cconcannon/gpx"
	"github.com/cconcannon/gpx2czml"
)

func main() {
	gpxFile, czmlFile := validateArgs()

	data, err := readFile(gpxFile)

	if err != nil {
		fmt.Println(err)
		return
	}

	var g gpx.Gpx

	err = gpx.Unmarshal(data, &g)

	if err != nil {
		fmt.Println(err)
		return
	}

	c := gpx2czml.CreatePath("path", "testing-path", g)

	transform, err := czml.Marshal(c)

	if err != nil {
		fmt.Println(err)
	}

	ioutil.WriteFile(czmlFile, transform, 0644)

	if err != nil {
		fmt.Println(err)
	}

}

func validateArgs() (in, out string) {
	args := os.Args

	if len(args) != 3 {
		fmt.Println("Please specify the input filepath and output filepath")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	return inputFile, outputFile
}

func readFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
		return []byte{}, err
	}

	defer file.Close()

	data, err := ioutil.ReadAll(file)

	if err != nil {
		fmt.Println(err)
		return []byte{}, err
	}

	return data, nil
}
