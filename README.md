# Tickli 📋✨

> A modern command line interface for TickTick task management

![Tickli Demo](assets/tickli-demo.gif)

## What is Tickli?

Tickli is a powerful CLI tool that lets you manage your TickTick tasks and projects directly from your terminal. Stay productive without leaving your command line!

## Features

- 🚀 Create and manage tasks right from your terminal
- 📂 Organize tasks into projects
- 📅 Set dates, priorities, and tags
- 🔄 Complete and uncomplete tasks
- 🔍 Filter and search your tasks
- 🔐 Secure OAuth authentication

## Installation

### Using Homebrew

```bash
brew tap sho0pi/brew
brew install tickli
```

### Download from Releases

You can also download prebuilt binaries from the [GitHub releases page](https://github.com/Sho0pi/tickli/releases).

## Quick Start

```bash
# Initialize and authenticate with TickTick
tickli init

# List available projects
tickli project list

# Switch to a project
tickli project use "Work Tasks"

# Add a new task
tickli task add "Finish documentation for project X"

# Add a high priority task due tomorrow
tickli task add "Important meeting" --priority high --date "tomorrow at 2pm"

# List your tasks
tickli task list

# Complete a task
tickli task complete <task-id>
```

## Key Commands

| Command                | Description                         |
| ---------------------- | ----------------------------------- |
| `tickli init`          | Set up authentication with TickTick |
| `tickli project list`  | Show all your projects              |
| `tickli project use`   | Switch active project context       |
| `tickli add`           | Quickly add a new task              |
| `tickli task list`     | List tasks in current project       |
| `tickli task show`     | View task details                   |
| `tickli task complete` | Mark a task as complete             |

## Interactive TUI Experience (Coming Soon!)

![Tickli TUI Demo](assets/tickli-tui-demo.gif)

## Roadmap 🗺️

- [x] Basic task management
- [x] Project management
- [x] Authentication
- [ ] Advanced date/time handling and timezone support
- [ ] Interactive modes for all commands
- [ ] Subtask management
- [ ] TUI interface with bubbletea
- [ ] Task filtering by multiple criteria
- [ ] Offline mode and syncing
- [ ] Custom views (Kanban, etc.)

## Documentation

For complete documentation on all available commands:

```bash
tickli --help
```

Or check out the [full documentation](docs/README.md).

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.