/**
 * Move a number of characters from the beginning of a filename
 * to the end of the filename.
 *
 * Files to rename must be placed in the current working directory.
 *
 * Use of this source code is governed by the GNU General Public License v 3.0+
 * which can be found in the LICENSE file.
 * 
 * Copyright (C) 2019 Rikard Hevosmaa
 */

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func visit(files *[]string) filepath.WalkFunc {
	// Step through the files in the working directory,
	// store the filenames in "files"
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if info.IsDir() {
			// Do not save a file if it's a directory
			return nil
		} else if filepath.Ext(path) == ".exe" || filepath.Ext(path) == ".go" {
			// Do not save executables
			return nil
		} else {
			*files = append(*files, info.Name())
		}
		return nil
	}
}

func getNumChar() int {
	var numChar int
	fmt.Printf("Number of characters to move (including spaces if any): ")
	n, err := fmt.Scanf("%d", &numChar)
	if err != nil || n != 1 {
		// Handle invalid input
		fmt.Println(n, err)
	}
	return numChar
}

func changeNames(files []string, numChar int) error {
	for i, file := range files {
		extension := filepath.Ext(file)
		filename := file[0:len(file)-len(extension)]

		// Create the new name. Trim leading and trailing whitespaces,
		// add the file extension.
		charToMove := filename[:numChar]
		filename = filename[numChar:]
		newFile := filename + "_" + charToMove
		newFile = strings.TrimSpace(newFile)
		newFile = newFile + extension

		// Rename the files
		err := os.Rename(file, newFile)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d\t%s renamed to %s\n", i+1, file, filename)
	}
	return nil
}

// Exit gracefully
func exit() {
	fmt.Println("\nProgram executed successfully.")
	time.Sleep(2*time.Second)
	os.Exit(0)
}

func main() {
	// Get the current working directory, save to "root"
	root, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// Get the name of the files in "root"
	var files []string
	err = filepath.Walk(root, visit(&files))
	if err != nil {
		log.Fatal(err)
	}
	// Change the filenames
	err = changeNames(files, getNumChar())
	if err != nil {
		log.Fatal(err)
	}
	exit()
}
