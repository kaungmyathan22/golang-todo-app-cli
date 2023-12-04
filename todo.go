package todo

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type Item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type Todos []Item

func (pointerTodos *Todos) Add(task string) {
	todo := Item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Now(),
	}
	*pointerTodos = append(*pointerTodos, todo)
}

func (pointerTodos *Todos) MarkAsComplete(index int) error {
	todos := *pointerTodos
	if index <= 0 || index > len(todos) {
		return errors.New("invalid index")
	}
	todos[index-1].CompletedAt = time.Now()
	todos[index-1].Done = true
	return nil
}

func (pointerTodos *Todos) Delete(index int) error {
	todos := *pointerTodos
	if index <= 0 || index > len(todos) {
		return errors.New("invalid index")
	}

	*pointerTodos = append(
		todos[:index-1],
		todos[index:]...,
	)
	return nil
}

func (pointerTodos *Todos) Load(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if len(file) == 0 {
		return err
	}
	err = json.Unmarshal(file, pointerTodos)
	if err != nil {
		return err
	}
	return nil
}

func (pointerTodos *Todos) Store(filename string) error {
	data, err := json.Marshal(pointerTodos)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}

func (pointerTodos *Todos) Print() error {
	for i, item := range *pointerTodos {
		i++
		fmt.Printf("%d - %s \n", i, item.Task)
	}
	return nil
}

func getInput(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}
	if len(scanner.Text()) == 0 {
		return "", errors.New("empty todo is not allowed")
	}
	return scanner.Text(), nil
}
