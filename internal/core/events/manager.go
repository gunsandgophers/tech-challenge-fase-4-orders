package events

type ManagerEvent interface {
	Add(n string, l func())
}
