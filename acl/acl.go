/*
author: foolbread
file: acl/acl.go
date: 2017/8/10
*/
package acl

import (
	"github.com/foolbread/fbcommon/golog"
)

func InitAcl(){
	golog.Info("fbftp acl initing......")
}

const (
	ONLY_READ = 1
	ONLY_WRITE = 2
	READ_WRITE = ONLY_READ&ONLY_WRITE
)
