package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"math"
)

// GenerateUniqueID generates a unique 8-digit integer
func GenerateUniqueID() int {
	// Generate a random 64-bit integer
	var randNum int64
	err := binary.Read(rand.Reader, binary.BigEndian, &randNum)
	if err != nil {
		return 0
	}

	// Hash the random number to ensure uniqueness
	hash := hashInt64(randNum)

	// Take the last 8 digits to ensure the ID is 8 digits long
	uniqueID := int(hash % 100000000)

	return uniqueID
}

// hashInt64 hashes a 64-bit integer using a cryptographic hash function
func hashInt64(value int64) int64 {
	// Create a new hash.Hash object
	h := sha256.New()

	// Convert the int64 value to a byte slice
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(value))

	// Write the byte slice to the hash object
	h.Write(b)

	// Sum returns the hash of the data written to the hash object so far
	// It doesn't need any arguments as it incorporates all data written to the hash
	sum := h.Sum(nil)

	// Convert the first 8 bytes of the hash sum to an int64 value
	// You can adjust the byte range based on your specific requirements
	result := int64(binary.BigEndian.Uint64(sum[:8]))

	// Take the absolute value to ensure it's positive
	result = int64(math.Abs(float64(result)))

	return result
}
