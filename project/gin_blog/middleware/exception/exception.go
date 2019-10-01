package exception

import (
	"runtime/debug"

	"gin_blog/pkg/e"
	"gin_blog/pkg/util"

	"github.com/gin-gonic/gin"
)



func SetUp() {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				DebugStack := ""
				for _, v := range strings.Split(string(debug.Stack()), "\n") {
					DebugStack += v + "<br>"
				}

				var buffer bytes.Buffer


				buffer.WriteString("err:" + fmt.Sprintf("%s", err))
				buffer.WriteString("\\n")
				//buffer.WriteString("RequestTime:" + util.GetCurrentDate())
				buffer.WriteString("RequestURL:" + c.Request.Method + "  " + c.Request.Host + c.Request.RequestURI)
				buffer.WriteString("RequestUA:" + c.Request.UserAgent())
				buffer.WriteString("RequestIP:" + c.ClientIP())
				buffer.WriteString("DebugStack:" + DebugStack)
				buffer.WriteString("RequestTime:" + util.GetCurrentDate())
				fmt.Print(buffer.String())

				fmt.Fprint(out, buffer.String())

				util.Output(c, http.StatusInternalServerError, e.ERROR, nil)
			}
		}()
		c.Next()
	}
}