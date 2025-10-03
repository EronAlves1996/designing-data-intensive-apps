### Learning Kata: Designing a Partitioning Strategy for "LogSentry"

**Time Limit:** 30 minutes

**Scenario:**
You are the lead engineer for "LogSentry," a new service that ingests and queries security event logs from thousands of servers. Each log entry has the following structure:

- `event_id` (UUID, Primary Key)
- `server_id` (String, e.g., "web-03", "db-primary-01")
- `timestamp` (ISO 8601 DateTime)
- `event_type` (String, e.g., "login_failure", "file_access", "network_alert")
- `severity` (Integer, 1-10)
- `raw_message` (Text)

The anticipated workload is:

- **High write throughput:** A constant stream of new log events.
- **Two main read patterns:**
  1.  Fetch all logs for a specific `server_id` over a recent time window (e.g., "show me all logs for `web-03` from the last hour").
  2.  Find all high-severity events (`severity >= 8`) across the entire system from the last 15 minutes.

Your initial single-node database is struggling. You have decided to scale out by partitioning your data across a cluster of 10 database nodes.

**Your Task:**

Design a partitioning scheme for LogSentry. Answer the following questions. Think about the trade-offs between different approaches.

1.  **Primary Key Partitioning:** How will you partition the main log data? You have two main candidates for your partition key: the `event_id` (UUID) or the `server_id`.
    - Which one would you choose and why?
    - Would you use **Key Range** or **Hash of Key** partitioning for your chosen key? Justify your decision based on the read patterns.

2.  **Secondary Index Challenge:** The query for high-severity events is a problem because it relies on a secondary index (`severity`).
    - Which secondary index partitioning method would be more suitable for this use case: **Partitioning by Document** or **Partitioning by Term**?
    - Explain the pros and cons of your choice in the context of the LogSentry workload. How would a client query for `severity >= 8` under your chosen scheme?

3.  **Hot Spot Mitigation:** Imagine one of your servers, `load-balancer-01`, generates 100x more logs than any other server. How could your partitioning scheme from Question 1 lead to a hot spot, and what is one technique you could use to relieve this pressure?

**Goal:**
The goal of this kata is not to write code, but to think through the architectural trade-offs described in the chapter. There is no single "correct" answer, but there are well-justified ones. Use the concepts from the chapter to defend your choices.
