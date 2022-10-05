package accounts

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	db "github.com/Xebec19/shiny-parakeet/db/sqlc"
	"github.com/Xebec19/shiny-parakeet/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// create a account
func createAccount(c *gin.Context) {
	var req createAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}
	userId := uuid.MustParse(c.MustGet("userId").(string))
	args := db.CreateAccountParams{
		AccountName: req.AccountName,
		Dob:         time.Time(req.DOB),
		Address:     sql.NullString{String: req.Address, Valid: true},
		Description: sql.NullString{String: req.Description, Valid: true},
		CreatedBy:   userId,
	}
	account, err := db.DBQuery.CreateAccount(c, args)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}
	payload := util.ResponseParams{
		Status: true,
		Data:   account,
	}
	c.JSON(http.StatusCreated, payload)
}

func readOneAccount(c *gin.Context) {
	var req readOneAccountRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	userId := uuid.MustParse(c.MustGet("userId").(string))
	args := db.ReadAccountParams{
		AccountID: uuid.MustParse(req.AccountId),
		CreatedBy: userId,
	}
	account, err := db.DBQuery.ReadAccount(c, args)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}
	payload := util.ResponseParams{
		Status: true,
		Data:   account,
	}
	c.JSON(http.StatusOK, payload)
}

func readManyAccounts(c *gin.Context) {
	var req readManyAccountsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	userId := uuid.MustParse(c.MustGet("userId").(string))
	page, _ := strconv.ParseInt(req.Page, 10, 32)
	limit, _ := strconv.ParseInt(req.Limit, 10, 32)
	offset := page * limit
	args := db.ReadAllAccountParams{
		CreatedBy: userId,
		Offset:    int32(offset),
		Limit:     int32(limit),
	}
	accounts, err := db.DBQuery.ReadAllAccount(c, args)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}
	payload := util.ResponseParams{
		Status: true,
		Data:   accounts,
	}
	c.JSON(http.StatusFound, payload)
}

func updateAccount(c *gin.Context) {
	var req updateAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}
	userId := uuid.MustParse(c.MustGet("userId").(string))

	args := db.UpdateAccountParams{
		AccountName: req.AccountName,
		Dob:         time.Time(req.DOB),
		Address:     sql.NullString{String: req.Address, Valid: true},
		Description: sql.NullString{String: req.Description, Valid: true},
		AccountID:   uuid.MustParse(req.AccountId),
		CreatedBy:   userId,
	}
	account, err := db.DBQuery.UpdateAccount(c, args)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}
	payload := util.ResponseParams{
		Status: true,
		Data:   account,
	}
	c.JSON(http.StatusOK, payload)
}

func deleteAccount(c *gin.Context) {
	var req deleteAccountRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	userId := uuid.MustParse(c.MustGet("userId").(string))
	args := db.DeleteAccountParams{
		AccountID: uuid.MustParse(req.AccountId),
		CreatedBy: userId,
	}

	db.DBQuery.DeleteAccount(c, args)
	payload := util.ResponseParams{
		Status: true,
		Data:   args.AccountID,
	}
	c.JSON(http.StatusOK, payload)
}
