# Todo CLI App

A beautiful, personalized command-line todo application written in Go. Manage your tasks with style and get motivated!

## Features

- **Personalized Experience**
  - Personal welcome message
  - Name-based task management
  - Motivational feedback
  - Progress tracking
  - Persistent user preferences

- **Beautiful CLI Interface**
  - Colorful output with intuitive visual cues
  - Clear task status indicators (/)
  - Task creation dates
  - Summary statistics

- **Core Functionality**
  - Add new tasks
  - Mark tasks as completed
  - Delete tasks
  - List all tasks with their status
  - Persistent storage using JSON

- **User Experience**
  - Simple and intuitive commands
  - Helpful error messages
  - Customizable storage location
  - Encouraging messages based on progress
  - Automatic task list updates after actions

## Demo

Here's how the app looks in action:

### First Time Usage
```bash
$ todo
Welcome to Todo App!
What's your name? Alice

Nice to meet you, Alice! 

Alice's Todo List
Your Personal Command Line Task Manager
────────────────────────────────────────
```

### Adding and Listing Todos
```bash
$ todo -add "Buy groceries"
Added todo: Buy groceries

Alice's Todos:
────────────────────────────────────────
[1] Buy groceries (added Nov 23)
────────────────────────────────────────
Total: 1  Completed: 0  Pending: 1
```

### Completing Tasks with Motivation
```bash
$ todo -complete 1
Completed todo #1

Alice's Todos:
────────────────────────────────────────
[1] Buy groceries (added Nov 23)
────────────────────────────────────────
Total: 1  Completed: 1  Pending: 0

Amazing job Alice! All tasks completed! 
```

### Multiple Tasks Progress
```bash
$ todo -list
Alice's Todos:
────────────────────────────────────────
[1] Buy groceries (added Nov 23)
[2] Call mom (added Nov 23)
[3] Write documentation (added Nov 23)
────────────────────────────────────────
Total: 3  Completed: 1  Pending: 2

Keep going Alice! You're making great progress!
```

### Help Screen
```bash
Alice's Todo List
Your Personal Command Line Task Manager
────────────────────────────────────────

Usage:
  todo [command] [arguments]

Available Commands:
  -list                  Show all your todos
  -add [todo text]       Create a new todo
  -complete [number]     Mark a todo as done
  -del [number]          Remove a todo

Examples:
  todo -add "Buy groceries"
  todo -list
  todo -complete 1
  todo -del 2
```

## Installation

1. Make sure you have Go 1.19 or later installed
2. Clone this repository:
   ```bash
   git clone https://github.com/arnoldadero/todo.git
   cd todo
   ```
3. Build the application:
   ```bash
   go build ./cmd/todo
   ```

## Configuration

The app stores your preferences in two files:
- `.todo_user`: Stores your name for personalized interactions
- `.todos.json`: Stores your todo items

You can customize the storage location by setting the `TODO_FILE` environment variable:
```bash
export TODO_FILE="/path/to/my/todos.json"
```

## Features Coming Soon

- Task categories and tags
- Due dates and reminders
- Priority levels
- Task notes and descriptions
- Weekly/monthly progress reports

## Contributing

Contributions are welcome! Feel free to submit issues and pull requests.

## License

This project is licensed under the MIT License - see the LICENSE file for details.