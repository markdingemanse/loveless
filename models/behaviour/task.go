package models;

import (
    "github.com/jinzhu/gorm"
);

type Task struct {
  id uint `gorm:"primary_key"`
  task string
}
