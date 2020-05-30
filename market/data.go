package goinbase

// https://docs.pro.coinbase.com/#orders
type Order struct {
	ClientOid   string `json:"client_oid"`
	Type        string `json:"type"`
	Side        string `json:"side"`
	ProductId   string `json:"product_id"`
	STP         string `json:"stp"`
	Stop        string `json:"stop"`
	StopPrice   string `json:"stop_price"`
	Price       string `json:"price"`
	Size        string `json:"size"`
	TimeInForce string `json:"time_in_force"`
	CancelAfter string `json:"cancel_after"`
	PostOnly    string `json:"post_only"`
	Funds       string `json:"funds"`
}

type Account struct {
	Id        string `json:"id"`
	Balance   string `json:"balance"`
	Holds     string `json:"holds"`
	Available string `json:"available"`
	Currency  string `json:"currency"`
}

type Market struct {
	Accounts map[string]Account `json: accounts"`
}

type Product struct {
	Id             string `json:"id"`
	BaseCurrency   string `json:"base_currency"`
	QuoteCurrency  string `json:"quote_currency"`
	BaseMinSize    string `json:"base_min_size"`
	BaseMaxSize    string `json:"base_max_size"`
	QuoteIncrement string `json:"quote_increment"`
}

type Currency struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	MinSize string `json:"min_size"`
}

type Time struct {
	Iso   string `json:"iso"`
	Epoch string `json:"epoch"`
}

type Received struct {
	Type      string `json:"type"`
	Time      string `json:"time"`
	ProductId string `json:"product_id"`
	Sequence  int    `json:"sequence"`
	OrderId   string `json:"order_id"`
	Size      string `json:"size"`
	Price     string `json:"price"`
	Side      string `json:"side"`
	OrderType string `json:"order_type"`
}
