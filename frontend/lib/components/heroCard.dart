import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

class HeroCard extends StatelessWidget {
  final String id;
  late final String name;

  HeroCard({Key? key, required this.id, required this.name})
      : super(key: key) {}

  @override
  Widget build(BuildContext context) {
    return Center(
        child: Card(
          child: SizedBox(
            width: 300,
            height: 100,
            child: Text(this.name),
          )
        )
      );
  }
}
