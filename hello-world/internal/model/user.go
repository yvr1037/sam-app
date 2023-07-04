package model

import (
	"hello-world/global"
	"time"
)

type Gender int8

const (
	Man   Gender = 0
	Woman Gender = 1
)

type Status int8

const (
	Normal Status = 0
	Banned Status = 1
)

const TableClient = "Client"

type Client struct {
	Name       string    `dynamo:"name" json:"name"`
	Gender     Gender    `dynamo:"gender" json:"gender"`
	Phone      string    `dynamo:"phone,hash" json:"phone"`
	Password   string    `dynamo:"password" json:"password"`
	Birth      string    `dynamo:"birth" json:"birth"`
	Bio        string    `dynamo:"bio" json:"bio"`
	Coin       float64   `dynamo:"coin" json:"coin"`
	Status     Status    `dynamo:"status" json:"status"`
	Advisers   []string  `dynamo:"advisers" json:"advisers"`
	CreateTime time.Time `dynamo:"create_time" json:"create_time"`
}

func TableClientCreate() error {
	return global.DB.CreateTable(TableClient, &Client{}).Run()
}

func (this *Client) Insert() error {
	table := global.DB.Table(TableClient)
	return table.Put(this).Run()
}

func DeleteClient(name string) error {
	table := global.DB.Table(TableClient)
	return table.Delete("name", name).Run()
}
