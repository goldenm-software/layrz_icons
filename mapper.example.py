from pathlib import Path

from generator.class_enum_generator import generate_class_enum
from generator.font_awesome_flutter import scan_font_awesome_flutter
from generator.icon_enum_generator import generate_icon_enum
from generator.iconsax import scan_iconsax_bold, scan_iconsax_broken, scan_iconsax_linear
from generator.ionicons import scan_ionicons
from generator.mapping_generator import generate_mapping
from generator.material_design_icons_flutter import scan_material_design_icons_flutter
from generator.solar_icons import scan_solar_icons_bold, scan_solar_icons_outline

BASE_DIR = Path(__file__).resolve().parent
PUB_PATH = Path('path/to/your/pub/cache/hosted/directory')
# Windows example: PUB_PATH = Path('C:\\Users\\myuser\\AppData\\Local\\Pub\\Cache\\hosted\\pub.dev')

print('Generating icon mappings...')
mdi_icon_mapping = scan_material_design_icons_flutter(PUB_PATH)
solar_icon_bold_mapping = scan_solar_icons_bold(PUB_PATH)
solar_icon_outline_mapping = scan_solar_icons_outline(PUB_PATH)
font_awesome_mapping = scan_font_awesome_flutter(PUB_PATH)
ionicons_mapping = scan_ionicons(PUB_PATH)
iconsax_bold_mapping = scan_iconsax_bold(PUB_PATH)
iconsax_broken_mapping = scan_iconsax_broken(PUB_PATH)
iconsax_linear_mapping = scan_iconsax_linear(PUB_PATH)
# Create your own library mappings here
print('Icon mappings generated.')

print('Creating class_enum.dart ...')
generate_class_enum(
  base_dir=BASE_DIR,
  mdi_icon_mapping=mdi_icon_mapping,
  solar_icon_bold_mapping=solar_icon_bold_mapping,
  solar_icon_outline_mapping=solar_icon_outline_mapping,
  font_awesome_mapping=font_awesome_mapping,
  ionicons_mapping=ionicons_mapping,
  iconsax_bold_mapping=iconsax_bold_mapping,
  iconsax_broken_mapping=iconsax_broken_mapping,
  iconsax_linear_mapping=iconsax_linear_mapping,
  # Add your new mapping here
)
print('class_enum.dart created.')
print('Creating icon_enum.dart ...')
generate_icon_enum(
  base_dir=BASE_DIR,
  mdi_icon_mapping=mdi_icon_mapping,
  solar_icon_bold_mapping=solar_icon_bold_mapping,
  solar_icon_outline_mapping=solar_icon_outline_mapping,
  font_awesome_mapping=font_awesome_mapping,
  ionicons_mapping=ionicons_mapping,
  iconsax_bold_mapping=iconsax_bold_mapping,
  iconsax_broken_mapping=iconsax_broken_mapping,
  iconsax_linear_mapping=iconsax_linear_mapping,
  # Add your new mapping here
)
print('icon_enum.dart created.')
print('Creating mapping.dart ...')
generate_mapping(
  base_dir=BASE_DIR,
  mdi_icon_mapping=mdi_icon_mapping,
  solar_icon_bold_mapping=solar_icon_bold_mapping,
  solar_icon_outline_mapping=solar_icon_outline_mapping,
  font_awesome_mapping=font_awesome_mapping,
  ionicons_mapping=ionicons_mapping,
  iconsax_bold_mapping=iconsax_bold_mapping,
  iconsax_broken_mapping=iconsax_broken_mapping,
  iconsax_linear_mapping=iconsax_linear_mapping,
  # Add your new mapping here
)
print('mapping.dart created.')
