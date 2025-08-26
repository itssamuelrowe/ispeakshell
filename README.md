# ispeakshell

A natural language to shell command translator powered by AI. Convert your plain English requests into safe, executable bash commands with interactive confirmation and rich terminal output.

## Features

- **AI-Powered**: Uses OpenAI GPT-3.5-turbo or local Ollama models
- **Rich UI**: Colored output with formatted tables and syntax highlighting
- **Safety First**: Interactive command confirmation before execution
- **Command History**: Save and track generated commands
- **Dual Mode**: Command-line arguments or interactive terminal mode
- **Template System**: Customizable prompt templates for different use cases

## Setup

### Prerequisites

- Go 1.18 or higher
- OpenAI API key (recommended) or Ollama for local inference

### Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd ispeakshell
```

2. Install dependencies:
```bash
go mod tidy
```

3. Build the application:
```bash
go build -o i
```

### OpenAI Setup (Recommended)

1. Get your OpenAI API key from [OpenAI Platform](https://platform.openai.com/api-keys)

2. Create a `.env` file in the project root:
```bash
echo "OPENAI_KEY=your_api_key_here" > .env
```

3. Run the application:
```bash
./i "list all files in current directory"
```

### Alternative: Ollama Setup (Local)

If you prefer to run models locally:

```bash
# Install Ollama
curl -fsSL https://ollama.com/install.sh | sh

# Pull a code model
ollama pull codegemma:7b

# Start Ollama server
ollama serve
```

## Usage

### Command Line Mode

Pass your request as command line arguments:

```bash
./i "find all Python files modified in the last 7 days"
./i "compress all log files in /var/log"
./i "show disk usage for each directory"
```

### Interactive Mode

Run without arguments for interactive mode:

```bash
./i
```

Then type your request and press Enter twice to execute.

## Example Usage

### Example 1: File Operations

**Input:**
```bash
./i "find all JavaScript files larger than 1MB"
```

**Output:**
```
Find JavaScript files larger than 1MB in current directory

#    Task                     Description
1    Find large JS files      Searches for JavaScript files (.js) larger than 1MB using find command with size filter

[1] find . -name "*.js" -size +1M -type f
Run this command? [y/n/save]: y
./src/bundle.js
./dist/app.min.js
```

### Example 2: System Information

**Input:**
```bash
./i "show me system memory usage and top 5 processes"
```

**Output:**
```
Display system memory usage and top resource-consuming processes

#    Task                     Description
1    Show memory usage        Displays current memory usage statistics
2    List top processes       Shows top 5 processes by memory consumption

[1] free -h
Run this command? [y/n/save]: y
               total        used        free      shared  buff/cache   available
Mem:            15Gi       8.2Gi       2.1Gi       1.2Gi       5.4Gi       5.8Gi
Swap:          2.0Gi          0B       2.0Gi

[2] ps aux --sort=-%mem | head -6
Run this command? [y/n/save]: y
USER         PID %CPU %MEM    VSZ   RSS TTY      STAT START   TIME COMMAND
chrome      1234  15.2  8.4 2847364 1347234 ?   Sl   09:15   2:45 /opt/google/chrome/chrome
code        5678   5.1  4.2 1654321 672345 ?    Sl   08:30   1:23 /usr/share/code/code
```

### Example 3: Git Operations

**Input:**
```bash
./i "show git status and recent commits"
```

**Output:**
```
Display current git repository status and recent commit history

#    Task                     Description
1    Check git status         Shows current working directory status
2    Show recent commits      Displays last 5 commits with oneline format

[1] git status
Run this command? [y/n/save]: y
On branch main
Your branch is up to date with 'origin/main'.
nothing to commit, working tree clean

[2] git log --oneline -5
Run this command? [y/n/save]: y
2435535 feat(ui): implement rich terminal interface
ad8b932 feat(core): implement AI integration and command execution
4f3f5a8 feat(prompts): add basic command generation template
```

## Command Confirmation Options

When prompted to run a command, you have three options:

- **y/yes**: Execute the command immediately
- **n/no**: Skip the command and continue
- **save**: Save the command to history without executing

## Command History

Commands are automatically saved to `~/.ispeakshell/history.json` when you choose the "save" option, allowing you to build a library of useful commands for future reference.

## Safety Features

- Interactive confirmation before executing any command
- Automatic quoting of file patterns (e.g., `"*.log"`)
- Interactive flags for destructive operations (e.g., `rm -i`)
- No sudo commands unless explicitly requested
- Scope validation to reject non-shell-related requests

## Configuration

### Environment Variables

- `OPENAI_KEY`: Your OpenAI API key for GPT-3.5-turbo
- `OLLAMA_HOST`: Ollama server host (default: 127.0.0.1:11434)

### Customizing Prompts

Edit `prompts/basic.template` to customize the AI prompt template for your specific needs.

## Contributing

Contributions are welcome! Please feel free to submit issues and pull requests.

## License

Copyright 2025 Samuel Rowe <samuelrowe1999@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at