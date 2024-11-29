library layrz_icons;

import 'package:flutter/widgets.dart';

part 'src/class_enum.dart';
part 'src/icon_enum.dart';
part 'src/mapping.dart';

class LayrzIcon {
  final String name;
  final int codePoint;
  final LayrzFamily family;

  const LayrzIcon({
    required this.name,
    required this.codePoint,
    required this.family,
  });

  IconData get iconData => IconData(codePoint, fontFamily: family.fontFamily, fontPackage: family.fontPackage);
}

enum LayrzFamily {
  materialDesignIcons,
  solarIconsBold,
  solarIconsOutline,
  fontAwesomeBrands,
  fontAwesomeSolid,
  fontAwesomeRegular,
  fontAwesomeLight,
  fontAwesomeDuotone,
  fontAwesomeThin,
  fontAwesomeSharpThin,
  fontAwesomeSharpLight,
  fontAwesomeSharpRegular,
  fontAwesomeSharpSolid,
  ionicons,
  iconsaxPlusBold,
  iconsaxPlusBroken,
  iconsaxPlusLinear,
  ;

  String get fontFamily {
    switch (this) {
      case LayrzFamily.materialDesignIcons:
        return 'Material Design Icons';
      case LayrzFamily.solarIconsBold:
        return 'SolarIconsBold';
      case LayrzFamily.solarIconsOutline:
        return 'SolarIconsOutline';
      case LayrzFamily.fontAwesomeBrands:
        return 'FontAwesomeBrands';
      case LayrzFamily.fontAwesomeSolid:
        return 'FontAwesomeSolid';
      case LayrzFamily.fontAwesomeRegular:
        return 'FontAwesomeRegular';
      case LayrzFamily.fontAwesomeLight:
        return 'FontAwesomeLight';
      case LayrzFamily.fontAwesomeDuotone:
        return 'FontAwesomeDuotone';
      case LayrzFamily.fontAwesomeThin:
        return 'FontAwesomeThin';
      case LayrzFamily.fontAwesomeSharpThin:
        return 'FontAwesomeSharpThin';
      case LayrzFamily.fontAwesomeSharpLight:
        return 'FontAwesomeSharpLight';
      case LayrzFamily.fontAwesomeSharpRegular:
        return 'FontAwesomeSharpRegular';
      case LayrzFamily.fontAwesomeSharpSolid:
        return 'FontAwesomeSharpSolid';
      case LayrzFamily.ionicons:
        return 'Ionicons';
      case LayrzFamily.iconsaxPlusBold:
        return 'IconsaxPlusBold';
      case LayrzFamily.iconsaxPlusBroken:
        return 'IconsaxPlusBroken';
      case LayrzFamily.iconsaxPlusLinear:
        return 'IconsaxPlusLinear';
    }
  }

  String get fontPackage {
    switch (this) {
      case LayrzFamily.materialDesignIcons:
        return 'material_design_icons_flutter';
      case LayrzFamily.solarIconsBold:
      case LayrzFamily.solarIconsOutline:
        return 'solar_icons';
      case LayrzFamily.fontAwesomeBrands:
      case LayrzFamily.fontAwesomeSolid:
      case LayrzFamily.fontAwesomeRegular:
      case LayrzFamily.fontAwesomeLight:
      case LayrzFamily.fontAwesomeDuotone:
      case LayrzFamily.fontAwesomeThin:
      case LayrzFamily.fontAwesomeSharpThin:
      case LayrzFamily.fontAwesomeSharpLight:
      case LayrzFamily.fontAwesomeSharpRegular:
      case LayrzFamily.fontAwesomeSharpSolid:
        return 'font_awesome_flutter';
      case LayrzFamily.ionicons:
        return 'ionicons';
      case LayrzFamily.iconsaxPlusBold:
      case LayrzFamily.iconsaxPlusBroken:
      case LayrzFamily.iconsaxPlusLinear:
        return 'iconsax_plus';
    }
  }
}
