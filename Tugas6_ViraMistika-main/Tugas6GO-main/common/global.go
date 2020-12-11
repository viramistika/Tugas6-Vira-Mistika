package common

//Struct API
// Order struct (Model) ...

//Message struct
type Message struct {
	Code    int     `json:"code"`
	Remark  string  `json:"remark"`
	OrderID string  `json:"orderID"`
	Orders  *Orders `json:"orders,omitempty"`
	Result  *Result `json:"result,omitempty"`
}

//Orders Model
type Orders struct {
	OrderID    string         `json:"orderID"`
	CustomerID string         `json:"customerID"`
	EmployeeID string         `json:"employeeID"`
	OrderDate  string         `json:"orderDate"`
	OrdersDet  []OrdersDetail `json:"ordersDetail"`
}

//OrdersDetail Model
type OrdersDetail struct {
	OrderID     string  `json:"orderID"`
	ProductID   string  `json:"ProductID"`
	ProductName string  `json:"ProductName"`
	UnitPrice   float64 `json:"UnitPrice"`
	Quantity    int     `json:"Quantity"`
}

//Result Model
type Result struct {
	Code   int    `json:"code"`
	Remark string `json:"remark,omitempty"`
}

//Customer Model
type Customer struct {
	CustomerID   string `json:"customerID"`
	CompanyName  string `json:"companyName"`
	ContactName  string `json:"contactName"`
	ContactTitle string `json:"contactTitle"`
	Address      string `json:"address"`
	City         string `json:"city"`
	Country      string `json:"country"`
	Phone        string `json:"phone"`
	PostalCode   string `json:"postalCode"`
}

//Product Model
type Product struct {
	ProductID       int    `json:"productID"`
	ProductName     string `json:"productName"`
	SupplierID      string `json:"supplierID"`
	CategoryID      string `json:"sategoryID"`
	QuantityPerUnit string `json:"quantityPerUnit"`
	UnitPrice       string `json:"unitPrice"`
	UnitsInStock    string `json:"unitsInStock"`
	UnitsOnOrder    string `json:"unitsOnOrder"`
}

type FastPayRequest struct {
	Merchant   string `json:"merchant"`
	MerchantID string `json:"merchant_id"`
	Request    string `json:"request"`
	Signature  string `json:"signature"`
}

type FastPayResponse struct {
	Merchant       string           `json:"merchant"`
	MerchantID     string           `json:"merchant_id"`
	PaymentChannel []PaymentChannel `json:"payment_channel"`
	Response       string           `json:"response"`
	ResponseCode   string           `json:"response_code"`
	ResponseDesc   string           `json:"response_desc"`
}

type PaymentChannel struct {
	PgCode string `json:"pg_code"`
	PgName string `json:"pg_name"`
}

type InquiryStatusReq struct {
	Request       string `json:"request"`
	TransactionID string `json:"trx_id"`
	MerchantID    string `json:"merchant_id"`
	BillNO        string `json:"bill_no"`
}

type InquiryStatusRes struct {
	Response          string `json:"response"`
	TransactionID     string `json:"trx_id"`
	MerchantID        string `json:"merchant_id"`
	Merchant          string `json:"merchant"`
	BillNO            string `json:"bill_no"`
	PaymentRefference string `json:"payment_reff"`
	PaymentDate       string `json:"payment_date"`
	PaymentStatusCode string `json:"payment_status_code"`
	PaymentStatusDesc string `json:"payment_status_desc"`
	ResponseCode      string `json:"response_code"`
	ResponseDesc      string `json:"response_desc"`
}

type TripRequest struct {
	DepartureDate1 string `json:"depature_date_1"`
	DepartureDate2 string `json:"depature_date_2"`
	Provinsi       int    `json:"provinsi"`
}

type TripResponse struct {
	Message    string       `json:"message"`
	Status     string       `json:"status"`
	TripDetail []TripDetail `json:"data"`
}

type TripDetail struct {
	AirlineName      string `json:"AirlineName"`
	AirportName      string `json:"AirportName"`
	CityName         string `json:"CityName"`
	Currency         string `json:"Currency"`
	DepartureDate    string `json:"DepartureDate"`
	Description      string `json:"Description"`
	Destination      string `json:"Destination"`
	DoubleType       string `json:"DoubleType"`
	Duration         string `json:"Duration"`
	Goods            string `json:"Goods,omitempty"`
	HotelName        string `json:"HotelName,omitempty"`
	HotelRating      string `json:"HotelRating,omitempty"`
	Lat              string `json:"Lat,omitempty"`
	LicenseNumber    string `json:"LicenseNumber,omitempty"`
	Logo             string `json:"Logo,omitempty"`
	Long             string `json:"Long,omitempty"`
	Origin           string `json:"Origin,omitempty"`
	OriginCity       string `json:"OriginCity,omitempty"`
	Price            string `json:"Price,omitempty"`
	PromoCode        string `json:"PromoCode,omitempty"`
	PromoDescription string `json:"PromoDescription,omitempty"`
	Provinsi         string `json:"Provinsi,omitempty"`
	QuadType         string `json:"QuadType,omitempty"`
	Rating           string `json:"Rating,omitempty"`
	ReturnDate       string `json:"ReturnDate,omitempty"`
	TermCondition    string `json:"TermCondition,omitempty"`
	Transit          string `json:"Transit,omitempty"`
	TravelID         string `json:"TravelID,omitempty"`
	TravelName       string `json:"TravelName,omitempty"`
	TripID           string `json:"TripID,omitempty"`
	TripleType       string `json:"TripleType,omitempty"`
}

//End Struct API
