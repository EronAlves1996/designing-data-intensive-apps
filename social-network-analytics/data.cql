CREATE (alice:USER {name: "Alice"})
CREATE (bob:USER {name:"Bob"})
CREATE (charlie:USER {name:"Charlie"})

MERGE (charlie)-[:FOLLOWS]->(alice)-[:FOLLOWS]->(bob)-[:FOLLOWS]->(charlie)
