# Todo CLI

A feature-rich command-line task management application written in Go. Keep track of your tasks, mark them as complete/incomplete, and manage them efficiently from your terminal.

## Features

- Create and manage tasks from the command line
- List all tasks, or filter by completion status
- Mark tasks as complete or incomplete
- Delete tasks when they're no longer needed
- Persistent storage in your home directory
- Simple and intuitive command structure

## Installation

1. Make sure you have Go installed on your system
2. Clone the repository:
   ```bash
   git clone https://github.com/ammaarM/todo-cli.git
   ```
3. Build the application:
   ```bash
   cd todo-cli
   go build
   ```
4. Optionally, move the binary to your PATH to use it from anywhere

## Usage

```bash
todoctl [command] [arguments]
```

### Available Commands

- `list [filter]` - Display all tasks
  - `list` - Show all tasks
  - `list completed` - Show only completed tasks
  - `list uncompleted` - Show only uncompleted tasks
- `add <task name>` - Add a new task
- `complete <task id>` - Mark a task as completed
- `uncomplete <task id>` - Mark a task as not completed
- `delete <task id>` - Delete a task
- `--help, -h` - Show help message

### Examples

```bash
# Add a new task
todoctl add "Complete the project documentation"

# List all tasks
todoctl list

# List only completed tasks
todoctl list completed

# Mark a task as complete
todoctl complete 1

# Mark a task as incomplete
todoctl uncomplete 1

# Delete a task
todoctl delete 1
```

## Data Storage

Tasks are stored in a JSON file located in your home directory at `~/.todo-cli/tasks.json`. The application will automatically create this directory and file if they don't exist.

## Task Structure

Each task contains the following information:
- ID: Unique identifier for the task
- Name: The task description
- Type: Task type (if specified)
- Completed: Boolean status of completion
- DateStarted: Timestamp when the task was created
- DateCompleted: Timestamp when the task was marked as complete

## Development

The project is structured as follows:
```
todo-cli/
├── cmd/
│   ├── commands.go    # Command handling
│   └── handlers.go    # Command implementations
├── task/
│   └── tasks.go       # Task data structure and storage
├── main.go           # Application entry point
└── README.md         # This file
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.