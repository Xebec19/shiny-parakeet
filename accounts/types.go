package accounts

import "time"

type createAccountRequest struct {
	AccountName string    `json:"accountName" binding:"required"`
	DOB         time.Time `json:"dob" binding:"required"`
	Address     string    `json:"address"`
	Description string    `json:"description"`
}

type readOneAccountRequest struct {
	AccountId string `form:"accountId" binding:"required"`
}
