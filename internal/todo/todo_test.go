package todo_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/arnoldadero/todo/internal/todo"
)

func TestAdd(t *testing.T) {
	todos := &todo.Todos{}
	taskText := "Test task"
	todos.Add(taskText)

	if len(*todos) != 1 {
		t.Errorf("Expected todos to have 1 item, got %d", len(*todos))
	}

	todoItem := (*todos)[0]
	if todoItem.Task != taskText {
		t.Errorf("Expected task text %q, got %q", taskText, todoItem.Task)
	}

	if todoItem.Done {
		t.Error("Expected new task to be not done")
	}

	if todoItem.CreatedAt.IsZero() {
		t.Error("Expected CreatedAt to be set")
	}
}

func TestComplete(t *testing.T) {
	todos := &todo.Todos{}
	taskText := "Test task"
	todos.Add(taskText)

	if err := todos.Complete(1); err != nil {
		t.Errorf("Error completing todo: %v", err)
	}

	todoItem := (*todos)[0]
	if !todoItem.Done {
		t.Error("Expected task to be done")
	}
}

func TestDelete(t *testing.T) {
	todos := &todo.Todos{}
	taskText := "Test task"
	todos.Add(taskText)

	if err := todos.Delete(1); err != nil {
		t.Errorf("Error deleting todo: %v", err)
	}

	if len(*todos) != 0 {
		t.Errorf("Expected todos to be empty, got %d items", len(*todos))
	}
}

func TestStoreAndLoad(t *testing.T) {
	// Create a temporary directory for test files
	tmpDir, err := os.MkdirTemp("", "todo-test")
	if err != nil {
		t.Fatalf("Could not create temp directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create a todo file in the temporary directory
	todoFile := filepath.Join(tmpDir, "todos.json")

	// Create some test todos
	todos := &todo.Todos{}
	todos.Add("Task 1")
	todos.Add("Task 2")
	if err := todos.Complete(1); err != nil {
		t.Fatalf("Error completing todo: %v", err)
	}

	// Store todos
	if err := todos.Store(todoFile); err != nil {
		t.Fatalf("Error storing todos: %v", err)
	}

	// Load todos into a new slice
	loadedTodos := &todo.Todos{}
	if err := loadedTodos.Load(todoFile); err != nil {
		t.Fatalf("Error loading todos: %v", err)
	}

	// Compare original and loaded todos
	if len(*todos) != len(*loadedTodos) {
		t.Errorf("Expected %d todos, got %d", len(*todos), len(*loadedTodos))
	}

	for i := range *todos {
		orig := (*todos)[i]
		loaded := (*loadedTodos)[i]

		if orig.Task != loaded.Task {
			t.Errorf("Task %d: expected %q, got %q", i+1, orig.Task, loaded.Task)
		}

		if orig.Done != loaded.Done {
			t.Errorf("Task %d: expected done=%v, got done=%v", i+1, orig.Done, loaded.Done)
		}

		if !orig.CreatedAt.Equal(loaded.CreatedAt) {
			t.Errorf("Task %d: expected created_at=%v, got created_at=%v", i+1, orig.CreatedAt, loaded.CreatedAt)
		}
	}
}
