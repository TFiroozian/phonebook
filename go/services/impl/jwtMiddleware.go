package impl

import (
	// Go native packages
	"errors"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	// Our packages
	"github.com/tfiroozian/phonebook/go/env"

	// Dep packages
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// This is just a persian poem!
const secretKey = "ﺲﺒﺗ ﺲﻠﻣی ﺐﺻﺪﻏیﻩﺍ ﻑﺅﺍﺩی"

type CustomClaims struct {
	RandomNumber int `json:"random_number"`
	jwt.StandardClaims
}

type MiddlewaresImpl struct{}

func (m *MiddlewaresImpl) GenerateToken(userId int64) (string, error) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// Expire token every week
	claims := CustomClaims{
		r1.Intn(10000),
		jwt.StandardClaims{
			Subject:   strconv.FormatInt(userId, 10),
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	return token.SignedString([]byte(secretKey))
}

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		st := strings.Split(token, "Bearer ")

		if token == "" || len(st) == 1 {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		t, err := jwt.ParseWithClaims(st[1], &CustomClaims{},
			func(token *jwt.Token) (interface{}, error) {
				if token.Method.Alg() != jwt.SigningMethodHS256.Name {
					return nil, errors.New("Not valid signing method for token")
				}

				return []byte(secretKey), nil

			})

		claims, ok := t.Claims.(*CustomClaims)
		if t.Valid && ok {
			userId, err2 := strconv.ParseInt(claims.StandardClaims.Subject, 10, 64)
			if err2 != nil {
				params := make(map[string]interface{})
				params["user-id"] = claims.StandardClaims.Subject
				msg := "strconv.ParseInt 64: " + err2.Error()
				env.Environment.Logger.WithFields(logrus.Fields{"params": params}).Error(msg)
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}

			c.Set("user_id", userId)
			c.Next()
		} else if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&(jwt.ValidationErrorExpired) != 0 {
				if time.Now().AddDate(0, 0, -7).Unix() > claims.StandardClaims.ExpiresAt {
					env.Environment.Logger.Info("Token is expired after a week")
					c.AbortWithStatus(http.StatusUnauthorized)
					return
				}
			}
		} else {
			env.Environment.Logger.Error(err.Error())
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}
}
