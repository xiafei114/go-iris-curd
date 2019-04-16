package routes

import (
	"go-iris-curd/main/middleware/jwts"
	"go-iris-curd/main/utils"
	modeSys "go-iris-curd/main/web/models/system"
	"go-iris-curd/main/web/supports"
	"go-iris-curd/main/web/supports/vo"
	"strconv"
	"strings"
	"time"

	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
)

// UserHub 用户路由
func UserHub(party iris.Party) {
	var user = party.Party("/user")
	//p.Use(middleware.BasicAuth)s
	user.Post("/registe", hero.Handler(Registe))
	user.Post("/login", hero.Handler(Login))
	user.Post("/logout", hero.Handler(LoginOut))
}

// Registe 注册
func Registe(ctx iris.Context) {
	var (
		err    error
		user   = new(modeSys.User)
		effect int64
	)

	if err = ctx.ReadJSON(&user); err != nil {
		ctx.Application().Logger().Errorf("用户[%s]注册失败。%s", user.Username, err.Error())
		goto FAIL
	}

	user.Password = utils.AESEncrypt([]byte(user.Password))
	user.Enable = true
	user.CreateTime = time.Now()

	effect, err = modeSys.CreateUser(user)
	if effect <= 0 || err != nil {
		ctx.Application().Logger().Errorf("用户[%s]注册失败。%s", user.Username, err.Error())
		goto FAIL
	}

	supports.OkR(ctx, supports.RegisteSuccess)
	return

FAIL:
	supports.Error(ctx, iris.StatusInternalServerError, supports.RegisteFailur, nil)
	return
}

// LoginOut 登出
func LoginOut(ctx iris.Context) {
	supports.Ok(ctx, supports.LoginSuccess, nil)
	return
}

// Login 登录
func Login(ctx iris.Context) {
	var (
		err        error
		user       = new(modeSys.User)
		mUser      = new(modeSys.User) // 查询数据后填充了数据的user
		ckPassword bool
		token      string
	)

	if err = ctx.ReadJSON(&user); err != nil {
		ctx.Application().Logger().Errorf("用户[%s]登录失败。%s", "", err.Error())
		supports.Error(ctx, iris.StatusBadRequest, supports.LoginFailur, nil)
		return
	}

	mUser.Username = user.Username
	has, err := modeSys.GetUserByUsername(mUser)
	if err != nil {
		ctx.Application().Logger().Errorf("用户[%s]登录失败。%s", user.Username, err.Error())
		supports.Error(ctx, iris.StatusInternalServerError, supports.LoginFailur, nil)
		return
	}

	if !has { // 用户名不正确
		supports.Unauthorized(ctx, supports.UsernameFailur, nil)
		return
	}

	ckPassword = utils.CheckPWD(user.Password, mUser.Password)
	if !ckPassword {
		supports.Unauthorized(ctx, supports.PasswordFailur, nil)
		return
	}

	token, err = jwts.GenerateToken(mUser)
	golog.Infof("用户[%s], 登录生成token [%s]", mUser.Username, token)
	if err != nil {
		ctx.Application().Logger().Errorf("用户[%s]登录，生成token出错。%s", user.Username, err.Error())
		supports.Error(ctx, iris.StatusInternalServerError, supports.TokenCreateFailur, nil)
		return
	}

	supports.Ok(ctx, supports.LoginSuccess, vo.BuildUserVO(token, mUser))
	return
}

// UserDepTree 部门树
func UserDepTree(ctx iris.Context) {
	var (
		err  error
		deps []*modeSys.Dep
	)

	if deps, err = modeSys.GetAllDep(); err != nil {
		supports.Error(ctx, iris.StatusInternalServerError, supports.OptionFailur, nil)
		return
	}

	supports.Ok(ctx, supports.OptionSuccess, vo.BuildDepTreeForUser(deps))
	return
}

