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

	//Map Part
	//Creating a intermediate value to store map result: [[This: 1, is:1, my:1, name:1], [my: 1, name: 1, is:1....]...]
	mapResult := &[]map[string]int{}
	wg := &sync.WaitGroup{} //Pointer for threads
	mut := &sync.RWMutex{}  //Locking the common mapResult while each thread updates it

	for _, each := range listOfSentence {
		wg.Add(1)
		//Call a go routine for each sentense and update final result to mapResult
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
	finalResult := make(map[string]int) //Store the final result
	for _, each := range *mapResult {
		wg.Add(1)
		//call a go routine to update each map result
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
