package models

type Product struct {
	Id    int64   `gorm:"primaryKey;autoIncrement" json:"id"`
	Nama  string  `gorm:"type:varchar(255)" json:"nama"`
	Stok  int32   `gorm:"type:int(5)" json:"stok"`
	Harga float64 `gorm:"type:decimal(14,2)" json:"harga"`
}
