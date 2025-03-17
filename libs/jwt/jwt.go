package jwt

// import (
// 	"errors"

// 	"github.com/golang-jwt/jwt/v5"
// )

// func ValidateToken(tokenString string) (*Claims, error) {
// 	//	appConfig := config.GetAppConfig()

// 	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
// 		return []byte("PI App"), nil
// 	})

// 	if err != nil {
// 		if ve, ok := err.(*jwt.ValidationError); ok {
// 			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
// 				//	jwtLogger.LogError("malformed token")
// 				return nil, errors.New("malformed token")
// 			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
// 				//	jwtLogger.LogError("token has expired")
// 				return nil, errors.New("token has expired")
// 			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
// 				//	jwtLogger.LogError("token not valid yet")
// 				return nil, errors.New("token not valid yet")
// 			} else {
// 				//	jwtLogger.LogError("token validation error")
// 				return nil, errors.New("token validation error")
// 			}
// 		}
// 		return nil, err
// 	}

// 	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
// 		return claims, nil
// 	}

// 	return nil, errors.New("invalid token")
// }
