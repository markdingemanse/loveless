package models;

import (
    "github.com/jinzhu/gorm"
);

type Word struct {
  id uint `gorm:"primary_key"`
  word string
}
