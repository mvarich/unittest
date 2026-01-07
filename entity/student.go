package entity

import(
	"gorm.io/gorm"
)

type Student struct{
	gorm.Model
	StudentID string `valid:"matches(^[BMD]\\d{7}$)~StudentID is invalid,required~StudentID is required"`
	Name string `valid:"required~Name is invalid"`
	GPA float32 `valid:"range(0|4)~GPA is invalid,required~GPA is required"`
}