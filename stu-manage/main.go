package main

import (
	"fmt"
	"os"
)

//面向对象的实现逻辑
//学生管理系统
// 1. 它保存了一些数据		---->结构体字段
// 2. 它有三个功能		---->结构体的方法

type student struct {
	id int64
	name string
}
//造一个学生管理者
type studentMgr struct {
	allStudent map[int64]student
}

//创建一个学生管理对象
var smr studentMgr

//查看学生
func (s studentMgr)showStudent()  {
	for _, stu := range s.allStudent {
		fmt.Printf("学号：%d\t姓名：%s\n",stu.id, stu.name)
	}
}

//添加学生
func (s studentMgr)addStudent()  {
	var(
		id int64
		name string
	)
	//获取用户输入----id和name
	fmt.Print("（添加学生信息）请输入学号：")
	fmt.Scanln(&id)
	fmt.Print("（添加学生信息）请输入姓名：")
	fmt.Scanln(&name)
	//根据用户输入创建结构体对象
	newStu := student{
		id: id,
		name: name,
	}
	//2.把新的学生放到s.allStudent里面
	s.allStudent[newStu.id] = newStu
	fmt.Println("（添加学生信息）添加学生成功！！！")
}

//修改学生
func (s studentMgr)modifyStudent()  {
	//获取用户输入的学号
	var stuID int64
	var stuName string
	fmt.Print("（修改学生信息）请输入学号：")
	fmt.Scanln(&stuID)
	//2.展示该学号的学生，如果没有提示查无此人
	//stu := s.allStudent[stuID] 如果map中没有这个键，则返回对应类型的0值
	stu, ok := s.allStudent[stuID]//如果map中没有这个键，则ok为false
	if ok == false {
		fmt.Println("（修改学生信息）查无此人！！！")
		return
	}
	fmt.Printf("（修改学生信息）你要修改的学生信息：学号：%d\t姓名：%s\n",stu.id,stu.name)
	//3.输入修改后的学生名
	fmt.Print("（修改学生信息）请输入学生的新名字：")
	fmt.Scanln(&stuName)
	//4.更新学生姓名
	stu.name = stuName
	s.allStudent[stuID] = stu
	fmt.Println("（修改学生信息）修改成功！！！")
}

//删除学生
func (s studentMgr)deleteStudent()  {
	//1.输入要删除的学生的id
	var (
		stuID int64
		i string
	)
	fmt.Print("（删除学生信息）请输入学号：")
	fmt.Scanln(&stuID)
	//2.查找并展示该学生的信息，并确认是否删除，如果没有则提示查无此人
	stu, ok := s.allStudent[stuID]
	if !ok{
		fmt.Println("（删除学生信息）查无此人！！！")
		return
	}
	fmt.Printf("（删除学生信息）你要删除的学生信息：学号：%d\t姓名：%s\n",stu.id,stu.name)
	//3.删除该学生信息  map键值对删除：delete(map type, key type)
	fmt.Print("（删除学生信息）删除后，系统将无法恢复原数据，是否删除(y/n):")
	fmt.Scanln(&i)
	switch i {
	case "y":
		delete(s.allStudent, stuID)
		fmt.Println("（删除学生信息）删除成功！！！")
	default:
		fmt.Println("（删除学生信息）选择不删除！！！")
	}
}
//菜单显示
func showMenu() {
	fmt.Println("--------------------welcome sms!!!--------------------")
	fmt.Println(`
	1. 查看所有的学生
	2. 添加学生
	3. 修改学生
	4. 删除学生
	5. 退出`)
}

func main() {
	smr = studentMgr{
		allStudent: make(map[int64]student, 100),
	}
	var choise = int(0)
	for  {
		showMenu()
		fmt.Print("请输入您的选择项：")
		fmt.Scanln(&choise)

		switch choise {
		case 1:
			smr.showStudent()
		case 2:
			smr.addStudent()
		case 3:
			smr.modifyStudent()
		case 4:
			smr.deleteStudent()
		case 5:
			os.Exit(1)
		default:
			fmt.Println("错误，重新选择！")
			continue
		}
	}
}
