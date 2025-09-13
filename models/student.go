package models

import (
    "time"
)

// Student مدل دانشجو
type Student struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    DeletedAt time.Time `gorm:"index" json:"deleted_at,omitempty"`
    Name      string    `json:"name" gorm:"not null"`
    Age       int       `json:"age" gorm:"not null"`
    Email     string    `json:"email" gorm:"unique"`
    Major     string    `json:"major" gorm:"default:''"`
}

// StudentResponse مدل پاسخ برای ارسال اطلاعات دانشجو
type StudentResponse struct {
    ID    uint   `json:"id"`
    Name  string `json:"name"`
    Age   int    `json:"age"`
    Email string `json:"email"`
    Major string `json:"major"`
}

// StudentRequest مدل درخواست برای ایجاد/بروزرسانی دانشجو
type StudentRequest struct {
    Name  string `json:"name" binding:"required,min=2,max=50"`
    Age   int    `json:"age" binding:"required,min=16,max=100"`
    Email string `json:"email" binding:"required,email"`
    Major string `json:"major"`
}

// UpdateStudentRequest مدل درخواست برای بروزرسانی دانشجو (فیلدهای اختیاری)
type UpdateStudentRequest struct {
    Name  *string `json:"name,omitempty" binding:"omitempty,min=2,max=50"`
    Age   *int    `json:"age,omitempty" binding:"omitempty,min=16,max=100"`
    Email *string `json:"email,omitempty" binding:"omitempty,email"`
    Major *string `json:"major,omitempty"`
}