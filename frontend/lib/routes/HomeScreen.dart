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
      body: buildHeroCards(),
      floatingActionButton: FloatingActionButton.extended(
        onPressed: () {
          Navigator.pushNamed(
            context,
            "a",
          );
        },
        label: const Text('Create New Hero'),
        icon: const Icon(Icons.person),
        backgroundColor: Color(0xff214375),
      ),
    );
  }

  Widget buildHeroCards() {
    return FutureBuilder<List<MyHero>>(
        future: heroes,
        builder: (context, snapshot) {
          if (snapshot.hasData) {
            List<MyHero> her = snapshot.data!;
            return GridView.builder(
                gridDelegate: const SliverGridDelegateWithMaxCrossAxisExtent(
                    maxCrossAxisExtent: 200,
                    childAspectRatio: 0.75,
                    crossAxisSpacing: 20,
                    mainAxisSpacing: 20),
                itemCount: her.length,
                itemBuilder: (BuildContext context, int index) {
                  return HeroCard(id: her[index].heroId, name: her[index].name);
                });
          } else if (snapshot.hasError && snapshot.data == null) {
            return Center(child: Text("No heroes."));
          } else if (snapshot.hasError) {
            return Center(child: Text("Error getting heroes."));
          } else {
            return Center(child: CircularProgressIndicator());
          }
        });
  }
}
