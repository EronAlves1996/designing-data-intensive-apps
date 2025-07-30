docker cp create_book.js library-inventory-system-mongo-1:/create_book.js 
docker exec -i library-inventory-system-mongo-1 mongosh --file create_book.js

