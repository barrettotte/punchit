package main

import (
	"fmt"
	"github.com/barrettotte/punchit"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func guard(err error) {
	if err != nil {
		panic(err)
	}
}

func processFile(filePath string, baseOut string) {
	cards, err := punchit.PunchFile(filePath)
	guard(err)
	guard(os.MkdirAll(baseOut, os.ModePerm))

	split := strings.Split(filePath, "/")
	outPath := baseOut + string(os.PathSeparator) + split[len(split)-1] + ".txt"

	data := ""
	for _, card := range cards {
		data += card.String() + "\n\n"
	}
	guard(writeToFile(outPath, data))

	fmt.Printf("Punched %d card(s) to %s\n", len(cards), outPath)
}

func writeToFile(filePath string, data string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}

func main() {
	var files []string
	root, out := "src", "cards"

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	guard(err)

	for _, file := range files {
		processFile(file, out)
	}
}
