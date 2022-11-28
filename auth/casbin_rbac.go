package auth

//func Casbin() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		_claims, ok := c.Get("claims")
//		if !ok {
//			app.NewResponse().Error(c, errcode.ErrInvalidPermission)
//			c.Abort()
//			return
//		}
//		claims := _claims.(*app.Payload)
//		sub := claims.AuthorityID
//		obj := c.Request.URL.RequestURI()
//		act := c.Request.Method
//
//		// Check policy exists in db
//		e := dao.AuthDao.NewCasbin()
//		status, _ := e.Enforce(sub, obj, act)
//		if status {
//			c.Next()
//		} else {
//			app.NewResponse().Error(c, errcode.ErrInvalidPermission)
//			c.Abort()
//			return
//		}
//	}
//}
