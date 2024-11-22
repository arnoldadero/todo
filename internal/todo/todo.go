package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

// Todo represents a single todo item
type Todo struct {
	Task      string    `json:"task"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
}

// Todos is a slice of Todo items
type Todos []Todo

// Add creates a new todo item
func (t *Todos) Add(task string) {
	todo := Todo{
		Task:      task,
		Done:      false,
		CreatedAt: time.Now(),
	}
	*t = append(*t, todo)
}

// Complete marks a todo as completed
func (t *Todos) Complete(index int) error {
	if err := t.validateIndex(index); err != nil {
		return err
	}

	// Convert to 0-based index
	index--
	(*t)[index].Done = true
	return nil
}

// Delete removes a todo
func (t *Todos) Delete(index int) error {
	if err := t.validateIndex(index); err != nil {
		return err
	}

	// Convert to 0-based index
	index--
	*t = append((*t)[:index], (*t)[index+1:]...)
	return nil
}

// validateIndex checks if the given index is valid
func (t *Todos) validateIndex(index int) error {
	if index <= 0 || index > len(*t) {
		return errors.New("invalid todo number")
	}
	return nil
}

// Store saves todos to a file
func (t *Todos) Store(filename string) error {
	data, err := json.Marshal(t)
	if err != nil {
		return fmt.Errorf("failed to marshal todos: %v", err)
	}

	return os.WriteFile(filename, data, 0644)
}

// Load reads todos from a file
func (t *Todos) Load(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // Return empty todos if file doesn't exist
		}
		return fmt.Errorf("failed to read todo file: %v", err)
	}

	if len(data) == 0 {
		return nil // Empty file, return empty todos
	}

	if err := json.Unmarshal(data, t); err != nil {
		return fmt.Errorf("failed to unmarshal todos: %v", err)
	}

	return nil
}
