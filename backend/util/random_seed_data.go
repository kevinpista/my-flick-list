package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet ="abcdefghijklmnopqrstuvwxyz"

// Ensures everytime we run function, the generated values are different
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Generates a random integer between min & max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max - min +1)
}

// Generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// Generates random name between length 4 and 24 characters
func RandomName() string {
	length := rand.Intn(24 - 4) + 4
	return RandomString(length)
}

// Generates a random paragraph of made-up words of at least 10-14 words; word lengths between 2-14 chars
func RandomParagraph() string {
	// Mnimum and maximum number of sentences
	minSentences := 1
	maxSentences := 3
  
	// Minimum and maximum number of words per sentence
	minWords := 10
	maxWords := 14
  
	// Generate random number of sentences
	numSentences := rand.Intn(maxSentences - minSentences +1) + minSentences
  
	paragraph := ""
	for i := 0; i < numSentences; i++ {
	  // Generate random number of words for the current sentence
	  numWords := rand.Intn(maxWords - minWords +1) + minWords
  
	  sentence := ""
	  for j := 0; j < numWords; j++ {
		// Generate random word length between 2 and 14 characters
		wordLength := rand.Intn(14-2) + 2
		word := RandomString(wordLength)
		sentence += word + " "
	  }
  
	  // Remove trailing space from the sentence
	  sentence = strings.TrimSpace(sentence)
  
	  // Add a period and a space to the end of the sentence
	  sentence += "."
  
	  paragraph += sentence + " "
	}
  
	// Remove trailing space from the paragraph
	return strings.TrimSpace(paragraph)
}

/*
func RandomInt64() int64 {
	return RandomInt(0, 1000)
}
*/

