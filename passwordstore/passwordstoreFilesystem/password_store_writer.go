package passwordstoreFilesystem

import (
	"log"
	"os"
)

func WriteFileContents(file File) {
	file.SetContent(file.GetContent() + "\n")
	f, err := os.OpenFile(file.GetAbsolutePath(), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err = f.WriteString(file.GetContent())
	if err != nil {
		log.Fatal(err)
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

func RemoveDirectory(dir Directory) {
	err := os.RemoveAll(dir.GetAbsoluteDirectoryPath())
	if err != nil {
		log.Default().Println("Directory: " + dir.GetAbsoluteDirectoryPath() + " could not be created because " + err.Error())
	} else {
		log.Default().Println("Directory: " + dir.GetAbsoluteDirectoryPath() + " written successfully")
	}
}
