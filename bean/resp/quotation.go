package Resp

import (
	"github.com/jmoiron/sqlx"
	"fmt"
)

type Quotation struct {
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
	ItemCode          string `json:"item_code" db:"ItemCode"`
	ItemName          string `json:"item_name" db:"ItemName"`
	Qty               float64 `json:"qty" db:"Qty"`
	Price             float64 `json:"price" db:"Price"`
	DisCountWordSub   string `json:"dis_count_word_sub" db:"DisCountWordSub"`
	DisCountAmountSub float64 `json:"dis_count_amount_sub" db:"DisCountAmountSub"`
	UnitCode          string `json:"unit_code" db:"UnitCode"`
	NetAmount         float64 `json:"net_amount" db:"NetAmount"`
	Amount            float64 `json:"amount" db:"Amount"`
	ItemDescription   string `json:"item_description" db:"ItemDescription"`
	IsConditionSend   int `json:"is_condition_send" db:"IsConditionSend"`
	Iscancel          int `json:"is_cancel" db:"Iscancel"`
	PackingRate       float64 `json:"packing_Rate" db:"PackingRate"`
	LineNumber        int `json:"line_number" db:"LineNumber"`
}

func (q *Quotation) GetByDocno(docno string,db *sqlx.DB)error {

	fmt.Println(q.DocNo)

	lcCommand := "select isnull(a.docno,'') as DocNo,a.DocDate,isnull(a.DueDate,'') as DueDate,isnull(a.DeliveryDate,'') as DeliveryDate" +
		",isnull(a.arcode,'') as ArCode,isnull(b.name1,'') as ArName,isnull(b.billAddress,'') as ArAddress" +
		",isnull(b.Telephone,'') as ArTelephone,isnull(b.fax,'') as ArFax,isnull(a.salecode,'') as SaleCode" +
		",isnull(c.name,'') as SaleName,'' as RefNo,isnull(a.taxrate,7) as TaxRate,a.TaxType" +
		",isnull(a.MyDescription1,a.MyDescription2) as MyDescription,isnull(a.sumofitemamount,0) as SumItemAmount" +
		",isnull(a.discountword,'') as DisCountWord,isnull(a.discountamount,0) as DisCountAmount" +
		",isnull(a.AfterDiscount,0) as AfterDiscountAmount,isnull(a.BeforeTaxAmount,0) as BeforeTaxAmount" +
		",isnull(a.TaxAmount,0) as TaxAmount,isnull(a.TotalAmount,0) as TotalAmount,a.Iscancel,a.IsConfirm,a.BillStatus" +
		",a.BillType,isnull(a.CreditDay,0) as CreditDay,isnull(ContactCode,'') as ContactCode,isnull(d.name,'') as ContactName" +
		",isnull(a.ProjectCode,'') as ProjectCode,isnull(f.name,'') as ProjectName,isnull(a.AllocateCode,'') as AllocateCode" +
		",isnull(g.name,'') as AllocateName,isnull(a.CreatorCode,'') as CreatorCode,isnull(a.CreateDateTime,'') as CreateDateTime" +
		",isnull(a.lastEditorCode,'') as EditorCode,isnull(a.lasteditdatet,'') as EditDateTime,isnull(a.ConfirmCode,'') as ConfirmCode" +
		",isnull(a.ConfirmDateTime,'') as ConfirmDataTime,isnull(a.CancelCode,'') as  CancelCode" +
		",isnull(a.CancelDateTime,'') as CancelDateTime" +
		" from bcnp.dbo.bcquotation as a" +
		" left join bcnp.dbo.bcar as b on a.arcode=b.code" +
		" left join bcnp.dbo.bcsale as c on a.salecode=c.code" +
		" left join bcnp.dbo.BCContactList as d on a.ContactCode=d.code and a.arcode=d.ParentCode" +
		" left join bcnp.dbo.BCProject as f on a.ProjectCode=f.code" +
		" left join bcnp.dbo.BCAllocate as g on a.AllocateCode=g.code" +
		" where a.docno='"+q.DocNo+"'"
	fmt.Println(lcCommand)
	// Get saleorder from Database by docno
	//ss = []Saleorder{}
	err := db.Get(q, lcCommand)
	if err != nil {
		return err
	}
	fmt.Println(q)
	// todo: add Node sub details
	qtsub := `select a.ItemCode,b.name1 as ItemName,a.Qty,a.Price,isnull(a.discountword,'') as DisCountWordSub
			,isnull(a.discountamount,0) as DisCountAmountSub,a.UnitCode,a.NetAmount,a.Amount
			,isnull(a.mydescription,'') as ItemDescription,a.IsConditionSend,a.Iscancel,isnull(a.PackingRate1,0) as PackingRate,a.LineNumber
		from bcnp.dbo.bcQuotationsub as a
		left join bcnp.dbo.bcitem as b on a.itemcode=b.code
		where a.docno=?`
	fmt.Println(qtsub)
	err = db.Select(&q.Subs, qtsub, docno)
	return err
}
