package main

import (
	"sync"
)

// Transaction represents a transfer between two accounts.
type Transaction struct {
	SenderID   int
	ReceiverID int
	Amount     int
}

// processTransactions processes transactions concurrently and safely.
//
// Rules handled:
//  1. If a transaction causes insufficient balance, it is paused.
//  2. The next transaction is processed immediately.
//  3. If the next transaction increases the paused sender’s balance,
//     retry the paused transaction once.
//  4. Ensures race-free balance updates using mutex locking.
//
// Returns two channels:
//   - transactions: to send transactions for processing
//   - done: closed when all transactions are processed
func processTransactions(balances map[int]int) (transactions chan Transaction, done chan struct{}) {
	transactions = make(chan Transaction)
	done = make(chan struct{})

	var mu sync.Mutex

	go func() {
		defer close(done)

		var pending *Transaction // holds one paused transaction at a time

		for tx := range transactions {
			mu.Lock()

			// If a pending transaction exists, try it first
			if pending != nil && balances[pending.SenderID] >= pending.Amount {
				balances[pending.SenderID] -= pending.Amount
				balances[pending.ReceiverID] += pending.Amount
				pending = nil
			}

			// Process current transaction
			if balances[tx.SenderID] >= tx.Amount {
				// Enough balance — execute
				balances[tx.SenderID] -= tx.Amount
				balances[tx.ReceiverID] += tx.Amount
			} else {
				// Insufficient funds — hold transaction
				pending = &tx
			}

			mu.Unlock()
		}

		// After all transactions sent — final retry if pending
		if pending != nil {
			mu.Lock()
			if balances[pending.SenderID] >= pending.Amount {
				balances[pending.SenderID] -= pending.Amount
				balances[pending.ReceiverID] += pending.Amount
			}
			mu.Unlock()
		}
	}()

	return transactions, done
}

/*
// Use Case: Concurrent transaction processing with pause/retry on insufficient funds.| Step  | Behavior                                                                                                         |
| ----- | ---------------------------------------------------------------------------------------------------------------- |
| **1** | Creates `transactions` (input channel) and `done` (completion signal).                                           |
| **2** | Spawns a goroutine that continuously listens for new transactions.                                               |
| **3** | Protects shared state (`balances`) with a `sync.Mutex` for concurrency safety.                                   |
| **4** | If a transaction fails due to low balance, it is stored as `pending`.                                            |
| **5** | The *next* transaction is processed, and then the pending one is retried once if possible.                       |
| **6** | When all transactions are sent and the input channel closes, one last retry is done for the pending transaction. |
| **7** | Finally, `done` is closed to notify that all work is complete.                                                   |
*/
/*
   Implement a function that processes banking transactions concurrently while maintaining data integrity. The system must handle transfers between accounts, ensuring that race conditions are avoided and that transactions are processed correctly even when they arrive out of order.



Your task is to implement a processTransactions function in Go that uses goroutines and channels to handle concurrent banking transactions safely. The function should be able to handle cases where a transaction might temporarily fail due to insufficient funds, but could succeed once another transaction increases the sender's balance. Each transaction will consist of a sender ID, a receiver ID, and an amount to be transferred between their accounts." since this clearly mentions what a transaction consists of



The function must:

Process transactions concurrently
Handle potential race conditions
Consider the next transaction when a transfer would cause a negative balance
Maintain accurate account balances throughout


Transaction Handling Logic

When processing transactions:

If a transfer would cause a negative balance, pause it.
Check if the next transaction increases the sender's balance.
If so, retry the paused transaction after the next one completes.
Only consider the immediate next transaction, not any beyond that.


Function Description

Complete the function processTransactions in the editor with the following parameter(s):

    balances map[int]int: the initial balances for each account ID



Return

The function returns two channels:

The transactions channel through which transactions can be sent from main.
The done channel that processTransactions will close when it receives the notification from main that all transactions have been sent.


Constraints

Each account ID is a unique positive integer.
The amount to be transferred is a non-negative integer.
The number of transactions will not exceed 10000
The balance of each account will not exceed 1000000.*/
