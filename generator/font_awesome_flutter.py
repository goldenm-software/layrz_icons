import re
from pathlib import Path
from typing import Dict

from generator.utils import to_kebab_case


def scan_font_awesome_flutter(pub_path: Path) -> Dict[str, str]:
  dart_file = pub_path / 'font_awesome_flutter-10.8.0' / 'lib'
  dart_file = dart_file / 'font_awesome_flutter.dart'

  with open(dart_file, 'r', encoding='utf-8') as f:
    lines = f.readlines()

  font_awesome_mapping = {}
  modes = [
    'Brands',
    'Solid',
    'Regular',
    'Light',
    'Duotone',
    'Thin',
    'SharpThin',
    'SharpLight',
    'SharpRegular',
    'SharpSolid',
  ]

  for line in lines:
    line = line.strip()
    if not line.startswith('static const IconData'):
      continue
    line = line.replace('static const IconData', '').strip()
    name = line.split('=')[0].strip()
    name = to_kebab_case(name)
    code_point = line.split('=')[1].strip().replace(';', '')
    if not code_point.startswith('IconData'):
      continue
    mode = code_point.split('(')[0].replace('IconData', '')
    if mode not in modes:
      continue

    code_point = code_point.split('(')[1].split(')')[0].strip()
    if mode not in font_awesome_mapping:
      font_awesome_mapping[mode] = {}
    if name in font_awesome_mapping[mode]:
      print(f'Font Awesome Flutter: Duplicated icon name: {name}')
      continue

    font_awesome_mapping[mode][name] = code_point

  print(f'Mapped {sum([len(v) for v in font_awesome_mapping.values()])} icons from Font Awesome Flutter')

  return font_awesome_mapping
