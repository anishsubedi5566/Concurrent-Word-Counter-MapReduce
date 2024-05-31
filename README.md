Basic MapReduce in Go

This repository implements a basic MapReduce framework in Go to demonstrate the core concepts of MapReduce programming.

Understanding MapReduce

MapReduce is a programming paradigm designed for processing and generating large datasets. It divides the work into two phases:

Map: Processes each data element independently, transforming it into intermediate key-value pairs.
Reduce: Groups intermediate key-value pairs with the same key and aggregates the values.
Functionality

This code performs a word count by splitting sentences into words, creating key-value pairs for each word with its count, and then summing the counts for each unique word.

Code Structure

main.go: The main program logic.
listOfSentence: A sample list of sentences to process.
mapResult: A pointer to a slice of maps to store intermediate results from the Map phase. This uses a pointer to avoid copying the entire data structure within goroutines.
wg: A sync.WaitGroup object to synchronize goroutines.
mut: A sync.RWMutex object to manage concurrent access to the mapResult slice.
Map Phase

Iterates through each sentence in listOfSentence.
Launches a goroutine for each sentence.
Splits the sentence into words using strings.Split.
Creates a temporary map hashmap to store word counts.
Increments the count for each word in hashmap.
Acquires a write lock on mut to ensure safe concurrent access.
Appends the hashmap to the mapResult slice using a pointer dereference.
Releases the lock.
Reduce Phase

Creates an empty map finalResult to store the final word counts.
Iterates through each intermediate map in mapResult.
Launches a goroutine for each intermediate map.
Iterates through each key-value pair in the current intermediate map.
Acquires a write lock on mut for safe concurrent access.
Updates the count for the key in finalResult, adding the current value if it exists.
Releases the lock.
Wait Groups and Mutexes

sync.WaitGroup ensures all goroutines complete before printing the final result.
sync.RWMutex provides synchronized access to the shared mapResult slice, preventing data races during concurrent updates.
Running the Code

Clone this repository.
Open a terminal in the project directory.
Run the Go program: go run main.go
The output will display the word counts:
