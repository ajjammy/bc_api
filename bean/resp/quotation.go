package Resp

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	//	"debug/dwarf"
	//	"go/doc"
)

type Qt struct {
	DocNo               string `json:"doc_no" db:"DocNo"`
	DocDate             string `json:"doc_date" db:"DocDate"`
	ArCode              string `json:"ar_code" db:"ArCode"`
	SaleCode            string `json:"sale_code"  db:"SaleCode"`
	BillType            int `json:"bill_type"`
	BillStatus          int `json:"bill_status" db:"BillStatus"`
	TaxType             int `json:"tax_type" db:"TaxType"`
	TaxRate             float64 `json:"tax_rate" db:"TaxRate"`
	SumOfItemAmount     float64 `json:"sum_of_item_amount" db:"sumofitemamount"`
	DisCountWord        string `json:"dis_count_word" db:"DisCountWord"`
	DisCountAmount      float64 `json:"dis_count_amount" db:"DisCountAmount"`
	AfterDiscountAmount float64 `json:"after_discount_amount" db:"AfterDiscountAmount"`
	BeforeTaxAmount     float64 `json:"before_tax_amount" db:"BeforeTaxAmount"`
	TaxAmount           float64 `json:"tax_amount" db:"TaxAmount"`
	TotalAmount         float64 `json:"total_amount" db:"TotalAmount"`
	NetAmount           float64 `json:"net_amount" db:"NetAmount"`
	CreditDay           int `json:"credit_day" db:"creditday"`
	DueDate             string `json:"due_date" db:"duedate"`
	Validity            int `json:"validity" db:"Validity"`
	AssertStatus        int `json:"assert_status" db:"AssertStatus"`
	MyDescription1      string `json:"my_description" db:"MyDescription1"`
	IsCancel            int `json:"is_cancel" db:"IsCancel"`
	Subs                []*QtSub `json:"subs"`
}

type QtSub struct {
	DocNo     string `json:"doc_no"`
	ItemCode  string `json:"item_code"`
	WhCode    string `json:"wh_code"`
	Qty       int64 `json:"qty"`
	UnitCode  string `json:"unit_code"`
	RemainQty float64 `json:"remain_qty" db:"RemainQty"`
	ItemName  string `json:"item_name"`
	Amount    float64 `json:"amount"`
}

type Quotation struct {
	Id                  string `json:"id" db:"roworder"`
	DocNo               string `json:"doc_no" db:"DocNo"`
	DocDate             string `json:"doc_date" db:"DocDate"`
	DueDate             string `json:"due_date" db:"DueDate"`
	DeliveryDate        string `json:"delivery_date" db:"DeliveryDate"`
	ArCode              string `json:"ar_code" db:"ArCode"`
	ArName              string `json:"ar_name" db:"ArName"`
	ArAddress           string `json:"ar_address" db:"ArAddress"`
	ArTelephone         string `json:"ar_telephone" db:"ArTelephone"`
	ArFax               string `json:"ar_fax" db:"ArFax"`
	SaleCode            string `json:"sale_code"  db:"SaleCode"`
	SaleName            string `json:"sale_name" db:"SaleName"`
	TaxRate             float64 `json:"tax_rate" db:"TaxRate"`
	TaxType             int `json:"tax_type" db:"TaxType"`
	MyDescription       string `json:"my_description" db:"MyDescription"`
	SumItemAmount       float32 `json:"sum_item_amount" db:"sumofitemamount"`
	DisCountWord        string `json:"dis_count_word" db:"DisCountWord"`
	DisCountAmount      float64 `json:"dis_count_amount" db:"DisCountAmount"`
	AfterDiscountAmount float64 `json:"after_discount_amount" db:"AfterDiscountAmount"`
	BeforeTaxAmount     float64 `json:"before_tax_amount" db:"BeforeTaxAmount"`
	TaxAmount           float64 `json:"tax_amount" db:"TaxAmount"`
	TotalAmount         float64 `json:"total_amount" db:"TotalAmount"`
	Iscancel            int `json:"is_cancel" db:"Iscancel"`
	IsConfirm           int `json:"is_confirm" db:"IsConfirm"`
	BillStatus          int `json:"bill_status" db:"BillStatus"`
	BillType            int `json:"bill_type" db:"BillType"`
	CreditDay           int `json:"credit_day" db:"CreditDay"`
	ContactCode         string `json:"contact_code" db:"ContactCode"`
	ContactName         string `json:"contact_name" db:"ContactName"`
	ProjectCode         string `json:"project_code" db:"ProjectCode"`
	ProjectName         string `json:"project_name" db:"ProjectName"`
	AllocateCode        string `json:"allocate_code" db:"AllocateCode"`
	AllocateName        string `json:"allocate_name" db:"AllocateName"`
	CreatorCode         string `json:"creator_code" db:"CreatorCode"`
	CreateDateTime      string `json:"create_date_time" db:"CreateDateTime"`
	EditorCode          string `json:"editor_code" db:"EditorCode"`
	EditDateTime        string `json:"edit_date_time" db:"EditDateTime"`
	ConfirmCode         string `json:"confirm_code" db:"ConfirmCode"`
	ConfirmDataTime     string `json:"confirm_data_time"  db:"ConfirmDataTime"`
	CancelCode          string `json:"cancel_code" db:"CancelCode"`
	CancelDateTime      string `json:"cancel_date_time"  db:"CancelDateTime"`
	IsConditionSend     int `json:"is_condition_send"`
	DepartCode          string `json:"depart_code" db:"departcode"`
	DeliveryDay         int `json:"delivery_day"`
	//	RefNo 		string `json:"ref_no" db:"refno"`
	Subs []*QuotationSub `json:"subs"`
}

