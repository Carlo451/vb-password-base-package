package passwordstoreFilesystem

import (
	"os"
	"strings"
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

func ReadDirDownFromPath(path string) PasswordStoreDir {
	entrys, err := os.ReadDir(path)
	dirPathSplit := strings.Split(path, "/")
	dirName := dirPathSplit[len(dirPathSplit)-1]
	dirPathSplit = dirPathSplit[:len(dirPathSplit)-1]
	dir := CreateNewEmptyStoreDir(dirName, strings.Join(dirPathSplit, "/"))
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
