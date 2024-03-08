var db = connect("mongodb://localhost:27017/?authMechanism=DEFAULT");

db = db.getSiblingDB('fupitz'); // we can not use "use" statement here to switch db

db.createUser(
    {
        user: "fupitz",
        pwd: "12345678",
        roles: [ { role: "readWrite", db: "fupitz"} ],
        passwordDigestor: "server",
    }
)