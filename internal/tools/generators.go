package tools

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// AllMappings holds all scanned icon mappings passed to generators.
type AllMappings struct {
	MDI              []IconEntry
	SolarBold        []IconEntry
	SolarOutline     []IconEntry
	SolarBroken      []IconEntry
	FontAwesome      *FontAwesomeMapping
	Ionicons         []IconEntry
	IconsaxBold      []IconEntry
	IconsaxBroken    []IconEntry
	IconsaxLinear    []IconEntry
	FluttySolarIcons []IconEntry
}

type fluttySolarWeight struct {
	GetterPrefix string // e.g. "fluttySolarBold"
	KeyPrefix    string // e.g. "flutty-solar-bold"
	Family       string // e.g. "fluttySolarBold"
}

var fluttySolarWeights = []fluttySolarWeight{
	{"fluttySolarBold", "flutty-solar-bold", "fluttySolarBold"},
	{"fluttySolarBroken", "flutty-solar-broken", "fluttySolarBroken"},
	{"fluttySolarLinear", "flutty-solar-linear", "fluttySolarLinear"},
	{"fluttySolarOutline", "flutty-solar-outline", "fluttySolarOutline"},
	// {"fluttySolarBoldDuotone", "flutty-solar-bold-duotone", "fluttySolarBoldDuotone"},
	// {"fluttySolarLineDuotone", "flutty-solar-line-duotone", "fluttySolarLineDuotone"},
}

