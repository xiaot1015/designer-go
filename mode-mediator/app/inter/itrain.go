package inter

type ITrain interface {
	Arrive()
	Depart()
	PreArrival()
}