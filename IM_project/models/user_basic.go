package models

import (
	"IM_project/utils"
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name     string
	PassWord string
	//验证身份信息
	Identity string
	//这里的有些内容是有现成的检验方式,比如email,手机号就没有,得自己写正则表达式
	//方式就是`valid:""`
	Phone string `valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email string `valid:"email"`
	Salt  string
	//客户可能用多个设备登陆
	ClientIp   string
	ClientPort string
	LoginTime  time.Time
	//心跳时间(超时自动下线)
	HeartbeatTime time.Time
	LoginOutTime  time.Time `gorm:"column:login_out_time" json:"login_out_time"`
	//判断是否下线
	IsLogout   bool
	DeviceInfo string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 20)
	utils.DB.Find(&data)
	return data
}

// 创建用户
func CreateUser(user UserBasic) error {
	result := utils.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// 删除用户
func DeleteUser(user UserBasic) (int, error) {
	//当删除不存在的列时,RowsAffected值是0,如果正常删除,比如成功删除第3行,那么RowsAffected的值就是3
	result := utils.DB.Delete(&user)
	if result.Error != nil {
		//这里的result.RowsAffected的类型是error,所以要int()强制转换一下
		return int(result.RowsAffected), result.Error
	}
	return int(result.RowsAffected), nil
}

// 用户修改信息
func UpdateUser(user UserBasic) error {
	result := utils.DB.Model(&user).Updates(UserBasic{
		Name:     user.Name,
		PassWord: user.PassWord,
		Phone:    user.Phone,
		Email:    user.Email,
		Salt:     user.Salt,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateLoginTime(user UserBasic) error {
	result := utils.DB.Model(&user).Updates(UserBasic{
		LoginTime: time.Now(),
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func FindUserByName(name string) (bool, *gorm.DB, UserBasic) {
	//true是存在,即已经被占用
	//false是不存在,即未被占用,可以正常注册
	user := UserBasic{}
	//通用语法,记住就行
	//"name = ?"中的name是数据库中的列名,?是占位符,可能类似%d这样的,实际参数是name
	//假如传进来的参数name是xiaoming,那就是在name列中找值为xiaoming的返回
	//这里的.First是只取第一个值,并赋值给user
	data := utils.DB.Where("name = ?", name).First(&user)
	if user.Name != "" {
		return true, data, user
	}
	return false, data, user
}

func FindUserByPhone(phone string) (bool, UserBasic) {
	//true是存在,即已经被占用
	//false是不存在,即未被占用,可以正常注册
	user := UserBasic{}
	utils.DB.Where("phone = ?", phone).First(&user)
	if user.Phone != "" {
		return true, user
	}
	return false, user
}

func FindUserByEmail(email string) (bool, UserBasic) {
	//true是存在,即已经被占用
	//false是不存在,即未被占用,可以正常注册
	user := UserBasic{}
	utils.DB.Where("name = ?", email).First(&user)
	if user.Email != "" {
		return true, user
	}
	return false, user
}

func FindUserById(id uint) (bool, UserBasic) {
	user := UserBasic{}
	utils.DB.Where("id = ?", id).First(&user)
	if user.ID != 0 {
		return true, user
	}
	return false, user
}

func FindUserAndRefreshToken(user *gorm.DB, token string) {
	user.Update("identity", token)
}

func IsExit(user UserBasic) (string, bool) {
	NameIsExit, _, _ := FindUserByName(user.Name)
	PhoneIsExit, _ := FindUserByPhone(user.Phone)
	EmailIsExit, _ := FindUserByEmail(user.Email)
	if NameIsExit {
		return "该用户名已被注册!", true
	} else if PhoneIsExit {
		return "该手机号已被注册", true
	} else if EmailIsExit {
		return "该邮箱已被注册", true
	} else {
		return "", false
	}
}
