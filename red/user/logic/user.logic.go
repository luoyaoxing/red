package logic

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"gitlab.kay.com/config"
	"gitlab.kay.com/logger"
	"gitlab.kay.com/red/common"
	"gitlab.kay.com/red/po/user"
	"gitlab.kay.com/red/proto/user"
	"xorm.io/core"
)

type UserLogic struct {
	engine *xorm.Engine
}

func NewUserLogic() (*UserLogic, error)  {
	logic := new(UserLogic)

	engine, err := xorm.NewEngine("mysql", config.Get("mysql", "xproject"))
	if err != nil {
		logger.Errorf("xorm NewEngine err:%s", err.Error())
		return nil, err
	}

	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "t_")
	fieldMapper := core.NewPrefixMapper(core.SnakeMapper{}, "f")

	engine.SetTableMapper(tbMapper)
	engine.SetColumnMapper(fieldMapper)

	logic.engine = engine
	return logic, nil
}

func (u *UserLogic) CreateUser(ctx context.Context, req *red_proto_user.CreateUserRequest, resp *red_proto_user.CreateUserResponse) error  {
	logger.Infof("UserLogic CreateUser start req:%s", req.String())

	header := new(red_proto_user.ResponseHeader)
	header.Code = common.ERR_SUCCESS.GetErrCode()
	header.Msg = common.ERR_SUCCESS.GetErrMsg()
	resp.RespHeader = header

	user := new(user.User)
	user.UserName = req.GetUserName()
	user.NickName = req.GetNickName()

	uid, err := u.engine.Table("t_user").Insert(user)
	if err != nil {
		logger.Errorf("CreateUser Insert err:%s", err.Error())
		header.Code = common.ERR_OPERAROR.GetErrCode()
		header.Msg = common.ERR_OPERAROR.GetErrMsg()
		return nil
	}

	resp.UserId = uint64(uid)
	logger.Infof("UserLogic CreateUser success")
	return nil
}

func (u *UserLogic) GetUser(ctx context.Context, req *red_proto_user.GetUserRequest, resp *red_proto_user.GetUserResponse) error  {
	logger.Infof("handler GetUser start req:%s", req.String())

	header := new(red_proto_user.ResponseHeader)
	header.Code = common.ERR_SUCCESS.GetErrCode()
	header.Msg = common.ERR_SUCCESS.GetErrMsg()
	resp.RespHeader = header

	user := new(user.User)
	user.UserId = req.GetUserId()

	ok, err := u.engine.Table("t_user").Get(user)
	if !ok || err != nil {
		logger.Errorf("GetUser Get err:%s", err.Error())
		header.Code = common.ERR_OPERAROR.GetErrCode()
		header.Msg = common.ERR_OPERAROR.GetErrMsg()
		return nil
	}

	resp.User = user.ToPROTO()
	logger.Infof("handler GetUser success")
	return nil
}