type QuotationSub struct {
	Id              int64  `json:"id" db:"roworder"`
	DocNo           string `json:"doc_no"`
	Taxtype         int `json:"taxtype"`
	ItemCode        string `json:"item_code" db:"ItemCode"`
	DocDate         string `json:"doc_date"`
	Arcode          string `json:"arcode"`
	Departcode      string `json:"departcode"`
	Salecode        string `json:"salecode"`
	Mydescription   string `json:"mydescription"`
	ItemName        string `json:"item_name" db:"ItemName"`
	Whcode          string `json:"wh_code"`
	Shelfcode       string `json:"shelf_code"`
	Qty             float64 `json:"qty" db:"Qty"`
	Remainqty       float32 `json:"remain_qty"`
	Price           float64 `json:"price" db:"Price"`
	DisCountWord    string `json:"dis_count_word_sub" db:"DisCountWord"`
	DisCountAmount  float64 `json:"dis_count_amount_sub" db:"DisCountAmount"`
	UnitCode        string `json:"unit_code" db:"UnitCode"`
	NetAmount       float64 `json:"net_amount" db:"NetAmount"`
	Amount          float64 `json:"amount" db:"Amount"`
	Homeamount      float32 `json:"home_amount"`
	ItemDescription string `json:"item_description" db:"ItemDescription"`
	IsConditionSend int `json:"is_condition_send" db:"IsConditionSend"`
	Iscancel        int `json:"is_cancel" db:"Iscancel"`
	LineNumber      int `json:"line_number" db:"LineNumber"`
	PackingRate     float64 `json:"packing_Rate" db:"PackingRate"`
	Categorycode    string `json:"category_code"`
	Groupcode       string `json:"group_code"`
	Brandcode       string `json:"brand_code"`
	Typecode        string `json:"type_code"`
	Formatcode      string `json:"format_code"`
	Barcode         string `json:"barcode"`
	Taxrate         float32 `json:"tax_rate"`
	Packingrate1    float32 `json:"packing_rate_1"`
	Packingrate2    float32 `json:"packing_rate_2"`
}

