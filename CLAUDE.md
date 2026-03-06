# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

layrz_icons is a Flutter package that unifies 5 icon libraries (Material Design Icons, Solar Icons, Font Awesome, Ionicons, Iconsax Plus) into a single API. Published on [pub.dev](https://pub.dev/packages/layrz_icons), hosted at [icons.layrz.com](https://icons.layrz.com).

## Build & Development Commands

```bash
# Flutter version management (uses FVM, pinned to 3.32.7)
fvm use

# Get dependencies
flutter pub get

# Analyze (lint) - note: lib/src/*.dart excluded from analysis (auto-generated)
flutter analyze

# Run tests
flutter test

# Run example app
cd example && flutter run

# Build example for web (used in CI for icons.layrz.com)
cd example && flutter build web --wasm --base-href /layrz_icons/
```

**Important:** Apps using this package must pass `--no-tree-shake-icons` to Flutter build commands.

## Architecture

### Code Generation Pipeline

The core icon definitions in `lib/src/` are **auto-generated** — never edit them manually. They are produced by Python scripts in `generator/`:

1. **Scanners** (`generator/*.py`) parse icon definitions from Flutter package cache (pub cache)
2. **Generators** produce three Dart files:
   - `lib/src/class_enum.dart` — `LayrzIconsClasses` static class with `LayrzIcon` getters
   - `lib/src/icon_enum.dart` — `LayrzIcons` static class with `IconData` getters
   - `lib/src/mapping.dart` — `iconMapping` global `Map<String, LayrzIcon>` for runtime lookup
3. Entry point: copy `generator/mapper.example.py` to `mapper.py`, set `PUB_PATH` to your pub cache, and run it

Python tooling requires Python 3.12+ with `ruff` for formatting/linting.

### Public API (lib/layrz_icons.dart)

Three access patterns:
- **Static IconData:** `LayrzIcons.mdiAccount` — returns `IconData` directly
- **Static LayrzIcon:** `LayrzIconsClasses.mdiAccount` — returns `LayrzIcon` with name, codePoint, family
- **String lookup:** `iconMapping["mdi-account"]` — runtime lookup by kebab-case name

### Key Types

- `LayrzIcon` — holds icon name (kebab-case), codePoint, and `LayrzFamily`; has `.iconData` getter
- `LayrzFamily` — enum of 17 font families with `fontFamily` and `fontPackage` getters
- Naming convention: `{prefix}{IconName}` where prefix is `mdi`, `solarBold`, `solarOutline`, `fa{Weight}`, `ionicons`, `iconsax{Variant}`

### Adding a New Icon Library

1. Write a scanner in `generator/` to extract names and codepoints
2. Add a `LayrzFamily` enum entry with fontFamily/fontPackage
3. Update generator functions to include the new library
4. Regenerate all three Dart files

## Commit Conventions

This project uses [Conventional Commits](https://www.conventionalcommits.org/). Allowed types: `feat`, `fix`, `docs`, `perf`, `refactor`, `test`, `sec`, `lab`, `exp`, `deps`, `revert`, `chore`, `style`. Breaking changes use `!` suffix (e.g., `feat!: ...`).

## CI/CD

- **PR checks** (`checks.yaml`): lint, test with coverage, changelog validation
- **Publish** (`publish.yaml`, on `vX.Y.Z` tags): publish to pub.dev, create GitHub release, deploy example app to GitHub Pages
