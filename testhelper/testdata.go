package testhelper

import (
	"portfolio-backend/model"
	"time"
)

var TestUsers = map[string]model.User{
	"testuser1": {ID: "testuser1", UserName: "TestUser1"},
	"testuser2": {ID: "testuser2", UserName: "TestUser2"},
	"testuser3": {ID: "testuser3", UserName: "TestUser3"},
}

var TestUserAuths = map[string]model.UserAuth{
	"testuser1": {UserID: "testuser1"},
	"testuser2": {UserID: "testuser2"},
	"testuser3": {UserID: "testuser3"},
}

var TestComic = model.Comic{
	ID:            "123",
	Title:         "Sample Comic",
	Author:        "Sample Author",
	ItemCaption:   "あらすじ",
	LargeImageURL: "http://example.com/comic.jpg",
	SalesDate:     "2023年01月01日",
}

var TestComics = []model.Comic{
	{ID: "1", Title: "Comic 1"},
	{ID: "2", Title: "Comic 2"},
}

var TestCompletedComic = model.CompletedComic{
	ID:        "12345",
	Rating:    5,
	Review:    "記録用",
	CreatedAt: time.Now(),
	User:      TestUsers["testuser1"],
	UserID:    TestUsers["testuser1"].ID,
	Comic:     TestComic,
	ComicID:   TestComic.ID,
}

var TestComicResponseJSON = `{
	"Items": [
		{
			"Item": {
				"itemNumber": "123",
				"Title": "Sample Comic",
				"Author": "Sample Author",
				"ItemCaption": "あらすじ",
				"LargeImageURL": "http://example.com/comic.jpg",
				"SalesDate": "2023年01月01日"
			}
		}
	]
}`
