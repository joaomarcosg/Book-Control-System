package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/joaomarcosg/Book-Control-System/internal/loans/models"
)

type LoanController struct {
	loanService models.LoanService
}

func NewLoanController(loanService models.LoanService) *LoanController {
	return &LoanController{
		loanService: loanService,
	}
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
