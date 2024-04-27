package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {

	listOfSentence := []string{"This is my name",
		"my name is Anish",
		"My name is Ashay bhai",
		"my name is Anish",
	}

	//Creating a intermediate value to store map result: [[This: 1, is:1, my:1, name:1], [my: 1, name: 1, is:1....]...]

	//Map Part
	mapResult := &[]map[string]int{}
	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	for _, each := range listOfSentence {
		wg.Add(1)
		go func(wg *sync.WaitGroup, mut *sync.RWMutex, mapResult *[]map[string]int, each string) {
			defer wg.Done()
			eachList := strings.Split(each, " ")
			hashmap := make(map[string]int)

			for _, item := range eachList {
				hashmap[item] += 1
			}

			mut.Lock()
			*mapResult = append(*mapResult, hashmap)
			mut.Unlock()

		}(wg, mut, mapResult, each)
	}
	wg.Wait()

	// Reducer part
	finalResult := make(map[string]int)
	for _, each := range *mapResult {
		wg.Add(1)
		go func(wg *sync.WaitGroup, mut *sync.RWMutex, each map[string]int, finalResult map[string]int) {
			defer wg.Done()
			for key, value := range each {
				mut.Lock()
				finalResult[key] += value
				mut.Unlock()
			}
		}(wg, mut, each, finalResult)
	}

	wg.Wait()
	fmt.Println("This is the final result", finalResult)
}
