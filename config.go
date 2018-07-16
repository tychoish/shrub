package shrub

type Configuration struct{}
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

func (c *Configuration) AddTask(name string) *Task  { return &Task{} }
func (t *Task) Commmand(cmd *Command) *Task         { return t }
func (t *Task) Dependency(dep TaskDependency) *Task { return t }
func (t *Task) Function(fn string) *Task            { return t }
func (t *Task) Priority(pri int) *Task              { return t }

func (c *Configuration) AddVariant(id string) *Variant                       { return &Variant{} }
func (v *Variant) Name(id string) *Variant                                   { return v }
func (v *Variant) DisplayName(id string) *Variant                            { return v }
func (v *Variant) RunOn(distro string) *Variant                              { return v }
func (v *Variant) Expansions(m map[string]interface{}) *Variant              { return v }
func (v *Variant) Expansion(k string, v interface{}) *Variant                { return v }
func (v *Variant) AddTask(name string) *Variant                              { return v }
func (v *Variant) AddTaskWithOptions(name string, opts TaskOptions) *Variant { return v }
func (v *Variant) DisplayTasks(def DisplayTaskDefinition) *Variant           { return v }

func (c *Configuration) TaskGroup(name string) *TaskGroup      { return &TaskGroup{} }
func (t *TaskGroup) AddTask(tasks ...[]string) *TaskGroup      { return t }
func (t *TaskGroup) SetMaxHosts(num int) *TaskGroup            { return t }
func (t *TaskGroup) AddSetupTask(cmd *Command) *TaskGroup      { return t }
func (t *TaskGroup) AddSetupGroup(cmd *Command) *TaskGroup     { return t }
func (t *TaskGroup) AddTeardownTask(cmd *Command) *TaskGroup   { return t }
func (t *TaskGroup) AddTeardownGroup(cmd *Command) *TaskGroup  { return t }
func (t *TaskGroup) AddTimeoutHandler(cmd *Command) *TaskGroup { return t }

func (c *Configuration) AddFunction(name string) *Function       { return &Function{} }
func (f *Function) Command(name string, args *Command) *Function { return f }
