from pathlib import Path
from typing import Dict

from generator.utils import to_kebab_case


def scan_iconsax_bold(pub_path: Path) -> Dict[str, str]:
  dart_file = pub_path / 'iconsax_plus-1.0.0' / 'lib' / 'src'
  dart_file = dart_file / 'iconsax_plus_bold.dart'

  with open(dart_file, 'r', encoding='utf-8') as f:
    lines = f.readlines()

  iconsax_bold_mapping = {}
  buffer = ''
  for line in lines:
    line = line.strip()
    if line.startswith('static const IconData'):
      buffer = line
      continue

    if buffer:
      complete_line = buffer + line
      buffer = ''
      complete_line = complete_line.replace('static const IconData', '').strip()
      name = complete_line.split('=')[0].strip()
      name = to_kebab_case(name)
      code_point = complete_line.split('=')[1].replace('IconData(', '').split(',')[0].strip()
      if name in iconsax_bold_mapping:
        print(f'Iconsax Plus Bold: Duplicated icon name: {name}')

      iconsax_bold_mapping[name] = code_point

  print(f'Mapped {len(iconsax_bold_mapping)} icons from IconsaxPlusBold')
  return iconsax_bold_mapping


def scan_iconsax_broken(pub_path: Path) -> Dict[str, str]:
  dart_file = pub_path / 'iconsax_plus-1.0.0' / 'lib' / 'src'
  dart_file = dart_file / 'iconsax_plus_broken.dart'

  with open(dart_file, 'r', encoding='utf-8') as f:
    lines = f.readlines()

  iconsax_broken_mapping = {}
  buffer = ''
  for line in lines:
    line = line.strip()
    if line.startswith('static const IconData'):
      buffer = line
      continue

    if buffer:
      complete_line = buffer + line
      buffer = ''
      complete_line = complete_line.replace('static const IconData', '').strip()
      name = complete_line.split('=')[0].strip()
      name = to_kebab_case(name)
      code_point = complete_line.split('=')[1].replace('IconData(', '').split(',')[0].strip()
      if name in iconsax_broken_mapping:
        print(f'Iconsax Plus Broken: Duplicated icon name: {name}')
        continue

      iconsax_broken_mapping[name] = code_point

  print(f'Mapped {len(iconsax_broken_mapping)} icons from IconsaxPlusBroken')
  return iconsax_broken_mapping


def scan_iconsax_linear(pub_path: Path) -> Dict[str, str]:
  dart_file = pub_path / 'iconsax_plus-1.0.0' / 'lib' / 'src'
  dart_file = dart_file / 'iconsax_plus_linear.dart'

  with open(dart_file, 'r', encoding='utf-8') as f:
    lines = f.readlines()

  iconsax_linear_mapping = {}
  buffer = ''
  for line in lines:
    line = line.strip()
    if line.startswith('static const IconData'):
      buffer = line
      continue

    if buffer:
      complete_line = buffer + line
      buffer = ''
      complete_line = complete_line.replace('static const IconData', '').strip()
      name = complete_line.split('=')[0].strip()
      name = to_kebab_case(name)
      code_point = complete_line.split('=')[1].replace('IconData(', '').split(',')[0].strip()
      if name in iconsax_linear_mapping:
        print(f'Iconsax Plus Linear: Duplicated icon name: {name}')
        continue

      iconsax_linear_mapping[name] = code_point

  print(f'Mapped {len(iconsax_linear_mapping)} icons from IconsaxPlusLinear')
  return iconsax_linear_mapping
