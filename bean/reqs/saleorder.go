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
	Docno string
	Taxtype int
	Itemcode string
	Docdate string
	Arcode string
	Departcode string
	Salecode string
	Mydescription string
	Itemname string
	Whcode string
	Shelfcode string
	Qty float32
	Remainqty float32
	Price float32
	Discountword string
	Discountamount float32
	Amount float32
	Netamount float32
	Homeamount float32
	Unitcode string
	Iscancel int
	Linenumber int
	Categorycode string
	Groupcode string
	Brandcode string
	Typecode string
	Formatcode string
	Barcode string
	Taxrate float32
	Packingrate1 float32
	Packingrate2 float32
}