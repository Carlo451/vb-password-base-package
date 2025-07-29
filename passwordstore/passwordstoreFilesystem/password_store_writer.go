package passwordstoreFilesystem

import (
	"log"
	"os"
)

func WriteFileContents(file File) {
	file.SetContent(file.GetContent() + "\n")
	err := os.WriteFile(file.GetAbsolutePath(), []byte(file.GetContent()), 0644)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Default().Println("File: " + file.GetAbsolutePath() + " written successfully")
	}

}

func WriteDirectory(dir Directory) {
	err := os.Mkdir(dir.GetAbsoluteDirectoryPath(), 0775)
	if err != nil {
		log.Default().Println("Directory: " + dir.GetAbsoluteDirectoryPath() + " could not be created because " + err.Error())
	} else {
		log.Default().Println("Directory: " + dir.GetAbsoluteDirectoryPath() + " written successfully")
	}

}
