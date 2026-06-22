library;

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

  // ignore: non_const_argument_for_const_parameter
  IconData get iconData => IconData(codePoint, fontFamily: family.fontFamily, fontPackage: family.fontPackage);
}

enum LayrzFamily {
  materialDesignIcons,
  fontAwesomeBrands,
  fontAwesomeSolid,
  fontAwesomeRegular,
  solarBold,
  solarBroken,
  solarLinear,
  solarOutline,
  ;

  String get fontFamily {
    switch (this) {
      case LayrzFamily.materialDesignIcons:
        return 'Material Design Icons';
      case LayrzFamily.fontAwesomeBrands:
        return 'FontAwesomeBrands';
      case LayrzFamily.fontAwesomeSolid:
        return 'FontAwesomeSolid';
      case LayrzFamily.fontAwesomeRegular:
        return 'FontAwesomeRegular';
      case LayrzFamily.solarBold:
        return 'SolarBold';
      case LayrzFamily.solarBroken:
        return 'SolarBroken';
      case LayrzFamily.solarLinear:
        return 'SolarLinear';
      case LayrzFamily.solarOutline:
        return 'SolarOutline';
    }
  }

  String get fontPackage {
    switch (this) {
      case LayrzFamily.materialDesignIcons:
        return 'flutter_material_design_icons';
      case LayrzFamily.fontAwesomeBrands:
      case LayrzFamily.fontAwesomeSolid:
      case LayrzFamily.fontAwesomeRegular:
        return 'font_awesome_flutter';
      case LayrzFamily.solarBold:
      case LayrzFamily.solarBroken:
      case LayrzFamily.solarLinear:
      case LayrzFamily.solarOutline:
        return 'flutty_solar_icons';
    }
  }
}
