package shrub

import (
	"encoding/json"
	"errors"
	"time"
)

type Command interface {
	Resolve() *CommandDefinition
	Validate() error
}

type CommandDefinition struct {
	FunctionName  string                 `json:"func,omitempty"`
	ExecutionType string                 `json:"type,omitempty"`
	DisplayName   string                 `json:"display_name,omitempty"`
	CommandName   string                 `json:"command,omitempty"`
	RunVariants   []string               `json:"variants"`
	TimeoutSecs   int                    `json:"timeout_secs,omitempty"`
	Params        map[string]interface{} `json:"params,omitempty"`
	Vars          map[string]string      `json:"vars,omitempty"`
}

func (c *CommandDefinition) Validate() error                      { return nil }
func (c *CommandDefinition) Resolve() *CommandDefinition          { return c }
func (c *CommandDefinition) Function(n string) *CommandDefinition { c.FunctionName = n; return c }
func (c *CommandDefinition) Type(n string) *CommandDefinition     { c.ExecutionType = n; return c }
func (c *CommandDefinition) Name(n string) *CommandDefinition     { c.DisplayName = n; return c }
func (c *CommandDefinition) Command(n string) *CommandDefinition  { c.CommandName = n; return c }
func (c *CommandDefinition) Timeout(s time.Duration) *CommandDefinition {
	c.TimeoutSecs = int(s.Seconds())
	return c
}
func (c *CommandDefinition) Variants(vs ...string) *CommandDefinition {
	c.RunVariants = append(c.RunVariants, vs...)
	return c
}
func (c *CommandDefinition) ResetVars() *CommandDefinition                      { c.Vars = nil; return c }
func (c *CommandDefinition) ResetParams() *CommandDefinition                    { c.Params = nil; return c }
func (c *CommandDefinition) ReplaceVars(v map[string]string) *CommandDefinition { c.Vars = v; return c }
func (c *CommandDefinition) ReplaceParams(v map[string]interface{}) *CommandDefinition {
	c.Params = v
	return c
}

func (c *CommandDefinition) Param(k string, v interface{}) *CommandDefinition {
	if c.Params == nil {
		c.Params = make(map[string]interface{})
	}

	c.Params[k] = v

	return c
}

func (c *CommandDefinition) ExtendParams(p map[string]interface{}) *CommandDefinition {
	if c.Params == nil {
		c.Params = p
		return c
	}

	for k, v := range p {
		c.Params[k] = v
	}

	return c
}

func (c *CommandDefinition) Var(k, v string) *CommandDefinition {
	if c.Vars == nil {
		c.Vars = make(map[string]string)
	}

	c.Vars[k] = v
	return c
}

func (c *CommandDefinition) ExtendVars(vars map[string]string) *CommandDefinition {
	if c.Vars == nil {
		c.Vars = vars
		return c
	}

	for k, v := range vars {
		c.Vars[k] = v

	}

	return c
}

type CommandSequence []*CommandDefinition

func (s CommandSequence) Command() *CommandDefinition {
	c := &CommandDefinition{}
	s = append(s, c)
	return c
}

func (s CommandSequence) Append(c ...*CommandDefinition) CommandSequence {
	s = append(s, c...)
	return s
}

func (s CommandSequence) Add(cmd Command) CommandSequence { s = append(s, cmd.Resolve()); return s }

func (s CommandSequence) Extend(cmds ...Command) CommandSequence {
	for _, cmd := range cmds {
		s = append(s, cmd.Resolve())
	}
	return s
}

////////////////////////////////////////////////////////////////////////
//
// Specific Command Implementations

func exportCmd(cmd Command) map[string]interface{} {
	if err := cmd.Validate(); err != nil {
		panic(err)
	}

	jsonStruct, err := json.Marshal(cmd)
	if err != nil {
		panic(err)
	}

	out := map[string]interface{}{}
	if err := json.Unmarshal(jsonStruct, &out); err != nil {
		panic(err)
	}

	return out
}

type CmdExec struct {
	Background       bool   `json:"background"`
	Silent           bool   `json:"silent"`
	ContinueOnError  bool   `json:"continue_on_err"`
	SystemLog        bool   `json:"system_log"`
	CombineOuutput   bool   `json:"redirect_standard_error_to_output"`
	IgnoreStdError   bool   `json:"ignore_standard_error"`
	IgnoreStdOut     bool   `json:"ignore_standard_out"`
	KeepEmptyArgs    bool   `json:"keep_empty_args"`
	WorkingDirectory string `json:"working_dir"`
	Command          string
	Binary           string
	Args             []string
	Env              map[string]string
}

func (c CmdExec) Validate() error { return nil }
func (c CmdExec) Resolve() *CommandDefinition {
	return &CommandDefinition{
		CommandName: "sub process.exec",
		Params:      exportCmd(c),
	}
}

