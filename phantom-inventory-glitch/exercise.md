### Kata: **"The Phantom Inventory Glitch"**

**Business Context:** You are a backend engineer at "ArtisanAura," a fast-growing online marketplace for handmade goods. The platform uses a distributed microservices architecture. A critical service is the **Inventory Service**, which is responsible for ensuring sellers don't oversell their unique, handcrafted items.

Recently, a flash sale for a popular ceramic artist caused a major issue. The system showed 10 "Forest Whisper" vases in stock. In the final second of the sale, 5 orders were successfully placed, but the system ended up selling **7 vases**, resulting in 2 angry customers whose orders had to be cancelled. The post-mortem points to a race condition in the distributed inventory check.

**Your Mission:** You are tasked with writing the core logic for the `placeOrder` function within the Inventory Service. This function must be resilient to the distributed system problems discussed in Chapter 8.

**Technical Setup & Constraints:**

- The inventory count for each item is stored in a central database (simulated here).
- The `placeOrder` function is called concurrently from multiple application server instances.
- You cannot assume a perfectly synchronized network or system clocks.
- The function must be _safe_ (never oversell) and have _high availability_ (it shouldn't fail just because of temporary network slowness).

**Core Requirements:**

1.  **Unreliable Networks & Timeouts:** Simulate a random network delay or process pause before the inventory deduction. Your function must handle the case where it holds a reservation for an item but the final confirmation is delayed. How do you prevent another request from selling the same unit?
2.  **Unreliable Clocks:** You are not allowed to use the system's time-of-day clock (e.g., `System.currentTimeMillis()` or `Instant.now()`) for making logical decisions about the order sequence or inventory reservation expiry. You must find an alternative.
3.  **Knowledge & Truth:** Implement a mechanism to establish a single "truth" about the inventory level. You need a way to perform a check and a deduction in a way that appears atomic to other concurrent processes, even in the face of partial failures.

**Kata Tasks (30-minute simulation):**

**Part 1: The Naive Implementation (5 mins)**
Write a simple version of `placeOrder(itemId, quantity)` that:

1.  Reads the current inventory.
2.  If `currentInventory >= quantity`, it deducts the quantity and returns `"ORDER_PLACED"`.
3.  Otherwise, it returns `"OUT_OF_STOCK"`.

**Revised Part 2: Introduce Chaos & Observe (5 mins)**

**Simulate Distributed Service Instances:** Instead of concurrent goroutines accessing a single in-memory map, we'll simulate multiple independent service instances. Create a simple `InventoryDB` interface that represents your database. Then create two different implementations:

1.  **`InstantDB`**: A perfect, instantaneous database (simulates a local in-memory store, which won't show the problem).
2.  **`NetworkLagDB`**: A database wrapper that simulates real-world network issues. It should:
    - Add a random delay (e.g., 50-200ms) before executing any query.
    - **Crucially:** For the "read inventory" operation, it should return the value it has at the moment it processes the request, not necessarily the latest value. This simulates the fact that in a real distributed system with replication, reads might be stale.

**The Test Scenario:**

- Initial stock: 10 vases
- Run 5 concurrent order requests, each for 1 vase, using the `NetworkLagDB`.
- Use your naive implementation from Part 1.

**What to look for:**
The `NetworkLagDB` should now properly demonstrate the problem. Because each request reads the inventory with random delays, multiple requests might all see "10" in stock at roughly the same time, and all proceed to deduct, causing overselling. This accurately simulates multiple application servers talking to a database with network latency and potential replication lag.

**Part 3: Build Resilience (20 mins)**
Refactor your `placeOrder` function to be robust. You must now:

- **Prevent Overselling:** The system must never sell more items than are in stock. This is your highest priority.
- **Handle Reservation Expiry:** If an order reservation is held for too long (e.g., the user's payment fails), that inventory must be released back to the pool. Implement this without relying on system clock time for the core logic. (Hint: Think about state machines and sequence numbers).
- **Define the Truth:** Your solution must have a single source of truth for the inventory level at the moment of deduction. How will you achieve this without a distributed lock? (Hint: The book often points to consensus or atomic operations at the database level).

**Success Criteria:**
You will know you are successful when you can run a simulated flash sale (e.g., 15 concurrent orders for 10 items) and the final number of successful orders is always less than or equal to the initial stock, with no two orders receiving the same physical unit of inventory.

This kata forces you to grapple with the core tensions of the chapter: managing state without a global clock, making reliable decisions over an unreliable network, and maintaining a consistent truth in a system with no central authority.
