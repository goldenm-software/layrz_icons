# Layrz Icons

Managed by Golden M, Inc.

## Motivation
Sometimes, you need to use icons in your project, but using a single library? Maybe too boring, using multiple? Well, works but can be a mess. So, we created Layrz Icons, a simple and easy to use icon library, designed to combine multiple libraries in a single one, also adding a mapping object to convert any `IconData` into a string.

## Supported libraries
Right now, `layrz_icons` integrates three libraries:
- [material_design_icons_flutter](https://pub.dev/packages/material_design_icons_flutter)
- [font_awesome_flutter](https://pub.dev/packages/font_awesome_flutter)
- [solar_icons](https://pub.dev/packages/solar_icons)

Many thanks to the developers of these libraries, they are awesome!

## How to add a new icon library?
You have two options to do that, opening an Issue and we'll add it for you, or you can do it yourself, just follow the steps below:

1. Duplicate the file `mapper.example.py` and rename it to `mapper.py`
2. Find the `<INSERT HERE YOUR NEW LIBRARY>` string and replace it with a similar workflow used on the other libraries
3. Run the script `python3 mapper.py` to generate the new mapping object.
4. Create a PR to check on our end if everyting is OK. If is the case, we'll merge it.

## Disclaimer
We don't attempt to copy or steal any icon, we just want to make it easier to use them in your projects. Also, if you want to add a new icon library, please make sure to check the license of the library, we don't want to have any legal issues.

## Other projects
By us, we have multiple open source libraries on multiple languages, for Dart and Flutter:
- [layrz_theme](https://pub.dev/packages/layrz_theme),a library to create themes for your Flutter app, like a classic CSS framework
- [layrz_models](https://pub.dev/packages/layrz_models), handle the entities from Layrz API in a simple way.
- [layrz_logging](https://pub.dev/packages/layrz_logging), logging never was so easy. A simple and easy to use logging library.
- [wiatag_kit](https://pub.dev/packages/wiatag_kit), convert any mobile device into a GPS tracker, using the WiaTag from Gurtam, Inc.
- [wialon](https://pub.dev/packages/wialon), a library to connect to Wialon Remote API, from Gurtam, Inc.

Also, check out our other packages on [PyPi of Golden M](https://pypi.org/user/goldenm/), [PyPi of Layrz](https://pypi.org/user/layrz-software/), and [NPM of Golden M](https://www.npmjs.com/~goldenm) or [RubyGems of Golden M](https://rubygems.org/profiles/goldenm).

## Interested in our work?
Golden M is a software/hardware development company what is working on a new, innovative and disruptive technologies.
For more information, contact us at [sales@layrz.com](mailto:sales@layrz.com)

## License
This project is under MIT License, for more information, check out the `LICENCE`
