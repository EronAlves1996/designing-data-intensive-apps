## Part 1

1. Alice sees a problem caused by read-after-write. When she comment a post, she's requesting a write that is routed to the leader node. The leader receives the write and proceed to asynchronously replicate the write to followers node. Followers can lag behind a little bit (replication lag), normally a fraction of second, but can be more. It's a problem of eventual consistency class.
2. This is another type of anomally, where the application, for the user, appears to go backwards in time, but occurs only because the first request is routed to a more fresh database, and the second request then is routed to a stale database, and she sees stale data.

## Part 2

3. The first strategy is to read your own write. That means to route the reads for a resource that was recently updated for the leader, because it's supposed to have the more fresh data (the writes go for it), and define some until the read for the resource go for the follower. Another strategy is to route contents that are editable by the user always to the leader and other content can be routed to the followers.
4. A simple strategy to employ is to classify routing by the resources being read. That way, if the client is reading the posts, all posts reading go to follower A and all comments reading go to follower B.

## Part 3

5. If a system has a considerable number of writes, the leader database can be overheaded by the writes plus reads it routes to them, making the complexities replication and distribution of the database worse than the benefits of making it, and your followers will be freed and not used at all. This type of strategy only works if the application is a typical read more than write.
