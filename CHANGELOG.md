# Changelog

## 1.1.1

- Bump the example app's `layrz_theme` to `^7.6.0` and `layrz_models` to `^3.11.0`, which adopt the renamed `solar*` icon getters so the demo site builds against `layrz_icons` 1.1.x.

## 1.1.0

- **BREAKING:** Replaced `material_design_icons_flutter` with `flutter_material_design_icons` as the Material Design Icons source (now 7447 icons). The `mdi*` getters and `mdi-*` keys are unchanged.
- **BREAKING:** Removed the `solar_icons`, `ionicons`, and `iconsax_plus` libraries. The `ionicons*`, `iconsax*`, and the old `solar_icons`-backed `solarBold`/`solarOutline`/`solarBroken` getters and their keys are gone.
- **BREAKING:** Upgraded `font_awesome_flutter` to `^11.0.0`. Only the free families remain: Brands, Regular and Solid. The Light, Duotone, Thin, and Sharp families have been removed.
- **BREAKING:** Renamed the `flutty_solar_icons` output from `fluttySolar*` to `solar*` (getters `solarBold*`/`solarBroken*`/`solarLinear*`/`solarOutline*`, keys `solar-bold-*`/`solar-broken-*`/`solar-linear-*`/`solar-outline-*`). These now back the `solar` namespace.

## 1.0.7

- Build site with Flutter 3.41.4

## 1.0.6

- Added new icon library `flutty_solar_icons` to get more variants of the solar icons, now you can choose between `solar_icons` and `flutty_solar_icons` for the solar icons.

## 1.0.5

- Updated README.md

## 1.0.4

- Fixed an issue related to Font Awesome icons mapping

## 1.0.3

- Optimizations
- Added `ionicons` to the integrated icons
- Added `iconsax_plus` to the integrated icons

## 1.0.2

- Added `LayrzIconsClasses` to simplify the class calling.

## 1.0.1

- Added `homepage` on `pubspec.yaml`
- Updated `README.md`

## 1.0.0

- Initial release
