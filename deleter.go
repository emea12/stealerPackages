package functions

import (
	"fmt"
	"os"
)

func Deleter(filePath string) {
	// Delete the file
	err := os.Remove(filePath)
	if err != nil {
		fmt.Println(fmt.Errorf("Error deleting file: %v", err))
		return
	}
	fmt.Printf("File %s deleted successfully\n", filePath)
}
