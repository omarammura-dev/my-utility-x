package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"

	// "myutilityx.com/errors"
	"myutilityx.com/errors"
	"myutilityx.com/models"
	"myutilityx.com/utils"
)



func addExpense(ctx *gin.Context){

	userId,exist := ctx.Get("userId")

	if !exist  {
		ctx.JSON(http.StatusInternalServerError,gin.H{"err": "not exist"})
		return
	}

	var expense models.Expense

	err := ctx.ShouldBindJSON(&expense)

	expense.Price = utils.RoundFloat(expense.Price,2)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError,"Error1")
		return
	}
	expense.UserId, _ = userId.(primitive.ObjectID)
	expense.ExpenseDate = time.Now()

	err = expense.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,errors.ErrInvalidExpenseType)
	}
	ctx.JSON(http.StatusNoContent,"")	
}

func updateExpense(ctx *gin.Context){
	expenseId := ctx.Param("id")

	
	id, err := primitive.ObjectIDFromHex(expenseId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errors.ErrSomethingWentWrong)
		return
	}

	var payload map[string]interface{}
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
	}
	payload["id"] = id

	

	
    if err := models.UpdateExpense(payload); err != nil {
        ctx.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
        return
    }
}

func GetAllExpenses(ctx *gin.Context){
	userId,exist := ctx.Get("userId")

	if !exist {
		ctx.JSON(http.StatusInternalServerError,errors.ErrSomethingWentWrong)
		return
	}

	var expense models.Expense
	expenses,err := expense.GetAllExpenses(userId.(primitive.ObjectID))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError,errors.ErrSomethingWentWrong)
		return		
	}

	ctx.JSON(http.StatusOK,expenses)
}