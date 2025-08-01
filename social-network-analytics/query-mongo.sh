docker cp mongo_query.js social-network-analytics-mongo-1:/mongo_query.js 
docker exec -i social-network-analytics-mongo-1 mongosh --file mongo_query.js

