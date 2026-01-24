package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/monochromegane/go-gitignore"
)

type Matcher struct {
	gitIgnore gitignore.IgnoreMatcher
	excludes  []string
	root      string
}

func NewMatcher(root string, excludes []string) *Matcher {
	baseDir := root
	if info, err := os.Stat(root); err == nil && !info.IsDir() {
		baseDir = filepath.Dir(root)
	}

	ignorePath := filepath.Join(baseDir, ".gitignore")

	gitIgnore, _ := gitignore.NewGitIgnore(ignorePath, baseDir)

	return &Matcher{
		gitIgnore: gitIgnore,
		excludes:  excludes,
		root:      root,
	}
}

func (m *Matcher) ShouldSkip(path string, info os.DirEntry) bool {
	if m.isHidden(info.Name()) {
		return true
	}

	if m.isGitDir(info) {
		return true
	}

	if m.gitIgnore != nil && m.gitIgnore.Match(path, info.IsDir()) {
		return true
	}

	if m.matchExcludes(path, info) {
		return true
	}

	return false
}

func (m *Matcher) matchExcludes(path string, info os.DirEntry) bool {
	if len(m.excludes) == 0 {
		return false
	}

	rel, err := filepath.Rel(m.root, path)
	if err != nil {
		rel = path
	}

	for _, pattern := range m.excludes {
		if matched, _ := filepath.Match(pattern, rel); matched {
			return true
		}

		if matched, _ := filepath.Match(pattern, info.Name()); matched {
			return true
		}
	}

	return false
}

func (m *Matcher) isHidden(name string) bool {
	return len(name) > 1 && strings.HasPrefix(name, ".")
}

func (m *Matcher) isGitDir(info os.DirEntry) bool {
	return info.IsDir() && info.Name() == ".git"
}
