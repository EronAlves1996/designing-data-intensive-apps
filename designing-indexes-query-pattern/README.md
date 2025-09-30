# Database Index Design Kata

This is a practical kata focused on designing database indexes for a specific query pattern, based on the "Other Indexing Structures" chapter from _Designing Data-Intensive Applications_. The goal is to think about multi-column indexes, covering indexes, and the performance trade-offs involved, without writing any code.

## The Problem

You're given an `orders` table from an e-commerce platform and four critical queries. The task is to design the most effective indexes to make these queries fast, considering the real-world constraints of a database system.

## What I Learned

- **Clustering is a DB-specific choice.** The primary key is always a unique index, but whether it's clustered (like in MySQL/InnoDB) or non-clustered (like in PostgreSQL) depends on the database. This decision impacts how data is physically stored and retrieved.
- **Column order in a composite index is everything.** An index on `(A, B)` is not the same as `(B, A)`. The leading column is the primary gatekeeper for the query.
- **`ORDER BY` can be optimized with an index.** If you need sorted results, having the sort column in the index itself (especially in the correct order, like `DESC`) can completely avoid an expensive in-memory sort operation.
- **A "covering index" is a powerful tool.** When an index contains all the columns needed for a query, the database never has to read the main table. This is a huge performance win.
- **Index design is about trade-offs.** You can't have a perfect index for every query. Sometimes you need to make a conscious business decision: which query should be faster, and which one can be a little slower?

## Rationale Behind the Solution

Looking at my proposed indexes, the main logic was:

1.  **For Query 1 (`customer_id, order_date DESC`):** You must lead with the filter (`customer_id`). Adding the `order_date DESC` directly in the index definition makes the "get latest 50" operation trivial for the database.
2.  **For Query 2 (`status, shipping_country, order_date, total_amount`):** The goal was to create a covering index. The columns are ordered from the most general filter (`status`) to the most specific ones, and it includes the aggregated column (`total_amount`).
3.  **For Query 3 (`status, total_amount DESC`):** This index is designed to both filter and pre-sort the data. It's essentially the perfect index for that specific admin panel query, making the `ORDER BY` a no-op.

The hardest part was the trade-off in the last question. Choosing one index for two different queries forces you to think about their relative importance to the business.

## Questions for Further Exploration

- How would the index choices change if the table had billions of rows instead of just millions?
- What monitoring tools would you use in production to confirm that your indexes are actually being used and are effective?
- How does an `INCLUDE` clause (to add non-key columns to an index) change the covering index strategy compared to just adding them to the key columns?
