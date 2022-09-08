package setting

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	extJSON = "json"
)

type AppSetting struct {
	DebugInfo              bool              `json:"debug_info"`
	UDPServer              SocketSetting     `json:"udpServer"`
	DisableCaptureMessage  bool              `json:"disable_capture_message"`
	EventMessageIgnoreKeys []string          `json:"event_message_ignore_keys"`
	EventMessages          []UDPEventSetting `json:"event_messages"`
}

type SocketSetting struct {
	Host      string `json:"host"`
	Port      int    `json:"port"`
	BlockSize int    `json:"blockSize"`
}

type UDPEventSetting struct {
	CaptureReg string `json:"captureReg"`
	Content    string `json:"content"`
	Value      string `json:"value"`
}

// LoadSetting -- Load settings from config file.
func LoadSetting(configPath string, settings *AppSetting) error {
	// get config file extension
	fileExt := strings.ToLower(filepath.Ext(configPath))
	if fileExt == "" {
		return errors.New("invalid file extension")
	}

	// get file name without extension
	fileName := strings.TrimSuffix(filepath.Base(configPath), fileExt)
	fileExt = fileExt[1:]

	if fileName == "" {
		return errors.New("invalid config file name")
	}

	// LoadSettings from config file
	content, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Println("Error occurs while reading config file, please make sure config file exists!")
		return err
	}

	switch fileExt {
	case extJSON:
		if err := json.Unmarshal(content, settings); err != nil {
			return err
		}
	default:
		return errors.New("invalid extension for config file:" + fileExt)
	}

	return initSetting(settings)
}

func initUDPServer(settings *AppSetting) error {
	if len(settings.UDPServer.Host) < 7 {
		settings.UDPServer.Host = "0.0.0.0"
	}
	if settings.UDPServer.Port == 0 {
		settings.UDPServer.Port = 996
	}
	if settings.UDPServer.BlockSize == 0 {
		settings.UDPServer.BlockSize = 1024
	}
	//TODO：Use reg check host and port.
	//return errors.New("invalid udpServer config.")
	return nil
}

// init here. Check methods are put the util  package is a best practices.
func initSetting(settings *AppSetting) error {
	if err := initUDPServer(settings); err != nil {
		return err
	}

	return loadSecretsFromFile(settings)
}

// read password from a special file .
func loadSecretsFromFile(settings *AppSetting) error {
	//var err error

	// if settings.Password, err = readSecretFromFile(
	// 	settings.PasswordFile,
	// 	settings.Password,
	// ); err != nil {
	// 	return fmt.Errorf("failed to load password from file: %w", err)
	// }

	return nil
}

// func readSecretFromFile(source, value string) (string, error) {
// 	if source == "" {
// 		return value, nil
// 	}

// 	content, err := os.ReadFile(source)

// 	if err != nil {
// 		return value, err
// 	}

// 	return strings.TrimSpace(string(content)), nil
// }
