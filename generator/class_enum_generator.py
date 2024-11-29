from pathlib import Path
from typing import Dict

from generator.utils import to_camel_case, to_kebab_case


def generate_class_enum(
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
  output_file = base_dir / 'lib' / 'src' / 'class_enum.dart'
  output_file.unlink(missing_ok=True)

  with open(output_file, 'w', encoding='utf-8') as f:
    f.write('part of "../layrz_icons.dart";\n\n')
    f.write('/// This is a auto-generated file. Do not modify it manually.\n\n')
    f.write('class LayrzIconsClasses {\n')
    for name, code_point in mdi_icon_mapping.items():
      f.write(f'  static LayrzIcon get mdi{to_camel_case(name, capitalize=True)} => ')
      f.write(f'LayrzIcon(codePoint: {code_point}, name: "mdi-{to_kebab_case(name)}", ')
      f.write('family: LayrzFamily.materialDesignIcons);\n')

    for name, code_point in solar_icon_bold_mapping.items():
      f.write(f'  static LayrzIcon get solarBold{to_camel_case(name, capitalize=True)} => ')
      f.write(f'LayrzIcon(codePoint: {code_point}, name: "solar-bold-{to_kebab_case(name)}", ')
      f.write('family: LayrzFamily.solarIconsBold);\n')

    for name, code_point in solar_icon_outline_mapping.items():
      f.write(f'  static LayrzIcon get solarOutline{to_camel_case(name, capitalize=True)} => ')
      f.write(f'LayrzIcon(codePoint: {code_point}, name: "solar-outline-{to_kebab_case(name)}", ')
      f.write('family: LayrzFamily.solarIconsOutline);\n')

    for mode, icons in font_awesome_mapping.items():
      for name in icons.keys():
        f.write(f'  static LayrzIcon get fa{mode.capitalize()}{to_camel_case(name, capitalize=True)} => ')
        f.write(f'LayrzIcon(codePoint: {code_point}, name: "fa-{to_kebab_case(mode)}-{to_kebab_case(name)}", ')
        f.write(f'family: LayrzFamily.fontAwesome{mode});\n')

    for name, code_point in ionicons_mapping.items():
      f.write(f'  static LayrzIcon get ionicons{to_camel_case(name, capitalize=True)} => ')
      f.write(f'LayrzIcon(codePoint: {code_point}, name: "ionicons-{to_kebab_case(name)}", ')
      f.write('family: LayrzFamily.ionicons);\n')

    for name, code_point in iconsax_bold_mapping.items():
      f.write(f'  static LayrzIcon get iconsaxBold{to_camel_case(name, capitalize=True)} => ')
      f.write(f'LayrzIcon(codePoint: {code_point}, name: "iconsax-bold-{to_kebab_case(name)}", ')
      f.write('family: LayrzFamily.iconsaxPlusBold);\n')

    for name, code_point in iconsax_broken_mapping.items():
      f.write(f'  static LayrzIcon get iconsaxBroken{to_camel_case(name, capitalize=True)} => ')
      f.write(f'LayrzIcon(codePoint: {code_point}, name: "iconsax-broken-{to_kebab_case(name)}", ')
      f.write('family: LayrzFamily.iconsaxPlusBroken);\n')

    for name, code_point in iconsax_linear_mapping.items():
      f.write(f'  static LayrzIcon get iconsaxLinear{to_camel_case(name, capitalize=True)} => ')
      f.write(f'LayrzIcon(codePoint: {code_point}, name: "iconsax-linear-{to_kebab_case(name)}", ')
      f.write('family: LayrzFamily.iconsaxPlusLinear);\n')

    f.write('}\n')
