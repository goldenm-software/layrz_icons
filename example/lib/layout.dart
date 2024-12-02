import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:layrz_icons/layrz_icons.dart';
import 'package:layrz_models/layrz_models.dart';
import 'package:layrz_theme/layrz_theme.dart';
import 'package:url_launcher/url_launcher_string.dart';

class Layout extends StatelessWidget {
  final Widget body;
  const Layout({
    super.key,
    required this.body,
  });

  @override
  Widget build(BuildContext context) {
    String currentPath = '';
    try {
      currentPath = GoRouterState.of(context).fullPath ?? '';
    } catch (e) {
      currentPath = '/';
    }

    return ThemedLayout(
      currentPath: currentPath,
      body: body,
      appTitle: 'layrz_icons',
      homePath: '/allIcons',
      enableNotifications: false,
      logo: const AppThemedAsset(
        normal: 'https://cdn.layrz.com/resources/layrz/logo/normal.png',
        white: 'https://cdn.layrz.com/resources/layrz/logo/white.png',
      ),
      favicon: const AppThemedAsset(
        normal: 'https://cdn.layrz.com/resources/layrz/favicon/normal.png',
        white: 'https://cdn.layrz.com/resources/layrz/favicon/white.png',
      ),
      onNavigatorPop: context.pop,
      onNavigatorPush: context.go,
      items: [
        ThemedNavigatorAction(
          labelText: 'GitHub repository',
          icon: LayrzIcons.faBrandsGithub,
          onTap: () => launchUrlString(
            'https://github.com/goldenm-software/layrz_icons',
            mode: LaunchMode.externalApplication,
          ),
        ),
        ThemedNavigatorSeparator(type: ThemedSeparatorType.dots),
        ThemedNavigatorPage(
          path: '/allIcons',
          labelText: 'All icons',
          icon: LayrzIcons.solarOutlineAlbum,
        ),
        ThemedNavigatorPage(
          labelText: 'Material design icons',
          path: '/${LayrzFamily.materialDesignIcons.fontPackage}',
          icon: LayrzIcons.mdiHomeVariant,
          children: [LayrzFamily.materialDesignIcons].map((family) {
            return ThemedNavigatorPage(
              path: '/${family.fontPackage}/${family.fontFamily.replaceAll(' ', '_')}',
              labelText: family.fontFamily,
              icon: family.icon,
            );
          }).toList(),
        ),
        ThemedNavigatorPage(
          labelText: 'Solar icons',
          icon: LayrzIcons.solarOutlineHomeAngle,
          path: '/${LayrzFamily.solarIconsBold.fontPackage}',
          children: [LayrzFamily.solarIconsBold, LayrzFamily.solarIconsOutline].map((family) {
            return ThemedNavigatorPage(
              path: '/${family.fontPackage}/${family.fontFamily.replaceAll(' ', '_')}',
              labelText: family.fontFamily,
              icon: family.icon,
            );
          }).toList(),
        ),
        ThemedNavigatorPage(
          labelText: 'Ionicons',
          path: '/${LayrzFamily.ionicons.fontPackage}',
          icon: LayrzFamily.ionicons.icon,
          children: [LayrzFamily.ionicons].map((family) {
            return ThemedNavigatorPage(
              path: '/${family.fontPackage}/${family.fontFamily.replaceAll(' ', '_')}',
              labelText: family.fontFamily,
              icon: family.icon,
            );
          }).toList(),
        ),
        ThemedNavigatorPage(
          labelText: 'Iconsax Plus',
          path: '/${LayrzFamily.iconsaxPlusBold.fontPackage}',
          icon: LayrzFamily.iconsaxPlusBold.icon,
          children:
              [LayrzFamily.iconsaxPlusBold, LayrzFamily.iconsaxPlusBroken, LayrzFamily.iconsaxPlusLinear].map((family) {
            return ThemedNavigatorPage(
              path: '/${family.fontPackage}/${family.fontFamily.replaceAll(' ', '_')}',
              labelText: family.fontFamily,
              icon: family.icon,
            );
          }).toList(),
        ),
        ThemedNavigatorPage(
          labelText: 'Font awesome',
          path: '/${LayrzFamily.fontAwesomeBrands.fontPackage}',
          icon: LayrzFamily.fontAwesomeSolid.icon,
          children: [LayrzFamily.fontAwesomeBrands, LayrzFamily.fontAwesomeSolid, LayrzFamily.fontAwesomeRegular]
              .map((family) {
            return ThemedNavigatorPage(
              path: '/${family.fontPackage}/${family.fontFamily.replaceAll(' ', '_')}',
              labelText: family.fontFamily,
              icon: family.icon,
            );
          }).toList(),
        ),
      ],
    );
  }
}

extension Icons on LayrzFamily {
  IconData get icon {
    switch (this) {
      case LayrzFamily.materialDesignIcons:
        return LayrzIcons.mdiHomeVariant;

      case LayrzFamily.solarIconsBold:
        return LayrzIcons.solarBoldHomeAngle;

      case LayrzFamily.solarIconsOutline:
        return LayrzIcons.solarOutlineHomeAngle;

      case LayrzFamily.ionicons:
        return LayrzIcons.ioniconsHome;

      case LayrzFamily.iconsaxPlusBold:
        return LayrzIcons.iconsaxBoldHome;

      case LayrzFamily.iconsaxPlusBroken:
        return LayrzIcons.iconsaxBrokenHome;

      case LayrzFamily.iconsaxPlusLinear:
        return LayrzIcons.iconsaxLinearHome;

      case LayrzFamily.fontAwesomeBrands:
        return LayrzIcons.faBrandsChrome;

      case LayrzFamily.fontAwesomeSolid:
        return LayrzIcons.faSolidHouse;

      case LayrzFamily.fontAwesomeRegular:
        return LayrzIcons.faRegularBuilding;

      default:
        return LayrzIcons.mdiHomeVariant;
    }
  }
}
