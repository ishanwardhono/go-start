package repo

import "app/entity"

var mockData = []entity.Article{
	{
		Title:   "test title 1",
		Content: "test article content 1",
		Author:  "test author 1",
	},
	{
		Title:   "test title 2",
		Content: "test article content 2",
		Author:  "test author 2",
	},
}

func GetMockData() []entity.Article {
	return mockData
}
