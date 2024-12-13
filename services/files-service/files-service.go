package filesservice

import (
	iofs "io/fs"
	"os"
)

type NodeType int

const (
	Directory NodeType = iota
	File
)

type Node struct {
	Name     string   `json:"name"`
	Type     NodeType `json:"type"`
	Children []*Node  `json:"children,omitempty"`
}

type FilesService struct {
	Directory string
}

func NewFilesService(rootPath string) *FilesService {
	return &FilesService{
		Directory: rootPath,
	}
}

/*
 * File operations
 * - CreateFile
 * - ReadFileContent
 * - DeleteFile
 * - writeFileContent
 */

func (fs *FilesService) CreateFile(relPath string, data []byte) (bool, error) {
	// create a file path from the relative path
	path := fs.Directory + "/" + relPath

	// write file content to disk
	err := os.WriteFile(path, data, 0644)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (fs *FilesService) ReadFileContent(relPath string) (bool, []byte, error) {
	// create a file path from the relative path
	path := fs.Directory + "/" + relPath

	_fs := os.DirFS(fs.Directory)

	// read file content from disk
	data, err := iofs.ReadFile(_fs, path)
	if err != nil {
		return false, nil, err
	}

	return true, data, nil

}

func (fs *FilesService) WriteFileContent(relPath string, data []byte) (bool, error) {
	// create a file path from the relative path
	path := fs.Directory + "/" + relPath

	// open file for writing
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return false, err
	}

	// write file content to disk
	_, err = file.Write(data)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (fs *FilesService) DeleteFile(relPath string) (bool, error) {
	// create a file path from the relative path
	path := fs.Directory + "/" + relPath

	// delete file from disk
	err := os.Remove(path)
	if err != nil {
		return false, err
	}

	return true, nil
}

/*
 * Directory operations
 * - CreateDirectory
 * - ListDirectory (returns a the children of the directory)
 * - DeleteDirectory
 */

func (fs *FilesService) CreateDirectory(relPath string) (bool, error) {
	// create a directory path from the relative path
	path := fs.Directory + "/" + relPath

	// create directory on disk
	err := os.Mkdir(path, 075)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (fs *FilesService) ListDirectory(relPath string) {
	// create a directory path from the relative path
	path := fs.Directory + "/" + relPath


	// list directory contents
	

}

func (fs *FilesService) DeleteDirectory(relPath string) (bool, error) {
	// create a directory path from the relative path
	path := fs.Directory + "/" + relPath

	// delete directory from disk
	err := os.RemoveAll(path)
	if err != nil {
		return false, err
	}

	return true, nil
}
