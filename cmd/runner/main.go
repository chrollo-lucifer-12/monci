package main

import (
	"fmt"
	"os"

	"github.com/monci/executor"
	"github.com/monci/workflow"
)

const YAML_PATH = "ymls/test.yml"

func main() {

	data, err := os.ReadFile(YAML_PATH)
	if err != nil {
		panic(err)
	}

	workflow, err := workflow.ParseYaml(data)
	if err != nil {
		panic(err)
	}

	for _, job := range workflow.Jobs {
		fmt.Printf("running job: %s\n", job.Name)

		err := executor.RunJob(job)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("pipeline succeeded")
}
