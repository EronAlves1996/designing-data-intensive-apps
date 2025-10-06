# ACID Transactions Kata

This is a simple demonstration of ACID transaction properties from "Designing Data Intensive Applications" - Chapter 7. The focus is understanding what happens when Atomicity and Isolation guarantees are violated in database transactions.

## What I Built

I implemented a basic banking system simulation to demonstrate two critical transaction problems:

1. **Atomicity Violation**: A funds transfer that deducts money but fails before crediting the other account
2. **Isolation Violation**: Reading inconsistent state during a concurrent transfer operation

## Key Observations

The atomicity problem shows why we need "all or nothing" guarantees - money disappeared from the system when the transfer partially completed. The isolation problem revealed how concurrent operations can see temporary inconsistent states that should never be visible.

I used Go's `big.Float` for precise decimal arithmetic since financial calculations can't afford floating-point errors. The concurrency example uses goroutines and wait groups to simulate real database contention scenarios.

## Lessons Learned

Atomicity isn't just about operations completing - it's about the system maintaining consistency even during failures. Isolation levels matter because without proper concurrency control, applications can see "in-flight" data that breaks business logic.

The kata made me realize why databases need complex machinery like transaction logs and locking - these aren't academic concepts but solutions to real problems I could reproduce in just 50 lines of code.

## Questions for Further Exploration

One thing I'm still thinking about: should I have used channels instead of shared memory for the concurrency test? Also, how would this change if I implemented proper rollback mechanisms? The current simulation just shows the problems - next step would be implementing the solutions with actual transaction boundaries.
