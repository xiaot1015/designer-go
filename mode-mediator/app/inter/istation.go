package inter

type IStation interface {
	CanArrive(ITrain) bool
	NotifyAboutDeparture()
}