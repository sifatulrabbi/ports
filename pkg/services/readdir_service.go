package services

import (
	"errors"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

type Entry struct {
	Name string `json:"name"`
	Size int64  `json:"size"`
	Type string `json:"type"`
	Ext  string `json:"ext"`
}

// Get all the home directory's directories and files information.
func GetHomeDir() ([]Entry, error) {
	results := []Entry{}
	u, err := user.Current()
	if err != nil {
		return results, err
	}

	entries, err := os.ReadDir(u.HomeDir)
	for _, entry := range entries {
		info, _ := entry.Info()
		item := Entry{
			Name: entry.Name(),
			Size: info.Size(),
		}
		if info.IsDir() {
			item.Type = "directory"
		} else {
			item.Type = "file"
			item.Ext = filepath.Ext(u.HomeDir + "/" + item.Name)
		}
		results = append(results, item)
	}

	return results, err
}

func GetDirInfo(path string) ([]Entry, error) {
	results := []Entry{}
	u, err := user.Current()
	if err != nil {
		return results, err
	}

	p := u.HomeDir + "/" + path
	entries, err := os.ReadDir(p)
	// Incase the path points to a file send the file's information.
	if err != nil && strings.Contains(err.Error(), "not a directory") {
		file, _ := os.Stat(p)
		item := Entry{
			Name: file.Name(),
			Size: file.Size(),
			Type: "file",
			Ext:  filepath.Ext(p),
		}
		results = append(results, item)
		return results, nil
	}

	for _, entry := range entries {
		info, _ := entry.Info()
		item := Entry{
			Name: entry.Name(),
			Size: info.Size(),
		}
		if info.IsDir() {
			item.Type = "directory"
		} else {
			item.Type = "file"
			item.Ext = filepath.Ext(p + "/" + item.Name)
		}
		results = append(results, item)
	}

	return results, err
}

func GetFile(path string) ([]byte, error) {
	content := []byte("")
	u, err := user.Current()
	if err != nil {
		return content, err
	}
	fullPath := u.HomeDir + "/" + path
	info, err := os.Stat(fullPath)
	if err != nil {
		return content, err
	}
	if info.IsDir() {
		return content, errors.New("this is a directory")
	}
	content, err = os.ReadFile(fullPath)
	return content, err
}
