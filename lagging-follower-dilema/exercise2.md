### Kata: Designing Consistency Guarantees for "TaskFlow"

**Objective:** To practice designing specific, precise technical strategies to enforce "Read Your Writes" and "Monotonic Reads" consistency guarantees.

**Scenario: "TaskFlow" - A Project Management Tool**

You are designing the backend for "TaskFlow," a project management application where users collaborate on projects. The database uses **single-leader replication** with one leader and three asynchronous followers to scale read capacity. A user can:

1.  **Change the status of a task** (e.g., from "To Do" to "In Progress").
2.  **View the project's task board**, which shows all tasks and their statuses.

Users are distributed globally, and read requests are load-balanced across all available follower nodes to reduce latency.

---

**Your Task (30 minutes)**

**Part 1: Diagnosing the Anomalies (5 minutes)**

1.  A user in Tokyo, Carlos, changes the status of a task. The write goes to the leader in Virginia, USA. A moment later, he views the task board, but his read request is served by a follower in Singapore that is experiencing replication lag. What consistency guarantee is violated? What does Carlos see?

2.  Later, Carlos refreshes the task board twice in a row. The first refresh is served by the follower in Singapore, which has now received the update. He sees his status change. The second refresh is, by chance, served by a follower in Brazil, which is lagging further behind. What consistency guarantee is violated now? What does Carlos see on the second refresh?

**Part 2: Designing the "Read Your Writes" Guarantee (10 minutes)**

3.  The product team says, "A user must _always_ see their own task status changes immediately." Describe a **concrete, implementable strategy** to achieve this. Your description should be specific enough that a developer could start coding it.
    - _Hint: Think about what information the client (web browser/app) and server need to track and exchange._

**Part 3: Designing the "Monotonic Reads" Guarantee (10 minutes)**

4.  The product team now says, "It's confusing for users if the task board appears to go backwards in time. Once a user sees a change, they should never see it revert on a subsequent read." Describe a **concrete, implementable strategy** to achieve Monotonic Reads for a logged-in user.
    - _Hint: The solution is not based on the type of data being read, but on the identity of the person reading it._

**Part 4: Trade-offs and Limitations (5 minutes)**

5.  What is a potential limitation or downside of the strategy you designed for Monotonic Reads (Question 4)? How might it affect system performance or reliability during a partial failure?

---

**Success Criteria**

You will know you have successfully completed this kata if your answers for Part 2 and 3 are:

- **Specific:** They mention concrete components like "client-side timestamp," "routing layer," or "user session."
- **Actionable:** A developer could understand the steps needed to implement them (e.g., "The client must send X header with each request").
- **Precise:** They correctly tie the solution directly to the user's identity and their sequence of reads, not just to the type of data.

Take your time, and focus on the technical mechanics of the solutions. Good luck!
