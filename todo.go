package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type Todo []item

// Add adds a new task to the Todo list.
//
// It takes a string parameter 'task' which represents the task to be added.
// It does not return anything.
func (t *Todo) Add(task string) {
	// Create a new item struct with the provided task, 'false' for the 'Done'
	// field, and the current time for the 'CreatedAt' field. The 'CompletedAt'
	// field is set to the zero value of time.Time.
	todo := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	// Append the new item to the Todo list. The '*t' dereferences the Todo
	// pointer to get the underlying slice and then appends the new item to it.
	*t = append(*t, todo)
}

// Complete marks a task as completed by setting the 'Done' field to true and the 'CompletedAt' field to the current time.
//
// It takes an integer parameter 'index' which represents the index of the task to be completed.
// The index is zero-based, meaning the first task is at index 0.
//
// It returns an error if the index is invalid (less than or equal to 0 or greater than the length of the Todo list).
func (t *Todo) Complete(index int) error {
	// Get a copy of the Todo list. This is done to avoid modifying the original list directly.
	ls := *t

	// Check if the index is valid. If it's not, return an error.
	if index <= 0 || index > len(ls) {
		return fmt.Errorf("invalid index: %d", index)
	}

	// Mark the task at the specified index as completed by setting the 'Done'
	// field to true and the 'CompletedAt' field to the current time.
	ls[index-1].Done = true
	ls[index-1].CompletedAt = time.Now()

	// Return nil to indicate that the operation was successful.
	return nil
}

// Delete removes a task from the Todo list.
//
// It takes an integer parameter 'index' which represents the index of the task
// to be removed. The index is zero-based, meaning the first task is at index 0.
//
// It returns an error if the index is invalid (less than or equal to 0 or
// greater than the length of the Todo list).
func (t *Todo) Delete(index int) error {
	// Get a copy of the Todo list. This is done to avoid modifying the original list directly.
	ls := *t

	// Check if the index is valid. If it's not, return an error.
	if index <= 0 || index > len(ls) {
		return fmt.Errorf("invalid index: %d", index)
	}

	// Remove the task at the specified index from the Todo list by creating a
	// new slice that contains all elements before the index and all elements
	// after the index.
	*t = append(ls[:index-1], ls[index:]...)

	// Return nil to indicate that the operation was successful.
	return nil
}

// Load loads a Todo list from a JSON file.
//
// It takes a string parameter 'filename' which represents the name of the file
// to be loaded.
//
// It returns an error if there was a problem reading the file or unmarshaling
// the JSON data.
func (t Todo) Load(filename string) error {
	// Read the contents of the file.
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		// If the error is related to the file not existing, return nil to
		// indicate that the operation was successful.
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		// If the error is not related to the file not existing, return the
		// error.
		return err
	}

	// If the file is empty, return an error.
	if len(file) == 0 {
		return err
	}

	// Unmarshal the JSON data into the Todo list.
	err = json.Unmarshal(file, t)
	if err != nil {
		// If there was an error unmarshaling the JSON data, return the error.
		return err
	}

	// Return nil to indicate that the operation was successful.
	return nil
}
