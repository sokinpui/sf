package main

import (
	"os"
	"path/filepath"
	"sync"
)

type Engine struct {
	matcher     *Matcher
	concurrency chan struct{}
	wg          sync.WaitGroup
}

func NewEngine(root string, maxConcurrency int) *Engine {
	return &Engine{
		matcher:     NewMatcher(root),
		concurrency: make(chan struct{}, maxConcurrency),
	}
}

func (e *Engine) Walk(root string, results chan<- string) {
	e.wg.Add(1)
	go e.walkDir(root, results)

	go func() {
		e.wg.Wait()
		close(results)
	}()
}

func (e *Engine) walkDir(path string, results chan<- string) {
	defer e.wg.Done()

	entries, err := os.ReadDir(path)
	if err != nil {
		return
	}

	for _, entry := range entries {
		fullPath := filepath.Join(path, entry.Name())

		if e.matcher.ShouldSkip(fullPath, entry) {
			continue
		}

		results <- fullPath

		if entry.IsDir() {
			e.spawnWorker(fullPath, results)
		}
	}
}

func (e *Engine) spawnWorker(path string, results chan<- string) {
	e.wg.Add(1)
	select {
	case e.concurrency <- struct{}{}:
		go func() {
			e.walkDir(path, results)
			<-e.concurrency
		}()
	default:
		e.walkDir(path, results)
	}
}
