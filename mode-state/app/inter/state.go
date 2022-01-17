package inter

type State interface {
	Ordered()
	OrderPaying()
	OrderPay()
	OrderCancel()
	OrderRefund()
}
