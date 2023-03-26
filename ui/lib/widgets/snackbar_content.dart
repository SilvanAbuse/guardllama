import 'package:flutter/material.dart';

class SnackbarContent extends StatelessWidget {
  const SnackbarContent({
    Key? key,
    required this.message,
    required this.backgroundColor,
    required this.iconData,
  }) : super(key: key);

  final String message;
  final Color backgroundColor;
  final IconData iconData;

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: const EdgeInsets.all(8),
      decoration: BoxDecoration(
        color: backgroundColor,
        borderRadius: BorderRadius.circular(10),
      ),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        crossAxisAlignment: CrossAxisAlignment.center,
        children: [
          Expanded(
            child: Row(
              mainAxisAlignment: MainAxisAlignment.start,
              children: [
                const SizedBox(width: 4),
                Icon(iconData, color: Colors.white),
                const SizedBox(width: 12),
                Flexible(
                  child: Column(
                    mainAxisSize: MainAxisSize.min,
                    crossAxisAlignment: CrossAxisAlignment.start,
                    mainAxisAlignment: MainAxisAlignment.center,
                    children: [
                      Text(
                        message,
                        style: const TextStyle(
                          color: Colors.white,
                          fontSize: 16,
                        ),
                      ),
                      // Text(
                      //   message,
                      //   style: const TextStyle(color: Colors.white),
                      // ),
                    ],
                  ),
                ),
              ],
            ),
          ),
          TextButton(
              onPressed: () => _closeSnackbar(context),
              child: const Icon(Icons.close_rounded))
        ],
      ),
    );
  }

  void _closeSnackbar(BuildContext context) {
    ScaffoldMessenger.of(context).hideCurrentSnackBar();
  }
}
