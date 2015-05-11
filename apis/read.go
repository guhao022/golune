package apis

import (
	"fmt"
	"golune/models"
)

type ReadHandler struct {
	BaseHandler
}

func (c *ReadHandler) Login() {
	mgo, err := models.NewDB()
	if err != nil {
		c.Ctx.Redirect(500, "/Err500")
	}
	defer mgo.Close()
	username := c.GetString("username")
	password := c.GetString("password")
	if err := mgo.CheckAdmin(username, password); err == nil {
		c.Data["json"] = map[string]string{"Recode": "200"}
	} else {
		c.Data["json"] = map[string]string{"Recode": "400", "Error": fmt.Sprintln(err)}
	}
	c.ServeJson()
}
