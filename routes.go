package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func initializeRoutes() {
	r.GET("/weapons", func(ctx *gin.Context) {
		weapons, err := AllWeapons()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"weapons": weapons})
	})
	r.GET("/weapons/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		idNum, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		weapon, err := GetWeapon(idNum)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"weapon": weapon})
	})
	r.POST("/weapons", func(ctx *gin.Context) {
		var newWeapon Weapon

		if err := ctx.ShouldBindJSON(&newWeapon); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		AddWeapon(newWeapon)
		ctx.JSON(http.StatusCreated, gin.H{"message": "Weapon created"})
	})
}
