package jwt

func UserAuthMiddleware(a auth.Auther, skipper ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userID string
		if t := ginplus.GetToken(c); t != "" {
			id, err := a.ParseUserID(t)
			if err != nil {
				if err == auth.ErrInvalidToken {
					ginplus.ResError(c, errors.ErrNoPerm)
					return
				}
				ginplus.ResError(c, errors.WithStatck(err))
				return
			}
			userID = id
		}
	}
}
