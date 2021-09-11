package middleware

//func JwtBlackList() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		jwtToken := c.GetHeader("Authorization")
//		ok, err := service.AuthSvc.CheckJwtIsBlackList(c, jwtToken)
//		if err == nil && ok {
//			app.NewResponse().Error(c, errcode.ErrInvalidToken)
//			c.Abort()
//			return
//		}
//	}
//}