// UserTable 用户分页查询
func UserTable(ctx iris.Context) {
	var (
		err   error
		page  *supports.Pagination
		total int64
		// userList []*modeSys.User
		userList = make([]*modeSys.User, 0)
		//depId    int
	)

	//depId, err = ctx.URLParamInt("depId")
	page, err = supports.NewPagination(ctx)
	if err != nil {
		goto FAIL
	}

	total, err = supports.GetPagination(&userList, page)
	if err != nil {
		goto ERR
	}

	ctx.JSON(vo.TableVO{
		Total: total,
		Rows:  vo.BuildUserVOList(userList...),
	})
	return

FAIL:
	supports.Error(ctx, iris.StatusBadRequest, supports.ParseParamsFailur, nil)
	return
ERR:
	supports.Error(ctx, iris.StatusInternalServerError, supports.OptionFailur, nil)
	return
}

// UpdateUser 更新用户
func UpdateUser(ctx iris.Context) {
	user := new(modeSys.User)
	if err := ctx.ReadJSON(&user); err != nil {
		ctx.Application().Logger().Errorf("更新用户[%s]失败。%s", "", err.Error())
		supports.Error(ctx, iris.StatusBadRequest, supports.OptionFailur, nil)
		return
	}
	effect, err := modeSys.UpdateUserById(user)
	if err != nil {
		ctx.Application().Logger().Errorf("更新用户[%s]失败。%s", "", err.Error())
		supports.Error(ctx, iris.StatusInternalServerError, supports.OptionFailur, nil)
		return
	}
	supports.Ok(ctx, supports.OptionSuccess, effect)
}

// DeleteUser 删除用户
func DeleteUser(ctx iris.Context, uids string) {
	uidList := strings.Split(uids, ",")
	if len(uidList) == 0 {
		ctx.Application().Logger().Errorf("删除用户错误, 参数不对.")
		supports.Error(ctx, iris.StatusBadRequest, supports.ParseParamsFailur, nil)
		return
	}

	dUids := make([]int64, 0)
	for _, v := range uidList {
		if v == "" {
			continue
		}
		uid, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			ctx.Application().Logger().Errorf("删除用户错误, %s", err.Error())
			supports.Error(ctx, iris.StatusInternalServerError, supports.ParseParamsFailur, nil)
			return
		}
		dUids = append(dUids, uid)
	}

	effect, err := modeSys.DeleteByUsers(dUids)
	if err != nil {
		ctx.Application().Logger().Errorf("删除用户错误, %s", err.Error())
		supports.Error(ctx, iris.StatusInternalServerError, supports.DeleteUsersFailur, nil)
		return
	}
	supports.Ok(ctx, supports.DeleteUsersSuccess, effect)
}

// DeleteUsers 删除用户
func DeleteUsers(ctx iris.Context) {

	cr := new(vo.CommonRequest)

	if err := ctx.ReadJSON(cr); err != nil {
		supports.Error(ctx, iris.StatusInternalServerError, supports.OptionFailur, err.Error())
		return
	}

	uidList := strings.Split(cr.Ids, ",")
	if len(uidList) == 0 {
		ctx.Application().Logger().Errorf("删除用户错误, 参数不对.")
		supports.Error(ctx, iris.StatusBadRequest, supports.ParseParamsFailur, nil)
		return
	}

	dUids := make([]int64, 0)
	for _, v := range uidList {
		if v == "" {
			continue
		}
		uid, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			ctx.Application().Logger().Errorf("删除用户错误, %s", err.Error())
			supports.Error(ctx, iris.StatusInternalServerError, supports.ParseParamsFailur, nil)
			return
		}
		dUids = append(dUids, uid)
	}

	effect, err := modeSys.DeleteByUsers(dUids)
	if err != nil {
		ctx.Application().Logger().Errorf("删除用户错误, %s", err.Error())
		supports.Error(ctx, iris.StatusInternalServerError, supports.DeleteUsersFailur, nil)
		return
	}
	supports.Ok(ctx, supports.DeleteUsersSuccess, effect)
}
