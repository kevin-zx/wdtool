package cut

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
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
		i++
		record, err := csvwR.Read()
		if err != nil && err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if i == 1 {
			for i, field := range cutFields {
				found := false
				for j, name := range record {
					log.Println(name, field)
					if field == name {
						cutFiledsNum[i] = j
						found = true
						break
					}
				}
				if !found {
					log.Printf("field %s not found", field)
					return ErrorInvalidField
				}
			}

			for j, name := range record {
				if label == name {
					labelNum = j
					break
				}
			}
			if labelNum == -1 {
				return ErrorInvalidLabel
			}
			continue
		}

		content := ""
		for _, num := range cutFiledsNum {
			content += record[num] + ","
		}
		content = strings.TrimSuffix(content, ",")
		words := x.Cut(content, use_hmm)
		labels := []string{}
		if labelSplit != "" {
			labels = strings.Split(record[labelNum], labelSplit)
		} else {
			labels = append(labels, record[labelNum])
		}

		for _, word := range words {
			cutResult["all"][word]++
			for _, label := range labels {
				if label == "" {
					continue
				}
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
			if count < 2 {
				continue
			}
			if len(strings.Split(word, "")) < 2 {
				continue
			}
			csvW.Write([]string{label, word, strconv.Itoa(count)})
		}
	}

	return nil
}
