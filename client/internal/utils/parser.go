package utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

// CSVParseTestcaseFile accepts a testcase path and returns all testcases.
func CSVParseTestcaseFile(path string) ([]map[string]interface{}, error) {
	list := make([]map[string]interface{}, 0)

	// open CSV file
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open the CSV file %s: %v", path, err)
	}
	defer file.Close()

	// create CSV reader
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // allow variable number of fields per row

	// read all data
	data, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV file %s: %v", path, err)
	}

	// loop variables
	var (
		index      string
		servers    []string
		byzantines []string
		sets       []string
	)

	// read row by row
	for i, row := range data {
		if row[0] == "" { // old row
			tmp := strings.Split(strings.Replace(strings.Replace(row[1], "(", "", -1), ")", "", -1), ", ")
			sets = append(sets,
				strings.TrimSpace(tmp[0]),
				strings.TrimSpace(tmp[1]),
				strings.TrimSpace(tmp[2]),
			)
		} else {
			// save the current values
			if index != "" {
				list = append(list, map[string]interface{}{
					"servers":      servers,
					"byzantines":   byzantines,
					"transactions": sets,
				})
			}

			// reset values
			index = row[0]
			servers = make([]string, 0)
			byzantines = make([]string, 0)
			sets = make([]string, 0)

			// set servers and contact servers
			tmpS := append(servers, strings.Split(strings.Replace(strings.Replace(row[2], "[", "", -1), "]", "", -1), ", ")...)
			for _, item := range tmpS {
				if item != "" {
					servers = append(servers, item)
				}
			}
			tmpB := append(servers, strings.Split(strings.Replace(strings.Replace(row[3], "[", "", -1), "]", "", -1), ", ")...)
			for _, item := range tmpB {
				if item != "" {
					byzantines = append(byzantines, item)
				}
			}

			// process the first row transactions
			tmp := strings.Split(strings.Replace(strings.Replace(row[1], "(", "", -1), ")", "", -1), ", ")
			sets = append(sets,
				strings.TrimSpace(tmp[0]),
				strings.TrimSpace(tmp[1]),
				strings.TrimSpace(tmp[2]),
			)
		}

		// save the last set
		if i == len(data)-1 {
			list = append(list, map[string]interface{}{
				"servers":      servers,
				"byzantines":   byzantines,
				"transactions": sets,
			})
		}
	}

	return list, nil
}
