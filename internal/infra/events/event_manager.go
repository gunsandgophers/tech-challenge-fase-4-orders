package events

type EventManager struct {
	lst map[string][]func()
}

func NewEventManager() *EventManager {
	return &EventManager{
		lst: make(map[string][]func()),
	}
}

func (m *EventManager) Invoke(n string) {
	for _, ls := range m.lst[n] {
		ls()
	}
}

func (m *EventManager) Add(n string, l func()) {
	m.lst[n] = append(m.lst[n], l)
}
