## Part 1

1. This is a violation of read-your-writes consistency or read-after-write consistency. Carlos sees like his data is lost.
2. Now, the monotonic reads consistency is violated. Carlos sees like the data is going backwards in time.

## Part 2

3. The backend now will write to the database, and the request gonna go to a leader. The backend can take a logical timestamp and answer this timestamp in a response header and the client can store this logical timestamp in it's state. Every time the front-end makes request that aren't safe (in the http protocol perspective), the backend then sends this logical timestamp. Now, in the next read, safe requests (GET), the frontend will send the logical timestamp. The backend can read the logical timestamp and check if the replica where it gonna read is up to date. If not, route to a more up to date replica.

## Part 3

4. Logged users always have some id. In every session, make some affinity between the id of the user and a replica where he should read. He will always read from that replica. Let's say we have two node followers. We gonna say that the even ids goes for node A, and not-even ids goes for node B.
