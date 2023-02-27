package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/kevin-zx/wdtool/cut"
)

var (
	inputFileName  = flag.String("input", "", "Input file name")
	outputFileName = flag.String("output", "", "Output file name")
	cutFileds      = flag.String("cut", "", "Cut fields (comma separated)")
	label          = flag.String("label", "", "Label")
	labelSeparator = flag.String("label-separator", "", "Label separator")
)

func main() {
	flag.Parse()
	if *inputFileName == "" {
		panic("input file name is required")
	}
	if *outputFileName == "" {
		panic("output file name is required")
	}
	if *cutFileds == "" {
		panic("cut fields is required")
	}
	fmt.Println("label:", *label)

	fileds := strings.Split(*cutFileds, ",")

	err := cut.CutCsv(*inputFileName, *outputFileName, fileds, *label, *labelSeparator)
	if err != nil {
		panic(err)
	}
}
