### Partitioning Kata: LogSentry Design

This is a code kata to practice data partitioning concepts from "Designing Data-Intensive Applications" Chapter 6. The goal is to design a partitioning scheme for a high-throughput log system, focusing on the trade-offs between different strategies.

**The Problem**
Design a partitioning strategy for "LogSentry," a system that ingests massive amounts of server logs. The main challenge is handling a write-heavy workload while still supporting two key read patterns: fetching logs by server and finding high-severity events across the system.

**My Solution & Rationale**

I chose to partition by `event_id` using a hash function. The main reason is to distribute the write load evenly across all nodes. Since `event_id` is a random UUID, it naturally prevents hot spots caused by some servers generating more logs than others. The downside is that querying for a specific `server_id` becomes slower, requiring a scatter/gather operation, but this is an acceptable trade-off for write scalability.

For the secondary index on `severity`, I went with partitioning by document (local indexes). This keeps the write path simple and fast, which is critical for our ingest-heavy system. The con is that querying for high-severity events requires asking all nodes, but since reads are less frequent, this is a good compromise.

The chosen design automatically handles the potential hot spot of a single server generating excessive logs, as the hash of `event_id` spreads those logs evenly.

**Key Lessons Learned**

- The choice of partition key is the most important decision; it directly defines the write distribution and query capabilities.
- There's no perfect solution, only trade-offs. Optimizing for writes often means making reads more complex, and vice-versa.
- Partitioning secondary indexes forces a fundamental choice: slow, simple writes with fast complex reads (global/term), or fast simple writes with slow complex reads (local/document).
- A random hash on a unique key is the best defense against skewed workloads and write hot spots.

**Questions for Further Exploration**

- How would the design change if the main read query was by `server_id` 90% of the time? Would we then choose `server_id` as the partition key and find another way to deal with write skew?
- What if we needed to support complex analytical queries? Would a totally different architecture (like an MPP data warehouse) be better?
- How would we actually implement the scatter/gather for the secondary index query in code? What are the failure modes?
