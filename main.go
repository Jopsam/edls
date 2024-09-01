package main

import (
	"flag"
	"fmt"
)

func main() {
	// Define a flag named p with a default value of an empty string. The String method returns a string pointer.
	patternFlag := flag.String("p", "", "Filter by pattern.")
	allFlag := flag.Bool("a", false, "List all files, including the hidden ones.")
	numberRecordsFlag := flag.Int("n", 10, "Number of records to display.")
	hasOrderByTimeFlag := flag.Bool("t", false, "Order by time, oldest first.")
	hasOrderBySizeFlag := flag.Bool("s", false, "Order by size, smallest first.")
	hasOrderReverseFlag := flag.Bool("r", false, "Reverse the order of the records.")

	// Parse the command-line flags. The flag values are available after parsing.
	flag.Parse()

	// Print the flag value. We need to dereference the flag value to get the string value.
	fmt.Println("patternFlag:", *patternFlag)
	fmt.Println("allFlag:", *allFlag)
	fmt.Println("numberRecordsFlag:", *numberRecordsFlag)
	fmt.Println("hasOrderByTimeFlag:", *hasOrderByTimeFlag)
	fmt.Println("hasOrderBySizeFlag:", *hasOrderBySizeFlag)
	fmt.Println("hasOrderReverseFlag:", *hasOrderReverseFlag)
}
