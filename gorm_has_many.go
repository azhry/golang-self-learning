package main

import (
  "github.com/jinzhu/gorm"
  "time"
)

type (
  TicketOrders struct {
    basemodel.BaseModel
    DeletedAt *time.Time `json:"deleted_at" sql:"index"`
    TransactionID int `json:"transaction_id" gorm:"column:transaction_id"`
    TicketID int `json:"ticket_id" gorm:"column:ticket_id"`
    Code string `json:"code" gorm:"column:code"`
  }

  TicketTransactions struct {
    config.BaseModel
    UserID int `json:"user_id" gorm:"column:user_id"`
    TotalAmount int `json:"total_amount" gorm:"column:total_amount"`
    PaymentType string `json:"payment_type" gorm:"column:payment_type"`
    Orders []TicketOrders `json:"orders" gorm:"foreignkey:transaction_id"` // association field
    DeletedAt *time.Time `json:"deleted_at" sql:"index"`
  }
)

func SelectTicketTransactions() ([]TicketTransactions, error) {
  transactions := []TicketTransactions{}
  err := config.App.DB.Preload("Orders").Find(&transactions).Error

  return transactions, err
}
