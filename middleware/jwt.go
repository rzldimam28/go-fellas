package middleware

import (
	"context"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/rzldimam28/wlb-test/model/helper"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GenerateToken(userId primitive.ObjectID) (string, error) {

	err := godotenv.Load()
	helper.PanicIfError(err)

	sign := jwt.New(jwt.GetSigningMethod("HS256"))

	claims := sign.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(60 * time.Minute).Unix()

	token, err := sign.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", errors.New("could not generate token")
	}

	return token, nil
}

func Auth(next http.Handler) http.Handler {
	err := godotenv.Load()
	helper.PanicIfError(err)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Auth")
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if t.Method != jwt.GetSigningMethod("HS256") {
				return nil, errors.New("unexpected signing method")
			}
			return []byte(os.Getenv("SECRET_KEY")), nil
		})
		if token != nil && err == nil {
			myToken := token.Claims.(jwt.MapClaims)["userId"]
			uid := myToken.(string)
			userId, err := primitive.ObjectIDFromHex(uid)
			helper.PanicIfError(err)
			userCtx := context.WithValue(r.Context(), "userId", userId)
			next.ServeHTTP(w, r.WithContext(userCtx))
		} else {
			webResponse := helper.CreateWebResponse(http.StatusUnauthorized, "Unauth", nil)
			helper.WriteToResponseBody(w, webResponse)
		}
	})
}