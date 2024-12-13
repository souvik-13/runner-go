package main

import (
	filewatcher "github.com/souvik-13/runner-go/services/file-watcher"
	filesservice "github.com/souvik-13/runner-go/services/files-service"
	terminalservice "github.com/souvik-13/runner-go/services/terminal-service"
)

type ServiceManager struct {
	Terminalservice *terminalservice.TerminalService
	Filesservice    *filesservice.FilesService
	Filewatcher     *filewatcher.FileWatcher

	// channels

}

// NewServiceManager creates a new ServiceManager
func NewServiceManager() *ServiceManager {
	return &ServiceManager{
		Terminalservice: terminalservice.NewTerminalService(),
		Filesservice:    filesservice.NewFilesService("/home/souvik/WorkDir/Webdev/runner-go/workspace"),
		Filewatcher:     filewatcher.NewFileWatcher(),
	}
}
