package models

import "gorm.io/gorm"

type User struct {
	gorm.Model //We use this to transform everything into a table. It reads the struct and converts it into a table

	//The json tags must match how they are described in the post or they will not be read. And if you return them with an encode they will take that name
	//gorm adds by default an ID primary key, a created_at, updated_at and deleted_at fields.
	Username  string     `gorm:"not null;unique" json:"username"`
	Password  string     `gorm:"not null" json:"password"`
	Questions []Question `gorm:"foreignKey:CreatorID" json:"questions"`
}
