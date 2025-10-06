### Learning Kata: The Unreliable Bank Transfer

**Objective:** To understand the practical implications of ACID properties, specifically Atomicity and Isolation, by simulating a scenario where they are absent.

**Scenario:**
You are modeling a simple banking system with a `bank_accounts` table. You need to implement a money transfer function. However, the database you are using has two strange "modes" of operation that you must investigate.

**Your Task:**

1.  **Model the Data:** Design a simple in-memory structure (e.g., a dictionary, a list of objects) to represent the `bank_accounts` table. Each account should have at least an `account_id` and a `balance`.

2.  **Implement a Faulty Transfer (Violating Atomicity):**
    - Write a `transfer_funds_non_atomic(from_account, to_account, amount)` function.
    - The function should:
      1.  Deduct the `amount` from the `from_account` balance.
      2.  **Simulate a System Crash:** Immediately after the deduction, artificially throw an exception (e.g., `raise Exception("Database node crashed!")`).
      3.  If step 2 is skipped (i.e., no crash), add the `amount` to the `to_account` balance.
    - **Investigation:** Run this function. What is the state of the two accounts after the "crash"? What problem has occurred? How does this violate the principles of Atomicity?

3.  **Implement a Concurrent Transfer (Violating Isolation):**
    - Write a `get_total_balance()` function that reads the balances of two specified accounts and returns their sum.
    - Now, simulate a scenario where the total balance is read _during_ a transfer.
      - In one thread (or function call), run a `transfer_funds(...)` function that moves money from Account A to Account B. This function should have a small delay (e.g., `time.sleep(0.1)`) between deducting from one account and adding to the other.
      - Simultaneously, in another thread (or a subsequent, immediate function call), run the `get_total_balance()` function for Accounts A and B.
    - **Investigation:** What value does `get_total_balance()` return? Should the total amount of money in the two accounts ever change during a transfer? What anomaly have you observed, and how does this relate to Isolation?

**Success Criteria:**
You will know you have successfully completed the kata when you can:

- Clearly articulate the problem caused by the lack of Atomicity in Part 2.
- Clearly articulate the concurrency anomaly (a temporary inconsistency) caused by the lack of Isolation in Part 3.
- Explain how a real database transaction with ACID properties would prevent both of these issues.

**Timebox:** This kata is designed to be completed in 20-30 minutes. Focus on the core concepts rather than building a perfect simulation. The goal is to internalize the "why" behind transactions.
