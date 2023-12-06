package model

type Comic struct {
	ID            string `json:"itemNumber" gorm:"primaryKey"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	ItemCaption   string `json:"itemCaption"`
	LargeImageURL string `json:"largeImageUrl"`
	SalesDate     string `json:"salesDate"`
}
