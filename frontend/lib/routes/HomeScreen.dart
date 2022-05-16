import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:frontend/components/heroCard.dart';
import 'package:frontend/models/hero.dart';
import 'package:frontend/services/heroService.dart';

class HomeScreen extends StatefulWidget {
  HomeScreen({Key? key}) : super(key: key);

  @override
  _HomeScreenState createState() => _HomeScreenState();
}

class _HomeScreenState extends State<HomeScreen> {
  HeroService heroService = new HeroService();
  //late - Because variable is initialized later
  late Future<List<MyHero>> heroes;

  @override
  void initState() {
    super.initState();
    this.heroes = heroService.getHeroes();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        appBar: AppBar(
          title: Text("Heroes App"),
        ),
        body: buildHeroCards());
  }

  Widget buildHeroCards() {
    return FutureBuilder<List<MyHero>>(
        future: heroes,
        builder: (context, snapshot) {
          if (snapshot.hasData) {
            List<MyHero> her = snapshot.data!;
            return LayoutBuilder(builder: (context, constraints) {
              return GridView.builder(
                  gridDelegate: SliverGridDelegateWithFixedCrossAxisCount(
                    crossAxisCount: 3,
                    crossAxisSpacing: 5,
                    mainAxisSpacing: 5,
                    childAspectRatio: 0.75,
                  ),
                  itemCount: her.length,
                  itemBuilder: (BuildContext context, int index) {
                    return HeroCard(
                        id: her[index].heroId, name: her[index].name);
                  });
            });
          } else {
            return Center(child: CircularProgressIndicator());
          }
        });
  }
}
