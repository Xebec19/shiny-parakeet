package accounts

import (
	"net/http"

	"github.com/Xebec19/shiny-parakeet/util"
	"github.com/gin-gonic/gin"
)

func createAccount(c *gin.Context) {
	var req createAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}
	// args := db.CreateAccountParams{
	// 	AccountName: req.AccountName,
	// 	Dob:         time.Time(req.DOB),
	// 	Address:     sql.NullString{req.Address,true},
	// 	Description: sql.NullString{req.Description,true},
	// 	CreatedBy: ,
	// }
}
