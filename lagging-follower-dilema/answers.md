## Part 1

1. Alice sees a problem caused by read-after-write. When she comment a post, she's requesting a write that is routed to the leader node. The leader receives the write and proceed to asynchronously replicate the write to followers node. Followers can lag behind a little bit (replication lag), normally a fraction of second, but can be more. It's a problem of eventual consistency class.
2. This is another type of anomally, where the application, for the user, appears to go backwards in time, but occurs only because the first request is routed to a more fresh database, and the second request then is routed to a stale database, and she sees stale data.
