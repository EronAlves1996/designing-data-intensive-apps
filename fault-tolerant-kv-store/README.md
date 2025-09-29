# Shared Nothing Architecture Katas

This is a conceptual study on designing data-intensive applications, focused on the internals of a Shared Nothing Architecture. The goal was to design the logic for a distributed key-value store without writing skeleton code, focusing on data flow, failure handling, and consistency.

## What I Learned

The first kata helped me structure the high-level architecture: sharding with fixed partitions, using a routing tier, and leader-follower replication. I correctly identified the trade-offs between hashing and range partitioning and the need for a coordination service like Zookeeper for metadata.

The second kata was harder. It forced me to open the black boxes. I learned that the sequence of operations is critical for durability: a leader must write to its own Write-Ahead Log (WAL) _before_ acknowledging a write to the client, even in async replication. I also solidified that the replica with the highest Log Sequence Number (LSN) should win a leader election.

My biggest gap was in handling uncommitted writes after a leader failover. I initially thought the data would be temporarily lost, but the correct logic is that any write not replicated to a quorum (majority) of nodes must be deliberately discarded to maintain consistency across the cluster.

For "read-your-writes" consistency, the solution is to have the client track a version token (like an LSN) and send it with reads. The routing tier can then route the request to any replica that is at least as up-to-date as that token.

## Rationale for the Solutions

The thinking here is to prioritize clarity in the data flow and explicit handling of failure scenarios. The solutions avoid "magic" by explaining the mechanisms—like heartbeats for failure detection and LSNs for ordering—instead of just naming the components. The focus is on the algorithms that make the architecture work, not just the architecture itself.

## Questions for Further Exploration

I'm still thinking about the practical details of the failover process. How exactly does a follower know which WAL entries are committed? Is it just based on the LSN of the last known leader? Also, in the "read-your-writes" pattern, what's the best way to propagate the LSN back to the client? Is it a header, or part of the response body? And how do we prevent this from becoming a bottleneck for very fast, sequential writes?
