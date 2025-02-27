import "dart:convert";

class Host {
  const Host({
    required this.name,
    required this.location,
    required this.rating,
    required this.description,
    required this.capacity,
    required this.price,
  });

  Host.fromMap(Map<String, dynamic> map)
    : name = map["name"],
      location = map["location"],
      rating = map["rating"],
      description = map["description"],
      capacity = map["capacity"],
      price = map["price"];

  factory Host.fromJson(String json) => Host.fromMap(jsonDecode(json));

  final String name;
  final String location;
  final double rating;
  final String? description;
  final int capacity;
  final double price;
}
