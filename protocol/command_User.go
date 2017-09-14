/*
author: foolbread
file: protocol/command_User.go
date: 2017/9/12
*/
package protocol

import (
	"github.com/foolbread/fbftp/session"
)

type commandUser struct {

}

func (p *commandUser)IsExtend()bool{
	return false
}

func (p *commandUser)RequireAuth()bool{
	return false
}

func (p *commandUser)RequireParam()bool{
	return true
}

func (p *commandUser)Execute(sess *session.FTPSession, arg string)error{
	return nil
}