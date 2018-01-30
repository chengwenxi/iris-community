// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package authz

import (
	"net/http"

	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
	"github.com/irisnet/iris-community/models"
	"time"
)

// NewAuthorizer returns the authorizer, uses a Casbin enforcer as input
func NewAuthorizer(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		a := &BasicAuthorizer{enforcer: e}
		r := c.Request
		user, status := a.GetUserRole(r)
		method := r.Method
		path := r.URL.Path
		if !a.enforcer.Enforce(user, path, method) {
			if status == 401 {
				c.Abort() //stop the current handler and not execute next
				a.RequireAuthorized(c.Writer)
				return
			}
			c.Abort() //stop the current handler and not execute next
			a.RequirePermission(c.Writer)
			return
		}
	}
}

// BasicAuthorizer stores the casbin handler
type BasicAuthorizer struct {
	enforcer *casbin.Enforcer
}

// GetUserName gets the user name from the request.
// Currently, only HTTP basic authentication is supported
func (a *BasicAuthorizer) GetUserRole(r *http.Request) (string, int) {
	if authorization := r.Header.Get("Authorization"); authorization == "" {
		return "guest", 0 //访客用户
	} else {
		userAuth := &models.UserAuth{
			AuthCode: authorization,
		}
		userAuth.FindByAuth()
		if userAuth.Id != 0 {
			expiresIn := time.Now().Sub(userAuth.Updatetime).Seconds()
			if expiresIn > float64(userAuth.ExpiresIn) {
				return "guest", 401 //访客用户,授权超时
			}
			user := &models.Users{
				Id: uint(userAuth.UserId),
			}
			user.First()
			return "user", 0 //普通用户
		} else {
			return "guest", 401 //访客用户
		}
	}
	return "guest", 0 //访客用户
}

// RequirePermission returns the 403 Forbidden to the client
func (a *BasicAuthorizer) RequirePermission(w http.ResponseWriter) {
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte(http.StatusText(http.StatusForbidden)))
}

// Unauthorized returns the 401 Unauthorized to the client
func (a *BasicAuthorizer) RequireAuthorized(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
}
