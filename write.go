package csvreader

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
)

func WirteAndSave(data [][]string, filepath string) error {
	file := &os.File{}
	defer file.Close()
	_, err := os.Stat(filepath)
	if !os.IsExist(err) {
		file, err = os.Create(filepath)
		if err != nil {
			return errors.New(fmt.Sprintf("Error creating file:%s", err))
		}
	} else {
		file, err = os.Open(filepath)
	}
	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, record := range data {
		err := writer.Write(record)
		if err != nil {
			return errors.New(fmt.Sprintf("Error writing record to CSV:%s", err))
		}
	}
	return nil
}
