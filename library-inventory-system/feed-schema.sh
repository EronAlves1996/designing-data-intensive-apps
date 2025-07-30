docker cp schema.sql library-inventory-system-db-1:/tmp/schema.sql
docker exec -i library-inventory-system-db-1 psql -U postgres -d postgres < schema.sql

