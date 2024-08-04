# DB Locks

This Go program demonstrates the usage of different SQL locking mechanisms to handle concurrency when booking seats for multiple users. It includes three methods:

- Booking seats without any locks
- Booking seats using exclusive locks
- Booking seats using skip locks

The program fetches all users, resets the seat details, and then books seats for each user using the specified locking mechanism.
It also prints the seat details and seating arrangement after each method.

### Locking Mechanisms

1. Without Locks: This method does not use any locks, which can lead to race conditions and inconsistent data.
2. Exclusive Locks: This method uses `FOR UPDATE` to lock the selected row for the duration of the transaction, preventing other transactions from selecting the same seat until the current transaction is complete.
3. Skip Locks: This method uses `FOR UPDATE SKIP LOCKED` to skip rows that are currently locked by other transactions. However, this syntax is not supported in MySQL, so it will result in a syntax error.
