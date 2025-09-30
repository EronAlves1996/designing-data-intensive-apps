### Kata: Designing Indexes for a Query Pattern

**Objective:** Analyze a set of common queries for a hypothetical database table and design the most effective single- and multi-column indexes to optimize them. You will not write any code, but you will reason about index design, column order, and index type (clustered, covering, etc.).

**Timebox:** 30 minutes

**Scenario:**
You are the database architect for an e-commerce platform. You have a central `orders` table with the following columns:

- `order_id` (UUID, Primary Key)
- `customer_id` (UUID)
- `order_date` (TIMESTAMP)
- `status` (ENUM: 'pending', 'shipped', 'delivered', 'cancelled')
- `total_amount` (DECIMAL)
- `shipping_country` (VARCHAR)
- `payment_method` (VARCHAR)

After analyzing the application logs, you identify the following four queries as the most frequent and performance-critical.

**Critical Queries:**

1.  **Dashboard Lookup:** Retrieve the 50 most recent orders for a specific customer.

    ```sql
    SELECT * FROM orders
    WHERE customer_id = '...'
    ORDER BY order_date DESC
    LIMIT 50;
    ```

2.  **Analytics Report:** Count the total number and value of orders that were `'shipped'` to a specific `shipping_country` within a date range.

    ```sql
    SELECT COUNT(*), SUM(total_amount) FROM orders
    WHERE status = 'shipped'
    AND shipping_country = '...'
    AND order_date BETWEEN '...' AND '...';
    ```

3.  **Admin Panel Search:** Find all orders in a `'pending'` status, sorted by their `total_amount` in descending order (to see the most valuable pending orders first).

    ```sql
    SELECT order_id, customer_id, total_amount FROM orders
    WHERE status = 'pending'
    ORDER BY total_amount DESC;
    ```

4.  **Order Lookup:** Retrieve a full order by its `order_id`.
    ```sql
    SELECT * FROM orders WHERE order_id = '...';
    ```

---

**Your Tasks:**

1.  **Primary Key Strategy (5 minutes):**
    - The table already uses `order_id` as the primary key. What type of index is typically created automatically for a primary key?
    - Would you make this a clustered or non-clustered index? Justify your choice based on the query patterns.

2.  **Designing Multi-Column Indexes (15 minutes):**
    For queries 1, 2, and 3, propose a single, optimal multi-column index for each. For each index, specify:
    - **The columns** in the index.
    - **The order** of the columns. Explain _why_ you chose this specific order, linking it to the `WHERE` clause and `ORDER BY` clause of the query.
    - **The type** of index if applicable (e.g., could it be a covering index?).

    - **Index for Query 1:**
    - **Index for Query 2:**
    - **Index for Query 3:**

3.  **Trade-offs and Consolidation (10 minutes):**
    - Creating an index for every query can be expensive for write operations. Look at your proposed indexes for Queries 2 and 3. Is there a single multi-column index that could _adequately_ serve both queries, even if it's not perfect for both? What would it be, and what would the trade-offs be?
    - For Query 3, you designed an index that perfectly covers the query. If you only had a simple index on `(status)`, why would the `ORDER BY total_amount DESC` clause be a performance problem? What would the database have to do
