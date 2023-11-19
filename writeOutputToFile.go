package functions

import (
	"fmt"
	"os"
)

func WriteOutputToFile(file *os.File, format string, a ...interface{}) {
	data := fmt.Sprintf(format, a...)
	if _, err := file.WriteString(data + "\n"); err != nil {
		fmt.Printf("Failed to write to output file: %v\n", err)
	}
}
