package tools

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// PubspecLock holds resolved package versions from pubspec.lock.
type PubspecLock struct {
	Versions map[string]string // package name -> resolved version
}

// ParsePubspecLock parses a pubspec.lock file and extracts package versions.
func ParsePubspecLock(baseDir string) (*PubspecLock, error) {
	lockFile := filepath.Join(baseDir, "pubspec.lock")
	f, err := os.Open(lockFile)
	if err != nil {
		return nil, fmt.Errorf("opening pubspec.lock: %w", err)
	}
	defer f.Close()

	result := &PubspecLock{Versions: make(map[string]string)}
	scanner := bufio.NewScanner(f)

	var currentPackage string
	for scanner.Scan() {
		line := scanner.Text()

		// Package names are at 2-space indent, followed by ":"
		if len(line) > 2 && line[0] == ' ' && line[1] == ' ' && line[2] != ' ' && strings.HasSuffix(strings.TrimSpace(line), ":") {
			currentPackage = strings.TrimSpace(strings.TrimSuffix(strings.TrimSpace(line), ":"))
			continue
		}

		// Version lines are at 4-space indent: '    version: "x.y.z"'
		if currentPackage != "" && strings.HasPrefix(line, "    version:") {
			version := strings.TrimSpace(strings.TrimPrefix(line, "    version:"))
			version = strings.Trim(version, "\"")
			result.Versions[currentPackage] = version
			currentPackage = ""
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("reading pubspec.lock: %w", err)
	}

	return result, nil
}

// PackageDir returns the pub cache directory name for a package (e.g. "font_awesome_flutter-10.12.0").
func (p *PubspecLock) PackageDir(packageName string) (string, error) {
	version, ok := p.Versions[packageName]
	if !ok {
		return "", fmt.Errorf("package %q not found in pubspec.lock", packageName)
	}
	return packageName + "-" + version, nil
}
