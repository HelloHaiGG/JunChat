package servers

import (
	"JunChat/common"
	"JunChat/common/iredis"
	"JunChat/core/models"
	core "JunChat/core/protocols"
	"JunChat/utils"
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	jsoniter "github.com/json-iterator/go"
	"strconv"
	"time"
)

type UserController struct{}

func (p *UserController) UserLogin(ctx context.Context, in *core.LoginParams) (*core.LoginRsp, error) {
	user := models.UserInfo{}
	status := common.Success
	var err error
	if in.Uid == "" && in.UName == "" {
		return &core.LoginRsp{
			Code: common.MissUserInfo,
		}, nil
	}
	if in.UName == "" {
		user.Uid = in.Uid
		err = user.FindByUid()
	}
	if in.Uid == "" {
		user.UserName = in.UName
		err = user.FindByName()
	}

	if err == gorm.ErrRecordNotFound {
		status = common.UserDoesNotExist
	} else if err != nil {
		status = common.InternalErr
	}

	if status != common.Success {
		return &core.LoginRsp{
			Code: int32(status),
		}, nil
	}
	err = user.LoginByUid()
	if err == gorm.ErrRecordNotFound {
		status = common.UserDoesNotExistOrPasswordErr
	} else if err != nil {
		status = common.InternalErr
	}
	if status != common.Success {
		return &core.LoginRsp{
			Code: int32(status),
		}, nil
	}

	//调度
	serverId, err := Dispatch(user.Uid)
	if err != nil {
		return &core.LoginRsp{
			Code: common.InternalErr,
		}, nil
	}

	token := &models.TokenEntity{Info: &user, ServerId: serverId, TimeStamp: time.Now().Unix()}
	str, _ := jsoniter.MarshalToString(token)

	session,_ := utils.AesEncrypt([]byte(str),utils.KEY)

	err = SetToken(session, user.Uid)
	if err != nil {
		return &core.LoginRsp{
			Code: common.InternalErr,
		}, nil
	}

	return &core.LoginRsp{
		Code:     common.Success,
		ServerId: serverId,
		Name:     user.UserName,
		Token:    session,
	}, nil

}

func (p *UserController) RegisterUser(ctx context.Context, in *core.RegisterParams) (*core.RegisterRsp, error) {

	var err error
	if in.UName == "" || in.Password == "" {
		return &core.RegisterRsp{Code: common.MissUserInfo}, nil
	}
	user := &models.UserInfo{UserName: in.UName, Uid: in.Password}
	err = user.FindByName()
	if err != nil && err != gorm.ErrRecordNotFound {
		return &core.RegisterRsp{Code: common.InternalErr}, nil
	} else if err == nil {
		return &core.RegisterRsp{Code: common.UserAlreadyExists}, nil
	}
	sf := utils.SFIdTool.GetID()

	user.UserName = in.UName
	user.Password = in.Password
	user.Uid = strconv.FormatInt(sf, 10)
	user.IsLogin = false
	user.CreateAt = time.Now().Unix()

	err = user.Register()
	if err != nil {
		return &core.RegisterRsp{Code: common.InternalErr}, nil
	}
	return &core.RegisterRsp{
		Code: common.Success,
		Uid:  user.Uid,
	}, nil
}

func SetToken(session string, id string) error {
	_, err := iredis.RedisCli.Set(fmt.Sprintf("JUN:CHAT:SESSION:%s", id), session, 30*60*time.Second).Result()
	return err
}