func GenerateClassEnum(baseDir string, m *AllMappings) error {
	outputFile := filepath.Join(baseDir, "lib", "src", "class_enum.dart")
	f, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("creating class_enum.dart: %w", err)
	}
	defer f.Close()

	w := func(s string) { f.WriteString(s) }

	w("part of \"../layrz_icons.dart\";\n\n")
	w("/// This is a auto-generated file. Do not modify it manually.\n\n")
	w("class LayrzIconsClasses {\n")

	for _, e := range m.MDI {
		if e.CodePoint == "" {
			continue
		}
		camel := ToCamelCase(e.Name, true)
		kebab := ToKebabCase(e.Name)
		w(fmt.Sprintf("  static LayrzIcon get mdi%s => LayrzIcon(codePoint: %s, name: \"mdi-%s\", family: LayrzFamily.materialDesignIcons);\n", camel, e.CodePoint, kebab))
	}

	for _, e := range m.SolarBold {
		if e.CodePoint == "" {
			continue
		}
		camel := ToCamelCase(e.Name, true)
		kebab := ToKebabCase(e.Name)
		w(fmt.Sprintf("  static LayrzIcon get solarBold%s => LayrzIcon(codePoint: %s, name: \"solar-bold-%s\", family: LayrzFamily.solarIconsBold);\n", camel, e.CodePoint, kebab))
	}

	for _, e := range m.SolarOutline {
		if e.CodePoint == "" {
			continue
		}
		camel := ToCamelCase(e.Name, true)
		kebab := ToKebabCase(e.Name)
		w(fmt.Sprintf("  static LayrzIcon get solarOutline%s => LayrzIcon(codePoint: %s, name: \"solar-outline-%s\", family: LayrzFamily.solarIconsOutline);\n", camel, e.CodePoint, kebab))
	}

	for _, e := range m.SolarBroken {
		if e.CodePoint == "" {
			continue
		}
		camel := ToCamelCase(e.Name, true)
		kebab := ToKebabCase(e.Name)
		w(fmt.Sprintf("  static LayrzIcon get solarBroken%s => LayrzIcon(codePoint: %s, name: \"solar-broken-%s\", family: LayrzFamily.solarIconsBroken);\n", camel, e.CodePoint, kebab))
	}

	for _, mode := range m.FontAwesome.Modes {
		icons := m.FontAwesome.Icons[mode]
		for _, e := range icons {
			if e.CodePoint == "" {
				continue
			}
			camel := ToCamelCase(e.Name, true)
			kebab := ToKebabCase(e.Name)
			modeKebab := ToKebabCase(mode)
			modeCap := strings.ToUpper(mode[:1]) + mode[1:]
			w(fmt.Sprintf("  static LayrzIcon get fa%s%s => LayrzIcon(codePoint: %s, name: \"fa-%s-%s\", family: LayrzFamily.fontAwesome%s);\n", modeCap, camel, e.CodePoint, modeKebab, kebab, mode))
		}
	}

	for _, e := range m.Ionicons {
		if e.CodePoint == "" {
			continue
		}
		camel := ToCamelCase(e.Name, true)
		kebab := ToKebabCase(e.Name)
		w(fmt.Sprintf("  static LayrzIcon get ionicons%s => LayrzIcon(codePoint: %s, name: \"ionicons-%s\", family: LayrzFamily.ionicons);\n", camel, e.CodePoint, kebab))
	}

	for _, e := range m.IconsaxBold {
		if e.CodePoint == "" {
			continue
		}
		camel := ToCamelCase(e.Name, true)
		kebab := ToKebabCase(e.Name)
		w(fmt.Sprintf("  static LayrzIcon get iconsaxBold%s => LayrzIcon(codePoint: %s, name: \"iconsax-bold-%s\", family: LayrzFamily.iconsaxPlusBold);\n", camel, e.CodePoint, kebab))
	}

	for _, e := range m.IconsaxBroken {
		if e.CodePoint == "" {
			continue
		}
		camel := ToCamelCase(e.Name, true)
		kebab := ToKebabCase(e.Name)
		w(fmt.Sprintf("  static LayrzIcon get iconsaxBroken%s => LayrzIcon(codePoint: %s, name: \"iconsax-broken-%s\", family: LayrzFamily.iconsaxPlusBroken);\n", camel, e.CodePoint, kebab))
	}

	for _, e := range m.IconsaxLinear {
		if e.CodePoint == "" {
			continue
		}
		camel := ToCamelCase(e.Name, true)
		kebab := ToKebabCase(e.Name)
		w(fmt.Sprintf("  static LayrzIcon get iconsaxLinear%s => LayrzIcon(codePoint: %s, name: \"iconsax-linear-%s\", family: LayrzFamily.iconsaxPlusLinear);\n", camel, e.CodePoint, kebab))
	}

	for _, weight := range fluttySolarWeights {
		for _, e := range m.FluttySolarIcons {
			if e.CodePoint == "" {
				continue
			}
			camel := ToCamelCase(e.Name, true)
			kebab := ToKebabCase(e.Name)
			w(fmt.Sprintf("  static LayrzIcon get %s%s => LayrzIcon(codePoint: %s, name: \"%s-%s\", family: LayrzFamily.%s);\n", weight.GetterPrefix, camel, e.CodePoint, weight.KeyPrefix, kebab, weight.Family))
		}
	}

	w("}\n")
	return nil
}

