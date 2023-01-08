package middleware

import (
	"context"
	dto "dewetour/dto/result"
	jwtToken "dewetour/pkg/jwt"
	"encoding/json"
	"net/http"
	"strings"
)

// type Result struct {
// 	Code    int         `json:"code"`
// 	Data    interface{} `json:"data"`
// 	Message string      `json:"message"`
// }

func AuthAdmin(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")

		token := r.Header.Get("Authorization")

		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			response := dto.ErrorResult{Code: http.StatusUnauthorized, Message: "unauthorized"}
			json.NewEncoder(w).Encode(response)
			return
		}

		token = strings.Replace(token, "Bearer ", "", 100)
		claims, err := jwtToken.DecodeToken(token)

		// fmt.Println(claims["role"])

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			response := Result{Code: http.StatusUnauthorized, Message: "unauthorized"}
			json.NewEncoder(w).Encode(response)
			return
		}

		if claims["role"] != "admin" {
			w.WriteHeader(http.StatusUnauthorized)
			response := Result{Code: http.StatusUnauthorized, Message: "only admin can do this action"}
			json.NewEncoder(w).Encode(response)
			return
		}

		ctx := context.WithValue(r.Context(), "userInfo", claims)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
