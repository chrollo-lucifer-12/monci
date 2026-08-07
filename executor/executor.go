package executor

import (
	"fmt"

	"github.com/monci/workflow"
)

func RunJob(job workflow.Job) error {

	fmt.Printf("=== %s ===", job.Name)
	fmt.Println()
	fmt.Println()
	fmt.Println()

	containerName := "test"

	err := createContainer(containerName)
	if err != nil {
		return err
	}
	err = startContainer(containerName)
	if err != nil {
		return err
	}

	for i, step := range job.Steps {
		fmt.Printf("[%d/%d] %s", i+1, len(job.Steps), step.Name)
		fmt.Println()
		fmt.Println()

		fmt.Println(step.Run)

		err := execCommand(containerName, step.Run)

		if err != nil {
			fmt.Printf("step failed: %s\n", step.Name)
			err = stopContainer(containerName)
			if err != nil {
				return err
			}
			return err
		}

		fmt.Printf(">\n")
		fmt.Println()
	}

	err = stopContainer(containerName)
	if err != nil {
		return err
	}

	fmt.Printf("job succeeded\n")

	return nil
}
