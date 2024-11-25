package main

import (
    "fmt"
    "sync"
)

type Logger struct{}

var instance *Logger
var once sync.Once

func GetInstance() *Logger {
    once.Do(func() {
        instance = &Logger{}
    })
    return instance
}

func (l *Logger) Log(message string) {
    fmt.Println(message)
}

func main() {
    logger1 := GetInstance()
    logger2 := GetInstance()

    if logger1 == logger2 {
        fmt.Println("Both instances are the same.")
    } else {
        fmt.Println("Instances are different.")
    }

    logger1.Log("Logging a message")
    logger2.Log("Logging another message")
}