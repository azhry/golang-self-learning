package main

import (
  "github.com/thedevsaddam/govalidator"
  "github.com/jinzhu/gorm"
  "net/http"
)

type (
  TicketOrders struct {
    TransactionID int `json:"transaction_id" gorm:"column:transaction_id" example:"1"`
    TicketID int `json:"ticket_id" gorm:"column:ticket_id" example:"1"`
    Code string `json:"code" gorm:"column:code" example:"ABCDE"`
  }

  TicketTransactions struct {
    UserID int `json:"user_id" gorm:"column:user_id" example:"2"`
    TotalAmount int `json:"total_amount" gorm:"column:total_amount" example:"42069"`
    PaymentType string `json:"payment_type" gorm:"column:payment_type" example:"debit"`
    Orders []TicketOrders `json:"orders" gorm:"foreignkey:transaction_id"`
  }
)

func TicketTransactionAdd(c echo.Context) error {
  defer c.Request().Body.Close()

  transaction := raws.TicketTransactions{}
  payloadRules := govalidator.MapData{
    "user_id": []string{"required", "numeric"},
    "payment_type": []string{"required", "in:debit,credit,gopay,ovo,dana,linkaja"},
    "total_amount": []string{"required", "numeric"},
  }

  validate := config.ValidateRequestPayload(c, payloadRules, &transaction)
  if validate != nil {
    return config.ReturnInvalidResponse(http.StatusUnprocessableEntity, validate, "validation error")
  }

  tx := config.App.DB.Begin()
  if err := tx.Create(&transaction).Error; err != nil {
    tx.Rollback()
    return config.ReturnInvalidResponse(http.StatusUnprocessableEntity, err, "Create transaction failed")
  }
  tx.Commit()

  return c.JSON(http.StatusOK, transaction)
}
