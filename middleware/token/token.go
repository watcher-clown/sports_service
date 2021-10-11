package token

import (
	"github.com/gin-gonic/gin"
	"github.com/garyburd/redigo/redis"
	"net/http"
	"sports_service/server/global/app/errdef"
	"sports_service/server/global/consts"
	"sports_service/server/models/muser"
	"strings"
	"sports_service/server/global/app/log"
	"time"
	"fmt"
)

// token校验
func TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		reply := errdef.New(c)
		var userid, hashcode, auth string
		c.Set(consts.USER_ID, userid)
		val, err := c.Request.Cookie(consts.COOKIE_NAME)
		if err != nil {
			auth = c.Request.Header.Get("auth")
			if auth == "" {
				log.Log.Errorf("c.Request.Cookie() err is %s", err.Error())
				reply.Response(http.StatusUnauthorized, errdef.INVALID_TOKEN)
				c.Abort()
				return
			}
		} else {
			auth = val.Value
		}

		log.Log.Debugf("auth:%v", auth)
		//v := c.Request.Header.Get("auth")
		ks := strings.Split(auth, "_")
		if len(ks) != 2 {
			log.Log.Errorf("len(ks) != 2")
			ks = strings.Split(auth, "%09")
		}

		if len(ks) == 2 {
			userid = ks[0]
			hashcode = ks[1]
		}

		if len(hashcode) <= 0 {
			log.Log.Errorf("len(hashcode) <= 0")
			reply.Response(http.StatusUnauthorized, errdef.INVALID_TOKEN)
			c.Abort()
			return
		}

		log.Log.Debugf("token_trace: token userid:%s", userid)
		model := new(muser.UserModel)
		token, err := model.GetUserToken(userid)
		if err != nil && err == redis.ErrNil {
			log.Log.Errorf("token_trace: get user token by redis err:%s", err)
			reply.Response(http.StatusUnauthorized, errdef.INVALID_TOKEN)
			c.Abort()
			return
		}

		// 客户端token是否和redis存储的一致
		if res := strings.Compare(auth, token); res != 0 {
			log.Log.Errorf("token_trace: token not match, server token:%s, client token:%s", token, auth)
			reply.Response(http.StatusUnauthorized, errdef.INVALID_TOKEN)
			c.Abort()
			return
		}

		log.Log.Debugf("client token:%s, server token:%s", auth, token)

		if userid != "" {
			// 给token续约
			if err := model.SaveUserToken(userid, auth); err != nil {
				log.Log.Errorf("token_trace: save user token err:%s", err)
			}

			// todo:
			go RecordInfo(model, userid)
		}

		c.Set(consts.USER_ID, userid)
		c.Next()
	}
}

func RecordInfo(model *muser.UserModel, userid string) {
	condition := fmt.Sprintf("user_id=%s", userid)
	cols := "last_login_time, update_at"
	now := int(time.Now().Unix())
	model.User.LastLoginTime = now
	model.User.UpdateAt = now
	if _, err := model.UpdateUserInfos(condition, cols); err != nil {
		log.Log.Errorf("token_trace: update login time fail, userId:%s, err:%s", userid, err)
	}

	if _, err := model.AddActivityRecord(userid, now); err != nil {
		log.Log.Errorf("token_trace: record activity user fail, userId:%s, err:%s", userid, err)
	}
}
