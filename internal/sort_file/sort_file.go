package sort_file

import (
	"fmt"
	"go-auto-sorting-files/internal/config"
	"go-auto-sorting-files/internal/scan_folder"
	"os"
	"path"
	"strings"
)

func Sort(folder config.Folder, fileList scan_folder.FileList) error {
	for _, file := range fileList.Files {
		if err := sortByNameContains(folder, file); err != nil {
			return err
		}
		if err := sortByExtension(folder, file); err != nil {
			return err
		}
	}
	return nil
}

func sortByExtension(folder config.Folder, file scan_folder.File) error {
	for _, rule := range folder.Rules {
		for _, extension := range rule.FileExtensions {
			if file.Extension == extension {
				if err := moveFile(path.Join(folder.Path, file.Name), path.Join(rule.Directory, file.Name)); err != nil {
					return err
				}
				fmt.Printf("File :%v moving by extension\n", file.Name)
				return nil
			}
		}
	}
	return nil
}

func sortByNameContains(folder config.Folder, file scan_folder.File) error {
	for _, rule := range folder.Rules {
		for _, keyword := range rule.NameContains {
			if strings.Contains(file.Name, keyword) {
				if err := moveFile(path.Join(folder.Path, file.Name), path.Join(rule.Directory, file.Name)); err != nil {
					return err
				}
				fmt.Printf("File :%v moving by name contains\n", file.Name)
				return nil
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
