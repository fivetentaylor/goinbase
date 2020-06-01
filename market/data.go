package goinbase

// https://docs.pro.coinbase.com/#orders
type Order struct {
	ClientOid   string `json:"client_oid,omitempty"`
	Type        string `json:"type,omitempty"`
	Side        string `json:"side,omitempty"`
	ProductId   string `json:"product_id,omitempty"`
	STP         string `json:"stp,omitempty"`
	Stop        string `json:"stop,omitempty"`
	StopPrice   string `json:"stop_price,omitempty"`
	Price       string `json:"price,omitempty"`
	Size        string `json:"size,omitempty"`
	TimeInForce string `json:"time_in_force,omitempty"`
	CancelAfter string `json:"cancel_after,omitempty"`
	PostOnly    string `json:"post_only,omitempty"`
	Funds       string `json:"funds,omitempty"`
}

type Account struct {
	Id        string `json:"id,omitempty"`
	Balance   string `json:"balance,omitempty"`
	Holds     string `json:"holds,omitempty"`
	Available string `json:"available,omitempty"`
	Currency  string `json:"currency,omitempty"`
}

type Market struct {
	Accounts map[string]Account `json: accounts,omitempty"`
}

type Product struct {
	Id             string `json:"id,omitempty"`
	BaseCurrency   string `json:"base_currency,omitempty"`
	QuoteCurrency  string `json:"quote_currency,omitempty"`
	BaseMinSize    string `json:"base_min_size,omitempty"`
	BaseMaxSize    string `json:"base_max_size,omitempty"`
	QuoteIncrement string `json:"quote_increment,omitempty"`
}

type BookItemLevel2 struct {
	Price     string `json:"price,omitempty"`
	Size      string `json:"size,omitempty"`
	NumOrders int    `json:"num_orders,omitempty"`
}

type BookLevel2 struct {
	Sequence int              `json:"sequence,omitempty"`
	Asks     []BookItemLevel2 `json:"asks"`
	Bids     []BookItemLevel2 `json:"bids"`
}

type BookItemLevel3 struct {
	Price   string `json:"price,omitempty"`
	Size    string `json:"size,omitempty"`
	OrderId string `json:"order_id,omitempty"`
}

type BookLevel3 struct {
	Sequence int              `json:"sequence,omitempty"`
	Asks     []BookItemLevel3 `json:"asks"`
	Bids     []BookItemLevel3 `json:"bids"`
}

type Currency struct {
	Id      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	MinSize string `json:"min_size,omitempty"`
}

type Time struct {
	Iso   string `json:"iso,omitempty"`
	Epoch string `json:"epoch,omitempty"`
}

type Received struct {
	Type      string `json:"type,omitempty"`
	Time      string `json:"time,omitempty"`
	ProductId string `json:"product_id,omitempty"`
	Sequence  int    `json:"sequence,omitempty"`
	OrderId   string `json:"order_id,omitempty"`
	Funds     string `json:"funds,omitempty"`
	Size      string `json:"size,omitempty"`
	Price     string `json:"price,omitempty"`
	Side      string `json:"side,omitempty"`
	OrderType string `json:"order_type,omitempty"`
}

type Open struct {
	Type          string `json:"type,omitempty"`
	Time          string `json:"time,omitempty"`
	ProductId     string `json:"product_id,omitempty"`
	Sequence      int    `json:"sequence,omitempty"`
	OrderId       string `json:"order_id,omitempty"`
	Price         string `json:"price,omitempty"`
	RemainingSize string `json:"remaining_size,omitempty"`
	Side          string `json:"side,omitempty"`
}

type Done struct {
	Type          string `json:"type,omitempty"`
	Time          string `json:"time,omitempty"`
	ProductId     string `json:"product_id,omitempty"`
	Sequence      int    `json:"sequence,omitempty"`
	OrderId       string `json:"order_id,omitempty"`
	Price         string `json:"price,omitempty"`
	Reason        string `json:"reason,omitempty"`
	RemainingSize string `json:"remaining_size,omitempty"`
	Side          string `json:"side,omitempty"`
}

type Match struct {
	Type         string `json:"type,omitempty"`
	Time         string `json:"time,omitempty"`
	ProductId    string `json:"product_id,omitempty"`
	Sequence     int    `json:"sequence,omitempty"`
	TradeId      int    `json:"trade_id,omitempty"`
	Price        string `json:"price,omitempty"`
	Side         string `json:"side,omitempty"`
	Size         string `json:"size,omitempty"`
	MakerOrderId string `json:"maker_order_id,omitempty"`
	TakerOrderId string `json:"taker_order_id,omitempty"`
}

type Change struct {
	Type      string `json:"type,omitempty"`
	Time      string `json:"time,omitempty"`
	ProductId string `json:"product_id,omitempty"`
	Sequence  int    `json:"sequence,omitempty"`
	OrderId   string `json:"order_id,omitempty"`
	Price     string `json:"price,omitempty"`
	Side      string `json:"side,omitempty"`
	NewSize   string `json:"new_size,omitempty"`
	OldSize   string `json:"old_size,omitempty"`
	NewFunds  string `json:"new_funds,omitempty"`
	OldFunds  string `json:"old_funds,omitempty"`
}

type Activate struct {
	Type      string `json:"type,omitempty"`
	Timestamp string `json:"time,omitempty"`
	ProductId string `json:"product_id,omitempty"`
	Sequence  int    `json:"sequence,omitempty"`
	OrderId   int    `json:"order_id,omitempty"`
	StopPrice string `json:"stop_price,omitempty"`
	StopType  string `json:"stop_type,omitempty"`
	Side      string `json:"side,omitempty"`
	UserId    string `json:"user_id,omitempty"`
	ProfileId string `json:"profile_id,omitempty"`
	Size      string `json:"size,omitempty"`
	Funds     string `json:"funds,omitempty"`
	Private   bool   `json:"funds,omitempty"`
}
