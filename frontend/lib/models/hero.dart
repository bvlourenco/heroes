class MyHero {
  final String heroId;
  final String name;

  const MyHero({required this.heroId, required this.name});

  factory MyHero.fromJson(Map<String, dynamic> json) {
    return MyHero(heroId: json['id'], name: json['name']);
  }
}
