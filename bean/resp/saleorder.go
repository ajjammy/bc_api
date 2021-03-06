package Resp

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	//	"time"
	//	"github.com/gin-gonic/gin"
)

type Saleorder struct {
	Docno           string `json:"doc_no" db:"docno"`
	Docdate         string `json:"doc_date" db:"docdate"`
	Taxtype         int `json:"tax_type" db :"taxtype"`
	Billtype        int `json:"bill_type" db:"billtype"`
	Arcode          string `json:"ar_code" db:"arcode"`
	Departcode      string `json:"depart_code" db:"departcode"`
	Creditday       int `json:"credit_day" db:"creditday"`
	Duedate         string `json:"due_date" db:"duedate"`
	Salecode        string `json:"sale_code" db:"salecode"`
	Taxrate         float32 `json:"tax_rate" db:"taxrate"`
	Isconfirm       int `json:"is_confirm" db:"isconfirm"`
	Mydescription   string `json:"my_description" db:"mydescription"`
	Billstatus      int `json:"bill_status" db:"billstatus"`
	Sostatus        int `json:"so_status" db:"sostatus"`
	Holdingstatus   int `json:"holding_status" db:"holdingstatus"`
	Sumofitemamount float32 `json:"sum_of_item_amount" db:"sumofitemamount"`
	Discountword    string `json:"discount_word"`
	Discountamount  float32 `json:"discount_amount" DB:"discountamount" `
	Afterdiscount   float32 `json:"after_discount"`
	Beforetaxamount float32 `json:"before_tax_amount"`
	Taxamount       float32 `json:"tax_amount"`
	Totalamount     float32 `json:"total_amount"`
	Netamount       float32 `json:"net_amount"`
	Iscancel        int `json:"is_cancel"`
	Creatorcode     string `json:"creator_code"`
	Createdatetime  string `json:"create_datetime"`
	Lasteditorcode  string `json:"last_editor_code"`
	Lasteditdatet   string  `json:"last_edit_datet"` // must to convert to datetime type in sql server
	Confirmcode     string `json:"confirm_code"`
	Confirmdatetime string `json:"confirm_datetime"`
	Cancelcode      string `json:"cancel_code"`
	Canceldatetime  string `json:"cancel_datetime"`
	Isconditionsend int `json:"is_condition_send"`
	Deliveryday     int `json:"delivery_day"`
	Deliverydate    string `json:"delivery_date"`
						       //items []*Sosub

	Items           []*Saleordersub `json:"item"`
}

type Saleordersub struct {
	Id             int64  `json:"id"`
	Docno          string `json:"doc_no"`
	Taxtype        int `json:"tax_type"`
	Itemcode       string `json:"item_code"`
	Docdate        string `json:"docdate" db:"docdate"`
	Arcode         string `json:"ar_code"`
	Departcode     string `json:"depart_code"`
	Salecode       string `json:"sale_code"`
	Mydescription  string `json:"my_description"`
	Itemname       string `json:"item_name"`
	Whcode         string `json:"wh_code"`
	Shelfcode      string `json:"shelf_code"`
	Qty            float32 `json:"qty"`
	Remainqty      float32 `json:"remain_qty"`
	Price          float32 `json:"price"`
	Discountword   string `json:"discount_word"`
	Discountamount float32 `json:"discount amount"`
	Amount         float32 `json:"amount"`
	Netamount      float32 `json:"net_amount"`
	Homeamount     float32 `json:"home_amount"`
	Unitcode       string `json:"unit_code"`
	Iscancel       int `json:"is_cancel"`
	Linenumber     int `json:"line_number"`
	Categorycode   string `json:"category_code"`
	Groupcode      string `json:"group_code"`
	Brandcode      string `json:"brand_code"`
	Typecode       string `json:"type_code"`
	Formatcode     string `json:"format_code"`
	Barcode        string `json:"barcode"`
	Taxrate        float32 `json:"tax_rate"`
	Packingrate1   float32 `json:"packing_rate_1"`
	Packingrate2   float32 `json:"packing_rate_2"`
}

