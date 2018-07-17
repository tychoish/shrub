package shrub

type TaskGroup struct {
	Name string `json:"name"`
}

func (g *TaskGroup) AddTask(tasks ...string) *TaskGroup        { return g }
func (g *TaskGroup) SetMaxHosts(num int) *TaskGroup            { return g }
func (g *TaskGroup) SetupTask(cmds ...Command) *TaskGroup      { return g }
func (g *TaskGroup) AddSetupTask() *CommandDefinition          { return &CommandDefinition{} }
func (g *TaskGroup) SetupGroup(cmds ...Command) *TaskGroup     { return g }
func (g *TaskGroup) AddSetupGroup() *CommandDefinition         { return &CommandDefinition{} }
func (g *TaskGroup) TeardownTask(cmds ...Command) *TaskGroup   { return g }
func (g *TaskGroup) AddTeardownTask() *CommandDefinition       { return &CommandDefinition{} }
func (g *TaskGroup) TeardownGroup(cmds ...Command) *TaskGroup  { return g }
func (g *TaskGroup) AddTeardownGroup() *CommandDefinition      { return &CommandDefinition{} }
func (g *TaskGroup) TimeoutHandler(cmds ...Command) *TaskGroup { return g }
func (g *TaskGroup) AddTimeoutHandler() *CommandDefinition     { return &CommandDefinition{} }
