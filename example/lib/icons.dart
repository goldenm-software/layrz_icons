import 'package:flutter/material.dart';
import 'package:layrz_icons/layrz_icons.dart';
import 'package:layrz_icons_example/export_dialog.dart';
import 'package:layrz_icons_example/layout.dart';
import 'package:layrz_theme/layrz_theme.dart';

class IconView extends StatefulWidget {
  final LayrzFamily? family;
  const IconView({super.key, this.family});

  @override
  State<IconView> createState() => _IconViewState();
}

class _IconViewState extends State<IconView> {
  bool get _isMobile => MediaQuery.sizeOf(context).width < kSmallGrid;
  String _search = '';
  List<LayrzIcon> get rawIcons {
    if (widget.family == null) return iconMapping.values.toList();
    return iconMapping.values.where((icon) => icon.family == widget.family).toList();
  }

  List<LayrzIcon> get icons {
    return rawIcons.where((icon) {
      if (_search.isEmpty) return true;

      return icon.name.toLowerCase().contains(_search.toLowerCase());
    }).toList();
  }

  @override
  Widget build(BuildContext context) {
    return Layout(
      body: Column(
        children: [
          Row(
            children: [
              Expanded(
                child: Text(
                  "Total icons: ${rawIcons.length}",
                  style: Theme.of(context).textTheme.titleLarge?.copyWith(fontWeight: FontWeight.bold),
                ),
              ),
              ThemedSearchInput(
                value: _search,
                asField: !_isMobile,
                onSearch: (value) => setState(() => _search = value),
              ),
            ],
          ),
          Text("Clip to copy to clipboard", style: Theme.of(context).textTheme.bodyMedium),
          const SizedBox(height: 16),
          Expanded(
            child: LayoutBuilder(
              builder: (context, constraints) {
                int axisCount = 8;
                if (constraints.maxWidth < kExtraSmallGrid) {
                  axisCount = 4;
                } else if (constraints.maxWidth < kSmallGrid) {
                  axisCount = 7;
                } else if (constraints.maxWidth < kMediumGrid) {
                  axisCount = 10;
                } else if (constraints.maxWidth < kLargeGrid) {
                  axisCount = 14;
                }

                return GridView.builder(
                  gridDelegate: ThemedGridDelegateWithFixedHeight(
                    crossAxisCount: axisCount,
                    height: 80,
                  ),
                  itemCount: icons.length,
                  itemBuilder: (context, index) {
                    final icon = icons[index];

                    return InkWell(
                      onTap: () => showDialog(
                        context: context,
                        builder: (context) => ExportDialog(icon: icon),
                      ),
                      borderRadius: BorderRadius.circular(10),
                      child: ThemedTooltip(
                        message: icon.name,
                        child: Padding(
                          padding: const EdgeInsets.all(10),
                          child: Center(
                            child: Icon(icon.iconData, size: 30),
                          ),
                        ),
                      ),
                    );
                  },
                );
              },
            ),
          ),
          // Expanded(
          //   child: ListView.builder(
          //     itemCount: icons.length,
          //     itemExtent: 50,
          //     itemBuilder: (context, index) {
          //       final icon = icons[index];

          //       return InkWell(
          //         onTap: () => showDialog(
          //           context: context,
          //           builder: (context) => ExportDialog(icon: icon),
          //         ),
          //         child: Padding(
          //           padding: const EdgeInsets.all(5).add(const EdgeInsets.symmetric(horizontal: 10)),
          //           child: Row(
          //             children: [
          //               Icon(icon.iconData, size: 30),
          //               const SizedBox(width: 10),
          //               Expanded(child: Text(icon.name, style: Theme.of(context).textTheme.bodyMedium)),
          //             ],
          //           ),
          //         ),
          //       );
          //     },
          //   ),
          // ),
        ],
      ),
    );
  }
}
