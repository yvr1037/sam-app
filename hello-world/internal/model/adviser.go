package model

import "time"

type WorkStatus int8

const (
	Open  WorkStatus = 0
	Close WorkStatus = 0
)

const TableAdviser = "Adviser"

type Adviser struct {
	Name        string     `json:"name" dynamo:"name"`
	Gender      Gender     `json:"gender" dynamo:"gender"`
	Phone       string     `json:"phone" dynamo:"phone,hash"`
	Password    string     `json:"password" dynamo:"password"`
	Birth       string     `json:"birth" dynamo:"birth"`
	Bio         string     `json:"bio" dynamo:"bio"`
	Coin        float64    `json:"coin" dynamo:"coin"`
	Status      Status     `json:"status" dynamo:"status"`
	WorkStatus  WorkStatus `json:"work_status" dynamo:"work_status"`
	OrderNumber int64      `json:"order_number" dynamo:"order_number"`
	Star        float64    `json:"star" dynamo:"star"`
	CommentNum  int64      `json:"comment_num" dynamo:"comment_num"`
	Experience  int32      `json:"experience" dynamo:"experience"`
	CreateTime  time.Time  `json:"create_time" dynamo:"create_time"`
}

func TableAdviserCreate() error {
	return dbclient.CreateTable(TableClient, &Client{}).Run()
}

func (this *Adviser) Insert() error {
	table := dbclient.Table(TableAdviser)
	return table.Put(this).Run()
}

func DeleteAdviser(name string) error {
	table := dbclient.Table(TableClient)
	return table.Delete("name", name).Run()
}
