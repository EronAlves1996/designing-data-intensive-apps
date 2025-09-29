## Part 1

We have some approaches we can use for this task.
Modulo based hashing have the drawback when we need to add and remove nodes, because, when this happen, we need to move so much data between nodes in not an evenly manner, and, when this happens, we need the sharding strategy to be sufficiently optimized to make take the less time on I/O possible, by moving the minimal amount of data to achieve balance.
We can used fixed number paritioning to approach the problem. Let's say we define the application will have 50 partitions. We gonna distribute the data around them, and, when nodes are added/removed, we move entire partitions between nodes. This makes the data transfer consistent and the node distribution even.
To decide in which partition we gonna write a key, range key partitioning have the advantages to be an easy criteria and is predictable, but if the application is write heavy, can cluter a single partition with write, making it a hot spot.
Hash partitioning is good, because write is optimized and evenly distributed, but read suffers, and we loose the ordering.
It's down to the requirements. If the application is write heavy, i'll go to hash partitioning, using things like consistent hashing. If the application is read heavy, i'll stick to range-key partitioning.

## Part 2

1. The first thing is discovering where the client gonna write. I choose to have a routing tier between the client and the store, and the routing tier will be plugged in a service discovery layer (like zookeeper). Zookeeper will be booking each node and partition status, and all it's metadata. The routing tier will query zookeeper for the information and then route the request to the appropriate node, like a loading balancer (not a round robin LL)
2. The routing tier already will route the request to the correct node. The node will proceed the write.
3. We can use replicas to make the data durable. For each partition, we can make another two replica, and, in each replica-set, one of them is a leader an the other are followers. We distribute the replica-set around datacenters. If a node restart, we have another two replicas that can go through the failover process and indicate another leader.
4. Well, the routing tier will receive the failure and talk to zookeeper to start a failover process to elect another leader. We are following leader-follower replica schema, so, the election is deadly simple in this schema.

## Part 3

1. We evolved through here to a fixed partition with leader-follower replica schema. So, the request goes through routing tier, but for read, we have 3 replicas that can be approach the task. I think that here, we can take the guarantees for leader-follower schema and follow them. It's supposed for the routing tier to have the information of the user and the timestamp of last read for this user. It have to route to the right replica using this information paying attention to the monotonic reads and read-your-write guarantees.
2. For high throughput, it's entirely on how is the workload of application. High number of reads? Strong consistency can be better. High number of writes? Eventual consistency can be better. But generally, in eventual consistency systems, we use a timestamp of last read (like LSN WAL number) to compare our read against the replica, and choose another replica if they are falling behind to avoid the read-your-writes loose of guarantee.
3. Check answer in item 2
