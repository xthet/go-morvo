package auth

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/xthet/go-morvo/models"
	"github.com/xthet/go-morvo/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type contextKey string

const UserKey contextKey = "userID"

func JWTAuth(handler_func http.HandlerFunc, collection models.UserCollection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token_string := utils.GetTokenFromRequest(r)
		token, err := validate_JWT(token_string)
		if err != nil || !token.Valid {
			permission_denied(w)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		str := claims["userID"].(string)

		id, err := primitive.ObjectIDFromHex(str)
		if err != nil {
			permission_denied(w)
			return
		}

		u, err := collection.GetUserByID(id)
		if err != nil {
			permission_denied(w)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, UserKey, u.ID)
		r = r.WithContext(ctx)

		handler_func(w, r)
	}
}

func CreateJWT(secret []byte, userID primitive.ObjectID) (string, error) {
	expsec, err := strconv.ParseInt(os.Getenv("JWT_EXPIRATION_SECS"), 10, 64)
	if err != nil {
		return "", err
	}
	expiration := time.Second * time.Duration(expsec)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    userID.Hex(),
		"expiresAt": time.Now().Add(expiration).Unix(),
	})

	token_string, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return token_string, err
}

func validate_JWT(token_string string) (*jwt.Token, error) {
	return jwt.Parse(token_string, func(token *jwt.Token) (interface{}, error) { // this function returns the secret so it can be parsed by jwt.Parse()
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})
}

func permission_denied(w http.ResponseWriter) {
	utils.WriteError(w, http.StatusForbidden, fmt.Errorf("unauthorized"))
}

func GetUserIDFromContext(ctx context.Context) int {
	userID, ok := ctx.Value(UserKey).(int)
	if !ok {
		return -1
	}

	return userID
}
