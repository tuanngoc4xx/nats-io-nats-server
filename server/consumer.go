package server

// Note: This is a simplified representation of the fix in server/consumer.go.
// In the actual codebase, we ensure that after loading a message for redelivery (which releases the lock),
// we re-acquire the lock and verify that the message is still in c.pending before proceeding.
// If the message is no longer in c.pending, we skip it.

// We will modify the getNextMsg function to include the check:
// if _, ok := c.pending[sseq]; !ok {
//     continue
// }
// and ensure that sequence numbers are updated atomically.