package service

import (
	"admin/common"
	"admin/common/utils"
	"admin/dao"
	"admin/dto"
	"errors"
	"strconv"
	"time"
)

// 登录验证
func LoginCheck(username string, password string) (user *dao.SysUser, msg string, isPass bool) {
	password = utils.Md5(password + common.PWsalt)
	//fmt.Println(password)//b9d11b3be25f5a1a7dc8ca04cd310b28
	user, ok := CheckUserExist(username, password)

	if !ok {
		return user, "用户或密码不正确", false
	}
	return user, "登录成功", true
}

//password  ==19d6e6f9f868a4523c607f5b1c985691
// 验证用户是否存在
func CheckUserExist(username string, password string) (user *dao.SysUser, ok bool) {
	user = new(dao.SysUser)
	common.DB.Where(
		dao.SysUser{
			Username: username,
			Password: password,
			Status:   1}).
		Preload("SysRole").First(user)
	return user, user.ID > 0
}

func ResetPassword(userForm dto.PassWord) (err error) {
	password := utils.Md5(userForm.Password + common.PWsalt)
	newPassword := utils.Md5(userForm.NewPassword + common.PWsalt)
	user, ok := CheckUserExist(userForm.Username, password)
	if ok {
		return common.DB.Model(&user).Update("password", newPassword).Error
	}
	return errors.New("原密码不正确")
}

func GetUserList(q dto.QueryUser) (u []dto.UserInfoOut, total int, err error) {
	//查找所有user,所以定义为切片
	var users []dao.SysUser
	query := common.DB.Model(&users)
	if q.Phone != "" {
		query = query.Where("phone = ?", q.Phone).Count(&total)
	}
	if q.DeptId != "" {
		query = query.Where("dept = ?", q.DeptId).Count(&total)
	}

	if q.Username != "" {
		query = query.Where("username like ?", "%"+q.Username+"%").Count(&total)
	}

	if q.PageSize != "" && q.Page != "" {
		PageSize := utils.StrToInt(q.PageSize) //string转int
		Page := utils.StrToInt(q.Page)         //string转int
		//query = query.Limit(10).Offset( 0*10)        从索引为0的数据开始查,查10条数据
		query = query.Limit(PageSize).Offset((Page - 1) * PageSize)
	}

	//使用预加载
	err = query.Preload("SysRole").Find(&users).Error
	//fmt.Println(err)//<nil>

	for _, user := range users {
		var dtoUser dto.UserInfoOut
		dtoUser.ID = strconv.Itoa(user.ID)
		dtoUser.Username = user.Username
		dtoUser.Phone = user.Phone
		dtoUser.Email = user.Email
		dtoUser.Nickname = user.Nickname
		dtoUser.Dept = strconv.Itoa(user.Dept)
		dtoUser.CreateTime = utils.TimeParseStr(user.CreatedAt, "FULL") //讲time.Time时间转化为2022-05-31 15:04:05 的string类型
		dtoUser.Status = user.Status
		dtoUser.RoleId = strconv.Itoa(user.RoleId)
		dtoUser.Remark = user.Remark
		dtoUser.RoleName = user.SysRole.RoleName
		u = append(u, dtoUser)
	}
	return
}

func SaveUser(u dto.UserInfoIn) (err error) {
	var user dao.SysUser
	user.ID = utils.StrToInt(u.ID)
	user.Username = u.Username
	user.Status = utils.StrToInt(u.Status)

	res, ok := strconv.Atoi(u.RoleName)
	//RoleName比如是超级管理员、测试账号
	if ok == nil {
		user.RoleId = res
	} else {
		user.RoleId = utils.StrToInt(u.RoleId)
	}

	user.Remark = u.Remark
	user.CreatedAt = time.Now() //也需要
	user.AvatarUrl = u.AvatarUrl
	user.Dept = utils.StrToInt(u.Dept)
	user.Email = u.Email
	user.Status = 1
	user.Nickname = u.Nickname
	user.Phone = u.Phone
	if user.ID == 0 {
		//加密后再传入
		user.Password = utils.Md5(u.Password + common.PWsalt)
		err = common.DB.Save(&user).Error
	} else {
		//更新密码
		err = common.DB.Model(user).Update(&user).Error
	}
	return
}

//只用id号就可以删除
func DeleteUser(id int) error {
	var role dao.SysUser
	role.ID = id
	err := common.DB.Delete(&role).Error
	return err
}
