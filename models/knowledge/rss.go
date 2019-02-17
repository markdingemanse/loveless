package models;

import (
    "github.com/jinzhu/gorm"
);

type Rss struct {
  gorm.Model
  FirstPost string
}

type RedifyURI struct {
    gorm.Model
    Uri string
    Slug string
}

// get the first_post from the model Rss
func (rss Rss) GetFirstPost() string {
    return rss.FirstPost
}

func CreateRedifyRssModel() Rss {
    return Rss{};
}

func CreateRedifyUriModels() []RedifyURI {
    return []RedifyURI{};
}
