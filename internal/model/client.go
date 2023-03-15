package model

import "time"

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

const TableClient = "client"

type Client struct {
	Phone      string    `dynamo:"phone,hash" json:"phone"`
	Name       string    `dynamo:"name" json:"name"`
	Gender     Gender    `dynamo:"gender" json:"gender"`
	Password   string    `dynamo:"password" json:"password"`
	Birth      string    `dynamo:"birth" json:"birth"`
	Bio        string    `dynamo:"bio" json:"bio"`
	Coin       float64   `dynamo:"coin" json:"coin"`
	Status     Status    `dynamo:"status" json:"status"`
	Advisers   []string  `dynamo:"advisers" json:"advisers"`
	CreateTime time.Time `dynamo:"create_time" json:"create_time"`
}

func TableClientCreate() error {
	return dbclient.CreateTable(TableClient, &Client{}).Run()
}

func (this *Client) Insert() error {
	table := dbclient.Table(TableClient)
	return table.Put(this).Run()
}

func DeleteClient(name string) error {
	table := dbclient.Table(TableClient)
	return table.Delete("name", name).Run()
}

func LoginByPwd(pwd, phone string) (bool, error) {
	table := dbclient.Table(TableClient)

	// table.Get("phone", phone).Filter("'password' == ?", pwd).One()
	num, err := table.Scan().Filter("'phone' = ? AND 'password' = ? ", phone, pwd).Count()
	return num == 1, err
}

func ExistClient(phone string) (bool, error) {
	table := dbclient.Table(TableClient)
	num, err := table.Scan().Filter("'phone' = ?", phone).Count()
	return num == 1, err
}

func InfoClient(phone string) (*Client, error) {
	client := Client{}
	table := dbclient.Table(TableClient)
	err := table.Get("phone", phone).One(&client)
	return &client, err
}
