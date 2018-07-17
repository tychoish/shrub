package shrub

type Configuration struct {
	functions []*Function
	tasks     []*Task
	groups    []*TaskGroup
	variants  []*Variant
}

type Task struct{}
type Variant struct{}
type TaskGroup struct{}
type Function struct{}

type TaskOptions struct {
	Stepback bool
	Distro   string
}

type DisplayTaskDefinition struct {
	Name       string
	Components []string
}

type TaskDependency struct {
	Name    string
	Variant string
}

func (c *Configuration) Task(name string) *Task {
	for _, t := range c.tasks {
		if t.name == name {
			return t
		}
	}

	t = new(Task)
	t.name = name
	c.tasks = append(c.tasks, t)

	return t
}

func (c *Configuration) TaskGroup(name string) *TaskGroup {
	for _, g := range c.groups {
		if g.name == name {
			return g
		}
	}

	g = new(TaskGroup)
	c.name = name
	c.groups = append(c.groups, g)
	return g
}

func (c *Configuration) Function(name string) *Function {
	for _, f := range c.functions {
		if f.name == name {
			return f
		}
	}

	f = new(Function)
	f.name = name
	c.functions = append(c.functions, f)
	return f
}

func (c *Configuration) Variant(id string) *Variant {
	for _, v := range c.variants {
		if v.name == id {
			return v
		}
	}

	v = new(Variant)
	v.name = id
	c.variants = append(c.variants, v)
	return v
}

func (t *Task) Command(cmd Command) *Task                                    { return t }
func (t *Task) AddCommand() *CommandDefinition                               { return &CommandDefinition{} }
func (t *Task) Dependency(dep TaskDependency) *Task                          { return t }
func (t *Task) Function(fn string) *Task                                     { return t }
func (t *Task) Priority(pri int) *Task                                       { return t }
func (v *Variant) Name(id string) *Variant                                   { return v }
func (v *Variant) DisplayName(id string) *Variant                            { return v }
func (v *Variant) RunOn(distro string) *Variant                              { return v }
func (v *Variant) Expansions(m map[string]interface{}) *Variant              { return v }
func (v *Variant) Expansion(k string, v interface{}) *Variant                { return v }
func (v *Variant) AddTask(name string) *Variant                              { return v }
func (v *Variant) AddTaskWithOptions(name string, opts TaskOptions) *Variant { return v }
func (v *Variant) DisplayTasks(def DisplayTaskDefinition) *Variant           { return v }
func (g *TaskGroup) AddTask(tasks ...[]string) *TaskGroup                    { return g }
func (g *TaskGroup) SetMaxHosts(num int) *TaskGroup                          { return g }
func (g *TaskGroup) SetupTask(c Command) *TaskGroup                          { return g }
func (g *TaskGroup) AddSetupTask() *CommandDefinition                        { return &CommandDefinition{} }
func (g *TaskGroup) SetupGroup(c Command) *TaskGroup                         { return g }
func (g *TaskGroup) AddSetupGroup() *CommandDefinition                       { return &CommandDefinition{} }
func (g *TaskGroup) TeardownTask(c Command) *TaskGroup                       { return g }
func (g *TaskGroup) AddTeardownTask() *CommandDefinition                     { return &CommandDefinition{} }
func (g *TaskGroup) TeardownGroup(c Command) *TaskGroup                      { return g }
func (g *TaskGroup) AddTeardownGroup() *CommandDefinition                    { return &CommandDefinition{} }
func (g *TaskGroup) TimeoutHandler(c Command) *TaskGroup                     { return g }
func (g *TaskGroup) AddTimeoutHandler() *CommandDefinition                   { return &CommandDefinition{} }
func (f *Function) AddCommand() *CommandDefinition                           { return &CommandDefinition{} }
func (f *Function) Command(c Command) *Function                              { return f }