func GenerateIconEnum(baseDir string, m *AllMappings) error {
	outputFile := filepath.Join(baseDir, "lib", "src", "icon_enum.dart")
	f, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("creating icon_enum.dart: %w", err)
	}
	defer f.Close()

	w := func(s string) { f.WriteString(s) }

	w("part of \"../layrz_icons.dart\";\n\n")
	w("/// This is a auto-generated file. Do not modify it manually.\n\n")
	w("class LayrzIcons {\n")

	for _, e := range m.MDI {
		if e.CodePoint == "" {
			continue
		}
		camel := ToCamelCase(e.Name, true)
		w(fmt.Sprintf("  static IconData get mdi%s => LayrzIconsClasses.mdi%s.iconData;\n", camel, camel))
	}

	for _, e := range m.SolarBold {
		if e.CodePoint == "" {
			continue
		}
		camel := ToCamelCase(e.Name, true)
		w(fmt.Sprintf("  static IconData get solarBold%s => LayrzIconsClasses.solarBold%s.iconData;\n", camel, camel))
	}

	for _, e := range m.SolarOutline {
		if e.CodePoint == "" {
			continue
		}
		camel := ToCamelCase(e.Name, true)
		w(fmt.Sprintf("  static IconData get solarOutline%s => LayrzIconsClasses.solarOutline%s.iconData;\n", camel, camel))
	}

	for _, e := range m.SolarBroken {
		if e.CodePoint == "" {
			continue
		}
		camel := ToCamelCase(e.Name, true)
		w(fmt.Sprintf("  static IconData get solarBroken%s => LayrzIconsClasses.solarBroken%s.iconData;\n", camel, camel))
	}

	for _, mode := range m.FontAwesome.Modes {
		icons := m.FontAwesome.Icons[mode]
		modeCap := strings.ToUpper(mode[:1]) + mode[1:]
		for _, e := range icons {
			if e.CodePoint == "" {
				continue
			}
			camel := ToCamelCase(e.Name, true)
			w(fmt.Sprintf("  static IconData get fa%s%s => LayrzIconsClasses.fa%s%s.iconData;\n", modeCap, camel, modeCap, camel))
		}
	}

	for _, e := range m.Ionicons {
		if e.CodePoint == "" {
			continue
		}
		camel := ToCamelCase(e.Name, true)
		w(fmt.Sprintf("  static IconData get ionicons%s => LayrzIconsClasses.ionicons%s.iconData;\n", camel, camel))
	}

	for _, e := range m.IconsaxBold {
		if e.CodePoint == "" {
			continue
		}
		camel := ToCamelCase(e.Name, true)
		w(fmt.Sprintf("  static IconData get iconsaxBold%s => LayrzIconsClasses.iconsaxBold%s.iconData;\n", camel, camel))
	}

	for _, e := range m.IconsaxBroken {
		if e.CodePoint == "" {
			continue
		}
		camel := ToCamelCase(e.Name, true)
		w(fmt.Sprintf("  static IconData get iconsaxBroken%s => LayrzIconsClasses.iconsaxBroken%s.iconData;\n", camel, camel))
	}

	for _, e := range m.IconsaxLinear {
		if e.CodePoint == "" {
			continue
		}
		camel := ToCamelCase(e.Name, true)
		w(fmt.Sprintf("  static IconData get iconsaxLinear%s => LayrzIconsClasses.iconsaxLinear%s.iconData;\n", camel, camel))
	}

	for _, weight := range fluttySolarWeights {
		for _, e := range m.FluttySolarIcons {
			if e.CodePoint == "" {
				continue
			}
			camel := ToCamelCase(e.Name, true)
			w(fmt.Sprintf("  static IconData get %s%s => LayrzIconsClasses.%s%s.iconData;\n", weight.GetterPrefix, camel, weight.GetterPrefix, camel))
		}
	}

	w("}\n")
	return nil
}

