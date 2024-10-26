package app

import "time"

type FoodType string

const (
	Biscuits FoodType = "biscuits"
	Meat     FoodType = "meat"
)

type User struct {
	ID        string    `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name      string    `json:"name"`
	Email     string    `json:"email" gorm:"unique"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Cat struct {
	ID        string    `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Meal struct {
	ID        string    `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CatID     string    `json:"cat_id" gorm:"type:uuid"`
	FedByID   string    `json:"fed_by_id" gorm:"type:uuid"`
	Timestamp time.Time `json:"timestamp"`
	FoodTypes []string  `json:"food_types" gorm:"type:text[]"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Associations
	Cat   Cat  `json:"cat" gorm:"foreignKey:CatID"`
	FedBy User `json:"fed_by" gorm:"foreignKey:FedByID"`
}
