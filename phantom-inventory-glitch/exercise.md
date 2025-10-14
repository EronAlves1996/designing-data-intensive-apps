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

**Part 2: Introduce Chaos & Observe (5 mins)**
Simulate concurrency by calling your function multiple times in parallel for the same item, with a total quantity that exceeds the stock. You will almost certainly observe the overselling problem. This demonstrates the partial failure of correctness.

**Part 3: Build Resilience (20 mins)**
Refactor your `placeOrder` function to be robust. You must now:

- **Prevent Overselling:** The system must never sell more items than are in stock. This is your highest priority.
- **Handle Reservation Expiry:** If an order reservation is held for too long (e.g., the user's payment fails), that inventory must be released back to the pool. Implement this without relying on system clock time for the core logic. (Hint: Think about state machines and sequence numbers).
- **Define the Truth:** Your solution must have a single source of truth for the inventory level at the moment of deduction. How will you achieve this without a distributed lock? (Hint: The book often points to consensus or atomic operations at the database level).

**Success Criteria:**
You will know you are successful when you can run a simulated flash sale (e.g., 15 concurrent orders for 10 items) and the final number of successful orders is always less than or equal to the initial stock, with no two orders receiving the same physical unit of inventory.

This kata forces you to grapple with the core tensions of the chapter: managing state without a global clock, making reliable decisions over an unreliable network, and maintaining a consistent truth in a system with no central authority.
