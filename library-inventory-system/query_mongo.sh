docker exec -i library-inventory-system-mongo-1 mongosh -u root -p root --authenticationDatabase admin --eval 'db.book.find({tags: { $all: ["sci-fi"]}})'
docker exec -i library-inventory-system-mongo-1 mongosh -u root -p root --authenticationDatabase admin --eval 'db.book.aggregate([{$match:{title:"Dune"}},{$unwind:"$copies"},{$group:{_id:"$copies.branch",count:{$sum:1}}}])'
docker exec -i library-inventory-system-mongo-1 mongosh -u root -p root --authenticationDatabase admin --eval 'db.book.aggregate([{$unwind:"$authors"}, {$group:{_id:"$authors",count:{$sum:1}}}])'

