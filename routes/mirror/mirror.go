package mirror

import (
	"github.com/gogits/gogs/pkg/context"
)

const (
	MIRROR = "mirror/mirror"
)

func Mirror(c *context.Context) {
	c.Data["PageIsMirror"] = true
	c.Data["Name"] = "wuxian"
	c.Success(MIRROR)
}
