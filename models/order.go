package models

import (
	"errors"
	"strconv"
	"time"
)

type Amount struct {
	Subtotal uint `json:"subtotal"`
	Shipping uint `json:"shipping"`
	Total    uint `json:"total"`
}

type OrderInfo struct {
	OrderCode       string `json:"-" sql:"column:code"`
	CustomerName    string `json:"name" sql:"customer_name"`
	CustomerAddress string `json:"address" sql:"customer_address"`
	CustomerPhone   string `json:"phone" sql:"customer_phone"`
	CustomerEmail   string `json:"email" sql:"customer_email"`
	CustomerNote    string `json:"note" sql:"column:note"`
}

type CheckoutReturn struct {
	Id        string      `json:"id,omitempty"`
	Items     []OrderItem `json:"items,omitempty"`
	Amounts   Amount      `json:"amounts" sql:"-"`
	OrderInfo *OrderInfo  `json:"customer,omitempty" sql:"-"`
	Status    string      `json:"status"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CheckoutReturnSlice []CheckoutReturn
type OrderSlice []Order

type Alert struct {
	Type    string
	Message string
}

type Order struct {
	Id            uint   `json:"-"`
	AccessToken   string `json:"access_token,omitempty"`
	Status        string `json:"status,omitempty"`
	StatusUpdated bool   `json:"-" sql:"-"`

	Items              []OrderItem `json:"items" sql:"order_items"`
	ItemUpdateNew      bool        `json:"-" sql:"-"`
	ItemUpdateIdx      int         `json:"-" sql:"-"`
	ItemUpdateId       int         `json:"-" sql:"-"`
	ItemUpdateQuantity int         `json:"-" sql:"-"`

	OrderCodeRet string `json:"id,omitempty" sql:"-"`
	OrderInfo    `json:"-"`

	Amounts Amount  `json:"amounts" sql:"-"`
	Alerts  []Alert `json:"alerts,omitempty" sql:"-"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type OrderStatusLog struct {
	Id        uint
	Code      string
	Status    string
	CreatedAt time.Time
}

type OrderItem struct {
	Id                       uint    `json:"id,string" sql:"id"`
	OrderId                  uint    `json:"-" sql:"REFERENCES Orders(id)"`
	ProductId                uint    `json:"product_id,string" sql:"REFERENCES products(id)"`
	ProductName              string  `json:"name" sql:"product_name"`
	ProductImageThumbnailUrl *string `json:"image_thumbnail_url" sql:"-"`
	ProductPrice             int     `json:"price" sql:"product_price"`
	Quantity                 int     `json:"quantity"`
}

type OrderUrl struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
}

type OrderPage struct {
	Data   *CheckoutReturnSlice `json:"data"`
	Paging OrderUrl             `json:"paging"`
}

func getOrderPage(offset uint, limit uint, total uint, sort string) OrderUrl {
	prevNum := uint64(offset - limit)
	nextNum := uint64(offset + limit)
	if offset < limit {
		prevNum = 0
	}
	if total <= offset {
		if total > limit {
			prevNum = uint64(total - limit)
		} else {
			prevNum = 0
		}
	}
	next := "/orders?offset=" + strconv.FormatUint(nextNum, 10) + "&limit=" + strconv.FormatUint(uint64(limit), 10) + "&" + sort
	previous := "/orders?offset=" + strconv.FormatUint(prevNum, 10) + "&limit=" + strconv.FormatUint(uint64(limit), 10) + "&" + sort

	// Nothing to show on next_url
	if uint64(total) <= nextNum {
		// If offset already zero, not thing to show on previous_url also
		if offset == 0 {
			return OrderUrl{}
		}

		// At least, we have something to show on previous_url
		return OrderUrl{
			Previous: &previous,
		}
	}
	if offset == 0 {
		return OrderUrl{
			Next: &next,
		}
	}
	return OrderUrl{
		Next:     &next,
		Previous: &previous,
	}

}

func (orders OrderSlice) GetPaging(offset uint, limit uint, total uint, sort string) *OrderPage {
	orders_return := CheckoutReturnSlice{}
	for idx, _ := range orders {
		orders_return = append(orders_return, orders[idx].ReturnCheckoutRequest())
		orders_return[idx].Id = orders[idx].OrderCode
	}
	return &OrderPage{
		Data:   &orders_return,
		Paging: getOrderPage(offset, limit, total, sort),
	}
}

func (order *Order) UpdateItems(product_id *uint, item_id *uint, quantity int, product_name string, product_price int) error {
	// Note: This function only accept 1 item at a time

	// microserve stuff...
	order.ItemUpdateNew = false
	if item_id != nil {
		order.ItemUpdateId = int(*item_id)
	}
	order.ItemUpdateQuantity = quantity

	for idx, item := range order.Items {
		// update quantity base on offset {item_id, quantity}
		if product_id != nil {
			if item.ProductId == *product_id {
				order.ItemUpdateIdx = idx
				// The item already in the list, return (not check for new item)
				return nil
			}
		}

		// update quantity base on {product_id, quantity}
		if item_id != nil {
			if item.Id == *item_id {
				order.Items[idx].Quantity = quantity
				order.ItemUpdateIdx = idx
				return nil
			}
		}
	}

	if item_id != nil {
		return errors.New("Item ID not found")
	}

	if quantity < 0 {
		return errors.New("Quantity for item should bigger than 0")
	}

	// The item is new, add it to the []item list {product_id, quantity}
	if product_id != nil {
		order.Items = append(order.Items,
			OrderItem{
				ProductId:    *product_id,
				Quantity:     quantity,
				ProductName:  product_name,
				ProductPrice: product_price,
			})
	}
	order.ItemUpdateNew = true

	return nil
}

func (order *Order) CalculateAmount() {
	for _, item := range order.Items {
		order.Amounts.Subtotal += uint(item.ProductPrice) * uint(item.Quantity)
	}
	order.Amounts.Shipping = 0
	order.Amounts.Total = order.Amounts.Shipping + order.Amounts.Subtotal
}

func (order *Order) EraseAccessToken() {
	order.AccessToken = ""
}

func (order *Order) RemoveZeroQuantityItems() {
	for idx, _ := range order.Items {
		if order.Items[idx].Quantity <= 0 {
			order.Items = append(order.Items[:idx], order.Items[idx+1:]...)
			return
		}
	}
}

func (order *Order) ReturnCheckoutRequest() CheckoutReturn {
	ret := new(CheckoutReturn)
	ret.Id = order.OrderCodeRet
	ret.Items = order.Items
	ret.Amounts = order.Amounts
	if order.OrderInfo.CustomerName != "" {
		ret.OrderInfo = &order.OrderInfo
	}
	ret.Status = order.Status
	ret.UpdatedAt = order.UpdatedAt
	ret.CreatedAt = order.CreatedAt
	return *ret
}
