package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

func readFile(filename string) (map[string]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	popMap := make(map[string]int)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		capital := scanner.Text()
		if !scanner.Scan() {
			return nil, fmt.Errorf("missing population for capital: %s", capital)
		}
		popStr := scanner.Text()
		pop, err := strconv.Atoi(popStr)
		if err != nil {
			return nil, fmt.Errorf("invalid population for %s: %v", capital, err)
		}
		popMap[capital] = pop
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanner error: %w", err)
	}

	return popMap, nil
}

// note that it starts with lowercase 's', we don't want it to be accessed directly
// i.e. it's private to this package, use a factory function to initialise this
// the singletonDatabase should only be instantiated once
type singletonDatabase struct {
	populations map[string]int
}

var once sync.Once // construct used to initialise something only once,
// used to enforce thread safety and lazy instantiation

var instance *singletonDatabase

// GetPopulation and GetSingletonDatabase violate the dependancy inversion principle
// because they define rigid methods to access the data. We need to have abstractions/interfaces,
// this helps in unit tests where you may use a dummy database
func (db *singletonDatabase) GetPopulation(city string) (int, error) {
	return db.populations[city], nil
}

func GetSingletonDatabase() *singletonDatabase {
	once.Do(func() {
		pops, e := readFile("./population.txt")
		db := singletonDatabase{}
		if e != nil {
			panic("Error reading file")
		}
		db.populations = pops
		instance = &db
	})
	// if called for the first time, once.Do will run, create and return instance,
	// if multiple threads call it, the first thread instantiates, others get the instance var.
	return instance
}

func main() {
	db := GetSingletonDatabase()
	pop, _ := db.GetPopulation("Tokyo")
	fmt.Printf("Population of Tokyo: %d\n", pop)
}
