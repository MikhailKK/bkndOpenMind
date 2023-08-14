package models

type Request struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Description string `json:"description"`
	Tag         string `json:"tag"`
}

func (r Request) GetDescription() string {
	return r.Description
}
