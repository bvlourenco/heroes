import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
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
  List<MyHero> heroesLoaded = [];

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
        onPressed: () => showCreateDialog(),
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
            heroesLoaded = snapshot.data!;
            return GridView.builder(
                gridDelegate: const SliverGridDelegateWithMaxCrossAxisExtent(
                    maxCrossAxisExtent: 300,
                    childAspectRatio: 0.75,
                    crossAxisSpacing: 20,
                    mainAxisSpacing: 20),
                itemCount: heroesLoaded.length,
                itemBuilder: (BuildContext context, int index) {
                  return heroCard(
                      heroesLoaded[index].heroId, heroesLoaded[index].name);
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

  showDeleteDialog(BuildContext context, String id, String name) {
    return showDialog<String>(
      context: context,
      builder: (BuildContext context) => AlertDialog(
        title: const Text('Delete Hero'),
        content:
            Text('Are you sure you want to delete this hero?\nName: ${name}'),
        actions: <Widget>[
          TextButton(
            onPressed: () => Navigator.pop(context, 'Cancel'),
            child: const Text('Cancel'),
          ),
          TextButton(
            onPressed: () => {
              setState(() {
                heroService.deleteHero(id: id);
                heroesLoaded.removeWhere((hero) => hero.heroId == id);
                Navigator.pop(context, 'OK');
              }),
            },
            child: const Text('Yes'),
          ),
        ],
      ),
    );
  }

  showEditDialog(BuildContext context, String id, String name) {
    String updatedName = "";
    return showDialog<String>(
      context: context,
      builder: (BuildContext context) => AlertDialog(
        title: const Text('Edit Hero'),
        content: TextField(
          onChanged: (value) {
            setState(() {
              updatedName = value;
            });
          },
          decoration: InputDecoration(hintText: "New hero name"),
        ),
        actions: <Widget>[
          TextButton(
            onPressed: () => Navigator.pop(context, 'Cancel'),
            child: const Text('Cancel'),
          ),
          TextButton(
            onPressed: () => {
              setState(() {
                heroService.updateHero(id: id, name: updatedName);
                int index =
                    heroesLoaded.indexWhere((hero) => hero.heroId == id);
                if (index != -1) {
                  heroesLoaded[index].name = updatedName;
                }
                Navigator.pop(context, 'Confirm');
              }),
            },
            child: const Text('Confirm'),
          ),
        ],
      ),
    );
  }

  showCreateDialog() {
    String updatedName = "";
    return showDialog<String>(
      context: context,
      builder: (BuildContext context) => AlertDialog(
        title: const Text('Create Hero'),
        content: TextField(
          onChanged: (value) {
            setState(() {
              updatedName = value;
            });
          },
          decoration: InputDecoration(hintText: "Hero name"),
        ),
        actions: <Widget>[
          TextButton(
            onPressed: () => Navigator.pop(context, 'Cancel'),
            child: const Text('Cancel'),
          ),
          TextButton(
            onPressed: () => createHero(updatedName),
            child: const Text('Confirm'),
          ),
        ],
      ),
    );
  }

  createHero(String updatedName) async {
    String? res = await heroService.createHero(name: updatedName);
    String id = "";
    if (res != null) {
      //Response is of type Inserted a hero with ID: ...
      id = res.split(":")[1];
    }
    setState(() {
      heroesLoaded.add(new MyHero(heroId: id, name: updatedName));
      Navigator.pop(context, 'Confirm');
    });
  }

  Widget heroCard(String id, String name) {
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
            Text(name),
            SizedBox(height: 5),
            Column(
              mainAxisAlignment: MainAxisAlignment.center,
              children: <Widget>[
                ElevatedButton.icon(
                  icon: Icon(Icons.edit, color: Colors.white),
                  label: Text("Edit"),
                  onPressed: () => showEditDialog(context, id, name),
                ),
                SizedBox(height: 5),
                ElevatedButton.icon(
                  icon: Icon(Icons.delete, color: Colors.white),
                  label: Text("Clear"),
                  onPressed: () => showDeleteDialog(context, id, name),
                  style: ElevatedButton.styleFrom(primary: Colors.red),
                ),
              ],
            )
          ],
        ));
  }
}
