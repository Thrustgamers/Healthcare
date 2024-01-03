package utils

import (
	"sync"
	"time"
)

// UniqueIDGenerator generates unique 8-digit integers
type UniqueIDGenerator struct {
	counter int
	mutex   sync.Mutex
}

// NewUniqueIDGenerator initializes a new UniqueIDGenerator
func NewUniqueIDGenerator() *UniqueIDGenerator {
	return &UniqueIDGenerator{}
}

// GenerateID generates a unique 8-digit integer
func (ug *UniqueIDGenerator) GenerateID() int {
	ug.mutex.Lock()
	defer ug.mutex.Unlock()

	// Get current timestamp in nanoseconds
	timestamp := time.Now().UnixNano()

	// Combine timestamp with counter to ensure uniqueness
	uniqueID := timestamp + int64(ug.counter)

	// Increment counter for the next call
	ug.counter++

	// Extract the last 8 digits to ensure the ID is 8 digits long
	return int(uniqueID % 100000000)
}
