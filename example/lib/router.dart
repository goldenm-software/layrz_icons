import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:layrz_icons/layrz_icons.dart';
import 'package:layrz_icons_example/icons.dart';
import 'package:layrz_icons_example/not_found.dart';
import 'package:layrz_theme/layrz_theme.dart';

Page<void> customTransitionBuilder(BuildContext context, GoRouterState state, Widget child) {
  return CustomTransitionPage(
    key: state.pageKey,
    child: child,
    transitionDuration: kPageTransitionDuration,
    transitionsBuilder: (context, animation, secondaryAnimation, child) {
      return FadeTransition(
        opacity: animation,
        child: child,
      );
    },
  );
}

final goRoutes = [
  GoRoute(
    path: '/allIcons',
    pageBuilder: (context, state) => customTransitionBuilder(context, state, const IconView()),
  ),
  ...LayrzFamily.values.map((family) {
    return GoRoute(
      path: '/${family.fontPackage}/${family.fontFamily.replaceAll(' ', '_')}',
      pageBuilder: (context, state) => customTransitionBuilder(context, state, IconView(family: family)),
    );
  }),
];

final router = GoRouter(
  initialLocation: '/allIcons',
  errorPageBuilder: (context, state) => customTransitionBuilder(context, state, const NotFoundView()),
  routes: goRoutes,
);
