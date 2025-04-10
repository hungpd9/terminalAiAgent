# Terminal AI Agent

A terminal-based application that uses an AI Agent (e.g., Google Gemini API) to assist users in executing commands efficiently.

## Features
- Detects the user's operating system (Windows, Linux, macOS).
- Analyzes and executes terminal commands via AI.
- Stores command history in SQLite3.
- Beautiful TUI built with Bubbletea and Lipgloss.
- Configurable via `.env`.

## Prerequisites
- Go 1.21 or higher
- A Google Gemini API key (or similar AI API)

## Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/terminalAiAgent.git
   cd terminal-ai-agent
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Create a `.env` file based on `.env.example`:
   ```bash
   cp .env.example .env
   ```
   Edit `.env` and add your API key:
   ```
   GEMINI_API_KEY=your-api-key-here
   ```

4. Build and run:
   ```bash
   go run cmd/main.go
   ```

## Usage
- Type a command and press `Enter` to execute it via AI.
- View command history in the interface.
- Press `Ctrl+C` to exit.

## Contributing
Feel free to submit issues or pull requests!

## License
MIT