func GenerateMapping(baseDir string, m *AllMappings) error {
	outputFile := filepath.Join(baseDir, "lib", "src", "mapping.dart")
	f, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("creating mapping.dart: %w", err)
	}
	defer f.Close()

	w := func(s string) { f.WriteString(s) }

	w("part of \"../layrz_icons.dart\";\n\n")
	w("/// This is a auto-generated file. Do not modify it manually.\n\n")
	w("Map<String, LayrzIcon> iconMapping = {\n")

	w("  // Material Design Icons\n")
	for _, e := range m.MDI {
		if e.CodePoint == "" {
			continue
		}
		kebab := ToKebabCase(e.Name)
		camel := ToCamelCase(e.Name, true)
		w(fmt.Sprintf("  \"mdi-%s\": LayrzIconsClasses.mdi%s,\n", kebab, camel))
	}
	w("  // /Material Design Icons\n")

	w("  // Solar Icons Bold\n")
	for _, e := range m.SolarBold {
		if e.CodePoint == "" {
			continue
		}
		kebab := ToKebabCase(e.Name)
		camel := ToCamelCase(e.Name, true)
		w(fmt.Sprintf("  \"solar-bold-%s\": LayrzIconsClasses.solarBold%s,\n", kebab, camel))
	}
	w("  // /Solar Icons Bold\n")

	w("  // Solar Icons Outline\n")
	for _, e := range m.SolarOutline {
		if e.CodePoint == "" {
			continue
		}
		kebab := ToKebabCase(e.Name)
		camel := ToCamelCase(e.Name, true)
		w(fmt.Sprintf("  \"solar-outline-%s\": LayrzIconsClasses.solarOutline%s,\n", kebab, camel))
	}
	w("  // /Solar Icons Outline\n")

	w("  // Solar Icons Broken\n")
	for _, e := range m.SolarBroken {
		if e.CodePoint == "" {
			continue
		}
		kebab := ToKebabCase(e.Name)
		camel := ToCamelCase(e.Name, true)
		w(fmt.Sprintf("  \"solar-broken-%s\": LayrzIconsClasses.solarBroken%s,\n", kebab, camel))
	}
	w("  // /Solar Icons Broken\n")

	w("  // Font Awesome Flutter\n")
	for _, mode := range m.FontAwesome.Modes {
		icons := m.FontAwesome.Icons[mode]
		modeCap := strings.ToUpper(mode[:1]) + mode[1:]
		modeKebab := ToKebabCase(mode)
		for _, e := range icons {
			if e.CodePoint == "" {
				continue
			}
			kebab := ToKebabCase(e.Name)
			camel := ToCamelCase(e.Name, true)
			w(fmt.Sprintf("  \"fa-%s-%s\": LayrzIconsClasses.fa%s%s,\n", modeKebab, kebab, modeCap, camel))
		}
	}
	w("  // /Font Awesome Flutter\n")

	w("  // Ionicons\n")
	for _, e := range m.Ionicons {
		if e.CodePoint == "" {
			continue
		}
		kebab := ToKebabCase(e.Name)
		camel := ToCamelCase(e.Name, true)
		w(fmt.Sprintf("  \"ionicons-%s\": LayrzIconsClasses.ionicons%s,\n", kebab, camel))
	}
	w("  // /Ionicons\n")

	w("  // Iconsax Plus Bold\n")
	for _, e := range m.IconsaxBold {
		if e.CodePoint == "" {
			continue
		}
		kebab := ToKebabCase(e.Name)
		camel := ToCamelCase(e.Name, true)
		w(fmt.Sprintf("  \"iconsax-bold-%s\": LayrzIconsClasses.iconsaxBold%s,\n", kebab, camel))
	}
	w("  // /Iconsax Plus Bold\n")

	w("  // Iconsax Plus Broken\n")
	for _, e := range m.IconsaxBroken {
		if e.CodePoint == "" {
			continue
		}
		kebab := ToKebabCase(e.Name)
		camel := ToCamelCase(e.Name, true)
		w(fmt.Sprintf("  \"iconsax-broken-%s\": LayrzIconsClasses.iconsaxBroken%s,\n", kebab, camel))
	}
	w("  // /Iconsax Plus Broken\n")

	w("  // Iconsax Plus Linear\n")
	for _, e := range m.IconsaxLinear {
		if e.CodePoint == "" {
			continue
		}
		kebab := ToKebabCase(e.Name)
		camel := ToCamelCase(e.Name, true)
		w(fmt.Sprintf("  \"iconsax-linear-%s\": LayrzIconsClasses.iconsaxLinear%s,\n", kebab, camel))
	}
	w("  // /Iconsax Plus Linear\n")

	w("  // Flutty Solar Icons\n")
	for _, weight := range fluttySolarWeights {
		for _, e := range m.FluttySolarIcons {
			if e.CodePoint == "" {
				continue
			}
			kebab := ToKebabCase(e.Name)
			camel := ToCamelCase(e.Name, true)
			w(fmt.Sprintf("  \"%s-%s\": LayrzIconsClasses.%s%s,\n", weight.KeyPrefix, kebab, weight.GetterPrefix, camel))
		}
	}
	w("  // /Flutty Solar Icons\n")

	w("};\n\n")
	return nil
}
