package commands

import (
    "os/exec"
    "runtime"
    "strings"
)

type Executor struct{}

func NewExecutor() *Executor {
    return &Executor{}
}

func (e *Executor) Execute(cmd string) (string, error) {
    var shell, flag string
    switch runtime.GOOS {
    case "windows":
        shell = "cmd"
        flag = "/C"
    default: // linux, macos
        shell = "bash"
        flag = "-c"
    }

    out, err := exec.Command(shell, flag, cmd).CombinedOutput()
    if err != nil {
        return string(out), err
    }
    return string(out), nil
}