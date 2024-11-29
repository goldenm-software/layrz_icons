from pathlib import Path
from typing import Dict

from generator.utils import to_kebab_case


def scan_ionicons(pub_path: Path) -> Dict[str, str]:
  dart_file = pub_path / 'ionicons-0.2.2' / 'lib'
  dart_file = dart_file / 'ionicons.dart'

  with open(dart_file, 'r', encoding='utf-8') as f:
    lines = f.readlines()

  ionicons_mapping = {}
  for line in lines:
    line = line.strip()
    if not line.startswith('static const'):
      continue

    line = line.replace('static const', '').strip()
    name = line.split('=')[0].strip()
    name = to_kebab_case(name)

    code_point = line.split('=')[1].strip().replace(';', '')

    code_point = code_point.split('(')[1].split(')')[0].strip()
    ionicons_mapping[name] = code_point

  print(f'Mapped {len(ionicons_mapping)} icons from Ionicons')
  return ionicons_mapping
