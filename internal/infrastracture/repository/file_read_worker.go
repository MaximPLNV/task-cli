package repository

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"task-cli/internal/domain/entities"
)

func NewReadWorker(fileName string) *FileReadWorker {
	fw := FileReadWorker{}
	fw.fileName = fileName
	fw.fileAccessFlags = os.O_RDONLY
	fw.finishAction = make(chan bool)

	return &fw
}

type FileReadWorker struct {
	lineAction      func(*entities.Task)
	finishAction    chan (bool)
	fileName        string
	fileAccessFlags int
}

func (fw *FileReadWorker) ReadByLine() error {
	defer close(fw.finishAction)
	file, err := fw.openFile()

	if err != nil {
		return err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	fw.loopReadScanner(scanner)
	return nil
}

func (fw *FileReadWorker) openFile() (*os.File, error) {
	file, openErr := os.OpenFile(fw.fileName, fw.fileAccessFlags, 0644)

	if openErr != nil {
		return nil, fmt.Errorf("There is no %s file", fw.fileName)
	}

	return file, nil
}

func (fw *FileReadWorker) loopReadScanner(sc *bufio.Scanner) {
loop:
	for sc.Scan() {
		var t entities.Task
		line := sc.Bytes()

		if err := json.Unmarshal(line, &t); err != nil {
			fmt.Printf("Issue during parsing line \"%s\"\n", line)
			continue
		}

		fw.lineAction(&t)

		select {
		case <-fw.finishAction:
			break loop
		default:
			continue loop
		}
	}
}
