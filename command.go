package shrub

import "time"

type Command interface {
	Resolve() *CommandDefinition
	Validate() error
}

type CommandDefinition struct {
	Function    string                 `json:"func,omitempty"`
	Type        string                 `json:"type,omitempty"`
	DisplayName string                 `json:"display_name,omitempty"`
	Command     string                 `json:"command,omitempty"`
	Variants    []string               `json:"variants"`
	TimeoutSecs int                    `json:"timeout_secs,omitempty"`
	Params      map[string]interface{} `json:"params,omitempty"`
	Vars        map[string]string      `json:"vars,omitempty"`
}

func (c *CommandDefinition) Validate() error                         { return nil }
func (c *CommandDefinition) Resolve() *CommandDefinition             { return c }
func (c *CommandDefinition) Function(n string) *CommandDefinition    { c.Function = n; return c }
func (c *CommandDefinition) Type(n string) *CommandDefinition        { c.Type = n; return c }
func (c *CommandDefinition) DisplayName(n string) *CommandDefinition { c.DisplayName = n; return c }
func (c *CommandDefinition) Command(n string) *CommandDefinition     { c.Command = n; return c }
func (c *CommandDefinition) Timeout(s time.Duration) *CommandDefinition {
	c.TimeoutSecs = int(s.Seconds())
	return c
}
func (c *CommandDefinition) Variants(vs ...string) *CommandDefinition {
	c.Variants = append(c.Variants, vs...)
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

////////////////////////////////////////////////////////////////////////
//
// Specific Command Implementations

type CmdExec struct {
	Args             []string
	Env              map[string]string
	WorkingDirectory string
}

func (c CmdExec) Resolve() *CommandDefinition { return nil }
func (c CmdExec) Validate() error             { return nil }

type CmdExecShell struct{}

func (c CmdExecShell) Resolve() *CommandDefinition { return nil }
func (c CmdExecShell) Validate() error             { return nil }

type CmdS3Put struct{}

func (c CmdS3Put) Resolve() *CommandDefinition { return nil }
func (c CmdS3Put) Validate() error             { return nil }

type CmdS3Get struct{}

func (c CmdS3Get) Resolve() *CommandDefinition { return nil }
func (c CmdS3Get) Validate() error             { return nil }

type CmdS3Copy struct{}

func (c CmdS3Copy) Resolve() *CommandDefinition { return nil }
func (c CmdS3Copy) Validate() error             { return nil }

type CmdGetProject struct{}

func (c CmdGetProject) Resolve() *CommandDefinition { return nil }
func (c CmdGetProject) Validate() error             { return nil }

type CmdResultsJSON struct{}

func (c CmdResultsJSON) Resolve() *CommandDefinition { return nil }
func (c CmdResultsJSON) Validate() error             { return nil }

type CmdResultsXunit struct{}

func (c CmdResultsXunit) Resolve() *CommandDefinition { return nil }
func (c CmdResultsXunit) Validate() error             { return nil }

type CmdResultsGoTest struct{}

func (c CmdResultsGoTest) Resolve() *CommandDefinition { return nil }
func (c CmdResultsGoTest) Validate() error             { return nil }

type CmdArchiveCreate struct{}

func (c CmdArchiveCreate) Resolve() *CommandDefinition { return nil }
func (c CmdArchiveCreate) Validate() error             { return nil }

type CmdArchiveExtract struct{}

func (c CmdArchiveExtract) Resolve() *CommandDefinition { return nil }
func (c CmdArchiveExtract) Validate() error             { return nil }

type CmdAttachArtifacts struct{}

func (c CmdAttachArtifacts) Resolve() *CommandDefinition { return nil }
func (c CmdAttachArtifacts) Validate() error             { return nil }
