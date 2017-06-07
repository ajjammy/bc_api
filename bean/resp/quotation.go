package Resp

type Quotation struct {
	DocNo               string `json:"doc_no" db:"DocNo"`
	DocDate             string `json:"doc_date" db:"DocDate"`
	ArCode              string `json:"ar_code" db:"ArCode"`
	ArName              string `json:"ar_name" db:"ArName"`
	ArAddress           string `json:"ar_address" db:"ArAddress"`
	ArTelephone         string `json:"ar_telephone" db:"ArTelephone"`
	ArFax               string `json:"ar_fax" db:"ArFax"`
	SaleCode            string `json:"sale_code"  db:"SaleCode"`
	SaleName            string `json:"sale_name" db:"SaleName"`
	RefNo               string `json:"ref_no" db:"RefNo"`
	TaxRate             string `json:"tax_rate" db:"TaxRate"`
	TaxType             int `json:"tax_type" db:"TaxType"`
	MyDescription       string `json:"my_description" db:"MyDescription"`
	SumItemAmount       float32 `json:"sum_item_amount" db:"SumItemAmount"`
	DisCountWord        string `json:"dis_count_word" db:"DisCountWord"`
	DisCountAmount      float64 `json:"dis_count_amount" db:"DisCountAmount"`
	AfterDiscountAmount float64 `json:"after_discount_amount" db:"AfterDiscountAmount"`
	BeforeTaxAmount     float64 `json:"before_tax_amount" db:"BeforeTaxAmount"`
	TaxAmount           float64 `json:"tax_amount" db:"TaxAmount"`
	TotalAmount         float64 `json:"total_amount" db:"TotalAmount"`
	IsConditionSend     int `json:"is_condition_send" db:"IsConditionSend"`
	CreditDay           int `json:"credit_day" db:"CreditDay"`
	ContactCode         string `json:"contact_code" db:"ContactCode"`
	ContactName         string `json:"contact_name" db:"ContactName"`
	ApproveCode         string `json:"approve_code" db:"ApproveCode"`
	ApproveName         string `json:"approve_name"  db:"ApproveName"`
	ProjectCode         string `json:"project_code" db:"ProjectCode"`
	ProjectName         string `json:"project_name" db:"ProjectName"`
	AllocateCode        string `json:"allocate_code" db:"AllocateCode"`
	AllocateName        string `json:"allocate_name" db:"AllocateName"`
	CreatorCode         string `json:"creator_code" db:"CreatorCode"`
	CreateDateTime      string `json:"create_date_time" db:"CreateDateTime"`
	EditorCode          string `json:"editor_code" db:"EditorCode"`
	EditDateTime        string `json:"edit_date_time" db:"EditDateTime"`
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
	LineNumber        int `json:"line_number" db:"LineNumber"`
}
