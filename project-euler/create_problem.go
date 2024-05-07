package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a name as a parameter.")
		return
	}

	name := os.Args[1]

	// Create the folder
	err := os.Mkdir(name, 0755) // 0755: Standard permissions for directories
	if err != nil {
		fmt.Println("Error creating folder:", err)
		return
	}

	// Create the .go file
	filename := filepath.Join(name, name+".go")
	problemTemplate := `package main

import "fmt"

func main() {
  answer := ""
  fmt.Println(answer)
}
`

	err = os.WriteFile(filename, []byte(problemTemplate), 0644) // 0644: Usual permissions for files
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	fmt.Println("Folder and file created successfully!")
}