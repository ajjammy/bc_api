package Resp

import (
	"github.com/jmoiron/sqlx"
	"fmt"
//	"time"
//	"github.com/gin-gonic/gin"
)

type Saleorder struct {
	//Docno string`json:"doc_no"`
	//Arcode string `json:"ar_code"`
	//SumOfItemAmount float32 `json:"sum_of_item_amount"`
	//DiscountAmount float32 `json:"discount_amount"`
	//BeforeTaxAmount float32 `json:"before_tax_amount"`
	//TaxAmount float32 `json:"tax_amount"`
	//TotalAmount float32 `json:"total_amount"`
	Docno string `json:"doc_no"`
	Docdate string `json:"docdate"`
	Taxtype int `json:"taxtype"`
	Billtype int `json:"billtype"`
	Arcode string `json:"arcode"`
	Departcode string `json:"departcode"`
	Creditday int `json:"creditday"`
	Duedate string `json:"duedate"`
	Salecode string `json:"salecode"`
	Taxrate float32 `json:"taxrate"`
	Isconfirm int `json:"isconfirm"`
	Mydescription string `json:"mydescription"`
	Billstatus int `json:"billstatus"`
	Sostatus int `json:"so_status"`
	Holdingstatus int `json:"holding_status"`
	Sumofitemamount float32 `json:"sumofitemamount"`
	Discountword string `json:"discountword"`
	Discountamount float32 `json:"discount_amount" DB:"discountamount" `
	Afterdiscount float32 `json:"afterdiscount"`
	Beforetaxamount float32 `json:"beforetaxamount"`
	Taxamount float32 `json:"taxamount"`
	Totalamount float32 `json:"totalamount"`
	Netamount float32 `json:"netamount"`
	Iscancel int `json:"iscancel"`
	Creatorcode string `json:"creatorcode"`
	Createdatetime string `json:"createdatetime"`
	Lasteditorcode string `json:"lasteditorcode"`
	Lasteditdatet string  `json:"lasteditdatet"`// must to convert to datetime type in sql server
	Confirmcode string `json:"confirmcode"`
	Confirmdatetime string `json:"confirmdatetime"`
	Cancelcode string `json:"cancelcode"`
	Canceldatetime string `json:"canceldatetime"`
	Isconditionsend int `json:"iscondition_send"`
	Deliveryday int `json:"delivery_day"`
	Deliverydate string `json:"delivery_date"`
	//items []*Sosub

	Items  []*Saleordersub `json:"item"`
}

type Saleordersub struct {
	Id int64  `json:"id"`//roworder
	//Linenumber int64  `json:"line_number"`
	//Itemcode string `json:"item_code"`
	//Itemname string `json:"item_name"`
	//Qty float32`json:"qty"`
	//Unitcode string `json:"unit_code"`
	//Price  float32 `json:"price"`
	//Amount float32 `json:"amounts"`
	//Netamount float32 `json:"net_amount"`
	//Packingrate1 float32 `json:"packing_rate_1"`
	//Packingrate2 float32 `json:"packing_rate_2"`

	Docno string `json:"doc_no"`
	Taxtype int `json:"tax_type"`
	Itemcode string `json:"item_code"`
	Docdate string `json:"docdate" db:"docdate"`
	Arcode string `json:"ar_code"`
	Departcode string `json:"depart_code"`
	Salecode string `json:"sale_code"`
	Mydescription string `json:"my_description"`
	Itemname string `json:"item_name"`
	Whcode string `json:"wh_code"`
	Shelfcode string `json:"shelf_code"`
	Qty float32 `json:"qty"`
	Remainqty float32 `json:"remain_qty"`
	Price float32 `json:"price"`
	Discountword string `json:"discount_word"`
	Discountamount float32 `json:"discount amount"`
	Amount float32 `json:"amount"`
	Netamount float32 `json:"netamount"`
	Homeamount float32 `json:"home_amount"`
	Unitcode string `json:"unit_code"`
	Iscancel int `json:"is_cancel"`
	Linenumber int `json:"line_number"`
	Categorycode string `json:"category_code"`
	Groupcode string `json:"group_code"`
	Brandcode string `json:"brand_code"`
	Typecode string `json:"type_code"`
	Formatcode string `json:"format_code"`
	Barcode string `json:"barcode"`
	Taxrate float32 `json:"tax_rate"`
	Packingrate1 float32 `json:"packing_rate_1"`
	Packingrate2 float32 `json:"packing_rate_2"`
	
}

