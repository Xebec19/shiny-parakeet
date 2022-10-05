package accounts

import (
	"database/sql"
	"net/http"
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
	account, err := util.DBQuery.CreateAccount(c, args)
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
	account, err := util.DBQuery.ReadAccount(c, args)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}
	payload := util.ResponseParams{
		Status: true,
		Data:   account,
	}
	c.JSON(http.StatusFound, payload)
}
