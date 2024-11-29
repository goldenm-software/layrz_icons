from pathlib import Path
from typing import Dict

from generator.utils import to_camel_case


def generate_icon_enum(
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
  output_file = base_dir / 'lib' / 'src' / 'icon_enum.dart'
  output_file.unlink(missing_ok=True)

  with open(output_file, 'w', encoding='utf-8') as f:
    f.write('part of "../layrz_icons.dart";\n\n')
    f.write('/// This is a auto-generated file. Do not modify it manually.\n\n')
    f.write('class LayrzIcons {\n')
    for name in mdi_icon_mapping.keys():
      f.write(f'  static IconData get mdi{to_camel_case(name, capitalize=True)} => ')
      f.write(f'LayrzIconsClasses.mdi{to_camel_case(name, capitalize=True)}.iconData;\n')

    for name in solar_icon_bold_mapping.keys():
      f.write(f'  static IconData get solarBold{to_camel_case(name, capitalize=True)} => ')
      f.write(f'LayrzIconsClasses.solarBold{to_camel_case(name, capitalize=True)}.iconData;\n')

    for name in solar_icon_outline_mapping.keys():
      f.write(f'  static IconData get solarOutline{to_camel_case(name, capitalize=True)} => ')
      f.write(f'LayrzIconsClasses.solarOutline{to_camel_case(name, capitalize=True)}.iconData;\n')

    for mode, icons in font_awesome_mapping.items():
      for name in icons.keys():
        f.write(f'  static IconData get fa{mode.capitalize()}{to_camel_case(name, capitalize=True)} => ')
        f.write(f'LayrzIconsClasses.fa{mode.capitalize()}{to_camel_case(name, capitalize=True)}.iconData;\n')

    for name in ionicons_mapping.keys():
      f.write(f'  static IconData get ionicons{to_camel_case(name, capitalize=True)} => ')
      f.write(f'LayrzIconsClasses.ionicons{to_camel_case(name, capitalize=True)}.iconData;\n')

    for name in iconsax_bold_mapping.keys():
      f.write(f'  static IconData get iconsaxBold{to_camel_case(name, capitalize=True)} => ')
      f.write(f'LayrzIconsClasses.iconsaxBold{to_camel_case(name, capitalize=True)}.iconData;\n')

    for name in iconsax_broken_mapping.keys():
      f.write(f'  static IconData get iconsaxBroken{to_camel_case(name, capitalize=True)} => ')
      f.write(f'LayrzIconsClasses.iconsaxBroken{to_camel_case(name, capitalize=True)}.iconData;\n')

    for name in iconsax_linear_mapping.keys():
      f.write(f'  static IconData get iconsaxLinear{to_camel_case(name, capitalize=True)} => ')
      f.write(f'LayrzIconsClasses.iconsaxLinear{to_camel_case(name, capitalize=True)}.iconData;\n')

    f.write('}\n')
