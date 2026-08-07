package executor

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/monci/workflow"
)

func RunJob(job workflow.Job) error {

	fmt.Printf("=== %s ===", job.Name)
	fmt.Println()
	fmt.Println()
	fmt.Println()

	for i, step := range job.Steps {
		fmt.Printf("[%d/%d] %s", i+1, len(job.Steps), step.Name)
		fmt.Println()
		fmt.Println()

		cmd := exec.Command("bash", "-c", step.Run)

		fmt.Println(cmd.String())

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			fmt.Printf("step failed: %s\n", step.Name)
			return err
		}

		fmt.Printf(">\n")
		fmt.Println()
	}

	fmt.Printf("job succeeded\n")

	return nil
}
