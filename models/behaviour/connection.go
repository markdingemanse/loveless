package models;

import (
    "github.com/jinzhu/gorm"
);

type Connection struct {
  id uint `gorm:"primary_key"`
  baseline_type string
  baseline_id int
  association_type string
  association_id int
}
