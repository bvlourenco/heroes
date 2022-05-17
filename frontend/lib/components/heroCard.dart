import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

class HeroCard extends StatelessWidget {
  final String id;
  late final String name;

  HeroCard({Key? key, required this.id, required this.name})
      : super(key: key) {}

  @override
  Widget build(BuildContext context) {
    return Container(
        alignment: Alignment.center,
        decoration: BoxDecoration(
            color: Theme.of(context).cardColor,
            borderRadius: BorderRadius.circular(5),
            border: Border.all(color: Color(0xff214375))),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.spaceEvenly,
          children: <Widget>[
            Expanded(
                flex: 2,
                child: Image(
                    fit: BoxFit.contain, image: AssetImage('noImage.png'))),
            SizedBox(height: 5),
            Text(this.name),
            SizedBox(height: 5),
            Column(
              mainAxisAlignment: MainAxisAlignment.center,
              children: <Widget>[
                ElevatedButton.icon(
                  icon: Icon(Icons.edit, color: Colors.white),
                  label: Text("Edit"),
                  onPressed: () {},
                ),
                SizedBox(height: 5),
                ElevatedButton.icon(
                  icon: Icon(Icons.delete, color: Colors.white),
                  label: Text("Clear"),
                  onPressed: () {},
                  style: ElevatedButton.styleFrom(primary: Colors.red),
                ),
              ],
            )
          ],
        ));
  }
}
