## Part 1

We have some approaches we can use for this task.
Modulo based hashing have the drawback when we need to add and remove nodes, because, when this happen, we need to move so much data between nodes in not an evenly manner, and, when this happens, we need the sharding strategy to be sufficiently optimized to make take the less time on I/O possible, by moving the minimal amount of data to achieve balance.
We can used fixed number paritioning to approach the problem. Let's say we define the application will have 50 partitions. We gonna distribute the data around them, and, when nodes are added/removed, we move entire partitions between nodes. This makes the data transfer consistent and the node distribution even.
To decide in which partition we gonna write a key, range key partitioning have the advantages to be an easy criteria and is predictable, but if the application is write heavy, can cluter a single partition with write, making it a hot spot.
Hash partitioning is good, because write is optimized and evenly distributed, but read suffers, and we loose the ordering.
It's down to the requirements. If the application is write heavy, i'll go to hash partitioning, using things like consistent hashing. If the application is read heavy, i'll stick to range-key partitioning.

## Part 2
