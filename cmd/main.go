package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/goldenm-software/layrz-icons/internal/tools"
)

func defaultPubCachePath() string {
	if env := os.Getenv("PUB_CACHE"); env != "" {
		return filepath.Join(env, "hosted", "pub.dev")
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}

	if runtime.GOOS == "windows" {
		return filepath.Join(home, "AppData", "Local", "Pub", "Cache", "hosted", "pub.dev")
	}

	// Linux / macOS
	return filepath.Join(home, ".pub-cache", "hosted", "pub.dev")
}

func main() {
	pubPath := flag.String("pub-path", "", "Path to your Flutter pub cache hosted directory. If empty, resolves from PUB_CACHE env or the default platform location")
	baseDir := flag.String("base-dir", "", "Base directory of the layrz_icons project (defaults to current working directory)")
	flag.Parse()

	if *pubPath == "" {
		*pubPath = defaultPubCachePath()
		if *pubPath == "" {
			fmt.Fprintln(os.Stderr, "Error: could not determine pub cache path. Set -pub-path or the PUB_CACHE environment variable")
			flag.Usage()
			os.Exit(1)
		}
		fmt.Printf("Using pub cache: %s\n", *pubPath)
	}

	if *baseDir == "" {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatalf("Failed to get working directory: %v", err)
		}
		*baseDir = wd
	}

	absPubPath, err := filepath.Abs(*pubPath)
	if err != nil {
		log.Fatalf("Failed to resolve pub path: %v", err)
	}
	absBaseDir, err := filepath.Abs(*baseDir)
	if err != nil {
		log.Fatalf("Failed to resolve base dir: %v", err)
	}

	fmt.Println("Parsing pubspec.lock...")
	lock, err := tools.ParsePubspecLock(absBaseDir)
	if err != nil {
		log.Fatalf("Failed to parse pubspec.lock: %v", err)
	}

	fmt.Println("Generating icon mappings...")

	mdi, err := tools.ScanMaterialDesignIconsFlutter(absPubPath, lock)
	if err != nil {
		log.Fatalf("Failed to scan MDI: %v", err)
	}

	solarBold, err := tools.ScanSolarIconsBold(absPubPath, lock)
	if err != nil {
		log.Fatalf("Failed to scan Solar Bold: %v", err)
	}

	solarOutline, err := tools.ScanSolarIconsOutline(absPubPath, lock)
	if err != nil {
		log.Fatalf("Failed to scan Solar Outline: %v", err)
	}

	// solarBroken, err := tools.ScanSolarIconsBroken(absPubPath, lock)
	// if err != nil {
	// 	log.Fatalf("Failed to scan Solar Broken: %v", err)
	// }

	fontAwesome, err := tools.ScanFontAwesomeFlutter(absPubPath, lock)
	if err != nil {
		log.Fatalf("Failed to scan Font Awesome: %v", err)
	}

	ionicons, err := tools.ScanIonicons(absPubPath, lock)
	if err != nil {
		log.Fatalf("Failed to scan Ionicons: %v", err)
	}

	iconsaxBold, err := tools.ScanIconsaxBold(absPubPath, lock)
	if err != nil {
		log.Fatalf("Failed to scan Iconsax Bold: %v", err)
	}

	iconsaxBroken, err := tools.ScanIconsaxBroken(absPubPath, lock)
	if err != nil {
		log.Fatalf("Failed to scan Iconsax Broken: %v", err)
	}

	iconsaxLinear, err := tools.ScanIconsaxLinear(absPubPath, lock)
	if err != nil {
		log.Fatalf("Failed to scan Iconsax Linear: %v", err)
	}

	fluttySolarIcons, err := tools.ScanFluttySolarIcons(absPubPath, lock)
	if err != nil {
		log.Fatalf("Failed to scan Flutty Solar Icons: %v", err)
	}

	fmt.Println("Icon mappings generated.")

	m := &tools.AllMappings{
		MDI:          mdi,
		SolarBold:    solarBold,
		SolarOutline: solarOutline,
		// SolarBroken:      solarBroken,
		FontAwesome:      fontAwesome,
		Ionicons:         ionicons,
		IconsaxBold:      iconsaxBold,
		IconsaxBroken:    iconsaxBroken,
		IconsaxLinear:    iconsaxLinear,
		FluttySolarIcons: fluttySolarIcons,
	}

	fmt.Println("Creating class_enum.dart ...")
	if err := tools.GenerateClassEnum(absBaseDir, m); err != nil {
		log.Fatalf("Failed to generate class_enum.dart: %v", err)
	}
	fmt.Println("class_enum.dart created.")

	fmt.Println("Creating icon_enum.dart ...")
	if err := tools.GenerateIconEnum(absBaseDir, m); err != nil {
		log.Fatalf("Failed to generate icon_enum.dart: %v", err)
	}
	fmt.Println("icon_enum.dart created.")

	fmt.Println("Creating mapping.dart ...")
	if err := tools.GenerateMapping(absBaseDir, m); err != nil {
		log.Fatalf("Failed to generate mapping.dart: %v", err)
	}
	fmt.Println("mapping.dart created.")
}
