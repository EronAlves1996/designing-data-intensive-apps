MATCH (s:USER {name:"Alice"})<-[:FOLLOWS]-(a)-[:FOLLOWS]->(b) return b.name
