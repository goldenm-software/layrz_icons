import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:layrz_icons/layrz_icons.dart';
import 'package:layrz_theme/layrz_theme.dart';
import 'package:url_launcher/url_launcher_string.dart';

class Layout extends StatelessWidget {
  final Widget body;
  const Layout({super.key, required this.body});

  @override
  Widget build(BuildContext context) {
    String currentPath = '';
    try {
      currentPath = GoRouterState.of(context).fullPath ?? '';
    } catch (e) {
      currentPath = '/';
    }

    return ThemedLayout(
      style: .sidebar,
      mobileStyle: .appBar,
      userName: "Layrz Icons Showcase",
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
      isBackEnabled: false,
      onNavigatorPop: context.pop,
      onNavigatorPush: context.go,
      items: [
        ThemedNavigatorAction(
          labelText: 'GitHub repository',
          icon: LayrzIcons.faBrandsGithub,
          onTap: () =>
              launchUrlString('https://github.com/goldenm-software/layrz_icons', mode: LaunchMode.externalApplication),
        ),
        ThemedNavigatorSeparator(type: ThemedSeparatorType.dots),
        ThemedNavigatorPage(path: '/allIcons', labelText: 'All icons', icon: LayrzIcons.solarOutlineAlbum),
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
          path: '/${LayrzFamily.solarBold.fontPackage}',
          children:
              [
                LayrzFamily.solarBold,
                LayrzFamily.solarBroken,
                LayrzFamily.solarLinear,
                LayrzFamily.solarOutline,
              ].map((family) {
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
          children: [LayrzFamily.fontAwesomeBrands, LayrzFamily.fontAwesomeSolid, LayrzFamily.fontAwesomeRegular].map((
            family,
          ) {
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

      case LayrzFamily.fontAwesomeBrands:
        return LayrzIcons.faBrandsChrome;

      case LayrzFamily.fontAwesomeSolid:
        return LayrzIcons.faSolidHouseChimney;

      case LayrzFamily.fontAwesomeRegular:
        return LayrzIcons.faRegularBuilding;

      case LayrzFamily.solarBold:
        return LayrzIcons.solarBoldHomeInEssentionalUi;

      case LayrzFamily.solarBroken:
        return LayrzIcons.solarBrokenHomeInEssentionalUi;

      case LayrzFamily.solarLinear:
        return LayrzIcons.solarLinearHomeInEssentionalUi;

      case LayrzFamily.solarOutline:
        return LayrzIcons.solarOutlineHomeInEssentionalUi;
    }
  }
}
