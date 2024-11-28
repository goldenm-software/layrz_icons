import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:layrz_icons/layrz_icons.dart';
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
  String _search = '';
  List<LayrzIcon> get icons {
    return iconMapping.values.where((icon) {
      if (_search.isEmpty) return true;

      return icon.name.toLowerCase().contains(_search.toLowerCase());
    }).toList();
  }

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      theme: generateLightTheme(titleFont: titleFont, bodyFont: bodyFont),
      home: Scaffold(
        body: Padding(
          padding: const EdgeInsets.all(10),
          child: Column(
            children: [
              Row(
                children: [
                  Expanded(
                    child: Text(
                      "Total icons integrated: ${iconMapping.length}",
                      style: Theme.of(context).textTheme.titleLarge,
                    ),
                  ),
                  ThemedSearchInput(
                    value: _search,
                    onSearch: (value) => setState(() => _search = value),
                  ),
                ],
              ),
              Text("Clip to copy to clipboard", style: Theme.of(context).textTheme.bodyMedium),
              const SizedBox(height: 16),
              Expanded(
                child: LayoutBuilder(
                  builder: (context, constraints) {
                    int axisCount = 2;
                    if (constraints.maxWidth < kSmallGrid) {
                      axisCount = 3;
                    } else if (constraints.maxWidth < kMediumGrid) {
                      axisCount = 4;
                    } else if (constraints.maxWidth < kLargeGrid) {
                      axisCount = 5;
                    }

                    return GridView.builder(
                      gridDelegate: ThemedGridDelegateWithFixedHeight(crossAxisCount: axisCount, height: 40),
                      itemCount: icons.length,
                      itemBuilder: (context, index) {
                        final icon = icons[index];

                        return InkWell(
                          onTap: () {
                            Clipboard.setData(ClipboardData(text: 'LayrzIcons.${_toCamelCase(icon.name)}'));
                            ScaffoldMessenger.of(context).showSnackBar(
                              SnackBar(
                                content: Text('Copied LayrzIcons.${_toCamelCase(icon.name)} to clipboard'),
                              ),
                            );
                          },
                          child: Padding(
                            padding: const EdgeInsets.all(5),
                            child: Row(
                              children: [
                                Icon(icon.iconData, size: 20),
                                const SizedBox(width: 10),
                                Expanded(
                                  child: Text(
                                    icon.name,
                                    style: Theme.of(context).textTheme.bodySmall,
                                  ),
                                ),
                              ],
                            ),
                          ),
                        );
                      },
                    );
                  },
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }

  String _toCamelCase(String text) {
    final parts = text.split('-');
    final first = parts.removeAt(0);
    return first + parts.map((part) => part[0].toUpperCase() + part.substring(1)).join();
  }
}
