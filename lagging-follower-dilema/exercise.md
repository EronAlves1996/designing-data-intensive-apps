### Kata: The Lagging Follower Dilemma

**Objective:** To reason about the real-world implications of replication lag in a single-leader database system and design strategies to mitigate its effects.

**Scenario: "QuickPost" - A Social Media Application**

Imagine you are designing the backend for "QuickPost," a simple social media app. The primary database uses a single-leader replication setup with one leader (in the US East region) and two asynchronous followers (one in US East, one in Europe). Users can perform two main actions:

1.  **Post a comment.**
2.  **View a thread of comments.**

A user, Alice, who is traveling, posts a comment from her phone. Due to network conditions, her write request is routed to and handled by the leader in US East.

**Your Task (30 minutes):**

Spend 30 minutes analyzing the following user stories and answering the questions. Focus on the _why_ and the _design trade-offs_, not on writing code.

**Part 1: Identifying the Problem (10 minutes)**

1.  Immediately after posting, Alice refreshes her phone to see her comment in the thread. However, her read request is routed to the follower in Europe, which hasn't yet received the replication update from the US East leader. What does Alice see? What is this specific issue called (from the chapter)?

2.  Later, Alice reads comments again. First, her request goes to the European follower and she sees her comment. She refreshes the page a second later, but this time her request is routed to the US East follower. Due to a temporary network hiccup, the European follower is actually slightly _ahead_ of the US East follower. What anomaly might Alice observe? What is this issue called?

**Part 2: Designing Solutions (15 minutes)**

3.  **Reading Your Own Writes:** How could you ensure that Alice _always_ sees her own comment immediately after posting it, even if her subsequent read requests go to a lagging follower? Describe at least two technical strategies mentioned in the chapter.

4.  **Monotonic Reads:** The scenario in question 2 violates the principle of monotonic reads. What is a simple strategy you could implement on the backend to prevent Alice from experiencing this "going back in time" effect? (Hint: Think about how to route her requests).

**Part 3: Considering Trade-offs (5 minutes)**

5.  Every solution has a cost. What is the potential downside of implementing the "Reading Your Own Writes" guarantee? For example, how might it affect performance or system complexity?

**Success Criteria:**
You have successfully completed this kata if you can:

- Correctly identify the replication lag anomalies (Stale Read and violation of Monotonic Reads).
- Propose plausible implementation strategies for mitigating these issues (e.g., read-after-write consistency using a tracking token or routing based on user ID).
- Articulate the trade-off involved, such as reduced load-balancing flexibility or increased complexity.

This exercise will solidify your understanding of the practical challenges of replication lag and the common patterns used to build user-friendly applications on top of eventually consistent systems.
