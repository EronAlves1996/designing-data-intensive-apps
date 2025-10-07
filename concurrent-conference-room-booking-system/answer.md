## 1.

For both the teams, the following operations apply:

1. The transaction Alpha starts. By the time, the snapshot on bookings table is as following:

| id  | room_name | time_slot        | booked_by  |
| --- | --------- | ---------------- | ---------- |
| 1   | R1        | 2023-10-25 15:00 | Team Gama  |
| 2   | R2        | 2023-10-25 15:00 | Team Theta |

2. The transaction Beta starts. by the time, they can see the same snapshot as team Beat, actually

3. The Team Alpha actually checks the availability, using the query on process step 1. It returns `2` as a result

4. The team Alpha makes a book, by inserting into the bookings table, using the query from the process step 3.

5. Team Alpha commits

6. Team Beta checks availability, using the query on process step 1, returning the same `2` result

7. Team Beta makes a book, inserting into the books table, using the query from the process step 3

8. Team Beta commits

## 2.

This anomaly is a write skew, basically a concurrent write violation anomally, where the writes are made basing upon a stale state on database. There's likely a phantom, because the write of first transaction should modify the results in another transaction, but repeatable read only make guarantees on read operations, not on writes, and the phantom effect can affect writes. If the system states that a specific invariant cannot occur, and the write skew make it possible, then it's likely a bug.

## 3.

The technical solution would be to specifically locking the rows returned from the query. We gonna modify the query, not using `COUNT` function, but specifically querying the row info, and then, append the `FOR UPDATE` clause to lock the rows. Another alternative is materializing conflicts, using a control table to control all the slots for the booking system, and then, the writes goes to only one place. The control table gonna have rows for all the possibilities, making a matrix between the roms and the time slots, and the writes in the row gonna represent the actual booking.