type CmdExecShell struct {
	Background       bool   `json:"background"`
	Silent           bool   `json:"silent"`
	ContinueOnError  bool   `json:"continue_on_err"`
	SystemLog        bool   `json:"system_log"`
	CombineOuutput   bool   `json:"redirect_standard_error_to_output"`
	IgnoreStdError   bool   `json:"ignore_standard_error"`
	IgnoreStdOut     bool   `json:"ignore_standard_out"`
	WorkingDirectory string `json:"working_dir"`
	Script           string `json:"script"`
}

func (c CmdExecShell) Validate() error { return nil }
func (c CmdExecShell) Resolve() *CommandDefinition {
	return &CommandDefinition{
		FunctionName: "shell.exec",
		Params:       exportCmd(c),
	}
}

type CmdS3Put struct {
	Optional               bool     `json:"optional"`
	LocalFile              string   `json:"local_file"`
	LocalFileIncludeFilter []string `json:"local_files_include_filter"`
	Bucket                 string   `json:"bucket"`
	RemoteFile             string   `json:"remote_file"`
	DisplayName            string   `json:"display_name"`
	ContentType            string   `json:"content_type"`
	CredKey                string   `json:"aws_key"`
	CredSecret             string   `json:"aws_secret"`
	Permissions            string   `json:"permissions"`
	Visibility             string   `json:"visibility"`
	BuildVariants          []string `json:"build_variants"`
}

func (c CmdS3Put) Validate() error {
	switch {
	case c.CredKey == "", c.CredSecret == "":
		return errors.New("must specify aws credentials")
	case c.LocalFile == "" && len(c.LocalFileIncludeFilter) == 0:
		return errors.New("must specify a local file to upload")
	default:
		return nil
	}
}
func (c CmdS3Put) Resolve() *CommandDefinition {
	return &CommandDefinition{
		FunctionName: "s3.put",
		Params:       exportCmd(c),
	}
}

type CmdS3Get struct{}

func (c CmdS3Get) Validate() error { return nil }
func (c CmdS3Get) Resolve() *CommandDefinition {
	return &CommandDefinition{
		FunctionName: "s3.get",
		Params:       exportCmd(c),
	}
}

type CmdS3Copy struct{}

func (c CmdS3Copy) Validate() error { return nil }
func (c CmdS3Copy) Resolve() *CommandDefinition {
	return &CommandDefinition{
		FunctionName: "s3Copy.copy",
		Params:       exportCmd(c),
	}
}

type CmdGetProject struct{}

func (c CmdGetProject) Validate() error { return nil }
func (c CmdGetProject) Resolve() *CommandDefinition {
	return &CommandDefinition{
		FunctionName: "git.get_project",
		Params:       exportCmd(c),
	}
}

type CmdResultsJSON struct{}

func (c CmdResultsJSON) Validate() error { return nil }
func (c CmdResultsJSON) Resolve() *CommandDefinition {
	return &CommandDefinition{
		FunctionName: "attach.results",
		Params:       exportCmd(c),
	}
}

type CmdResultsXunit struct{}

func (c CmdResultsXunit) Validate() error { return nil }
func (c CmdResultsXunit) Resolve() *CommandDefinition {
	return &CommandDefinition{
		FunctionName: "attach.xunit_results",
		Params:       exportCmd(c),
	}
}

type CmdResultsGoTest struct {
	JSONFormat   bool `json:"-"`
	LegacyFormat bool `json:"-"`
}

func (c CmdResultsGoTest) Validate() error { return nil }
func (c CmdResultsGoTest) Resolve() *CommandDefinition {
	if c.JSONFormat {
		return &CommandDefinition{
			FunctionName: "gotest.parse_json",
			Params:       exportCmd(c),
		}
	}

	return &CommandDefinition{
		FunctionName: "gotest.parse_files",
		Params:       exportCmd(c),
	}
}

type CmdArchiveCreate struct {
	Auto bool `json:"-"`
}

func (c CmdArchiveCreate) Validate() error { return nil }
func (c CmdArchiveCreate) Resolve() *CommandDefinition {
	return &CommandDefinition{
		FunctionName: "",
		Params:       exportCmd(c),
	}
}

type CmdArchiveExtract struct {
	Auto bool `json:"-"`
}

func (c CmdArchiveExtract) Validate() error { return nil }
func (c CmdArchiveExtract) Resolve() *CommandDefinition {
	return &CommandDefinition{
		FunctionName: "",
		Params:       exportCmd(c),
	}
}

type CmdAttachArtifacts struct{}

func (c CmdAttachArtifacts) Validate() error { return nil }
func (c CmdAttachArtifacts) Resolve() *CommandDefinition {
	return &CommandDefinition{
		FunctionName: "",
		Params:       exportCmd(c),
	}
}
