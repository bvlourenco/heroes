import 'dart:convert';

import 'package:frontend/models/hero.dart';
import 'package:frontend/components/serviceException.dart';
import 'package:http/http.dart' as http;
import 'package:http/http.dart';

class HeroService {
  Map<String, String> reqHeaders = {'Content-type': 'application/json'};

  Future<List<MyHero>> getHeroes() async {
    final res = await http.get(Uri.parse("http://localhost:8080/hero"));
    if (res.statusCode == 200) {
      final jsonRes = json.decode(res.body) as List;
      List<MyHero> heroes = jsonRes.map((e) => MyHero.fromJson(e)).toList();
      return heroes;
    } else {
      throw ServiceException(
          "Error getting heroes:" + res.statusCode.toString() + res.body);
    }
  }

  Future<MyHero> getHero({required String id}) async {
    final res = await http.get(Uri.parse("http://localhost:8080/hero/" + id));
    if (res.statusCode == 200) {
      return MyHero.fromJson(jsonDecode(res.body));
    } else {
      throw ServiceException(
          "Error getting hero:" + res.statusCode.toString() + res.body);
    }
  }

  Future<String?> createHero({required String name}) async {
    var body = jsonEncode({'name': name});

    final res = await http.post(Uri.parse("http://localhost:8080/hero"),
        body: body, headers: reqHeaders);
    return getResponseACK(res);
  }

  Future<String?> deleteHero({required String id}) async {
    final res =
        await http.delete(Uri.parse("http://localhost:8080/hero/" + id));
    return getResponseACK(res);
  }

  Future<String?> updateHero({required String id, required String name}) async {
    var body = jsonEncode({'id': id, 'name': name});

    final res = await http.put(Uri.parse("http://localhost:8080/hero"),
        body: body, headers: reqHeaders);
    return getResponseACK(res);
  }

  String getResponseACK(Response res) {
    if (res.statusCode == 200) {
      return res.body;
    } else {
      throw ServiceException(
          "Error creating hero:" + res.statusCode.toString() + res.body);
    }
  }
}
