### Kata: Designing the Internals of a Sharded, Replicated Store

**Objective:** Deepen your understanding of the internal mechanisms of a Shared Nothing system. You will design the core algorithms for replication, failure handling, and consistency. Focus on the _how_, not just the _what_.

**Timebox:** 30 minutes

**Scenario:**
You are building the core of "KataKV v2." You have a fixed number of partitions (e.g., 10). Each partition is replicated across 3 nodes using a leader-follower model. A separate routing tier exists that knows the current leader for each partition.

---

### Part 1: The Write Path & Durability Guarantee (10 minutes)

A client sends a `PUT` request for key `K1` with value `V1` to the leader of its respective partition.

- **Task A:** Describe the precise, step-by-step process _on the leader node_ from receiving the request to sending a success response back to the client. Your description must ensure that the data is not lost if the leader node crashes and restarts one second after acknowledging the write.
- **Task B:** Now, describe how this write is asynchronously replicated to the two follower nodes. What data structure does the leader use to track which writes have been acknowledged by which followers? (Hint: Think about a sliding window of operations).

### Part 2: Leader Failure & The Failover Process (10 minutes)

The leader for Partition 5 suddenly crashes due to a hardware fault. The two followers (F1 and F2) for Partition 5 are still running.

- **Task A:** How do followers F1 and F2 detect that the leader has crashed? Be specific about the mechanism.
- **Task B:** Describe the steps F1 and F2 must take to elect a new leader. You cannot use Zookeeper for the election logic itself (you can use it only to publish the final result). Design a simple election protocol that ensures only one new leader is elected. (Hint: What information would a follower need to be eligible to become a leader?).
- **Task C:** What happens to client writes that were acknowledged by the old leader but not yet replicated to all followers? How does the new leader handle this during its recovery process?

### Part 3: Implementing "Read-Your-Writes" Consistency (10 minutes)

Your system uses leader-based replication but allows reads from followers to increase read throughput. This creates a potential staleness issue.

- **Task A:** A client writes `V2` to key `K1`. It immediately issues a `GET` request for `K1`. How can this read request be routed to ensure it sees `V2` and not a stale value? Design a mechanism that does **not** require the routing tier to store any persistent client state.
- **Task B:** How does your solution from Task A handle a scenario where the client's preferred replica (e.g., the leader) is temporarily unavailable? What is the fallback logic?

---

**Success Criteria:**
You have successfully completed this kata when you have:

1.  Designed a write-ahead log (WAL) mechanism for durability and described the replication state machine.
2.  Designed a heartbeat-based failure detector and a simple leader election protocol based on data freshness.
3.  Defined a mechanism for handling uncommitted writes after a leader failover.
4.  Designed a client-driven, stateful session mechanism for "read-your-writes" consistency without pushing state to the routing tier.