func (s *Saleorder)GetByDocno(docno string, db *sqlx.DB) (err error) {

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
		"from dbo.bcsaleorder where docno = '" + docno + "'"
	//fmt.Println(lcCommand)
	// Get saleorder from Database by docno
	//ss = []Saleorder{}
	err = db.Get(s, lcCommand)
	if err != nil {
		return err
	}
	//fmt.Println(s)
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
		from dbo.bcsaleordersub
			where docno=? `
	//fmt.Println(sosub)
	err = db.Select(&s.Items, sosub, docno)
	return err
}

func (s *Saleorder)GetByKeyWord(keyword string, db *sqlx.DB) (ss []Saleorder, err error) {
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
		"from dbo.bcsaleorder where docno like '%" + keyword + "%'"
	fmt.Println(lcCommand)
	// Get saleorder from Database by docno
	//ss = []Saleorder{}
	err = db.Select(&ss, lcCommand)
	if err != nil {
		return nil, err
	}
	// todo : add child node in for loop
	for i, so := range ss {

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
		from dbo.bcsaleordersub
			where docno=? `
		fmt.Println(sosub)
		err = db.Select(&ss[i].Items, sosub, so.Docno)

	}

	//fmt.Println(ss)
	return ss, nil
}

func (s *Saleorder)Insert(db *sqlx.DB) (NewSoNumber string, err error) {

	lccommand := `insert into dbo.bcsaleorder (
		docno,docdate,taxtype,billtype,arcode,
		departcode,creditday,duedate,salecode,taxrate,
		isconfirm,mydescription,billstatus,sostatus,holdingstatus,
		sumofitemamount,discountword,discountamount,afterdiscount,beforetaxamount,
		taxamount,totalamount,netamount,iscancel,creatorcode,
		createdatetime,lasteditorcode,lasteditdatet,confirmcode,confirmdatetime,
		cancelcode,canceldatetime,isconditionsend,deliveryday,deliverydate)
		values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
	_, err = db.Exec(lccommand,
		s.Docno, s.Docdate, s.Taxtype, s.Billtype, s.Arcode,
		s.Departcode, s.Creditday, s.Duedate, s.Salecode, s.Taxrate,
		s.Isconfirm, s.Mydescription, s.Billstatus, s.Sostatus, s.Holdingstatus,
		s.Sumofitemamount, s.Discountword, s.Discountamount, s.Afterdiscount, s.Beforetaxamount,
		s.Taxamount, s.Totalamount, s.Netamount, s.Iscancel, s.Creatorcode,
		s.Createdatetime, s.Lasteditorcode, s.Lasteditdatet, s.Confirmcode, s.Confirmdatetime,
		s.Cancelcode, s.Canceldatetime, s.Isconditionsend, s.Deliveryday, s.Deliverydate)

	//fmt.Println(lccommand)
	if err != nil {
		return s.Docno, err
	}

	// todo : insert sub
	err = s.InsertSub(s.Items, db)
	if err != nil {
		fmt.Println(err.Error)
		return s.Docno, err
	}
	return s.Docno, err
}

func (s *Saleorder)InsertSub(sb []*Saleordersub, db *sqlx.DB) (err error) {
	for _, k := range sb {

		lccommand := `insert into dbo.bcsaleordersub (
	docno,taxtype,itemcode,docdate,arcode,
	departcode,salecode,mydescription,itemname,whcode,
	shelfcode,qty,remainqty,price,discountword,
	discountamount,amount,netamount,homeamount,unitcode,
	iscancel,linenumber,categorycode,groupcode,brandcode,
	typecode,formatcode,barcode,taxrate,packingrate1,
	packingrate2
	) values (
	?,?,?,?,?,
	?,?,?,?,?,
	?,?,?,?,?,
	?,?,?,?,?,
	?,?,?,?,?,
	?,?,?,?,?,
	? )`

		_, err := db.Exec(lccommand,
			k.Docno, k.Taxtype, k.Itemcode, k.Docdate, k.Arcode,
			k.Departcode, k.Salecode, k.Mydescription, k.Itemname, k.Whcode,
			k.Shelfcode, k.Qty, k.Remainqty, k.Price, k.Discountamount,
			k.Discountamount, k.Amount, k.Netamount, k.Homeamount, k.Unitcode,
			k.Iscancel, k.Linenumber, k.Categorycode, k.Groupcode, k.Brandcode,
			k.Typecode, k.Formatcode, k.Barcode, k.Taxrate, k.Packingrate1,
			k.Packingrate2)

		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		//fmt.Println(lccommand)
	}

	return err
}

func(s *Saleorder)CheckExists(db *sqlx.DB, docno string) (bool) {
	fmt.Println("Begin CheckExists")
	lccommand := "select top 1 docno from dbo.bcsaleorder where docno = '" + docno + "'"
	rs, _ := db.Exec(lccommand)
	chkRow, _ := rs.RowsAffected()
	if chkRow > 0 {
		return true
		fmt.Println("data aleady exists!!! cannot insert this number : ", docno)
	}
	return false

}

func(s *Saleorder)Delete(db *sqlx.DB, docno string) (err error) {
	//todo : Delete Before Update Saleorder
	//todo : saleorder
	fmt.Println("begin Saleorder.Delete")
	lccommand := "delete from dbo.bcsaleorder where docno = '" + docno + "'"
	rs, _ := db.Exec(lccommand)
	_, err = rs.RowsAffected()
	if err != nil {
		return err
	}

	//todo : saleordersub
	lccommand = "delete from dbo.bcsaleordersub where docno = '" + docno + "'"
	_, err = db.Exec(lccommand)

	_, err = rs.RowsAffected()
	if err != nil {
		return err
	}

	return nil

}

func(s *Saleorder)Void(db *sqlx.DB, docno string, cancelcode string) (err error) {
	//todo : Delete Before Update Saleorder
	//todo : saleorder
	fmt.Println("begin Saleorder.Delete")
	lccommand := "update dbo.bcsaleorder set iscancel = 1,cancelcode=" + cancelcode + ",canceldatetime=getdate() where docno = '" + docno + "'"
	rs, _ := db.Exec(lccommand)
	_, err = rs.RowsAffected()
	if err != nil {
		return err
	}

	//todo : saleordersub
	lccommand = "Update dbo.bcsaleordersub set iscancel = 1 where docno = '" + docno + "'"
	_, err = db.Exec(lccommand)

	_, err = rs.RowsAffected()
	if err != nil {
		return err
	}

	return nil

}

func(s *Saleorder)Update(db *sqlx.DB) (msg string, err error) {
	// Update bcsaleorder

	// check so status - Update can use only new,onprocess only
	// status : done , cancel cannot update
	if s.Billstatus == 1 || s.Iscancel == 1 {
		msg = "Cannot update status (CANCEL,DONE) "
		return msg, err
	}

	lcCommand := "update dbo.bcsaleorder set " +
		"docno = ?,docdate=?,taxtype=?,billtype=?,arcode=?," +
		"departcode=?,creditday=?,duedate=?,salecode=?,taxrate=?," +
		"isconfirm=?,mydescription=?,billstatus=?," +
		"sostatus=?,holdingstatus=?,sumofitemamount=?,discountword=?,discountamount=?," +
		"afterdiscount=?,beforetaxamount=?,taxamount=?,totalamount=?,netamount=?," +
		"lasteditorcode = ?lasteditdatet = getdate(),isconditionsend=?,deliveryday=?,deliverydate=?" +
		"where docno = ?"
	_, err = db.Exec(lcCommand, s.Docno, s.Docdate, s.Taxtype, s.Billtype, s.Arcode,
		s.Departcode, s.Creditday, s.Duedate, s.Salecode, s.Taxrate,
		s.Isconfirm, s.Mydescription, s.Billstatus,
		s.Sostatus, s.Holdingstatus, s.Sumofitemamount, s.Discountword, s.Discountamount,
		s.Afterdiscount, s.Beforetaxamount, s.Taxamount, s.Totalamount, s.Netamount,
		s.Lasteditorcode, s.Isconditionsend, s.Deliveryday, s.Deliverydate, s.Docno)
	if err != nil {
		msg = "Update Error "
		return msg, err
	}

	// update bcsaleordersub --> use delete before insert

	lcCommand = "delete from dbo.bcsaleordersub where docno = '" + s.Docno + "'"
	_, err = db.Exec(lcCommand)
	if err != nil {
		msg = "delete Detail of order error ! "
		return msg, err
	}

	// todo : insert bcsaleordersub
	// todo : insert sub
	err = s.InsertSub(s.Items, db)
	if err != nil {
		fmt.Println(err.Error)
		msg = "insert Detail of order Error "
		return s.Docno, err
	}

	msg = "Completed updated"
	return msg, err
}