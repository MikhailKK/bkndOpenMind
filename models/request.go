package models

type Request struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Description string `json:"description"`
	Phone       string `json:"phone"`
}

func (r Request) GetDescription() string {
	return r.Description
}
