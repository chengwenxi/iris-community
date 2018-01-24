Gorm 操作指导

[TOC]

# Create record 

gorm 新建数据的方法有 `Create()`方法，但不能实现部分插入。它采用整体插入，struct 中声明变量都会转化为 sql 语句中的字段插入，对于未给赋值的变量采用对应类型的零值来插入。

具体例子如下：

```Go
// 声明 Users 结构
type Users struct {
	ID uint
	Email string
	Passwords string `gorm:"column:password"`
	IsFemale int // 数据库中默认值为 1
	IsActived int // 数据库中默认值为 1
}
user1 := {
  ID: 6,
  Email: "123",
}
user2 := {
  Email: "test",
  Passwords: "123",
}

```

当使用 `Create()` 方法来插入数据时

```Go
db.Create(&user2)
// generate sql is 
// UPDATE `users` SET `email` = '123', `password` = '', `is_female` = '0', `is_actived` = '0'  WHERE `users`.`id` = 6
```

而我们想要的 sql 为：

```sql
INSERT INTO `users` (`email`,`password`) VALUES ('test','123')
```

---

Issue：如何实现部分插入？ 或避免使用 gorm 插入时，struct 中的零值覆盖默认值呢？

Answer：gorm 文档中并没有提供部分插入的实现，我们只能采取全部插入的方式，但可以使用以下方式来避免零值覆盖默认值。

> 方式一
> ```Go
> // 此方法不需要显示的调用，会在程序调用 Create() 方法前自动调用
> func (user *User) BeforeCreate() error {
>   scope.SetColumn("IsFemale", 1)
>   scope.SetColumn("IsActived", 1)
>   return nil
> }
> ```
>
> 方式二
>
> ```Go
> // Omit() 中填入所有此次插入需要忽略的值，Users struct 中的其他值使用给定值或对应零值来更新
> db.Omit("IsFemale", "IsActived").Create(&user)
> ```
>
> 

# Update Record

gorm 中的 `Save()` 方法也不是部分更新，会更新 struct 中全部的字段，未赋值的字段使用零值来更新，所以一般不建议直接使用 `Save()` 来更新，因为极易覆盖数据库中原有的值。

```go
// Save will include all fields when perform the Updating SQL, even it is not changed
user1.IsFemale = 0
db.Save(&user)
// generate sql is 
// UPDATE users SET name='jinzhu 2', age=100, birthday='2016-01-01', updated_at = '2013-11-17 
```



## 更新一条记录中的多个值

1. 使用 Struct 结构来更新（`db.Model(&user).Updates(Users{...})` ）


> **正确用法**
> ```Go
> db.Model(&user1).Updates(Users{IsFemale: 1, Email: "test"})
> // generated sql is 
> // update users set is_female = 1, email = 'test' where id = 6
> ```
>
> ***错误用法***：`Model(...)` 方法中的实例没有包含主键，导致更新了所有记录（而且并不是按实例中所给的条件更新）。
>
> ```GO
> db.Model(&user2).Updates(Users{IsFemale: 1, Email: "test"})
> // or wanted sql: update users set is_female = 1, email = 'test' where email = '123'
> // but generate sql is
> // update users set is_female = 1, email = 'test'
> ```
>
> ***Note***：struct 类型并不能更新对应变量类型的零值
>
> ```Go
> db.Model(&user1).Updates(Users{IsFemale: 0, Email: "test"})
> // generate sql is 
> // update users set email = 'test' where id = 6
> ```

2. 使用 Map 结构来更新（`db.Model(&user).Updates(map[string]interface{}{...})`）

> **正确用法**
>
> ```go
> db.Model(&user1).Updates(map[string]interface{}{"IsFemale": 0, "Email": "test"})
> // generate sql is
> // update users set is_female = 0, email = 'test' where id = 6
> ```
>
> ***错误用法***：`Model(...)` 方法中的实例没有包含主键，导致更新了所有记录（而且并不是按实例中所给的条件更新）。
>
> ```go
> db.Model(&user2).Updates(Users{IsFemale: 1, Email: "test"})
> // or wanted sql: update users set is_female = 1, email = 'test' where email = '123'
> // but generate sql is
> // update users set is_female = 1, email = 'test'
> ```

## 更新多条记录

用法

```Go
db.Model(&Users{}).Where(&Users{Passwords:"123"}).UpdateColumns(Users{IsFemale: 1})
// generate sql is 
// update users set is_female = 1 where password = '123'
// note：struct 对零值的更新不起作用
```

更多 where 字句的用法详见 [query](http://jinzhu.me/gorm/crud.html#query)

# Delete Record

## 单条删除

使用 `Delete()` 方法进行单条删除时，在删除前必须确保传入主键必须存在否则会 **删除整张表**

```Go
db.Delete(&user1)
// gegerate sql is : delete from users where id = 6
db.Delete(&user2)
// generate sql is : delete from users
```

## 批量删除

```Go
db.Where(Users{Email: "test"}).Delete(Users{})
// generate sql is 
// delete from users where email = 'test'
```

更多 where 字句的用法详见 query

---

