func(s *Saleorder)GetByDocno(docno string,db *sqlx.DB)(err error){

	lcCommand := "select docno,docdate,taxtype,billtype," +
		"arcode,departcode,creditday,duedate,salecode,taxrate," +
		"isconfirm," +
		"isnull(mydescription,'') as mydescription ," +
		"billstatus," +
		"sostatus," +
		"holdingstatus," +
		"sumofitemamount," +
		"isnull(discountword,'') as discountword," +
		"isnull(discountamount,0) as discountamount," +
		"isnull(afterdiscount,0)as afterdiscount," +
		"isnull(beforetaxamount,0)as beforetaxamount," +
		"isnull(taxamount,0) as taxamount," +
		"isnull(totalamount,0) as totalamount," +
		"isnull(netamount,0) as netamount," +
		"isnull(iscancel,0)as iscancel," +
		"isnull(creatorcode,'') as creatorcode," +
		"isnull(createdatetime,'') as createdatetime," +
		"isnull(lasteditorcode,'') as lasteditorcode," +
		"isnull(lasteditdatet,'') as lasteditdatet," +
		"isnull(confirmcode,'') as confirmcode," +
		"isnull(confirmdatetime,'') as confirmdatetime," +
		"isnull(cancelcode,'') as cancelcode," +
		"isnull(canceldatetime,'') as canceldatetime," +
		"isnull(isconditionsend,0) as isconditionsend," +
		"isnull(deliveryday,0)as deliveryday," +
		"isnull(deliverydate,getdate()) as deliverydate " +
		"from bcnp.dbo.bcsaleorder where docno = '"+docno+"'"
	fmt.Println(lcCommand)
	// Get saleorder from Database by docno
	//ss = []Saleorder{}
	err = db.Get(s,lcCommand)
	if err !=nil{
		return err
	}
	fmt.Println(s)
	// todo: add Node sub details

	sosub := `select  roworder as id,
			isnull(docno,'') as docno,
			isnull(taxtype,0) as taxtype,
			isnull(itemcode,'') as itemcode,
			isnull(docdate,'') as docdate,
			isnull(arcode,'') as arcode,
			isnull(departcode,'') as departcode,
			isnull(salecode,'') as salecode,
			isnull(mydescription,'') as mydescription,
			isnull(itemname,'') as itemname,
			isnull(whcode,'') as whcode,
			isnull(shelfcode,'')as shelfcode,
			isnull(qty,0) as qty,
			isnull(remainqty,0) as remainqty,
			isnull(price,0) as price,
			isnull(discountword,'') as discountword,
			isnull(discountamount,0) as discountamount,
			isnull(amount,0) as amount,
			isnull(netamount,0) as netamount,
			isnull(homeamount,0) as homeamount,
			isnull(unitcode,'') as unitcode,
			isnull(iscancel,0)as iscancel,
			isnull(linenumber,0) as linenumber,
			isnull(categorycode,'') as categorycode,
			isnull(groupcode,'') as groupcode,
			isnull(brandcode,'') as brandcode,
			isnull(typecode,'') as typecode,
			isnull(formatcode,'') as formatcode,
			isnull(barcode,'') as barcode,
			isnull(taxrate,0) as taxrate,
			isnull(packingrate1,1) as packingrate1,
			isnull(packingrate2,1) as packingrate2
		from bcnp.dbo.bcsaleordersub
			where docno=? `
	fmt.Println(sosub)
	err = db.Select(&s.Items,sosub,docno)
	return err
}

