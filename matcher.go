package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/monochromegane/go-gitignore"
)

type Matcher struct {
	gitIgnore gitignore.IgnoreMatcher
}

func NewMatcher(root string) *Matcher {
	ignorePath := filepath.Join(root, ".gitignore")

	gitIgnore, _ := gitignore.NewGitIgnore(ignorePath, root)

	return &Matcher{
		gitIgnore: gitIgnore,
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

	return false
}

func (m *Matcher) isHidden(name string) bool {
	return len(name) > 1 && strings.HasPrefix(name, ".")
}

func (m *Matcher) isGitDir(info os.DirEntry) bool {
	return info.IsDir() && info.Name() == ".git"
}
