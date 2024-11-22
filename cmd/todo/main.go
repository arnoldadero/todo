package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/arnoldadero/todo/internal/todo"
	"github.com/fatih/color"
)

var todoFile = ".todos.json"
var userName string

func init() {
	if envFile := os.Getenv("TODO_FILE"); envFile != "" {
		todoFile = envFile
	}
	// Try to load username from file
	if data, err := os.ReadFile(".todo_user"); err == nil {
		userName = strings.TrimSpace(string(data))
	}
}

func getUserName() string {
	if userName != "" {
		return userName
	}
	
	color.New(color.FgHiCyan, color.Bold).Println("\n👋 Welcome to Todo App!")
	fmt.Print("What's your name? ")
	
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	
	if name == "" {
		name = "Friend"
	}
	
	// Save username for future sessions
	os.WriteFile(".todo_user", []byte(name), 0644)
	userName = name
	
	color.New(color.FgHiGreen).Printf("\nNice to meet you, %s! 🎉\n", name)
	time.Sleep(1 * time.Second)
	return name
}

func printHeader() {
	name := getUserName()
	color.New(color.FgHiCyan, color.Bold).Printf("\n📝 %s's Todo List\n", name)
	color.New(color.FgHiBlack).Println("Your Personal Command Line Task Manager")
	fmt.Println(strings.Repeat("─", 40))
}

func printUsage() {
	printHeader()
	
	color.New(color.FgHiWhite, color.Bold).Println("\n🔍 Usage:")
	fmt.Println("  todo [command] [arguments]")

	color.New(color.FgHiWhite, color.Bold).Println("\n🛠  Available Commands:")
	cmdColor := color.New(color.FgHiGreen, color.Bold)
	descColor := color.New(color.FgHiWhite)
	
	cmdColor.Print("  -list")
	descColor.Println("                  Show all your todos")
	cmdColor.Print("  -add [todo text]")
	descColor.Println("       Create a new todo")
	cmdColor.Print("  -complete [number]")
	descColor.Println("     Mark a todo as done")
	cmdColor.Print("  -del [number]")
	descColor.Println("          Remove a todo")

	color.New(color.FgHiWhite, color.Bold).Println("\n💡 Examples:")
	color.HiBlue("  todo -add \"Buy groceries\"")
	color.HiBlue("  todo -list")
	color.HiBlue("  todo -complete 1")
	color.HiBlue("  todo -del 2")
	fmt.Println()
}

func printTodoList(todos *todo.Todos) {
	if len(*todos) == 0 {
		color.Yellow("\n📝 %s, your todo list is empty! Add one with: todo -add \"your todo\"", userName)
		return
	}

	color.New(color.FgHiCyan, color.Bold).Printf("\n📋 %s's Todos:\n", userName)
	fmt.Println(strings.Repeat("─", 40))

	// Count completed and pending todos
	completed := 0
	for _, t := range *todos {
		if t.Done {
			completed++
		}
	}

	for i, todo := range *todos {
		checkbox := "☐"
		statusColor := color.New(color.FgHiWhite)
		if todo.Done {
			checkbox = "✓"
			statusColor = color.New(color.FgHiGreen)
		}

		// Format the todo item
		fmt.Printf("%s ", checkbox)
		statusColor.Printf("[%d] %s", i+1, todo.Task)
		
		// Add creation time if available
		if !todo.CreatedAt.IsZero() {
			color.New(color.FgHiBlack).Printf(" (added %s)", todo.CreatedAt.Format("Jan 2"))
		}
		fmt.Println()
	}

	// Print summary
	fmt.Println(strings.Repeat("─", 40))
	color.New(color.FgHiBlack).Printf("Total: %d  ", len(*todos))
	color.New(color.FgHiGreen).Printf("Completed: %d  ", completed)
	color.New(color.FgHiYellow).Printf("Pending: %d\n", len(*todos)-completed)
	
	// Add motivational message based on completion status
	if completed == len(*todos) && len(*todos) > 0 {
		color.HiGreen("\n🎉 Amazing job %s! All tasks completed! 🌟", userName)
	} else if completed > 0 {
		color.HiCyan("\n💪 Keep going %s! You're making great progress!", userName)
	}
}

func main() {
	// Initialize username at the start
	getUserName()

	add := flag.Bool("add", false, "add a new todo")
	complete := flag.Int("complete", 0, "complete a todo")
	del := flag.Int("del", 0, "delete a todo")
	list := flag.Bool("list", false, "list all todos")

	flag.Parse()

	todos := &todo.Todos{}

	if err := todos.Load(todoFile); err != nil {
		color.Red("Error: %v", err)
		os.Exit(1)
	}

	switch {
	case *list:
		printTodoList(todos)

	case *add:
		if len(flag.Args()) == 0 {
			color.Red("⚠️  Error: No todo text provided")
			os.Exit(1)
		}
		todoText := strings.Join(flag.Args(), " ")
		todos.Add(todoText)
		if err := todos.Store(todoFile); err != nil {
			color.Red("Error: %v", err)
			os.Exit(1)
		}
		color.Green("✨ Added todo: %s", todoText)
		// Show the updated list after adding
		printTodoList(todos)

	case *complete > 0:
		if err := todos.Complete(*complete); err != nil {
			color.Red("Error: %v", err)
			os.Exit(1)
		}
		if err := todos.Store(todoFile); err != nil {
			color.Red("Error: %v", err)
			os.Exit(1)
		}
		color.Green("✅ Completed todo #%d", *complete)
		// Show the updated list after completing
		printTodoList(todos)

	case *del > 0:
		if err := todos.Delete(*del); err != nil {
			color.Red("Error: %v", err)
			os.Exit(1)
		}
		if err := todos.Store(todoFile); err != nil {
			color.Red("Error: %v", err)
			os.Exit(1)
		}
		color.Yellow("🗑  Deleted todo #%d", *del)
		// Show the updated list after deleting
		printTodoList(todos)

	default:
		printUsage()
		os.Exit(0)
	}
}
