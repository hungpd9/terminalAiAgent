package ui

import (
    "fmt"
    "os"
    "terminal-ai-agent/internal/ai"
    "terminal-ai-agent/internal/commands"
    "terminal-ai-agent/internal/history"
    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"
)

type model struct {
    input    string
    output   string
    history  *history.History
    executor *commands.Executor
    aiClient *ai.GeminiClient
}

var style = lipgloss.NewStyle().
    BorderStyle(lipgloss.RoundedBorder()).
    BorderForeground(lipgloss.Color("63")).
    Padding(1, 2)

func NewProgram() *tea.Program {
    h, err := history.NewHistory("history.bolt")
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing history: %v\n", err)
        os.Exit(1)
    }
    return tea.NewProgram(model{
        history:  h,
        executor: commands.NewExecutor(),
        aiClient: ai.NewGeminiClient(),
    })
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.Type {
        case tea.KeyCtrlC:
            m.history.Close()
            return m, tea.Quit
        case tea.KeyEnter:
            m.history.Add(m.input)
            response, err := m.aiClient.AnalyzeCommand(m.input)
            if err != nil {
                m.output = fmt.Sprintf("AI Error: %v", err)
                return m, nil
            }
            out, err := m.executor.Execute(response)
            if err != nil {
                m.output = fmt.Sprintf("Exec Error: %v\nOutput: %s", err, out)
            } else {
                m.output = out
            }
            m.input = ""
            return m, nil
        case tea.KeyRunes:
            m.input += string(msg.Runes)
            return m, nil
        case tea.KeyBackspace:
            if len(m.input) > 0 {
                m.input = m.input[:len(m.input)-1]
            }
            return m, nil
        }
    }
    return m, nil
}

func (m model) View() string {
    commands, _ := m.history.GetAll()
    return style.Render(fmt.Sprintf(
        "Terminal AI Agent\n\nInput: %s\nOutput: %s\n\nHistory: %v\n\nPress Ctrl+C to quit",
        m.input, m.output, commands,
    ))
}