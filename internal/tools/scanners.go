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
	pkgDir, err := lock.PackageDir("material_design_icons_flutter")
	if err != nil {
		return nil, err
	}
	dartFile := filepath.Join(pubPath, pkgDir, "lib", "icon_map.dart")
	lines, err := readLines(dartFile)
	if err != nil {
		return nil, fmt.Errorf("reading MDI file: %w", err)
	}

	searchRegex := regexp.MustCompile(`^'(\w+)': _MdiIconData\(`)
	seen := make(map[string]bool)
	var result []IconEntry

	for _, line := range lines {
		line = strings.TrimSpace(line)
		matches := searchRegex.FindStringSubmatch(line)
		if matches == nil {
			continue
		}
		codePoint := strings.Split(strings.Split(line, "_MdiIconData(")[1], ")")[0]
		name := ToKebabCase(matches[1])

		if seen[name] {
			fmt.Printf("Material Design Icons Flutter: Duplicated icon name: %s\n", name)
			continue
		}
		seen[name] = true
		result = append(result, IconEntry{Name: name, CodePoint: codePoint})
	}

	fmt.Printf("Mapped %d icons from Material Design Icons Flutter\n", len(result))
	return result, nil
}

func scanSolarIcons(pubPath string, lock *PubspecLock, variant, dartFileName string) ([]IconEntry, error) {
	pkgDir, err := lock.PackageDir("solar_icons")
	if err != nil {
		return nil, err
	}
	dartFile := filepath.Join(pubPath, pkgDir, "lib", "src", dartFileName)
	lines, err := readLines(dartFile)
	if err != nil {
		return nil, fmt.Errorf("reading Solar Icons %s file: %w", variant, err)
	}

	seen := make(map[string]bool)
	var result []IconEntry
	var buffer string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "static const SolarIconsData") {
			buffer = line
			continue
		}
		if buffer != "" {
			completeLine := buffer + line
			buffer = ""
			completeLine = strings.TrimSpace(strings.Replace(completeLine, "static const SolarIconsData", "", 1))
			name := strings.TrimSpace(strings.Split(completeLine, "=")[0])
			name = ToKebabCase(name)
			codePointPart := strings.Split(completeLine, "=")[1]
			codePoint := strings.TrimSpace(strings.Split(strings.Replace(codePointPart, "SolarIconsData(", "", 1), ",")[0])

			if seen[name] {
				fmt.Printf("Solar Icons %s: Duplicated icon name: %s\n", variant, name)
				continue
			}
			seen[name] = true
			result = append(result, IconEntry{Name: name, CodePoint: codePoint})
		}
	}

	fmt.Printf("Mapped %d icons from Solar Icons %s\n", len(result), variant)
	return result, nil
}

func ScanSolarIconsBold(pubPath string, lock *PubspecLock) ([]IconEntry, error) {
	return scanSolarIcons(pubPath, lock, "Bold", "solar_icons_bold.dart")
}

func ScanSolarIconsOutline(pubPath string, lock *PubspecLock) ([]IconEntry, error) {
	return scanSolarIcons(pubPath, lock, "Outline", "solar_icons_outline.dart")
}

