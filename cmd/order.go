package cmd

import (
	"github.com/xerexchain/matching-engine/cmd/result_code"
	"github.com/xerexchain/matching-engine/order"
	"github.com/xerexchain/matching-engine/orderbook"
	"github.com/xerexchain/matching-engine/orderbook/event"
)

type OrderCommand interface {
}

type orderCommand struct {
	t         Type
	uid       int64
	orderId   int64
	orderType order.Type

	// required for PLACE_ORDER only;
	// for CANCEL/MOVE contains original order action (filled by orderbook)
	orderAction order.Action

	symbol int32
	price  int64
	size   int64

	// new orders INPUT - reserved price for fast moves of GTC bid orders in exchange mode
	reserveBidPrice int64

	timestamp  int64
	userCookie int32

	// filled by grouping processor:
	eventsGroup  int64
	serviceFlags int32

	// can also be used for saving intermediate state
	resultCode resultcode.ResultCode

	tradeEvent event.TradeEvent

	marketData orderbook.L2MarketData
	_          struct{}
}

// No removing/revoking
func (o *orderCommand) ProcessMatcherEvents(ch chan<- event.TradeEvent) {
	eve := o.tradeEvent

	for eve != nil {
		ch <- eve
		eve = eve.Next()
	}
}

func NewOrderCommand() OrderCommand {
	return &orderCommand{}
}
