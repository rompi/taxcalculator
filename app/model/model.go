package model

type Object struct {
	Id         int64   `json:"id"`
	Name       string  `json:"name"`
	TaxCode    int     `json:"tax_code"`
	Price      float64 `json:"price"`
	Type       string  `json:"type"`
	Refundable bool    `json:"refundable"`
	Tax        float64 `json:"tax"`
	Amount     float64 `json:"Amount"`
}

type ResponseList struct {
	Data  []*Object `json:"data"`
	Error []string  `json:"error"`
}

type Response struct {
	Data  *Object  `json:"data"`
	Error []string `json:"error"`
}

func BuildResponse(o *Object, msg []string) *Response {
	payload := &Response{
		Data:  o,
		Error: msg,
	}
	return payload
}

func BuildResponseList(o []*Object, msg []string) *ResponseList {
	payload := &ResponseList{
		Data:  o,
		Error: msg,
	}
	return payload
}

func BuildObject(name string, taxcode int, price float64) *Object {
	return &Object{
		Name:    name,
		TaxCode: taxcode,
		Price:   price,
	}
}

func BuildObjects(o *Object) []*Object {
	var objs []*Object
	objs = append(objs, o)
	return objs
}
