package db

import (
	"fmt"
	"os"
)

// Write to db file
func Write(line string) error {
	fileName := "db.jsonl"
	file, openErr := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, os.ModeAppend)

	if openErr != nil {
		fmt.Fprintf(os.Stderr, "xkcd:write:open %v\n", openErr)
		newFile, err := os.Create("./db.jsonl")

		if err != nil {
			return err
		}

		file = newFile
	}

	_, writeErr := file.WriteString(line + "\n")

	if writeErr != nil {
		fmt.Fprintf(os.Stderr, "xkcd:write:write %v\n", writeErr)
		return writeErr
	}

	syncErr := file.Sync()

	return syncErr
}
