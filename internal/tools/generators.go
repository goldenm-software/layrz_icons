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
	FontAwesome      *FontAwesomeMapping
	FluttySolarIcons []IconEntry
}

type fluttySolarWeight struct {
	GetterPrefix string // e.g. "solarBold"
	KeyPrefix    string // e.g. "solar-bold"
	Family       string // e.g. "solarBold"
}

var fluttySolarWeights = []fluttySolarWeight{
	{"solarBold", "solar-bold", "solarBold"},
	{"solarBroken", "solar-broken", "solarBroken"},
	{"solarLinear", "solar-linear", "solarLinear"},
	{"solarOutline", "solar-outline", "solarOutline"},
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

	w("  // Solar Icons\n")
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
	w("  // /Solar Icons\n")

	w("};\n\n")
	return nil
}
