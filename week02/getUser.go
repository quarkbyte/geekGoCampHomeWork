package main
import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

type User struct {
	uid int
}
// Q:我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，
// 是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
// A:不应该，ErrNoRows错误和业务无关不抛给上层，转换成了自定义项目中的业务错误USER_EMPTY抛给上层
func getUser() (User,error) {
	db, err := sql.Open("mysql", "root:user78@/test") //后面格式为"user:password@/dbname"
	defer db.Close()
	if err != nil{
		panic(err)
	}
	user:=User{}
	sqlStr:="select * from userinfo"
	err = db.QueryRow(sqlStr).Scan(&user.uid)
	switch {
	case err == sql.ErrNoRows:
		return user,errors.New( "USER_EMPTY")
	case err != nil:
		return user,errors.Wrap(err,sqlStr)
	default:
		return user,nil
	}
}