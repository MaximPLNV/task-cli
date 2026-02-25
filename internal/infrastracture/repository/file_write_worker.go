package repository

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"task-cli/internal/domain/entities"
)

func NewWriteWorker(fileName string) *FileWriteWorker {
	fw := FileWriteWorker{}
	fw.fileName = fileName
	fw.tempFilePattern = "temp-tasks-*"
	fw.fileAccessFlags = os.O_RDONLY | os.O_CREATE

	return &fw
}

type FileWriteWorker struct {
	lineAction      func(*entities.Task, *[]byte, *bufio.Writer) error
	postWriteAction func(*bufio.Writer) error
	fileName        string
	tempFilePattern string
	fileAccessFlags int
}

func (fw *FileWriteWorker) WriteByLine() error {
	file, err := fw.openFile()

	if err != nil {
		return err
	}
	defer file.Close()
	tmp, tmpErr := os.CreateTemp("", fw.tempFilePattern)

	if tmpErr != nil {
		return errors.New("File modification issues")
	}
	defer tmp.Close()
	defer os.Remove(tmp.Name())

	scanner := bufio.NewScanner(file)
	writer := bufio.NewWriter(tmp)

	if loopErr := fw.loopWriteScanner(scanner, writer); loopErr != nil {
		return loopErr
	}

	if fw.postWriteAction != nil {
		if postActionErr := fw.postWriteAction(writer); postActionErr != nil {
			return postActionErr
		}
	}
	return fw.saveChanges(file, writer)
}

func (fw *FileWriteWorker) openFile() (*os.File, error) {
	file, openErr := os.OpenFile(fw.fileName, fw.fileAccessFlags, 0644)

	if openErr != nil {
		return nil, fmt.Errorf("There is no %s file", fw.fileName)
	}

	return file, nil
}

func (fw *FileWriteWorker) loopWriteScanner(sc *bufio.Scanner, wr *bufio.Writer) error {
	for sc.Scan() {
		var t entities.Task
		line := sc.Bytes()

		if err := json.Unmarshal(line, &t); err != nil {
			fmt.Printf("Issue during parsing line \"%s\"\n", line)
			continue
		}

		if err := fw.lineAction(&t, &line, wr); err != nil {
			return err
		}
	}

	return nil
}

func (fw *FileWriteWorker) saveChanges(file *os.File, writer *bufio.Writer) error {
	writer.Flush()

	if renameErr := os.Rename(file.Name(), fw.fileName); renameErr != nil {
		return errors.New("File modification issues")
	}

	return nil
}
