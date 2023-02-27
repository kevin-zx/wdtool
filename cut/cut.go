package cut

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/yanyiwu/gojieba"
)

var (
	ErrorInvalidField = errors.New("invalid field")
	ErrorInvalidLabel = errors.New("invalid label")
)

func CutCsv(inputFileName string, outputFileName string, cutFields []string, label string, labelSplit string) error {
	inputFile, err := os.Open(inputFileName)
	if err != nil {
		return err
	}
	defer inputFile.Close()
	csvwR := csv.NewReader(inputFile)
	cutFiledsNum := make([]int, len(cutFields))
	labelNum := -1
	i := 0
	use_hmm := true
	x := gojieba.NewJieba()
	defer x.Free()
	cutResult := make(map[string]map[string]int)
	cutResult["all"] = make(map[string]int)
	// for _, field := range cutFields {
	//   cutResult[field] = make(map[string]int)
	// }

	for {
		record, err := csvwR.Read()
		if err != nil && err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if i == 0 {
			for i, field := range cutFields {
				found := false
				for j, name := range record {
					if field == name {
						cutFiledsNum[i] = j
						found = true
						break
					}
					if label == field {
						labelNum = i
					}
				}
				if !found {
					return ErrorInvalidField
				}
				if labelNum == -1 {
					return ErrorInvalidLabel
				}
			}
			continue
		}

		content := ""
		for _, num := range cutFiledsNum {
			content += record[num] + ","
		}
		content = strings.TrimSuffix(content, ",")
		words := x.Cut(content, use_hmm)
		labels := strings.Split(record[labelNum], labelSplit)
		for _, word := range words {
			cutResult["all"][word]++
			for _, label := range labels {
				if _, ok := cutResult[label]; !ok {
					cutResult[label] = make(map[string]int)
				}
				cutResult[label][word]++
			}
		}
	}
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		return err
	}
	defer outputFile.Close()
	csvW := csv.NewWriter(outputFile)
	defer csvW.Flush()
	csvW.Write([]string{
		"label", "word", "count",
	})
	for label, result := range cutResult {
		for word, count := range result {
			csvW.Write([]string{label, word, strconv.Itoa(count)})
		}
	}

	return nil
}
