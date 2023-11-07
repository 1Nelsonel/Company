package model

type Company struct {
    ID       uint   `gorm:"primaryKey" json:"id"`
    Name     string `json:"name"`
    Location string `json:"location"`
    Email    *string `json:"email"`
}

type Product struct {
    ID        uint    `gorm:"primaryKey" json:"id"`
    Name      string `json:"name"`
    Price     float64 `json:"price"`
    CompanyID uint   `json:"company_id"`
    Company   Company `gorm:"foreignKey:CompanyID" json:"company"`
}
