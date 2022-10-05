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

type readManyAccountsRequest struct {
	Page  string `json:"page" binding:"required"`
	Limit string `json:"limit" binding:"required"`
}
