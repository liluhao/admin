package dao

import "time"

type SysUser struct {
	ID        int       `gorm:"primary_key" json:"id"`
	Username  string    `json:"username" gorm:"unique_index;not null"`
	Nickname  string    `json:"nickname"` //通常情况下数据库里该字段都默认为空
	Password  string    `json:"password"`
	AvatarUrl string    `json:"avatar_url" gorm:"default:'static/upload/avatar/default.png'"` //通常情况下数据库里该字段都默认为：static/upload/avatar/default.png
	RoleId    int       `json:"role_id" `
	Status    int       `json:"status" gorm:"default:1"` //通常情况下数据库里该字段都默认为：1
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Dept      int       `json:"dept_id" ` //对应的是sys_deptbo表的id，所以是int型
	Phone     string    `json:"phone" `
	Email     string    `json:"email" `
	Remark    string    `json:"remark" `                       //通常情况下该字段都为空
	SysRole   SysRole   `gorm:"ForeignKey:RoleId" json:"role"` //为外键，创建sys_user表的同时，还会创建sys_role表
}
