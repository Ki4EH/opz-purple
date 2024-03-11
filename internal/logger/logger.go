package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

type Logger interface {
	Info(message string)
	Error(message string)
}

type FileLogger struct {
	file *os.File
}

// NewFileLogger Создаем лог файл на текущий день, куда будем записывать
// всю информацию за день
func NewFileLogger() (*FileLogger, error) {
	dir, _ := os.Getwd()
	logDir := filepath.Join(dir, "logs")

	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err = os.Mkdir(logDir, os.ModeDir)
		if err != nil {
			return nil, fmt.Errorf("error creating dir")
		}
	}

	today := time.Now().Format("2006-01-02")
	filePath := filepath.Join(logDir, today+".log")
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	return &FileLogger{file: file}, nil
}

// Info Добавляет отметку INFO в лог команды
func (l *FileLogger) Info(message string) {
	log.SetOutput(l.file)
	log.Println("[INFO]: ", message)
}

// Error Добавляет отметку ERROR в лог команды
func (l *FileLogger) Error(message string) {
	log.SetOutput(l.file)
	log.Println("[ERROR]: ", message)
}

// Close сохраняет лог и закрывает сервер
func (l *FileLogger) Close() {
	log.SetOutput(l.file)
	log.Println("[INFO]: [SERVER] Server closed")
	_ = l.file.Close()
}
