package repository

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"task-cli/internal/entities"
)

func NewJsonRepository(file string) *JsonRepository {
	jr := JsonRepository{}
	jr.fileName = file
	jr.tempFilePattern = "temp-tasks-*"
	jr.fileAccessFlags = os.O_RDONLY

	return &jr
}

type JsonRepository struct {
	fileName        string
	tempFilePattern string
	fileAccessFlags int
}

func (jr *JsonRepository) GetAll() (*[]entities.Task, error) {
	file, err := jr.openFile()

	if err != nil {
		return nil, err
	}
	defer file.Close()

	tasks := jr.parseTasks(file, "")
	return tasks, nil
}

func (jr *JsonRepository) GetByStatus(st string) (*[]entities.Task, error) {
	file, err := jr.openFile()

	if err != nil {
		return nil, err
	}
	defer file.Close()

	tasks := jr.parseTasks(file, st)
	return tasks, nil
}

func (jr *JsonRepository) Add(newT *entities.Task) error {
	jr.fileAccessFlags = jr.fileAccessFlags | os.O_CREATE

	firstFreeId := 0
	isInserted := false
	fn := func(line *[]byte, wr *bufio.Writer) error {
		var t entities.Task
		parsingErr := json.Unmarshal(*line, &t)

		if parsingErr != nil {
			return nil
		}

		if !isInserted && firstFreeId != t.Id {
			newT.Id = firstFreeId
			jsonData, jsonErr := json.Marshal(newT)

			if jsonErr != nil {
				return jr.generateModificationError()
			}

			wr.Write(jsonData)
			isInserted = true
		}

		firstFreeId++
		wr.Write(*line)
		return nil
	}

	afterFn := func(wr *bufio.Writer) error {
		if !isInserted {
			newT.Id = firstFreeId
			jsonData, jsonErr := json.Marshal(newT)

			if jsonErr != nil {
				return jr.generateModificationError()
			}

			wr.Write(jsonData)
			isInserted = true
		}

		return nil
	}

	return jr.updateFileByLine(fn, afterFn)
}

func (jr *JsonRepository) Update(updTask *entities.Task) error {
	isUpdated := false

	fn := func(line *[]byte, wr *bufio.Writer) error {
		var t entities.Task
		parsingErr := json.Unmarshal(*line, &t)

		if parsingErr != nil {
			return nil
		}

		if t.Id == updTask.Id {
			jsonData, jsonErr := json.Marshal(updTask)

			if jsonErr != nil {
				return jr.generateModificationError()
			}

			line = &jsonData
			isUpdated = true
		} else if t.Id > updTask.Id && !isUpdated {
			return errors.New("There is no task with this Id")
		}

		wr.Write(*line)
		return nil
	}

	return jr.updateFileByLine(fn, nil)
}

func (jr *JsonRepository) Delete(id int) error {
	fn := func(line *[]byte, wr *bufio.Writer) error {
		var t entities.Task
		parsingErr := json.Unmarshal(*line, &t)

		if parsingErr == nil && t.Id != id {
			wr.Write(*line)
		}

		return nil
	}

	return jr.updateFileByLine(fn, nil)
}

func (jr *JsonRepository) openFile() (*os.File, error) {
	file, openErr := os.OpenFile(jr.fileName, jr.fileAccessFlags, 0644)

	if openErr != nil {
		return nil, fmt.Errorf("There is no %s file", jr.fileName)
	}

	return file, nil
}

func (jr *JsonRepository) parseTasks(file *os.File, st string) *[]entities.Task {
	var tasks []entities.Task

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var t entities.Task
		line := scanner.Bytes()
		err := json.Unmarshal(line, &t)

		if err != nil {
			fmt.Printf("Issue during parsing line \"%s\"\n", line)
			continue
		}

		if st == "" || st == t.Status {
			tasks = append(tasks, t)
		}
	}

	return &tasks
}

func (jr *JsonRepository) generateModificationError() error {
	return errors.New("File modification issues")
}

func (jr *JsonRepository) updateFileByLine(fn func(*[]byte, *bufio.Writer) error, aftScanFn func(*bufio.Writer) error) error {
	file, err := jr.openFile()

	if err != nil {
		return err
	}
	defer file.Close()

	tmp, tmpErr := os.CreateTemp("", jr.tempFilePattern)

	if tmpErr != nil {
		return jr.generateModificationError()
	}
	defer os.Remove(tmp.Name())
	defer tmp.Close()

	scanner := bufio.NewScanner(file)
	writer := bufio.NewWriter(tmp)

	for scanner.Scan() {
		line := scanner.Bytes()

		if fnErr := fn(&line, writer); fnErr != nil {
			return err
		}

	}

	if aftScanFn != nil {
		if aftScanFnErr := aftScanFn(writer); aftScanFnErr != nil {
			return aftScanFnErr
		}
	}

	return jr.saveChanges(tmp, writer)
}

func (jr *JsonRepository) saveChanges(file *os.File, writer *bufio.Writer) error {
	writer.Flush()

	if renameErr := os.Rename(file.Name(), jr.fileName); renameErr != nil {
		return jr.generateModificationError()
	}

	return nil
}