func (q *Quotation) GetByDocno(docno string, db *sqlx.DB) error {
	fmt.Println(q.DocNo)
	lcCommand := "select a.roworder , " +
		" isnull(a.docno,'') as DocNo," +
		" isnull(a.DocDate,'') as DocDate," +
		" isnull(a.DueDate,'') as DueDate," +
		" isnull(a.DeliveryDate,'') as DeliveryDate," +
		" isnull(a.arcode,'') as ArCode," +
		" isnull(b.name1,'') as ArName," +
		" isnull(b.billAddress,'') as ArAddress," +
		" isnull(b.Telephone,'') as ArTelephone," +
		" isnull(b.fax,'') as ArFax," +
		" isnull(a.salecode,'') as SaleCode," +
		" isnull(c.name,'') as SaleName," +
	//" ' ' as RefNo," +
		" isnull(a.taxrate,7) as TaxRate," +
		" a.TaxType," +
		" isnull(a.MyDescription1,'') as MyDescription," +
		" isnull(a.sumofitemamount,0) as sumofitemamount," +
		" isnull(a.discountword,'') as DisCountWord," +
		" isnull(a.discountamount,0) as DisCountAmount," +
		" isnull(a.AfterDiscount,0) as AfterDiscountAmount," +
		" isnull(a.BeforeTaxAmount,0) as BeforeTaxAmount," +
		" isnull(a.TaxAmount,0) as TaxAmount," +
		" isnull(a.TotalAmount,0) as TotalAmount," +
		" a.Iscancel," +
		" a.IsConfirm," +
		" a.BillStatus," +
		" a.BillType," +
		" isnull(a.CreditDay,0) as CreditDay," +
		" isnull(ContactCode,'') as ContactCode," +
		" isnull(d.name,'') as ContactName," +
		" isnull(a.ProjectCode,'') as ProjectCode," +
		" isnull(f.name,'') as ProjectName," +
		" isnull(a.AllocateCode,'') as AllocateCode," +
		" isnull(g.name,'') as AllocateName," +
		" isnull(a.CreatorCode,'') as CreatorCode," +
		" isnull(a.CreateDateTime,'') as CreateDateTime," +
		" isnull(a.lastEditorCode,'') as EditorCode," +
		" isnull(a.lasteditdatet,'') as EditDateTime," +
		" isnull(a.ConfirmCode,'') as ConfirmCode," +
		" isnull(a.ConfirmDateTime,'') as ConfirmDataTime," +
		" isnull(a.CancelCode,'') as  CancelCode," +
		" isnull(a.CancelDateTime,'') as CancelDateTime" +
		" from dbo.bcquotation as a" +
		" left join dbo.bcar as b on a.arcode=b.code" +
		" left join dbo.bcsale as c on a.salecode=c.code" +
		" left join dbo.BCContactList as d on a.ContactCode=d.code and a.arcode=d.ParentCode" +
		" left join dbo.BCProject as f on a.ProjectCode=f.code" +
		" left join dbo.BCAllocate as g on a.AllocateCode=g.code" +
		" where a.docno='" + q.DocNo + "'"
	fmt.Println(lcCommand)
	// Get saleorder from Database by docno
	//ss = []Saleorder{}
	err := db.Get(q, lcCommand)
	if err != nil {
		return err
	}
	fmt.Println(q)
	// todo: add Node sub details
	qtsub := `select a.roworder
			,isnull(a.docno,'') as docno
			,isnull(a.docdate,'') as docdate
			,isnull(a.arcode,'') as arcode
			,isnull(a.salecode,'') as salecode
			,isnull(a.ItemCode,'') as ItemCode
			,isnull(a.departcode,'') as departcode
			,isnull(b.name1,'') as ItemName
			,isnull(a.Qty,0) as Qty
			,isnull(a.whcode,'') as whcode
			,isnull(a.shelfcode,'') as shelfcode
			,isnull(a.Price,0) as Price
			,isnull(a.discountword,'') as DisCountWord
			,isnull(a.discountamount,0) as DisCountAmount
			,a.UnitCode,a.NetAmount,a.Amount
			,isnull(a.mydescription,'') as ItemDescription
			,a.IsConditionSend
			,a.Iscancel
			,isnull(a.PackingRate1,0) as PackingRate,a.LineNumber
		  from dbo.bcQuotationsub as a
		  left join dbo.bcitem as b on a.itemcode=b.code
		  where a.docno=?`
	fmt.Println(qtsub)
	err = db.Select(&q.Subs, qtsub, docno)
	return err
}

// test insert only header
func (qh *Qt) InsQt(db *sqlx.DB) (err error) {
	fmt.Println("start InsQt ")
	lcCommand := "insert into dbo.bcquotation(" +
		"DocNo ,DocDate,ArCode,SaleCode,TotalAmount," +
		"billtype,billstatus,taxtype,taxrate,sumofitemamount," +
		"discountword,discountamount,afterdiscount,beforetaxamount,taxamount," +
		"netamount,creditday,duedate,validity,assertstatus," +
		"mydescription1,iscancel) " +
		"values(?,?,?,?,?," +
		"		?,?,?,?,?," +
		"		?,?,?,?,?," +
		"		?,?,?,?,?," +
		"		?,?)"
	_, err = db.Exec(
		lcCommand,
		qh.DocNo, qh.DocDate, qh.ArCode, qh.SaleCode, qh.TotalAmount,
		qh.BillType, qh.BillStatus, qh.TaxType, qh.TaxRate, qh.SumOfItemAmount,
		qh.DisCountWord, qh.DisCountAmount, qh.AfterDiscountAmount, qh.BeforeTaxAmount, qh.TaxAmount,
		qh.NetAmount, qh.CreditDay, qh.DueDate, qh.Validity, qh.AssertStatus,
		qh.MyDescription1, qh.IsCancel,

	)

	if err != nil {
		fmt.Println(error.Error)
		return err
	}
	// todo : insert sub
	err = qh.InsSub(qh.Subs, db)
	if err != nil {
		fmt.Println(err.Error)
		return err
	}
	return nil

}

