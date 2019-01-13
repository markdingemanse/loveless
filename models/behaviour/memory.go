package models;

import (
    "github.com/jinzhu/gorm"
);

type Memory struct {
  id uint `gorm:"primary_key"`
  initiator string
  payload string
  response string
  connections string
  collection string
}
