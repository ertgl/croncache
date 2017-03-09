package lib

type IoC interface {
	Node() Node
	SetNode(Node)
	Executer() Executer
	SetExecuter(Executer)
	TaskManager() TaskManager
	SetTaskManager(TaskManager)
}
