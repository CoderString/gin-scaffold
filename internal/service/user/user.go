package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"ship/internal/entities/dto"
	"ship/pkg/reply"
)

func (s *ServiceUser) Add() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user dto.UserDTO
		if err := c.BindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if err := s.BizUser.Add(context.Background(), &user); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		reply.Success(c, nil)
	}
}

func (s *ServiceUser) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user dto.UserDTO
		if err := c.BindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if err := s.BizUser.Delete(context.Background(), user.ID); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		reply.Success(c, nil)
	}
}
