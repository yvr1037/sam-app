package model

import "time"

type WorkStatus int8

const (
	Open  WorkStatus = 0
	Close WorkStatus = 1
)

const TableAdviser = "adviser"

type AdviserComment struct {
	Commenter string `json:"commenter" dynamo:"commenter"`
	Rate      int8   `json:"rate" dynamo:"rate"`
	Comment   string `json:"comment" dynamo:"comment"`
}

type Adviser struct {
	Phone       string           `json:"phone" dynamo:"phone,hash"`
	Name        string           `json:"name" dynamo:"name"`
	Gender      Gender           `json:"gender" dynamo:"gender"`
	Password    string           `json:"password" dynamo:"password"`
	Birth       string           `json:"birth" dynamo:"birth"`
	Bio         string           `json:"bio" dynamo:"bio"`
	Coin        float64          `json:"coin" dynamo:"coin"`
	Status      Status           `json:"status" dynamo:"status"`
	WorkStatus  WorkStatus       `json:"work_status" dynamo:"work_status"`
	OrderNumber int64            `json:"order_number" dynamo:"order_number"`
	OrderFinish int64            `json:"order_finish" dynamo:"order_finish"`
	Star        float64          `json:"star" dynamo:"star"`
	Comments    []AdviserComment `json:"comments" dynamo:"comments"`
	CommentNum  int64            `json:"comment_num" dynamo:"comment_num"`
	Experience  int32            `json:"experience" dynamo:"experience"`
	CreateTime  time.Time        `json:"create_time" dynamo:"create_time"`
}

func TableAdviserCreate() error {
	return dbclient.CreateTable(TableAdviser, &Adviser{}).Run()
}

func (this *Adviser) Insert() error {
	table := dbclient.Table(TableAdviser)
	return table.Put(this).Run()
}

func DeleteAdviser(name string) error {
	table := dbclient.Table(TableAdviser)
	return table.Delete("name", name).Run()
}

func LoginByPwdAdviser(pwd, phone string) (bool, error) {
	table := dbclient.Table(TableAdviser)
	// table.Get("phone", phone).Filter("'password' == ?", pwd).One()
	num, err := table.Scan().Filter("'phone' = ? AND 'password' = ? ", phone, pwd).Count()
	return num == 1, err
}

func ExistAdvister(phone string) (bool, error) {
	table := dbclient.Table(TableAdviser)
	num, err := table.Scan().Filter("'phone' = ?", phone).Count()
	return num == 1, err
}

func InfoAdviser(phone string) (*Adviser, error) {
	adviser := Adviser{}
	table := dbclient.Table(TableAdviser)
	err := table.Get("phone", phone).One(&adviser)
	return &adviser, err
}
