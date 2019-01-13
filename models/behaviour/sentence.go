package models;

import (
    "github.com/jinzhu/gorm"
);

type Sentence struct {
  id uint `gorm:"primary_key"`
  sentence string
}
