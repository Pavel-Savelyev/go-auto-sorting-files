package sort_file

import (
	"fmt"
	"go-auto-sorting-files/internal/config"
	"go-auto-sorting-files/internal/scan_folder"
	"os"
	"path"
)

func Sort(folder config.Folder, fileList scan_folder.FileList) error {
	for _, file := range fileList.Files {
		for _, rule := range folder.Rules {
			for _, extension := range rule.FileExtensions {
				if file.Extension == extension {
					if err := moveFile(path.Join(folder.Path, file.Name), path.Join(rule.Directory, file.Name)); err != nil {
						return err
					}
					fmt.Printf("File :%v moving \n\r", file.Name)
					break
				}
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
