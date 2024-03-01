package models

type Model struct {
	ID        int    `json:"id" gorm:"primarykey"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	DeletedAt string `json:"deleted_at,omitempty" gorm:"index"`
}

func (m *Model) ValidateBase() error {
	return nil
}

func (m *Model) ValidateInsert() error {
	return nil
}
