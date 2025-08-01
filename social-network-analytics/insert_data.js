const db = connect('mongodb://root:root@localhost/test?authSource=admin');
// contents of users_posts.json
db.users.insertMany(
  [
    { "id": 1, "name": "Alice", "posts": [{ "id": 101, "content": "Hello!" }] },
    {
      "id": 2,
      "name": "Bob",
      "posts": [{ "id": 102, "content": "Good morning!" }]
    }
  ]
)
