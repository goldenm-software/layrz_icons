import 'package:flutter/material.dart';
import 'package:layrz_icons/layrz_icons.dart';
import 'package:layrz_icons_example/router.dart';
import 'package:layrz_models/layrz_models.dart';
import 'package:layrz_theme/layrz_theme.dart';

const titleFont = AppFont(source: FontSource.google, name: 'Ubuntu');
const bodyFont = AppFont(source: FontSource.google, name: 'Ubuntu');

void main() async {
  WidgetsFlutterBinding.ensureInitialized();

  await ThemedFontHandler.preloadFont(titleFont);
  await ThemedFontHandler.preloadFont(bodyFont);

  runApp(const MyApp());
}

class MyApp extends StatefulWidget {
  const MyApp({super.key});

  @override
  State<MyApp> createState() => _MyAppState();
}

class _MyAppState extends State<MyApp> {
  @override
  Widget build(BuildContext context) {
    return MaterialApp.router(
      debugShowCheckedModeBanner: false,
      theme: generateLightTheme(titleFont: titleFont, bodyFont: bodyFont),
      routerConfig: router,
      builder: (context, child) {
        return ThemedSnackbarMessenger(
          child: child ?? const SizedBox(),
        );
      },
    );
  }
}
