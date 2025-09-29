### Kata: Design a Fault-Tolerant, Sharded Key-Value Store

**Objective:** Apply the principles of a Shared Nothing Architecture to design the logic for a simple, distributed key-value store. You will not write code, but you will design the data flow, failure handling, and key algorithms on a whiteboard or in text.

**Timebox:** 30 minutes

**Scenario:**
You are building "KataKV," a simple key-value store that must scale to handle millions of key-value pairs. You have decided on a Shared Nothing Architecture with 3 nodes (Node A, Node B, Node C).

**Part 1: Data Partitioning & Placement (10 minutes)**

- **Task:** You need to distribute your data across the 3 nodes. Describe your sharding strategy.
  - How do you decide which node is responsible for a given key (e.g., "user:123", "post:456")?
  - **Consider:** A simple modulo-based hashing (`hash(key) % 3`) has a major drawback. What happens when you add a fourth node (Node D)? Describe the problem and propose a better strategy (e.g., consistent hashing) at a high level. How does your chosen strategy minimize the amount of data that needs to be moved when nodes are added or removed?

**Part 2: Handling a `PUT` Request (10 minutes)**

- **Task:** A client wants to store a value `{ "name": "Alice" }` for the key `user:101`. Describe the step-by-step journey of this request through your system.
  - How does the client discover the correct node to talk to?
  - What happens on the target node when it receives the `PUT` request?
  - How do you ensure the data is durable (survives a node restart)?
  - **Consider:** What if the target node is down when the request arrives? Design a simple failover mechanism. (Hint: Think about using a lightweight consensus protocol or leader-election for coordination, but don't overcomplicate it).

**Part 3: Handling a `GET` Request & Consistency (10 minutes)**

- **Task:** Another client now wants to read the value for `user:101`.
  - Describe the step-by-step process for a successful read.
  - **Consider:** What if, due to a network partition, a replica on another node is slightly out of date? What is the consistency model of your system? (e.g., Strong Consistency, Eventual Consistency?). Justify your choice for a high-throughput key-value store.
  - How would a client read its own writes? Propose a simple mechanism to achieve read-your-writes consistency without making the entire system strongly consistent.

---

**Success Criteria:**
You have successfully completed this kata when you have:

1.  Defined a sharding strategy that handles node addition/removal efficiently.
2.  Mapped out the data flow for both write and read operations.
3.  Articulated a clear plan for handling a single node failure during a write.
4.  Defined the consistency model for your system and proposed a solution for a common consistency concern (read-your-writes).

This exercise will solidify your understanding of the trade-offs and design patterns inherent in a Shared Nothing Architecture.
