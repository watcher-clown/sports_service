package muser

import (
	"sports_service/server/dao"
	"sports_service/server/global/app/log"
	"sports_service/server/global/consts"
	"sports_service/server/util"
	"github.com/gin-gonic/gin"
	"sports_service/server/global/rdskey"
	"fmt"
	"time"
	"errors"
)

type base struct {

}

// 获取版本号
func (m *base) getVersion(c *gin.Context) string {
	return c.GetHeader("Version")
}

// 获取设备类型
func (m *base) getDeviceType(c *gin.Context) int {
	switch util.GetClient(c.Request.UserAgent()) {
	case util.IPHONE, util.IPad, util.Ios:
		return int(consts.IOS_PLATFORM)
	case util.ANDROID:
		return int(consts.ANDROID_PLATFORM)
	default:
		return int(consts.ANDROID_PLATFORM)
	}
}

// 默认头像
func (m *base) defaultAvatar() string {
	return consts.DEFAULT_AVATAR
}

// 默认昵称
func (m *base) getNickName(nickName string) string {
	if nickName == "" {
		nickName = "FPV用户"
	}

	nickNameNum, err := m.IncrNickNameNum()
	if err != nil {
		log.Log.Errorf("social_trace: user nickname incr err:%s", err)
	}

	if nickNameNum < 10 {
		return fmt.Sprintf("%s0%d%d", nickName, nickNameNum, util.GenerateRandnum(10000, 99999))
	}

	return fmt.Sprintf("%s%d%d", nickName, nickNameNum, util.GenerateRandnum(10000, 99999))
}

func (m *base) IncrNickNameNum() (int64, error) {
	rds := dao.NewRedisDao()
	return rds.INCR(rdskey.USER_NICKNAME_INCR)
}


// 注册
func (m *base) Register(u *UserModel, s *SocialModel, c *gin.Context, unionId, avatar, nickName string, socialType, gender int) error {
	key := rdskey.MakeKey(rdskey.LOGIN_REPEAT, socialType, unionId)
	ok, err:= IsReapeat(key)
	if err != nil {
		log.Log.Errorf("social_trace: redis err %s",err)
		return err
	}

	if !ok {
		log.Log.Errorf("social_trace: 用户重复注册 unionID:%s", unionId)
		return errors.New("social_trace: 用户重复注册")
	}

	rds:=dao.NewRedisDao()
	rds.EXPIRE64(key, rdskey.KEY_EXPIRE_MIN)
	m.newUser(u, c, avatar, nickName, gender)
	m.newSocialAccount(s, socialType, u.User.UserId, unionId)
	return nil
}

// 设置用户社交帐号信息
func (m *base) newSocialAccount(s *SocialModel, socialType int, userid, unionid string) {
	s.SetCreateAt(time.Now().Unix())
	s.SetSocialType(socialType)
	s.SetUnionId(unionid)
	s.SetUserId(userid)
	return
}

// 设置用户信息
func (m *base) newUser(u *UserModel, c *gin.Context, avatar, nickName string, gender int) {
	now := time.Now().Unix()
	m.setDefaultInfo(u, avatar, gender)
	u.SetUserType(consts.TYPE_WECHAT)
	u.SetDeviceType(m.getDeviceType(c))
	u.SetNickName(m.getNickName(nickName))
	u.SetLastLoginTime(now)
	// todo 暂时先使用时间 + 4位随机数 生成uid
	u.SetUid(m.getUserID())
	u.SetCreateAt(now)
	u.SetUpdateAt(now)
	u.SetPassword("")
	return
}

// 设置默认信息
func (m *base) setDefaultInfo(u *UserModel, avatar string, gender int) {
	u.SetAvatar(consts.DEFAULT_AVATAR)
	if avatar != "" {
		u.SetAvatar(avatar)
	}

	if gender != 0 {
		u.SetGender(gender)
	}
}

// uid 8位
func (m *base) getUserID() string {
  uid := fmt.Sprint(util.GetXID())
  if len(uid) == 8 {
    return uid
  }

  rds := dao.NewRedisDao()
  ok, err := rds.EXISTS(rdskey.USER_ID_INCR)
  if err != nil {
    log.Log.Errorf("user_trace: uid incr err:%s", err)
    return uid
  }

  if !ok {
    rds.Set(rdskey.USER_ID_INCR, 10240102)
  }

  num := util.GenerateRandnum(1, 33)
  incrUid, err := rds.INCRBY(rdskey.USER_ID_INCR, int64(num))
  if err != nil {
    log.Log.Errorf("user_trace: uid incr err:%s", err)
  }

  return fmt.Sprint(incrUid)
}
