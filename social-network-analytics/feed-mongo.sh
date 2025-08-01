docker cp insert_data.js social-network-analytics-mongo-1:/insert_data.js 
docker exec -i social-network-analytics-mongo-1 mongosh --file insert_data.js

