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
		if os.IsExist(err) {
			fmt.Println("Directory already exists, proceeding...")
		} else {
			fmt.Println("Error creating folder:", err)
			return
		}
	}

	// Check if the file already exists
	filePath := filepath.Join(name, name+".go")
	if _, err := os.Stat(filePath); err == nil {
		fmt.Println("File already exists. Aborting to prevent overwrite.")
		return
	} else if !os.IsNotExist(err) {
		fmt.Println("Error checking file existence:", err)
		return
	}

	// Create the .go file
	problemTemplate := `package main

import (
  "fmt"
  "time"
)

func main() {
	startTime := time.Now()
  answer := 0
  fmt.Println(answer)
	duration := time.Since(startTime)
	fmt.Println("Duration:", duration)
}
`

	err = os.WriteFile(filePath, []byte(problemTemplate), 0644)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	fmt.Println("Folder and file created successfully!")
}
