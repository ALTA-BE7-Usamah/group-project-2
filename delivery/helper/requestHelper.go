package helper

type OrderRequestFormat struct {
	CartID     []uint              `json:"cart_id" form:"cart_id"`
	Address    AddressRequesFormat `json:"address" form:"address"`
	CreditCard CreditRequestFormat `json:"credit_card" form:"credit_card"`
}

type AddressRequesFormat struct {
	Street  string `json:"street" form:"street"`
	City    string `json:"city" form:"city"`
	State   string `json:"state" form:"state"`
	ZipCode uint   `json:"zip_code" form:"zip_code"`
}

type CreditRequestFormat struct {
	Type   string `json:"type" form:"type"`
	Name   string `json:"name" form:"name"`
	Number string `json:"number" form:"number"`
	CVV    uint   `json:"cvv" form:"cvv"`
	Month  uint   `json:"month" form:"month"`
	Year   uint   `json:"year" form:"year"`
}
