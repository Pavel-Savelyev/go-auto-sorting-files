package scan_folder

import (
	"os"
	"path"
)

type FileList struct {
	Files []file
}

type file struct {
	Name      string
	Extension string
	IsDir     bool
}

func ScanFolder(pathFolder string) (FileList, error) {
	data, err := os.ReadDir(pathFolder)
	if err != nil {
		return FileList{}, err
	}

	files := make([]file, len(data))

	for index, value := range data {
		files[index] = newFile(value.Name(), path.Ext(value.Name()), value.IsDir())
	}

	return FileList{Files: files}, nil
}

func newFile(name string, extension string, isDir bool) file {
	return file{
		Name:      name,
		Extension: extension,
		IsDir:     isDir,
	}
}
