package handlers

import (
	"log"
	"net/http"

	"github.com/coffemanfp/yofiotest/assigners"
	"github.com/coffemanfp/yofiotest/database"
	"github.com/gin-gonic/gin"
)

// CreateAssignment handler to create a assignment.
func CreateAssignment(db database.AssignersDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var assigner assigners.Assigner

		err := c.Bind(&assigner)
		if err != nil || assigner.Investment == 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, newErrorMessage("invalid body"))
			return
		}

		assigner.Assign(assigner.Investment)

		newAssigner, err := db.Create(assigner)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, newErrorMessage("Oops!"))
			return
		}
		newAssigner = assigners.Assigner{
			ID:            newAssigner.ID,
			Success:       assigner.Success,
			Investment:    assigner.Investment,
			CreditType300: assigner.CreditType300,
			CreditType500: assigner.CreditType500,
			CreditType700: assigner.CreditType700,
		}

		if newAssigner.Success {
			c.JSON(http.StatusOK, newAssigner)
		} else {
			c.JSON(http.StatusBadRequest, newAssigner)
		}
	}
}

// GetStats handler to get the assignments statistics.
func GetStats(db database.AssignersDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		stats, err := db.GetStats()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, newErrorMessage("Oops!"))
			return
		}

		c.JSON(http.StatusOK, stats)
	}
}
