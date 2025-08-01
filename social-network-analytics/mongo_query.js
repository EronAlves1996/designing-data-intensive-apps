const db = connect('mongodb://root:root@localhost/test?authSource=admin');

const queryResult = db.users.find({ "posts.content": { $regex: "morning" } }, { "posts": 1 })
console.log(queryResult)
