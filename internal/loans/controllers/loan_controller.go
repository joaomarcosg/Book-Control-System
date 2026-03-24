package controllers

import "github.com/gin-gonic/gin"

type LoanController struct{}

func NewLoanController() *LoanController {
	return &LoanController{}
}

func (l *LoanController) RegisterRoutes(r *gin.Engine) {
	loans := r.Group("/loans")

	{
		loans.POST("", l.CreateLoan)
		loans.GET("/:id", l.GetLoan)
		loans.GET("", l.GetAllLoans)
	}

}

func (l *LoanController) CreateLoan(ctx *gin.Context) {
}

func (l *LoanController) GetLoan(ctx *gin.Context) {
}

func (l *LoanController) GetAllLoans(ctx *gin.Context) {
}
