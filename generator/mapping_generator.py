from pathlib import Path
from typing import Dict

from generator.utils import to_camel_case, to_kebab_case


def generate_mapping(
  base_dir: Path,
  mdi_icon_mapping: Dict[str, str],
  solar_icon_bold_mapping: Dict[str, str],
  solar_icon_outline_mapping: Dict[str, str],
  font_awesome_mapping: Dict[str, Dict[str, str]],
  ionicons_mapping: Dict[str, str],
  iconsax_bold_mapping: Dict[str, str],
  iconsax_broken_mapping: Dict[str, str],
  iconsax_linear_mapping: Dict[str, str],
  # You need to add your new mapping here
) -> None:
  output_file = base_dir / 'lib' / 'src' / 'mapping.dart'
  output_file.unlink(missing_ok=True)

  with open(output_file, 'w', encoding='utf-8') as f:
    f.write('part of "../layrz_icons.dart";\n\n')
    f.write('/// This is a auto-generated file. Do not modify it manually.\n\n')
    f.write('Map<String, LayrzIcon> iconMapping = {\n')

    f.write('  // Material Design Icons\n')
    for name, _ in mdi_icon_mapping.items():
      f.write(f'  "mdi-{to_kebab_case(name)}": LayrzIconsClasses.mdi{to_camel_case(name, capitalize=True)},\n')
    f.write('  // /Material Design Icons\n')
    f.write('  // Solar Icons Bold\n')
    for name, _ in solar_icon_bold_mapping.items():
      f.write(f'  "solar-bold-{to_kebab_case(name)}": LayrzIconsClasses.')
      f.write(f'solarBold{to_camel_case(name, capitalize=True)},\n')
    f.write('  // /Solar Icons Bold\n')

    f.write('  // Solar Icons Outline\n')
    for name, _ in solar_icon_outline_mapping.items():
      f.write(f'  "solar-outline-{to_kebab_case(name)}": LayrzIconsClasses.')
      f.write(f'solarOutline{to_camel_case(name, capitalize=True)},\n')
    f.write('  // /Solar Icons Outline\n')

    f.write('  // Font Awesome Flutter\n')
    for mode, icons in font_awesome_mapping.items():
      for name, _ in icons.items():
        f.write(f'  "fa-{to_kebab_case(mode)}-{to_kebab_case(name)}": LayrzIconsClasses.')
        f.write(f'fa{mode.capitalize()}{to_camel_case(name, capitalize=True)},\n')
    f.write('  // /Font Awesome Flutter\n')

    f.write('  // Ionicons\n')
    for name, _ in ionicons_mapping.items():
      f.write(f'  "ionicons-{to_kebab_case(name)}": LayrzIconsClasses.')
      f.write(f'ionicons{to_camel_case(name, capitalize=True)},\n')
    f.write('  // /Ionicons\n')

    f.write('  // Iconsax Plus Bold\n')
    for name, _ in iconsax_bold_mapping.items():
      f.write(f'  "iconsax-bold-{to_kebab_case(name)}": LayrzIconsClasses.')
      f.write(f'iconsaxBold{to_camel_case(name, capitalize=True)},\n')
    f.write('  // /Iconsax Plus Bold\n')

    f.write('  // Iconsax Plus Broken\n')
    for name, _ in iconsax_broken_mapping.items():
      f.write(f'  "iconsax-broken-{to_kebab_case(name)}": LayrzIconsClasses.')
      f.write(f'iconsaxBroken{to_camel_case(name, capitalize=True)},\n')
    f.write('  // /Iconsax Plus Broken\n')

    f.write('  // Iconsax Plus Linear\n')
    for name, _ in iconsax_linear_mapping.items():
      f.write(f'  "iconsax-linear-{to_kebab_case(name)}": LayrzIconsClasses.')
      f.write(f'iconsaxLinear{to_camel_case(name, capitalize=True)},\n')
    f.write('  // /Iconsax Plus Linear\n')

    f.write('};\n\n')