func(s *Saleorder)GetByKeyWord(keyword string,db *sqlx.DB)(ss []Saleorder,err error){
	lcCommand := "select top 20 " +
		"isnull(docno,'') as docno," +
		"isnull(docdate,'') as docdate," +
		"isnull(taxtype,0) as taxtype," +
		"isnull(billtype,0) as billtype," +
		"isnull(arcode,'') as arcode," +
		"isnull(departcode,'') as departcode," +
		"isnull(creditday,0) as creditday," +
		"isnull(duedate,'') as duedate," +
		"isnull(salecode,'') as salecode,taxrate," +
		"isconfirm," +
		"isnull(mydescription,'') as mydescription ," +
		"billstatus," +
		"sostatus," +
		"holdingstatus," +
		"sumofitemamount," +
		"isnull(discountword,'') as discountword," +
		"isnull(discountamount,0) as discountamount," +
		"isnull(afterdiscount,0)as afterdiscount," +
		"isnull(beforetaxamount,0)as beforetaxamount," +
		"isnull(taxamount,0) as taxamount," +
		"isnull(totalamount,0) as totalamount," +
		"isnull(netamount,0) as netamount," +
		"isnull(iscancel,0)as iscancel," +
		"isnull(creatorcode,'') as creatorcode," +
		"isnull(createdatetime,'') as createdatetime," +
		"isnull(lasteditorcode,'') as lasteditorcode," +
		"isnull(lasteditdatet,'') as lasteditdatet," +
		"isnull(confirmcode,'') as confirmcode," +
		"isnull(confirmdatetime,'') as confirmdatetime," +
		"isnull(cancelcode,'') as cancelcode," +
		"isnull(canceldatetime,'') as canceldatetime," +
		"isnull(isconditionsend,0) as isconditionsend," +
		"isnull(deliveryday,0)as deliveryday," +
		"isnull(deliverydate,'') as deliverydate " +
		"from bcnp.dbo.bcsaleorder where docno like '%"+keyword+"%'"
	fmt.Println(lcCommand)
	// Get saleorder from Database by docno
	//ss = []Saleorder{}
	err = db.Select(&ss,lcCommand)
	if err !=nil{
		return nil,err
	}
	// todo : add child node in for loop
	for i,so := range ss{

		sosub := `select  roworder as id,
			docno,taxtype,itemcode,docdate,arcode,
			isnull(departcode,'') as departcode,
			isnull(salecode,'') as salecode,
			isnull(mydescription,'') as mydescription,
			isnull(itemname,'') as itemname,
			isnull(whcode,'') as whcode,
			isnull(shelfcode,'')as shelfcode,
			isnull(qty,0) as qty,
			isnull(remainqty,0) as remainqty,
			isnull(price,0) as price,
			isnull(discountword,'') as discountword,
			isnull(discountamount,0) as discountamount,
			isnull(amount,0) as amount,
			isnull(netamount,0) as netamount,
			isnull(homeamount,0) as homeamount,
			isnull(unitcode,'') as unitcode,
			isnull(iscancel,0)as iscancel,
			isnull(linenumber,0) as linenumber,
			isnull(categorycode,'') as categorycode,
			isnull(groupcode,'') as groupcode,
			isnull(brandcode,'') as brandcode,
			isnull(typecode,'') as typecode,
			isnull(formatcode,'') as formatcode,
			isnull(barcode,'') as barcode,
			isnull(taxrate,0) as taxrate,
			isnull(packingrate1,1) as packingrate1,
			isnull(packingrate2,1) as packingrate2
		from bcnp.dbo.bcsaleordersub
			where docno=? `
		fmt.Println(sosub)
		err = db.Select(&ss[i].Items,sosub,so.Docno)

	}

	//fmt.Println(ss)
	return ss,nil
}

