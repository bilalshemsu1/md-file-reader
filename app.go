package main

import (
	"context"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx         context.Context
	startupFile string
	watcher     *fsnotify.Watcher
	watchMap    map[string]bool
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		watchMap: make(map[string]bool),
	}
}

// startup is called when the app starts
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// shutdown is called when the app closes
func (a *App) shutdown(ctx context.Context) {
	// close watcher cleanly when app exits
	if a.watcher != nil {
		a.watcher.Close()
		a.watcher = nil
	}
}

// StartWatching starts watching a file for changes
func (a *App) StartWatching(path string) error {
	if a.watchMap[path] {
		return nil
	}

	if a.watcher == nil {
		w, err := fsnotify.NewWatcher()
		if err != nil {
			return err
		}
		a.watcher = w

		// start background goroutine to listen for events
		go func() {
			for {
				select {
				case event, ok := <-a.watcher.Events:
					if !ok {
						return
					}
					// only care about write events
					if event.Has(fsnotify.Write) {
						runtime.EventsEmit(a.ctx, "fileChanged", event.Name)
					}
				case err, ok := <-a.watcher.Errors:
					if !ok {
						return
					}
					_ = err
				}
			}
		}()
	}

	// add file to watcher
	err := a.watcher.Add(path)
	if err != nil {
		return err
	}

	// mark as watched
	a.watchMap[path] = true
	return nil
}

// StopWatching stops watching a file for changes
func (a *App) StopWatching(path string) error {
	// not watching this file? skip
	if !a.watchMap[path] {
		return nil
	}

	// remove from watcher
	if a.watcher != nil {
		err := a.watcher.Remove(path)
		if err != nil {
			return err
		}
	}

	// remove from watchMap
	delete(a.watchMap, path)

	// if no more files watched → close watcher entirely
	if len(a.watchMap) == 0 && a.watcher != nil {
		a.watcher.Close()
		a.watcher = nil
	}

	return nil
}

func (a *App) GetStartupFile() string {
	return a.startupFile
}

// OpenFile opens a native file dialog and returns the selected file path
func (a *App) OpenFile() string {
	path, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Open Markdown File",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Markdown Files",
				Pattern:     "*.md;*.markdown;*.txt;*.mdx",
			},
		},
	})
	if err != nil {
		return ""
	}
	return path
}

// ReadFile reads a file from disk and returns its content
func (a *App) ReadFile(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// SaveFile saves content to an existing file path
func (a *App) SaveFile(path string, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}

// SaveFileAs opens a native save dialog and saves content to chosen path
func (a *App) SaveFileAs(content string) (string, error) {
	path, err := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		Title:           "Save Markdown File",
		DefaultFilename: "untitled.md",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Markdown Files",
				Pattern:     "*.md;*.markdown",
			},
		},
	})
	if err != nil || path == "" {
		return "", err
	}
	err = os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		return "", err
	}
	return path, nil
}
