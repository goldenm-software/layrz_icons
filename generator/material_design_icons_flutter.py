import re
from pathlib import Path
from typing import Dict

from generator.utils import to_kebab_case


def scan_material_design_icons_flutter(pub_path: Path) -> Dict[str, str]:
  dart_file = pub_path / 'material_design_icons_flutter-7.0.7296' / 'lib'
  dart_file = dart_file / 'icon_map.dart'

  with open(dart_file, 'r', encoding='utf-8') as f:
    lines = f.readlines()

  mdi_icon_mapping = {}
  search_regex = re.compile(r"'(\w+)': _MdiIconData\(")
  for line in lines:
    line = line.strip()

    if search_regex.match(line):
      code_point = line.split('_MdiIconData(')[1].split(')')[0]
      name = line.split("'")[1].split("'")[0]
      name = to_kebab_case(name)

      if name in mdi_icon_mapping:
        print(f'Material Design Icons Flutter: Duplicated icon name: {name}')
        continue

      mdi_icon_mapping[name] = code_point

  print(f'Mapped {len(mdi_icon_mapping)} icons from Material Design Icons Flutter')

  return mdi_icon_mapping
