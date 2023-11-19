package functions

import (
	"encoding/json"
	"fmt"
	"os"
)

func WriteDataToFile(file *os.File, dataType string, data interface{}) {
	dataBytes, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		fmt.Printf("Failed to marshal data: %v\n", err)
		return
	}

	WriteOutputToFile(file, fmt.Sprintf("%s:\n%s\n", dataType, string(dataBytes)))
}