func ScanSolarIconsBroken(pubPath string, lock *PubspecLock) ([]IconEntry, error) {
	return scanSolarIcons(pubPath, lock, "Broken", "solar_icons_broken.dart")
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

	validModes := map[string]bool{
		"Brands": true, "Solid": true, "Regular": true, "Light": true,
		"Duotone": true, "Thin": true, "SharpThin": true, "SharpLight": true,
		"SharpRegular": true, "SharpSolid": true,
	}

	result := &FontAwesomeMapping{
		Icons: make(map[string][]IconEntry),
	}
	seenPerMode := make(map[string]map[string]bool)
	modeOrder := make(map[string]bool)

	totalCount := 0
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if !strings.HasPrefix(line, "static const IconData") {
			continue
		}
		line = strings.TrimSpace(strings.Replace(line, "static const IconData", "", 1))
		name := strings.TrimSpace(strings.Split(line, "=")[0])
		name = ToKebabCase(name)
		codePointFull := strings.TrimSpace(strings.TrimSuffix(strings.Split(line, "=")[1], ";"))

		if !strings.HasPrefix(codePointFull, "IconData") {
			continue
		}

		mode := strings.Replace(strings.Split(codePointFull, "(")[0], "IconData", "", 1)
		if !validModes[mode] {
			continue
		}

		codePoint := strings.TrimSpace(strings.Split(strings.Split(codePointFull, "(")[1], ")")[0])

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

func ScanIonicons(pubPath string, lock *PubspecLock) ([]IconEntry, error) {
	pkgDir, err := lock.PackageDir("ionicons")
	if err != nil {
		return nil, err
	}
	dartFile := filepath.Join(pubPath, pkgDir, "lib", "ionicons.dart")
	lines, err := readLines(dartFile)
	if err != nil {
		return nil, fmt.Errorf("reading Ionicons file: %w", err)
	}

	var result []IconEntry
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if !strings.HasPrefix(line, "static const") {
			continue
		}
		line = strings.TrimSpace(strings.Replace(line, "static const", "", 1))
		name := strings.TrimSpace(strings.Split(line, "=")[0])
		name = ToKebabCase(name)
		codePointFull := strings.TrimSuffix(strings.TrimSpace(strings.Split(line, "=")[1]), ";")
		codePoint := strings.TrimSpace(strings.Split(strings.Split(codePointFull, "(")[1], ")")[0])
		result = append(result, IconEntry{Name: name, CodePoint: codePoint})
	}

	fmt.Printf("Mapped %d icons from Ionicons\n", len(result))
	return result, nil
}

func scanIconsax(pubPath string, lock *PubspecLock, variant, dartFileName string) ([]IconEntry, error) {
	pkgDir, err := lock.PackageDir("iconsax_plus")
	if err != nil {
		return nil, err
	}
	dartFile := filepath.Join(pubPath, pkgDir, "lib", "src", dartFileName)
	lines, err := readLines(dartFile)
	if err != nil {
		return nil, fmt.Errorf("reading Iconsax %s file: %w", variant, err)
	}

	seen := make(map[string]bool)
	var result []IconEntry
	var buffer string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "static const IconData") {
			buffer = line
			continue
		}
		if buffer != "" {
			completeLine := buffer + line
			buffer = ""
			completeLine = strings.TrimSpace(strings.Replace(completeLine, "static const IconData", "", 1))
			name := strings.TrimSpace(strings.Split(completeLine, "=")[0])
			name = ToKebabCase(name)
			codePointPart := strings.Split(completeLine, "=")[1]
			codePoint := strings.TrimSpace(strings.Split(strings.Replace(codePointPart, "IconData(", "", 1), ",")[0])

			if seen[name] {
				fmt.Printf("Iconsax Plus %s: Duplicated icon name: %s\n", variant, name)
				continue
			}
			seen[name] = true
			result = append(result, IconEntry{Name: name, CodePoint: codePoint})
		}
	}

	fmt.Printf("Mapped %d icons from IconsaxPlus%s\n", len(result), variant)
	return result, nil
}

func ScanIconsaxBold(pubPath string, lock *PubspecLock) ([]IconEntry, error) {
	return scanIconsax(pubPath, lock, "Bold", "iconsax_plus_bold.dart")
}

func ScanIconsaxBroken(pubPath string, lock *PubspecLock) ([]IconEntry, error) {
	return scanIconsax(pubPath, lock, "Broken", "iconsax_plus_broken.dart")
}

func ScanIconsaxLinear(pubPath string, lock *PubspecLock) ([]IconEntry, error) {
	return scanIconsax(pubPath, lock, "Linear", "iconsax_plus_linear.dart")
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
