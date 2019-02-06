package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	textPtr := flag.String("text", "", "Text to parse. (Required)")
	metricPtr := flag.String("metric", "chars", "Metric {chars|words|lines};.")
	uniquePtr := flag.Bool("unique", false, "Measure unique values of a metric.")
	flag.Parse()

	if *textPtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Printf("textPtr: %s, metricPtr: %s, uniquePtr: %t\n", *textPtr, *metricPtr, *uniquePtr)

	// Prints out arguments that come *after* flags
	for i, f := range flag.Args() {
		fmt.Printf("Arg %d %s\n", i, f)
	}
}
