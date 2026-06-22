package tools

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// IconEntry holds a single icon name and its code point, preserving insertion order.
type IconEntry struct {
	Name      string
	CodePoint string
}

// FontAwesomeMapping holds icons grouped by mode, preserving order.
type FontAwesomeMapping struct {
	Modes []string
	Icons map[string][]IconEntry
}

func readLines(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func ScanMaterialDesignIconsFlutter(pubPath string, lock *PubspecLock) ([]IconEntry, error) {
	pkgDir, err := lock.PackageDir("flutter_material_design_icons")
	if err != nil {
		return nil, err
	}
	dartFile := filepath.Join(pubPath, pkgDir, "lib", "src", "icons.enum.dart")
	lines, err := readLines(dartFile)
	if err != nil {
		return nil, fmt.Errorf("reading MDI file: %w", err)
	}

	// The new package declares each icon as a multi-line block:
	//   static const IconData abTesting = IconData(
	//     983497,
	//     fontFamily: 'Material Design Icons',
	//     fontPackage: 'flutter_material_design_icons',
	//   );
	// The header line names the icon (and may be preceded by @Deprecated / doc
	// comments, which we ignore); the next non-empty line is the decimal codepoint.
	// Dart reserved words (null, switch) are declared with a trailing "$"
	// (e.g. "null$", "switch$"); strip it so the kebab key matches the bare word.
	headerRegex := regexp.MustCompile(`^static const IconData (\w+)\$? = IconData\($`)
	seen := make(map[string]bool)
	var result []IconEntry
	var pendingName string

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if pendingName != "" {
			if line == "" {
				continue
			}
			codePoint := strings.TrimSpace(strings.TrimSuffix(line, ","))
			name := pendingName
			pendingName = ""

			if seen[name] {
				fmt.Printf("Material Design Icons: Duplicated icon name: %s\n", name)
				continue
			}
			seen[name] = true
			result = append(result, IconEntry{Name: name, CodePoint: codePoint})
			continue
		}

		matches := headerRegex.FindStringSubmatch(line)
		if matches == nil {
			continue
		}
		pendingName = ToKebabCase(matches[1])
	}

	fmt.Printf("Mapped %d icons from Material Design Icons\n", len(result))
	return result, nil
}

func ScanFontAwesomeFlutter(pubPath string, lock *PubspecLock) (*FontAwesomeMapping, error) {
	pkgDir, err := lock.PackageDir("font_awesome_flutter")
	if err != nil {
		return nil, err
	}
	dartFile := filepath.Join(pubPath, pkgDir, "lib", "font_awesome_flutter.dart")
	lines, err := readLines(dartFile)
	if err != nil {
		return nil, fmt.Errorf("reading Font Awesome file: %w", err)
	}

	// font_awesome_flutter v11 declares each icon as a nested block. The mode is
	// carried by the fontFamily string ('FontAwesome<Mode>'), not the constructor:
	//   static const FaIconData github = FaIconData(
	//     IconData(
	//       0xf09b,
	//       fontFamily: 'FontAwesomeBrands',
	//       fontPackage: 'font_awesome_flutter',
	//     ),
	//   );
	// The free package ships only Brands, Regular and Solid.
	validModes := map[string]bool{
		"Brands": true, "Regular": true, "Solid": true,
	}

	headerRegex := regexp.MustCompile(`^static const FaIconData (\w+) = FaIconData\($`)
	codePointRegex := regexp.MustCompile(`^(0x[0-9a-fA-F]+),?$`)
	familyRegex := regexp.MustCompile(`^fontFamily: 'FontAwesome(\w+)',?$`)

	result := &FontAwesomeMapping{
		Icons: make(map[string][]IconEntry),
	}
	seenPerMode := make(map[string]map[string]bool)
	modeOrder := make(map[string]bool)

	totalCount := 0
	var pendingName, pendingCodePoint string

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if matches := headerRegex.FindStringSubmatch(line); matches != nil {
			pendingName = ToKebabCase(matches[1])
			pendingCodePoint = ""
			continue
		}

		if pendingName == "" {
			continue
		}

		if pendingCodePoint == "" {
			if matches := codePointRegex.FindStringSubmatch(line); matches != nil {
				pendingCodePoint = matches[1]
			}
			continue
		}

		matches := familyRegex.FindStringSubmatch(line)
		if matches == nil {
			continue
		}
		mode := matches[1]
		name := pendingName
		codePoint := pendingCodePoint
		pendingName = ""
		pendingCodePoint = ""

		if !validModes[mode] {
			continue
		}

		if seenPerMode[mode] == nil {
			seenPerMode[mode] = make(map[string]bool)
		}
		if seenPerMode[mode][name] {
			fmt.Printf("Font Awesome Flutter: Duplicated icon name: %s\n", name)
			continue
		}
		seenPerMode[mode][name] = true

		if !modeOrder[mode] {
			modeOrder[mode] = true
			result.Modes = append(result.Modes, mode)
		}
		result.Icons[mode] = append(result.Icons[mode], IconEntry{Name: name, CodePoint: codePoint})
		totalCount++
	}

	fmt.Printf("Mapped %d icons from Font Awesome Flutter\n", totalCount)
	return result, nil
}

func ScanFluttySolarIcons(pubPath string, lock *PubspecLock) ([]IconEntry, error) {
	pkgDir, err := lock.PackageDir("flutty_solar_icons")
	if err != nil {
		return nil, err
	}
	dartFile := filepath.Join(pubPath, pkgDir, "lib", "src", "solar_icons.dart")
	lines, err := readLines(dartFile)
	if err != nil {
		return nil, fmt.Errorf("reading Flutty Solar Icons file: %w", err)
	}

	seen := make(map[string]bool)
	var result []IconEntry

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if !strings.HasPrefix(line, "static const") {
			continue
		}
		line = strings.TrimSpace(strings.Replace(line, "static const", "", 1))
		parts := strings.SplitN(line, "=", 2)
		if len(parts) < 2 {
			continue
		}
		name := ToKebabCase(strings.TrimSpace(parts[0]))
		codePointFull := strings.TrimSuffix(strings.TrimSpace(parts[1]), ";")
		openParen := strings.Index(codePointFull, "(")
		closeParen := strings.Index(codePointFull, ")")
		if openParen < 0 || closeParen < 0 {
			continue
		}
		codePoint := strings.TrimSpace(codePointFull[openParen+1 : closeParen])

		if seen[name] {
			fmt.Printf("Flutty Solar Icons: Duplicated icon name: %s\n", name)
			continue
		}
		seen[name] = true
		result = append(result, IconEntry{Name: name, CodePoint: codePoint})
	}

	fmt.Printf("Mapped %d icons from Flutty Solar Icons\n", len(result))
	return result, nil
}
