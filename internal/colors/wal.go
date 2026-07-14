package colors

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type PywalColors struct {
	Special struct {
		Background string `json:"background"`
		Foreground string `json:"foreground"`
		Cursor     string `json:"cursor"`
	} `json:"special"`
	Colors map[string]string `json:"colors"`
}

func GetPywalColors() PywalColors {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	walPath := filepath.Join(home, ".cache", "wal", "colors.json")

	file, err := os.ReadFile(walPath)
	if err != nil {
		panic(err)
	}

	var theme PywalColors
	if err := json.Unmarshal(file, &theme); err != nil {
		panic(err)
	}
	return theme
}
