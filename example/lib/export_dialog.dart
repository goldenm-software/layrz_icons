import 'package:flutter/material.dart';
import 'package:flutter/rendering.dart';
import 'package:flutter/services.dart';
import 'package:layrz_icons/layrz_icons.dart';
import 'package:layrz_theme/layrz_theme.dart';
import 'dart:ui' as ui;

class ExportDialog extends StatefulWidget {
  final LayrzIcon icon;

  const ExportDialog({
    super.key,
    required this.icon,
  });

  @override
  State<ExportDialog> createState() => _ExportDialogState();
}

class _ExportDialogState extends State<ExportDialog> {
  final GlobalKey _iconKey = GlobalKey();
  Color _color = Colors.black;
  @override
  Widget build(BuildContext context) {
    return Dialog(
      elevation: 0,
      backgroundColor: Colors.transparent,
      child: Container(
        constraints: const BoxConstraints(maxWidth: 400),
        padding: const EdgeInsets.all(16),
        decoration: generateContainerElevation(context: context, elevation: 5, radius: 16),
        child: Column(
          mainAxisSize: MainAxisSize.min,
          children: [
            Text(
              "Export ${widget.icon.name}",
              style: Theme.of(context).textTheme.bodyLarge?.copyWith(fontWeight: FontWeight.bold),
            ),
            const SizedBox(height: 8),
            RepaintBoundary(
              key: _iconKey,
              child: Icon(
                widget.icon.iconData,
                size: 60,
                color: _color,
              ),
            ),
            const SizedBox(height: 8),
            ThemedColorPicker(
              labelText: 'Color',
              value: _color,
              padding: EdgeInsets.zero,
              onChanged: (color) => setState(() => _color = color),
            ),
            const SizedBox(height: 8),
            Row(
              children: [
                Expanded(
                  child: ThemedButton(
                    labelText: 'Copy enum value',
                    icon: LayrzIcons.solarOutlineCopy,
                    color: Colors.blue,
                    onTap: () {
                      Clipboard.setData(ClipboardData(text: 'LayrzIcons.${_toCamelCase(widget.icon.name)}'));
                      ThemedSnackbarMessenger.of(context).showSnackbar(ThemedSnackbar(
                        color: Colors.green,
                        icon: LayrzIcons.solarOutlineCopy,
                        message: 'Copied to clipboard',
                      ));
                    },
                  ),
                ),
                const SizedBox(width: 8),
                Expanded(
                  child: ThemedButton(
                    labelText: 'Export as PNG',
                    icon: LayrzIcons.solarOutlineDownloadMinimalistic,
                    color: Colors.orange,
                    onTap: () async {
                      try {
                        RenderRepaintBoundary boundary =
                            _iconKey.currentContext!.findRenderObject() as RenderRepaintBoundary;
                        Size size = boundary.size;

                        double? targetRatio = await showDialog<double>(
                          context: context,
                          builder: (context) => const SelectAspectRatio(),
                        );

                        if (targetRatio == null) {
                          if (!context.mounted) return;
                          ThemedSnackbarMessenger.of(context).showSnackbar(ThemedSnackbar(
                            color: Colors.red,
                            icon: LayrzIcons.solarOutlineCloseSquare,
                            message: "You don't picked a target size",
                          ));
                          return;
                        }

                        double aspectRatio = targetRatio / size.width;
                        ui.Image image = await boundary.toImage(pixelRatio: aspectRatio);
                        ByteData? byteData = await image.toByteData(format: ui.ImageByteFormat.png);
                        Uint8List pngBytes = byteData!.buffer.asUint8List();
                        saveFile(
                          filename: '${widget.icon.name}.png',
                          bytes: pngBytes,
                        );
                      } catch (e) {
                        debugPrint('Failed to export image: $e');
                        if (!context.mounted) return;
                        ThemedSnackbarMessenger.of(context).showSnackbar(ThemedSnackbar(
                          color: Colors.red,
                          icon: LayrzIcons.solarOutlineCloseSquare,
                          message: 'Failed to export image',
                        ));
                        return;
                      }
                    },
                  ),
                ),
              ],
            ),
          ],
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

class SelectAspectRatio extends StatefulWidget {
  const SelectAspectRatio({super.key});

  @override
  State<SelectAspectRatio> createState() => _SelectAspectRatioState();
}

class _SelectAspectRatioState extends State<SelectAspectRatio> {
  List<double> get sizes => [4096, 2048, 1024, 512, 256, 128, 64];
  @override
  Widget build(BuildContext context) {
    return Dialog(
      elevation: 0,
      backgroundColor: Colors.transparent,
      child: Container(
        constraints: const BoxConstraints(maxWidth: 300),
        padding: const EdgeInsets.all(16),
        decoration: generateContainerElevation(context: context, elevation: 5, radius: 16),
        child: SingleChildScrollView(
          child: Column(
            children: [
              Text(
                "Select a target size",
                style: Theme.of(context).textTheme.bodyLarge?.copyWith(fontWeight: FontWeight.bold),
              ),
              const SizedBox(height: 8),
              ListView.builder(
                shrinkWrap: true,
                itemCount: sizes.length,
                itemBuilder: (context, index) {
                  final size = sizes[index];
                  return Material(
                    color: Colors.transparent,
                    child: InkWell(
                      onTap: () => Navigator.of(context).pop(size),
                      child: Padding(
                        padding: const EdgeInsets.all(8),
                        child: Text(
                          '${size.toInt()}px x ${size.toInt()}px',
                          style: Theme.of(context).textTheme.bodyMedium,
                          textAlign: TextAlign.center,
                        ),
                      ),
                    ),
                  );
                },
              ),
            ],
          ),
        ),
      ),
    );
  }
}
