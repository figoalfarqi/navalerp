package middleware

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

type ErrorLog struct {
	Timestamp string        `json:"timestamp"`
	Method    string        `json:"method"`
	Path      string        `json:"path"`
	Status    int           `json:"status"`
	Duration  time.Duration `json:"duration_ms"`
	IP        string        `json:"ip"`
}

func GetMonthlyLogFile(filename string) (string, error) {
	now := time.Now()
	dir := filepath.Join(
		"public",
		"log",
		now.Format("2006-01"),
	)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}

	return filepath.Join(dir, "error.json"), nil
}

func WriteErrorLog(entry ErrorLog) error {
	filePath, err := GetMonthlyLogFile("error.json")
	if err != nil {
		return err
	}
	file, err := os.OpenFile(filePath,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0644,
	)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	return encoder.Encode(entry) // 1 JSON per baris (JSONL)
}

func WriteSlowLog(entry ErrorLog) error {
	filePath, err := GetMonthlyLogFile("slow.json")
	if err != nil {
		return err
	}
	file, err := os.OpenFile(filePath,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0644,
	)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	return encoder.Encode(entry) // 1 JSON per baris (JSONL)
}
