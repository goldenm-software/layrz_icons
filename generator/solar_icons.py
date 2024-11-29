from pathlib import Path
from typing import Dict

from generator.utils import to_kebab_case


def scan_solar_icons_bold(pub_path: Path) -> Dict[str, str]:
  dart_file = pub_path / 'solar_icons-0.0.5' / 'lib' / 'src'
  dart_file = dart_file / 'solar_icons_bold.dart'

  with open(dart_file, 'r', encoding='utf-8') as f:
    lines = f.readlines()

  solar_icon_bold_mapping = {}
  buffer = ''
  for line in lines:
    line = line.strip()
    if line.startswith('static const SolarIconsData'):
      buffer = line
      continue

    if buffer:
      complete_line = buffer + line
      buffer = ''
      complete_line = complete_line.replace('static const SolarIconsData', '').strip()
      name = complete_line.split('=')[0].strip()
      name = to_kebab_case(name)
      code_point = complete_line.split('=')[1].replace('SolarIconsData(', '').split(',')[0].strip()
      if name in solar_icon_bold_mapping:
        print(f'Solar Icons Bold: Duplicated icon name: {name}')
        continue
      solar_icon_bold_mapping[name] = code_point

  print(f'Mapped {len(solar_icon_bold_mapping)} icons from Solar Icons Bold')

  return solar_icon_bold_mapping


def scan_solar_icons_outline(pub_path: Path) -> Dict[str, str]:
  dart_file = pub_path / 'solar_icons-0.0.5' / 'lib' / 'src'
  dart_file = dart_file / 'solar_icons_outline.dart'

  with open(dart_file, 'r', encoding='utf-8') as f:
    lines = f.readlines()

  solar_icon_outline_mapping = {}
  buffer = ''
  for line in lines:
    line = line.strip()
    if line.startswith('static const SolarIconsData'):
      buffer = line
      continue

    if buffer:
      complete_line = buffer + line
      buffer = ''
      complete_line = complete_line.replace('static const SolarIconsData', '').strip()
      name = complete_line.split('=')[0].strip()
      name = to_kebab_case(name)
      code_point = complete_line.split('=')[1].replace('SolarIconsData(', '').split(',')[0].strip()
      if name in solar_icon_outline_mapping:
        print(f'Solar Icons Outline: Duplicated icon name: {name}')
        continue

      solar_icon_outline_mapping[name] = code_point

  print(f'Mapped {len(solar_icon_outline_mapping)} icons from Solar Icons Outline')

  return solar_icon_outline_mapping
