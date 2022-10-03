package accounts

type createAccountRequest struct {
	AccountName string `json:"accountName" binding:"required"`
	DOB         string `json:"dob" binding:"required"`
	Address     string `json:"address"`
	Description string `json:"description"`
}
