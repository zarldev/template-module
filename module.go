package module

type Module struct {
	Toggled bool
}

func New(t bool) *Module {
	return &Module{
		Toggled: t,
	}
}

func (m *Module) Toggle() {
	if m == nil {
		return
	}
	m.Toggled = !m.Toggled
}

func (m *Module) State() bool {
	if m == nil {
		return false
	}
	return m.Toggled
}

func (m *Module) String() string {
	if m == nil {
		return "OFF"
	}
	if m.Toggled {
		return "ON"
	}
	return "OFF"
}
