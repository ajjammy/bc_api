package reqs

import "time"

type SaleorderReq struct {

	docno string
	docdate string
	taxtype int
	billtype int
	arcode string
	departcode string
	creditday int
	duedate string
	salecode string
	taxrate float32
	isconfirm int
	mydescription string
	billstatus int
	sostatus int
	holdingstatus int
	sumofitemamount float32
	discountword string
	discontamount float32
	afterdiscount float32
	beforetaxamount float32
	taxamount float32
	totalamount float32
	netamount float32
	iscancel int
	creatorcode string
	createdatetime string
	lasteditorcode string
	lasteditdatet string  // must to convert to datetime type in sql server
	confirmcode string
	confirmdatetime string
	cancelcode string
	canceldatetime string
	isconditionsend int
	deliveryday int
	deliverydate time.Time
	items []*Sosub
}

type Sosub struct {
	docno string
	taxtype int
	itemcode string
	docdate string
	arcode string
	departcode string
	salecode string
	mydescription string
	itemname string
	whcode string
	shelfcode string
	qty float32
	remainqty float32
	price float32
	discountword string
	discountamount float32
	amount float32
	netamount float32
	homeamount float32
	unitcode string
	iscancel int
	linenumber int
	categorycode string
	groupcode string
	brandcode string
	typecode string
	formatcode string
	barcode string
	taxrate float32
	packingrate1 float32
	packingrate2 float32
}