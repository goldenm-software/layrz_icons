import re
from pathlib import Path


def _to_kebab_case(name: str) -> str:
    return re.sub(r"(?<!^)(?=[A-Z])", "-", name).lower()


def _to_camel_case(name: str) -> str:
    return name[0].upper() + name[1:]


BASE_DIR = Path(__file__).resolve().parent

PUB_PATH = Path("<PATH_TO_YOUR_PUB_CACHE_DIR>")
DART_FILE = PUB_PATH / "material_design_icons_flutter-7.0.7296" / "lib"
DART_FILE = DART_FILE / "icon_map.dart"

with open(DART_FILE, "r", encoding="utf-8") as f:
    lines = f.readlines()

mdi_icon_mapping = {}
search_regex = re.compile(r"'(\w+)': _MdiIconData\(")
is_reading = False
for line in lines:
    line = line.strip()

    if search_regex.match(line):
        code_point = line.split("_MdiIconData(")[1].split(")")[0]
        name = line.split("'")[1].split("'")[0]
        mdi_icon_mapping[name] = code_point

print(f"Mapped {len(mdi_icon_mapping)} icons from Material Design Icons Flutter")

DART_FILE = PUB_PATH / "solar_icons-0.0.5" / "lib" / "src"
DART_FILE = DART_FILE / "solar_icons_bold.dart"

with open(DART_FILE, "r", encoding="utf-8") as f:
    lines = f.readlines()

solar_icon_bold_mapping = {}
buffer = ""
for line in lines:
    line = line.strip()
    if line.startswith("static const SolarIconsData"):
        buffer = line
        continue

    if buffer:
        complete_line = buffer + line
        buffer = ""
        complete_line = complete_line.replace("static const SolarIconsData", "").strip()
        name = complete_line.split("=")[0].strip()
        code_point = (
            complete_line.split("=")[1]
            .replace("SolarIconsData(", "")
            .split(",")[0]
            .strip()
        )
        solar_icon_bold_mapping[name] = code_point

print(f"Mapped {len(solar_icon_bold_mapping)} icons from Solar Icons Bold")

DART_FILE = PUB_PATH / "solar_icons-0.0.5" / "lib" / "src"
DART_FILE = DART_FILE / "solar_icons_outline.dart"

with open(DART_FILE, "r", encoding="utf-8") as f:
    lines = f.readlines()

solar_icon_outline_mapping = {}
buffer = ""
for line in lines:
    line = line.strip()
    if line.startswith("static const SolarIconsData"):
        buffer = line
        continue

    if buffer:
        complete_line = buffer + line
        buffer = ""
        complete_line = complete_line.replace("static const SolarIconsData", "").strip()
        name = complete_line.split("=")[0].strip()
        code_point = (
            complete_line.split("=")[1]
            .replace("SolarIconsData(", "")
            .split(",")[0]
            .strip()
        )
        solar_icon_outline_mapping[name] = code_point

print(f"Mapped {len(solar_icon_outline_mapping)} icons from Solar Icons Outline")

DART_FILE = PUB_PATH / "font_awesome_flutter-10.8.0" / "lib"
DART_FILE = DART_FILE / "font_awesome_flutter.dart"

with open(DART_FILE, "r", encoding="utf-8") as f:
    lines = f.readlines()

font_awesome_mapping = {}
modes = [
    "Brands",
    "Solid",
    "Regular",
    "Light",
    "Duotone",
    "Thin",
    "SharpThin",
    "SharpLight",
    "SharpRegular",
    "SharpSolid",
]

for line in lines:
    line = line.strip()
    if not line.startswith("static const IconData"):
        continue
    line = line.replace("static const IconData", "").strip()
    name = line.split("=")[0].strip()
    code_point = line.split("=")[1].strip().replace(";", "")
    if not code_point.startswith("IconData"):
        continue
    mode = code_point.split("(")[0].replace("IconData", "")
    if mode not in modes:
        continue

    code_point = code_point.split("(")[1].split(")")[0].strip()
    if mode not in font_awesome_mapping:
        font_awesome_mapping[mode] = {}
    font_awesome_mapping[mode][name] = code_point

print(
    f"Mapped {sum([len(v) for v in font_awesome_mapping.values()])} icons from Font Awesome Flutter"
)

# <INSERT HERE YOUR NEW LIBRARY>

OUTPUT_FILE = BASE_DIR / "lib" / "src" / "mapping.dart"

OUTPUT_FILE.unlink(missing_ok=True)

