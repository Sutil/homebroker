package entity

type OrdderQueue struct {
	Orders []*Order
}

func (oq *OrdderQueue) Less(i, j int) bool {
	return oq.Orders[i].Price < oq.Orders[j].Price
}

func (oq *OrdderQueue) Swap(i, j int) {
	oq.Orders[i], oq.Orders[j] = oq.Orders[j], oq.Orders[i]
}

func (oq *OrdderQueue) Len() int {
	return len(oq.Orders)
}

func (oq *OrdderQueue) Push(x interface{}) {
	oq.Orders = append(oq.Orders, x.(*Order))
}

func (oq *OrdderQueue) Pop() interface{} {
	old := oq.Orders
	n := len(old)
	item := old[n-1]
	oq.Orders = old[0 : n-1]
	return item
}

func NewOrderQueue() *OrdderQueue {
	return &OrdderQueue{}
}
