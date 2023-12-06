package model

import "time"

type CompletedComic struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Rating    uint      `json:"rating"`
	Review    string    `json:"review" gorm:"size:100"`
	CreatedAt time.Time `json:"createdAt"`
	User      User      `json:"user" gorm:"foreignKey:UserID; constraint:OnDelete:CASCADE"`
	UserID    string    `json:"userId" gorm:"not null;uniqueIndex:unique_user_comic"`
	Comic     Comic     `json:"comic" gorm:"foreignKey:ComicID; constraint:OnDelete:CASCADE"`
	ComicID   string    `json:"comicId" gorm:"not null;uniqueIndex:unique_user_comic"`
}