//  test insert sub
func (qh *Qt) InsSub(sub []*QtSub, db *sqlx.DB) (err error) {
	fmt.Println("start InsQt_SUB ")
	for _, k := range sub {
		lccommand := `insert into dbo.bcQuotationsub(
	DocNo,ItemCode,WhCode,Qty ,UnitCode	,
	Amount,	ArCode,taxtype,taxrate,remainqty,
	docdate,itemname
	) values
	(	?,?,?,?,?,
		?,?,?,?,?,
		?,?
	)`

	fmt.Println(qh.DocDate)
		_,err := db.Exec(lccommand,
			qh.DocNo, k.ItemCode, k.WhCode, k.Qty, k.UnitCode,
			k.Amount, qh.ArCode, qh.TaxType, qh.TaxRate, k.RemainQty,
			qh.DocDate, k.ItemName)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}

	return err
}

// ------------------------------------------------------------
func (q *Quotation) Insert(db *sqlx.DB) (NewQtNo string, err error) {

	lccommand := `insert into dbo.bcquotation (
		docno,docdate,taxtype,billtype,arcode,
		creditday,duedate,salecode,taxrate,isconfirm,
		mydescription1,billstatus,SumofItemAmount,discountword,discountamount,
		afterdiscount,beforetaxamount,taxamount,totalamount,iscancel,
		creatorcode,createdatetime,lasteditorcode,lasteditdatet,confirmcode,
		confirmdatetime,cancelcode,canceldatetime,deliverydate)
		values(  ?,?,?,?,?
			,?,?,?,?,?
			,?,?,?,?,?
			,?,?,?,?,?
			,?,?,?,?,?,
			?,?,?,?)`
	_, err = db.Exec(lccommand,
		q.DocNo, q.DocDate, q.TaxType, q.BillType, q.ArCode,
		q.CreditDay, q.DueDate, q.SaleCode, q.TaxRate, q.IsConfirm,
		q.MyDescription, q.BillStatus, q.SumItemAmount, q.DisCountWord, q.DisCountAmount,
		q.AfterDiscountAmount, q.BeforeTaxAmount, q.TaxAmount, q.TotalAmount, q.Iscancel,
		q.CreatorCode, q.CreateDateTime, q.EditorCode, q.EditDateTime, q.ConfirmCode,
		q.ConfirmDataTime, q.CancelCode, q.CancelDateTime, q.DeliveryDate)

	//fmt.Println(lccommand)
	if err != nil {
		return q.DocNo, err
	}

	// todo : insert sub
	err = q.InsertSub(q.Subs, db)
	if err != nil {
		fmt.Println(err.Error)
		return q.DocNo, err
	}
	return NewQtNo, err
}

func (q *Quotation) InsertSub(sub []*QuotationSub, db *sqlx.DB) (err error) {
	for _, k := range sub {

		lccommand := `
		insert into dbo.bcQuotationsub (
	docno,taxtype,itemcode,docdate,arcode,
	departcode,salecode,mydescription,itemname,whcode,
	shelfcode,qty,remainqty,price,discountword,
	discountamount,unitcode,netamount,amount,homeamount,
	isconditionsend,iscancel,linenumber,packingrate1,packingrate2 )
	values	(?,?,?,?,?,
		 ?,?,?,?,?,
		 ?,?,?,?,?,
		 ?,?,?,?,?,
		 ?,?,?,?,?
		 )`

		_, err := db.Exec(lccommand,
			q.DocNo, q.TaxType, k.ItemCode, q.DocDate, q.ArCode,
			q.DepartCode, q.SaleCode, k.Mydescription, k.ItemName, k.Whcode,
			k.Shelfcode, k.Qty, k.Remainqty, k.Price, k.DisCountWord,
			k.DisCountAmount, k.UnitCode, k.NetAmount, k.Amount, k.Homeamount,
			q.IsConditionSend, k.Iscancel, k.LineNumber, k.Packingrate1, k.Packingrate2)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		//fmt.Println(lccommand)
	}

	return err
}
func (q *Quotation) CheckExists(db *sqlx.DB, docno string) (bool) {
	fmt.Println("Begin CheckExists")
	lccommand := "select top 1 docno from dbo.bcquotation where docno = '" + docno + "'"
	rs, _ := db.Exec(lccommand)
	chkRow, _ := rs.RowsAffected()
	if chkRow > 0 {
		return true
		fmt.Println("data aleady exists!!! cannot insert this number : ", docno)
	}
	return false

}

