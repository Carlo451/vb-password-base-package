package passwordstoreFilesystem

import (
	"os"
)

func ReadFiles(contentDir PasswordStoreContentDir) []File {
	contents := make([]File, 0)
	entrys, err := os.ReadDir(contentDir.GetAbsoluteDirectoryPath())
	if err != nil {
		panic(err)
	}
	for _, entry := range entrys {
		if entry.IsDir() {
			panic("directory " + entry.Name() + " is not a file")
		} else {
			emptyFile := NewCleanPasswordStoreContentFile(entry.Name(), contentDir)
			contents = append(contents, ReadFile(emptyFile))
		}
	}
	return contents
}

func ReadFile(file File) File {
	bytes, err := os.ReadFile(file.GetAbsolutePath())
	if err != nil {
		return nil
	}
	stringContent := string(bytes)
	file.SetContent(stringContent)
	return file
}

func ReadDir(dir Directory) Directory {
	entrys, err := os.ReadDir(dir.GetAbsoluteDirectoryPath())
	if err != nil {
		panic(err)
	}
	for _, entry := range entrys {
		if entry.IsDir() {
			dir.ReadDirectoryRec(entry)
		} else {

		}
	}
	return dir
}
