package pathparser

import (
	"os"
	"strings"
)

type ParsedPath struct {
	SubDirectories   []string
	ContentDirectory string
}

func (p *ParsedPath) BuildPathWithoutContentDir() string {
	path := ""
	for _, subDir := range p.SubDirectories {
		path += subDir + "/"
	}
	return strings.TrimRight(path, "/")
}

func (p *ParsedPath) BuildPathCompletePath() string {
	path := ""
	for _, subDir := range p.SubDirectories {
		path += subDir + "/"
	}
	return path + p.ContentDirectory
}

func ParsePathWithContentDirectory(storePath, path string) ParsedPath {
	allDirectories := GetAllSubDirsOrdered(storePath, path)
	allDirectories = allDirectories[1:]
	contentDir := allDirectories[len(allDirectories)-1]
	subDirectories := allDirectories[:len(allDirectories)-1]
	return ParsedPath{subDirectories, contentDir}
}

func GetAllSubDirsOrdered(storePath, path string) []string {
	_, err := os.ReadDir(storePath)
	if err != nil {
		panic("Store does not exist")
	}
	dirSubPath, _ := strings.CutPrefix(path, storePath)
	return strings.Split(dirSubPath, "/")
}
