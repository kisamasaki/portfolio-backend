package model

import "time"

type User struct {
	ID         string      `json:"id" gorm:"primaryKey"`
	UserName   string      `json:"userName" gorm:"not null"`
	ImageURL   string      `json:"imageUrl" gorm:"not null"`
	Auth       UserAuth    `json:"auth" gorm:"foreignKey:UserID"`
	Followings []Following `json:"followings" gorm:"foreignKey:FollowingID"`
	Followers  []Following `json:"followers" gorm:"foreignKey:FollowerID"`
	CreatedAt  time.Time   `json:"createdAt"`
}

type UserAuth struct {
	ID       string `json:"id" gorm:"primaryKey"`
	UserID   string `json:"userId" gorm:"primaryKey;unique"`
	Password string `json:"password" gorm:"not null"`
}
