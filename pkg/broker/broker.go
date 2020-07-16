package broker
type Broker interface(
	Buy()
	Sell()
	SellShort()
	BuyToCover()
)