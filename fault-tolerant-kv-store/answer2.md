## Part 1

1. When the replica-set is configured to sync writes, the leader writes it on WAL and then talk to it's followers, waiting for the confirmation on each of them to return a success to the client. It can poses a significant overhead, and, if any one of the followers fail (like a two phase transaction), the request fail entirely. For async writes, the leader acknowledge the write only on it's own WAL, writes the data to your own database and then returns immediately to the client the success (if succeeded). The write in the replicas are async, then in some point in the future (in the majority of cases, it's a fraction of one second), the change is replicated to it's replicas.
2. Using the Write Ahead Log, the leader can keep track which writes is not replicated to the followers, and then transmit them. The WAL have a sequential number, called log sequential number, that can be compared against leader LSN to know which operations and which order have to be replicated between them.

## Part 2

1. Generally using heartbeats between them. Generally, is a simple tcp message with random bytes, since tcp have ways to ensure that a receiver received the message to a sender. If heartbeats starts to fail, it's a strong signal that leader is out of reach.
2. We can use a simple quorum protocol between them to see who can be a leader. I think, here, the leader should be the database with most fresh data between the two. This means, the database who WAL have the higher LSN. The two followers compare each other which have the higher LSN and then, the leader is elected.
3. This is a good question. I can only think that the data will be temporarily lost, until the old leader becomes available, but, after the LSN became advanced, the old leader will fall behind and maybe the conflict must be handled or entirely discarted.

---

## AI NOTES

    The handling of uncommitted writes is the most complex part of failover. Your answer that data is "temporarily lost" is a common initial thought, but it's incorrect for a system that promises durability. The correct logic is:

        The new leader (the one with the highest LSN, e.g., F1 with LSN 105) knows it is the new leader.

        As part of its recovery, it must determine which entries in its log are "committed" (and safe to apply) and which are "uncommitted."

        A common rule (from Raft) is: an entry is committed once it has been replicated to a quorum (majority) of the replica set. In our 3-node cluster, a quorum is 2 nodes.

        Therefore, any write with an LSN that existed on the old leader but did not make it to a quorum (i.e., only the old leader had it) is uncommitted and must be discarded. The new leader rolls back these entries.

        Client writes that were only on the old leader are indeed lost, and the client will receive an error on retry. This is a deliberate trade-off for consistency, ensuring all replicas agree on the same history.

---

## Part 3

1. When the client writes to the key, the database can return the LSN, and the client keeps track by forwarding it to the front-end, and the front-end keeps track of each, for each device/user. when the front-end requests a read, it sends the LSN or a timestamp in a header. The backend reads it and sends it to the routing tier, and the routing tier can compare the timestamp or logical timestamp against the results it reads from database, to decide if accepts the result or read from a more accurate replica.
2. the fallback logic is to read a replica from the more accurated value, with the timestamp in hands. If leader is down, then it can reads from another replica that is up to date.

---

## AI NOTES

The routing tier does not need to "compare the timestamp against the results it reads." This would be far too slow. The simpler, more efficient design is:

    Each replica (leader or follower) knows its own current LSN (the highest LSN it has applied).

    The routing tier simply knows the current LSN for each replica (this is lightweight metadata, updated via heartbeats).

    When a read request comes in with Client-LSN: 101, the routing tier can instantly route it to any replica whose current_LSN >= 101.

---
