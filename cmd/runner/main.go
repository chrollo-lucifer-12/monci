package main

import (
	"fmt"
	"os"
	"os/exec"

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

		for _, step := range job.Steps {
			fmt.Printf("running step: %s\n", step.Name)

			cmd := exec.Command("bash", "-c", step.Run)

			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			err := cmd.Run()
			if err != nil {
				panic(err)
			}
		}

	}
}
