import 'package:flutter/widgets.dart';

import 'family.dart';

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
