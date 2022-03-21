/*
 *
 *  * Copyright (C) 2022 The orange protocol Authors
 *  * This file is part of The orange library.
 *  *
 *  * The Orange is free software: you can redistribute it and/or modify
 *  * it under the terms of the GNU Lesser General Public License as published by
 *  * the Free Software Foundation, either version 3 of the License, or
 *  * (at your option) any later version.
 *  *
 *  * The orange is distributed in the hope that it will be useful,
 *  * but WITHOUT ANY WARRANTY; without even the implied warranty of
 *  * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *  * GNU Lesser General Public License for more details.
 *  *
 *  * You should have received a copy of the GNU Lesser General Public License
 *  * along with The orange.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/orange-protocol/orange-server-v1/jwt"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			// Allow unauthenticated users in
			if header == "" {
				next.ServeHTTP(w, r)
				//http.Error(w, "need login first", http.StatusForbidden)
				return
			}

			//validate jwt token
			tokenStr := header
			did, err := jwt.ParseToken(tokenStr)
			if err != nil {
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintln(w, `{"message":"invalid jwt token, please login again"}`)
				//http.Error(w, `{"message":"invalid jwt token, please login again"}`, http.StatusForbidden)
				return
			}

			// put it in context
			ctx := context.WithValue(r.Context(), userCtxKey, did)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) string {
	raw, _ := ctx.Value(userCtxKey).(string)
	return raw
}

func CheckLogin(ctx context.Context, did string) error {
	if did != ForContext(ctx) {
		return fmt.Errorf("did is not match,please login first")
	}
	return nil
}
