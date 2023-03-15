package model

import "time"

const TableOrder = "order"

type OrderStatus int8

const (
	OrderCreate OrderStatus = iota + 0
	OrderDeduction
	OrderCancel
	OrderFinish
)

type OrderComment struct {
	Commenter string `json:"commenter" dynamo:"commenter"`
	Comment   string `json:"comment" dynamo:"comment"`
}

type Order struct {
	ID           string         `json:"id" dynamo:"id,hash"`
	ClientPhone  string         `json:"client_phone" dynamo:"client_phone"`
	AdviserPhone string         `json:"adviser_phone" dynamo:"adviser_phone"`
	Coin         float64        `json:"coin" dynamo:"coin"`
	Comments     []OrderComment `json:"comments" dynamo:"comments"`
	Status       OrderStatus    `json:"status" dynamo:"status"`
	UpdateTime   time.Time      `json:"update_time" dynamo:"update_time"`
	CreateTime   time.Time      `json:"create_time" dynamo:"create_time"`
}

func TableOrderCreate() error {
	return dbclient.CreateTable(TableOrder, &Client{}).Run()
}

func (this *Order) Insert() error {
	table := dbclient.Table(TableOrder)
	return table.Put(this).Run()
}

func ChangeOrderStatus(ID string, status OrderStatus) error {
	order := Order{}
	table := dbclient.Table(TableOrder)
	err := table.Get("id", ID).One(&order)
	if err != nil {
		return err
	}
	order.Status = status
	err = order.Insert()
	return err
}

func AddOrderComment(ID string, comment OrderComment) error {
	order := Order{}
	table := dbclient.Table(TableOrder)
	err := table.Get("id", ID).One(&order)
	if err != nil {
		return err
	}
	order.Comments = append(order.Comments, comment)
	err = order.Insert()
	return nil
}

func InfoOrder(ID string) (*Order, error) {
	order := Order{}
	table := dbclient.Table(TableOrder)
	err := table.Get("id", ID).One(&order)
	return &order, err
}

func ListOrder(phone string) ([]Order, error) {
	list := []Order{}
	table := dbclient.Table(TableOrder)
	err := table.Scan().Filter("'client_phone' = ?", phone).All(&list)
	return list, err
}