func (q *Quotation) Void(db *sqlx.DB, docno string, cancelcode string) (msg string, result bool, err error) {
	//todo : Delete Before Update Saleorder
	//todo : saleorder
	fmt.Println("begin Quotation.Void")

	// todo : check status of quotation  cannot be  void if confirmed or refered
	result, err = q.isRefered(docno, db)
	if result == true {
		msg = "Document is confirmed  or refered or canceled "
		result = false
		return msg, result, err

	}

	// begin void document
	fmt.Println("begin void document ")
	// todo : update iscancel , cancelcode, canceldatetime to bcquotation
	lccommand := "update dbo.bcquotation set iscancel = 1,cancelcode=" + cancelcode + ",canceldatetime=getdate() where docno = '" + docno + "'"
	rs, _ := db.Exec(lccommand)
	_, err = rs.RowsAffected()
	if err != nil {
		result = false
		return "void header error ", result, err
	}

	//todo : quotationsub cancel by update iscacel field
	lccommand = "Update dbo.bcquotationsub set iscancel = 1 where docno = '" + docno + "'"
	_, err = db.Exec(lccommand)

	_, err = rs.RowsAffected()
	if err != nil {
		return "void Item error ", result, err
	}

	return "Void Completed", true, nil

}

// tempoary struct for check before void
type chkstatus struct {
	Docno      string
	Isconfirm  int32
	Billstatus int32
	Iscancel   int32
}

func (q *Quotation) isRefered(docno string, db *sqlx.DB) (result bool, err error) {
	lcCommand := "select top 1  docno,isconfirm,billstatus,iscancel from dbo.bcquotation where docno = '" + docno + "'"
	fmt.Println(lcCommand)
	fmt.Println("isRefered Check function begin")
	chk := chkstatus{}
	err = db.Get(&chk, lcCommand)
	if err != nil {
		fmt.Println("error query to check refered document quotation %s", err.Error())
		return false, err
	}
	if chk.Isconfirm == 1 || chk.Billstatus == 1 || chk.Iscancel == 1 {
		return true, nil // เอกสารอ้างอิงไปแล้ว
	}
	//fmt.Println("test")
	return false, nil
}

func (q *Quotation) Update(db *sqlx.DB) (msg string, err error) {
	// Update Quotation

	// check qt status - Update can use only new,onprocess only
	fmt.Println("model.quotation.Update() -> received : ", q.DocNo)

	// status : done , cancel cannot update
	if q.BillStatus == 1 || q.Iscancel == 1 {
		msg = "Cannot update status (CANCEL,DONE) "
		return msg, err
	}

	lcCommand := "update dbo.bcquotation set " +
		"docno = ?,docdate=?,taxtype=?,billtype=?,arcode=?," +
		"departcode=?,creditday=?,duedate=?,salecode=?,taxrate=?," +
		"isconfirm=?,mydescription1=?,billstatus=?,sumofitemamount=?,discountword=?," +
		"discountamount=?,afterdiscount=?,beforetaxamount=?,taxamount=?,totalamount=?," +
		"lasteditorcode = ?,lasteditdatet = getdate(),deliveryday=?,deliverydate=? " +
		" where docno = ?"
	_, err = db.Exec(lcCommand,
		q.DocNo, q.DocDate, q.TaxType, q.BillType, q.ArCode,
		q.DepartCode, q.CreditDay, q.DueDate, q.SaleCode, q.TaxRate,
		q.IsConfirm, q.MyDescription, q.BillStatus,
		q.SumItemAmount, q.DisCountWord, q.DisCountAmount,
		q.AfterDiscountAmount, q.BeforeTaxAmount, q.TaxAmount, q.TotalAmount,
		q.EditorCode, q.DeliveryDay, q.DeliveryDate,
		q.DocNo)
	if err != nil {
		msg = "Update Error "
		return msg, err
	}

	// update bcsaleordersub --> use delete before insert

	lcCommand = "delete from dbo.bcquotationsub where docno = '" + q.DocNo + "'"
	_, err = db.Exec(lcCommand)
	if err != nil {
		msg = "delete Detail of order error ! "
		return msg, err
	}

	// todo : insert bcsaleordersub
	// todo : insert sub
	err = q.InsertSub(q.Subs, db)
	if err != nil {
		fmt.Println(err.Error)
		msg = "insert Detail of order Error "
		return msg, err
	}

	msg = "Completed updated"
	return msg, nil

}
