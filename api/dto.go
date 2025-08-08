package api

type SendRequest struct {
	From   string  `json:"from" binding:"required"`
	To     string  `json:"to" binding:"required"`
	Amount float64 `json:"amount" binding:"required,gt=0"`
}

type SendResponse struct {
	Message string `json:"message"`
}

type BalanceResponse struct {
	Balance float64 `json:"balance"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
