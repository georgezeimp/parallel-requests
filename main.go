package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/georgezeimp/parallel-requests/hasher"
	"github.com/georgezeimp/parallel-requests/output"
	"github.com/georgezeimp/parallel-requests/processor"
	"github.com/georgezeimp/parallel-requests/request"
)

func main() {
	// Read and prepare input
	parallel := flag.String("parallel", "", "")
	flag.Parse()
	addresses := flag.Args()
	npr, err := strconv.Atoi(*parallel)
	if err != nil {
		npr = 10
	}

	// Prepare utils
	hasher := hasher.NewHasher()
	outputPresenter := output.NewPresenter()
	requestService := request.NewService(hasher, outputPresenter)
	processor := processor.NewProcessor(hasher, outputPresenter, requestService)
	result := processor.Process(addresses, npr)

	// Print output
	for _, item := range result {
		fmt.Println(item)
	}
}
