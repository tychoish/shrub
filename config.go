package shrub

type Configuration struct {
	Functions map[string]CommandSequence `json:"functions"`
	Tasks     []*Task                    `json:"tasks"`
	Groups    []*TaskGroup               `json:"groups"`
	Variants  []*Variant                 `json:"variants"`
	Pre       CommandSequence            `json:"pre"`
	Post      CommandSequence            `json:"post"`
	Timeout   CommandSequence            `json:"timeout"`
}

func (c *Configuration) Task(name string) *Task {
	for _, t := range c.Tasks {
		if t.Name == name {
			return t
		}
	}

	t := new(Task)
	t.Name = name
	c.Tasks = append(c.Tasks, t)

	return t
}

func (c *Configuration) TaskGroup(name string) *TaskGroup {
	for _, g := range c.Groups {
		if g.GroupName == name {
			return g
		}
	}

	g := new(TaskGroup)
	c.Groups = append(c.Groups, g)
	return g.Name(name)
}

func (c *Configuration) Function(name string) CommandSequence {
	if c.Functions == nil {
		c.Functions = make(map[string]CommandSequence)
	}

	seq, ok := c.Functions[name]
	if !ok {
		c.Functions[name] = seq
	}

	return seq
}

func (c *Configuration) Variant(id string) *Variant {
	for _, v := range c.Variants {
		if v.BuildName == id {
			return v
		}
	}

	v := new(Variant)
	c.Variants = append(c.Variants, v)
	return v.Name(id)
}
