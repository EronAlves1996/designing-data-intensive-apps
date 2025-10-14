# Distributed Inventory Kata

This is a practical implementation of distributed systems concepts from "Designing Data-Intensive Applications" Chapter 8. The kata simulates an inventory management system that must handle concurrent orders without overselling, despite network delays and partial failures.

## The Challenge

Build an inventory system that can handle flash sales with multiple concurrent orders, where the network is unreliable and processes can pause arbitrarily. The system must never sell more items than available in stock.

## Key Implementation

The solution uses a **reservation pattern** to prevent race conditions:

1. **Check availability** - Query current inventory
2. **Try reserve** - Atomically reserve items with a unique ID
3. **Process order** - Simulate payment processing with random delays
4. **Confirm or release** - Either confirm the reservation or release it back to inventory

The `NetworkLagDB` wrapper simulates real-world network conditions with random 50-200ms delays, while `InstantDB` provides the core inventory logic with proper synchronization.

## Lessons Learned

The reservation system effectively prevents overselling by creating temporary holds on inventory. This handles the distributed systems problem where multiple services might see stale inventory data due to network delays.

I used atomic operations for reservation IDs and mutexes for shared state protection. The tricky part was ensuring that reservations are properly cleaned up when confirmations fail - the defer approach was too aggressive, so I switched to explicit release calls only on failure.

## Questions for Further Exploration

One thing I considered but didn't implement was reservation timeouts. How should the system handle reservations that are never confirmed or released? Should there be a background process that cleans up stale reservations after some time?

Also, what happens if the confirmation succeeds but the release call fails due to network issues? Would this create permanently reserved inventory that's never available for sale?s
