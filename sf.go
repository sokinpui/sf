package sf

import "runtime"

func Run(roots []string, fileType string, excludes []string, showHidden bool) []string {
	if len(roots) == 0 {
		roots = []string{"."}
	}

	engine := NewEngine(runtime.NumCPU()*2, fileType, excludes, showHidden)
	resultsChan := make(chan string, 100)

	go engine.Walk(roots, resultsChan)

	var paths []string
	for path := range resultsChan {
		paths = append(paths, path)
	}

	return paths
}
