package models;

import (
    "github.com/jinzhu/gorm"
);

type Phrase struct {
  id uint `gorm:"primary_key"`
  phrase string
}