print("Writing mapping file...")
with open(OUTPUT_FILE, "w", encoding="utf-8") as f:
    f.write('part of "../layrz_icons.dart";\n\n')
    f.write("/// This is a auto-generated file. Do not modify it manually.\n\n")
    f.write("const Map<String, LayrzIcon> iconMapping = {\n")

    f.write("  // Material Design Icons\n")
    for name, code_point in mdi_icon_mapping.items():
        f.write(
            f'  "mdi-{_to_kebab_case(name)}": LayrzIcon(codePoint: {code_point}, name: "mdi-{_to_kebab_case(name)}", family: LayrzFamily.materialDesignIcons),\n'
        )
    f.write("  // /Material Design Icons\n")
    f.write("  // Solar Icons Bold\n")
    for name, code_point in solar_icon_bold_mapping.items():
        f.write(
            f'  "solar-bold-{_to_kebab_case(name)}": LayrzIcon(codePoint: {code_point}, name: "solar-bold-{_to_kebab_case(name)}", family: LayrzFamily.solarIconsBold),\n'
        )
    f.write("  // /Solar Icons Bold\n")

    f.write("  // Solar Icons Outline\n")
    for name, code_point in solar_icon_outline_mapping.items():
        f.write(
            f'  "solar-outline-{_to_kebab_case(name)}": LayrzIcon(codePoint: {code_point}, name: "solar-outline-{_to_kebab_case(name)}", family: LayrzFamily.solarIconsOutline),\n'
        )
    f.write("  // /Solar Icons Outline\n")

    f.write("  // Font Awesome Flutter\n")
    for mode, icons in font_awesome_mapping.items():
        for name, code_point in icons.items():
            f.write(
                f'  "fa-{_to_kebab_case(mode)}-{_to_kebab_case(name)}": LayrzIcon(codePoint: {code_point}, name: "fa-{_to_kebab_case(mode)}-{_to_kebab_case(name)}", family: LayrzFamily.fontAwesome{mode}),\n'
            )
    f.write("  // /Font Awesome Flutter\n")

    # <INSERT HERE YOUR NEW LIBRARY>

    f.write("};\n\n")

print("Creating enum file...")

OUTPUT_FILE = BASE_DIR / "lib" / "src" / "icon_enum.dart"
OUTPUT_FILE.unlink(missing_ok=True)

with open(OUTPUT_FILE, "w", encoding="utf-8") as f:
    f.write('part of "../layrz_icons.dart";\n\n')
    f.write("/// This is a auto-generated file. Do not modify it manually.\n\n")
    f.write("class LayrzIcons {\n")
    for name in mdi_icon_mapping.keys():
        f.write(
            f'  static IconData get mdi{_to_camel_case(name)} => iconMapping["mdi-{_to_kebab_case(name)}"]!.iconData;\n'
        )
    for name in solar_icon_bold_mapping.keys():
        f.write(
            f'  static IconData get solarBold{_to_camel_case(name)} => iconMapping["solar-bold-{_to_kebab_case(name)}"]!.iconData;\n'
        )
    for name in solar_icon_outline_mapping.keys():
        f.write(
            f'  static IconData get solarOutline{_to_camel_case(name)} => iconMapping["solar-outline-{_to_kebab_case(name)}"]!.iconData;\n'
        )
    for mode, icons in font_awesome_mapping.items():
        for name in icons.keys():
            f.write(
                f'  static IconData get fa{mode.capitalize()}{_to_camel_case(name)} => iconMapping["fa-{_to_kebab_case(mode)}-{_to_kebab_case(name)}"]!.iconData;\n'
            )

    # <INSERT HERE YOUR NEW LIBRARY>
    f.write("}\n")

OUTPUT_FILE = BASE_DIR / "lib" / "src" / "class_enum.dart"
OUTPUT_FILE.unlink(missing_ok=True)

with open(OUTPUT_FILE, "w", encoding="utf-8") as f:
    f.write('part of "../layrz_icons.dart";\n\n')
    f.write("/// This is a auto-generated file. Do not modify it manually.\n\n")
    f.write("class LayrzIconsClasses {\n")
    for name in mdi_icon_mapping.keys():
        f.write(
            f'  static LayrzIcon get mdi{_to_camel_case(name)} => iconMapping["mdi-{_to_kebab_case(name)}"]!;\n'
        )
    for name in solar_icon_bold_mapping.keys():
        f.write(
            f'  static LayrzIcon get solarBold{_to_camel_case(name)} => iconMapping["solar-bold-{_to_kebab_case(name)}"]!;\n'
        )
    for name in solar_icon_outline_mapping.keys():
        f.write(
            f'  static LayrzIcon get solarOutline{_to_camel_case(name)} => iconMapping["solar-outline-{_to_kebab_case(name)}"]!;\n'
        )
    for mode, icons in font_awesome_mapping.items():
        for name in icons.keys():
            f.write(
                f'  static LayrzIcon get fa{mode.capitalize()}{_to_camel_case(name)} => iconMapping["fa-{_to_kebab_case(mode)}-{_to_kebab_case(name)}"]!;\n'
            )

    # <INSERT HERE YOUR NEW LIBRARY>
    f.write("}\n")
