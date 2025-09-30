## 1.

- It's created a B-tree unique index. The clustering of index depends on the database.
- Generally it depends on database. MySQL, as an example, uses clustered index, but postgresql uses non-clustered. For that matter, the fourth query selects all the fields of table by order_id. Because it's critical, I would use a clustered index, seeing that I can take a hit because of write performance.

## 2.

- **Query 1**: (customer_id, order_date desc) => this would not be a covering index
- **Query 2**: (status, shipping_country, order_date, total_amount) => this gonna be a covering index
- **Query 3**: (status, total_amount desc) => this would not be a covering index
