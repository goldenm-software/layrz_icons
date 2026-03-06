import 'package:flutter/material.dart';
import 'package:layrz_icons_example/router.dart';
import 'package:layrz_models/layrz_models.dart';
import 'package:layrz_theme/layrz_theme.dart';

const kFont = AppFont(source: .google, name: 'Ubuntu Mono');

void main() async {
  WidgetsFlutterBinding.ensureInitialized();

  await ThemedFontHandler.preloadFont(kFont);

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
      theme: generateLightTheme(titleFont: kFont, bodyFont: kFont),
      routerConfig: router,
      builder: (context, child) {
        return ThemedSnackbarMessenger(child: child ?? const SizedBox());
      },
    );
  }
}
