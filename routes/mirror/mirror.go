package mirror

import (
	"fmt"

	"github.com/gogits/gogs/models"
	"github.com/gogits/gogs/pkg/context"
)

const (
	MIRROR = "mirror/mirror"
)

func Mirror(c *context.Context) {
	c.Data["PageIsMirror"] = true
	c.Data["Name"] = c.User.Name
	mirrors, err := models.GetMirrors(1)
	if err != nil {
		fmt.Errorf("can not get mirrors")
	}
	c.Data["Mirrors"] = mirrors

	c.Success(MIRROR)
}

func MirrorStatus(c *context.Context) {
	mirrorId := c.ParamsInt("mirrorId")
	c.JSONSuccess(mirrorId)
}
