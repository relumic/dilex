// Package client is the public API for DILEX.
//
// DILEX provides priority-ordered data delivery between two nodes across
// an unreliable connection. Highest-priority payloads transmit first,
// always. Data is persisted before transmission and survives crashes,
// drops, and restarts without loss.
//
// The entire public API surface is three methods:
//
//	Send(data []byte, priority int) error
//	Receive() ([]byte, error)
//	Status() (Status, error)
//
// All other implementation details are internal to this module.
package client
