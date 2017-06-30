package Resp

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	//"go/doc"
)

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
	RefNo               string `json:"ref_no" db:"RefNo"`
	TaxRate             float64 `json:"tax_rate" db:"TaxRate"`
	TaxType             int `json:"tax_type" db:"TaxType"`
	MyDescription       string `json:"my_description" db:"MyDescription"`
	SumItemAmount       float32 `json:"sum_item_amount" db:"SumItemAmount"`
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
	Subs                []*QuotationSub `json:"subs"`
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
		" a.DocDate," +
		" isnull(a.DueDate,'') as DueDate," +
		" isnull(a.DeliveryDate,'') as DeliveryDate," +
		" isnull(a.arcode,'') as ArCode," +
		" isnull(b.name1,'') as ArName," +
		" isnull(b.billAddress,'') as ArAddress," +
		" isnull(b.Telephone,'') as ArTelephone," +
		" isnull(b.fax,'') as ArFax," +
		" isnull(a.salecode,'') as SaleCode," +
		" isnull(c.name,'') as SaleName," +
		" '' as RefNo," +
		" isnull(a.taxrate,7) as TaxRate," +
		" a.TaxType," +
		" isnull(a.MyDescription1,'') as MyDescription," +
		" isnull(a.sumofitemamount,0) as SumItemAmount," +
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
		" from bcnp.dbo.bcquotation as a" +
		" left join bcnp.dbo.bcar as b on a.arcode=b.code" +
		" left join bcnp.dbo.bcsale as c on a.salecode=c.code" +
		" left join bcnp.dbo.BCContactList as d on a.ContactCode=d.code and a.arcode=d.ParentCode" +
		" left join bcnp.dbo.BCProject as f on a.ProjectCode=f.code" +
		" left join bcnp.dbo.BCAllocate as g on a.AllocateCode=g.code" +
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
			,a.docno
			,a.docdate
			,a.arcode
			,isnull(a.salecode,'') as salecode
			,a.ItemCode
			,isnull(a.departcode,'') as departcode
			,isnull(b.name1,'') as ItemName
			,a.Qty
			,isnull(a.whcode,'') as whcode
			,isnull(a.shelfcode,'') as shelfcode
			,a.Price
			,isnull(a.discountword,'') as DisCountWord
			,isnull(a.discountamount,0) as DisCountAmount
			,a.UnitCode,a.NetAmount,a.Amount
			,isnull(a.mydescription,'') as ItemDescription
			,a.IsConditionSend
			,a.Iscancel
			,isnull(a.PackingRate1,0) as PackingRate,a.LineNumber
		  from bcnp.dbo.bcQuotationsub as a
		  left join bcnp.dbo.bcitem as b on a.itemcode=b.code
		  where a.docno=?`
	fmt.Println(qtsub)
	err = db.Select(&q.Subs, qtsub, docno)
	return err
}

func (q *Quotation)Insert(db *sqlx.DB) (NewQtNo string, err error) {

	lccommand := `insert into bcnp.dbo.bcquotation (
		docno,docdate,taxtype,billtype,arcode,
		creditday,duedate,salecode,taxrate,isconfirm,
		mydescription,billstatus,SumItemAmount,discountword,discountamount,
		afterdiscount,beforetaxamount,taxamount,totalamount,beforetaxamount,
		iscancel,creatorcode,createdatetime,lasteditorcode,lasteditdatet,
		confirmcode,confirmdatetime,cancelcode,canceldatetime,deliverydate)
		values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
	_, err = db.Exec(lccommand,
		q.DocNo, q.DocDate, q.TaxType, q.BillType, q.ArCode,
		q.CreditDay, q.DueDate, q.SaleCode, q.TaxRate,q.IsConfirm,
		q.MyDescription, q.BillStatus,q.SumItemAmount, q.DisCountWord, q.DisCountAmount,
		q.AfterDiscountAmount, q.BeforeTaxAmount,q.TaxAmount, q.TotalAmount, q.BeforeTaxAmount,
		q.Iscancel, q.CreatorCode,q.CreateDateTime, q.EditorCode, q.EditDateTime,
		q.ConfirmCode, q.ConfirmDataTime,q.CancelCode, q.CancelDateTime, q.DeliveryDate)

	fmt.Println(lccommand)
	if err != nil {
		return q.DocNo, err
	}

	// todo : insert sub
	err = q.InsertSub(q.Subs,db)
	if err != nil {
		fmt.Println(err.Error)
		return q.DocNo,err
	}
	return NewQtNo, err
}

func (q *Quotation)InsertSub(sub []*QuotationSub, db *sqlx.DB) (err error) {
	for _,k :=  range sub{

	lccommand := `
		insert into bcnp.dbo.bcQuotationsub (
	docno,taxtype,itemcode,docdate,arcode,
	departcode,salecode,mydescription,itemname,whcode,
	shelfcode,qty,remainqty,price,discountword,
	discountamount,unitcode,netamount,amount,homeamount,
	itemdescription,isconditionsend,iscancel,linenumber,packingrate,
	packingrate1,packingrate2 )
	values (
	?,?,?,?,?
	?,?,?,?,?
	?,?,?,?,?
	?,?,?,?,?
	?,?,?,?,?
	?,? )
	`
		_, err := db.Exec(lccommand,
		k.DocNo,k.Taxtype,k.ItemCode,k.DocDate,k.Arcode,
		k.Departcode,k.Salecode,k.Mydescription,k.ItemName,
		k.Shelfcode, k.Qty, k.Remainqty, k.Price, k.DisCountWord,
		k.ItemDescription,k.IsConditionSend,k.Iscancel,k.LineNumber,k.PackingRate,
		k.Packingrate1,k.Packingrate2	)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		fmt.Println(lccommand)
	}

	return err
}
