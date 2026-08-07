package workflow

import (
	"gopkg.in/yaml.v3"
)

type Workflow struct {
	Name string `yaml:"name"`
	//	On   yaml.Node
	//	Env  map[string]string `yaml:"env"`
	Jobs map[string]Job `yaml:"jobs"`
}

type Job struct {
	Name string `yaml:"name,omitempty"`
	//	RunsOn string            `yaml:"runs-on"`
	//	Env   map[string]string `yaml:"env,omitempty"`
	Steps []Step `yaml:"steps"`
}

type Step struct {
	Name string `yaml:"name,omitempty"`
	//	Uses            string            `yaml:"uses,omitempty"`
	Run string `yaml:"run,omitempty"`
	//	With            map[string]any    `yaml:"with,omitempty"`
	//	Env             map[string]string `yaml:"env,omitempty"`
	//
	// /	ContinueOnError bool              `yaml:"continue-on-error,omitempty"`
}

func ParseYaml(data []byte) (Workflow, error) {
	var workflow Workflow
	if err := yaml.Unmarshal(data, &workflow); err != nil {
		return workflow, err
	}
	return workflow, nil
}
