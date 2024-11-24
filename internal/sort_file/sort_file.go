package sort_file

import (
	"fmt"
	"go-auto-sorting-files/internal/config"
	"go-auto-sorting-files/internal/scan_folder"
	"os"
	"path"
)

func Sort(config config.AppConfig, fileList scan_folder.FileList) error {
	for _, file := range fileList.Files {
		for _, rule := range config.Rules {
			if file.Extension == rule.Extension {
				if err := moveFile(path.Join(config.Path, file.Name), path.Join(rule.Directory, file.Name)); err != nil {
					return err
				}
				fmt.Printf("File :%v moving \n\r", file.Name)
				break
			}
		}
	}
	return nil
}

func moveFile(sourcePath string, destPath string) error {
	err := os.Rename(sourcePath, destPath)
	if err != nil {
		return err
	}
	return nil
}
