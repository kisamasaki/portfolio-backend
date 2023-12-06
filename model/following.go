package model

type Following struct {
	ID          string `json:"id" gorm:"primaryKey"`
	FollowingID string `json:"followingID" gorm:"not null"`
	FollowerID  string `json:"followerID" gorm:"not null"`
}
