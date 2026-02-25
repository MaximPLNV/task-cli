package repository

import (
	"bufio"
	"encoding/json"
	"errors"
	"task-cli/internal/domain/entities"
)

func NewJsonRepository(fileName string) *JsonRepository {
	jr := JsonRepository{}
	jr.fileName = fileName

	return &jr
}

type JsonRepository struct {
	fileName string
}

func (jr *JsonRepository) GetById(id int) (*entities.Task, error) {
	worker := NewReadWorker(jr.fileName)
	var taskById *entities.Task

	worker.lineAction = func(t *entities.Task) {
		if t.Id == id {
			taskById = t
			worker.finishAction <- true
		} else if t.Id > id {
			worker.finishAction <- true
		}
	}

	if err := worker.ReadByLine(); err != nil {
		return nil, err
	}

	if taskById == nil {
		return nil, errors.New("Task with the Id not found")
	}

	return taskById, nil
}

func (jr *JsonRepository) GetAll() (*[]entities.Task, error) {
	worker := NewReadWorker(jr.fileName)
	var tasks []entities.Task

	worker.lineAction = func(t *entities.Task) {
		tasks = append(tasks, *t)
	}

	if err := worker.ReadByLine(); err != nil {
		return nil, err
	}

	if len(tasks) < 1 {
		return nil, errors.New("There is no any task in the file")
	}

	return &tasks, nil
}

func (jr *JsonRepository) GetByStatus(st string) (*[]entities.Task, error) {
	worker := NewReadWorker(jr.fileName)
	var tasks []entities.Task

	worker.lineAction = func(t *entities.Task) {
		if t.Status == st {
			tasks = append(tasks, *t)
		}
	}

	if err := worker.ReadByLine(); err != nil {
		return nil, err
	}

	if len(tasks) < 1 {
		return nil, errors.New("There is no any task with this status")
	}

	return &tasks, nil
}

func (jr *JsonRepository) Add(newT *entities.Task) error {
	worker := NewWriteWorker(jr.fileName)
	firstFreeId := 0
	isInserted := false

	worker.lineAction = func(t *entities.Task, line *[]byte, wr *bufio.Writer) error {
		if !isInserted && firstFreeId != t.Id {
			newT.Id = firstFreeId
			jsonData, jsonErr := json.Marshal(newT)

			if jsonErr != nil {
				return errors.New("File modification issues")
			}

			wr.Write(jsonData)
			isInserted = true
		}

		firstFreeId++
		wr.Write(*line)
		return nil
	}

	worker.postWriteAction = func(wr *bufio.Writer) error {
		if !isInserted {
			newT.Id = firstFreeId
			jsonData, jsonErr := json.Marshal(newT)

			if jsonErr != nil {
				return errors.New("File modification issues")
			}

			wr.Write(jsonData)
			isInserted = true
		}

		return nil
	}

	return worker.WriteByLine()
}

func (jr *JsonRepository) Update(updTask *entities.Task) error {
	worker := NewWriteWorker(jr.fileName)
	isUpdated := false

	worker.lineAction = func(t *entities.Task, line *[]byte, wr *bufio.Writer) error {
		if t.Id == updTask.Id {
			jsonData, jsonErr := json.Marshal(updTask)

			if jsonErr != nil {
				return errors.New("File modification issues")
			}

			line = &jsonData
			isUpdated = true
		} else if t.Id > updTask.Id && !isUpdated {
			return errors.New("There is no task with this Id")
		}

		wr.Write(*line)
		return nil
	}
	return worker.WriteByLine()
}

func (jr *JsonRepository) Delete(id int) error {
	worker := NewWriteWorker(jr.fileName)

	worker.lineAction = func(t *entities.Task, line *[]byte, wr *bufio.Writer) error {
		if t.Id != id {
			wr.Write(*line)
		}

		return nil
	}
	return worker.WriteByLine()
}
