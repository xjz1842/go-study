package main

import (
	"fmt"
	"github.com/gin-gonic/gin"

	//"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

//用户
type User struct {
	ID int
	Name string
}

//数据源，类似MySQL中的数据
var users = []User{
	{ID: 1,Name: "张三"},
	{ID: 2,Name: "李四"},
	{ID: 3,Name: "王五"},
}

func listUser(c *gin.Context)  {
	c.JSON(200,users)
}

func handleUsers(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w,"ID:1,Name:张三")
	fmt.Fprintln(w,"ID:2,Name:李四")
	fmt.Fprintln(w,"ID:3,Name:王五")
}

func main() {
	//http.HandleFunc("/users",handleUsers)
	//http.ListenAndServe(":8080", nil)
	//省略没有改动的代码
	r:=gin.Default()
	r.GET("/users", listUser)
	//省略没有改动的代码
	r.GET("/users/:id", getUser)
	r.POST("/users",createUser)
	//省略没有修改的代码
	r.DELETE("/users/:id", deleteUser)
	//省略没有修改的代码
	r.PATCH("/users/:id",updateUserName)
	r.Run()
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")
	var user User
	found := false

	i := -1
	//类似于数据库的SQL查询
	for index, u := range users {
		if strings.EqualFold(id, strconv.Itoa(u.ID)) {
			user = u
			found = true
			i = index
			break
		}
	}
	if found {
		c.JSON(200, user)
		users = append(users[:i], users[i+1:]...)
	} else {
		c.JSON(404, gin.H{
			"message": "用户不存在",
		})
	}
}

//func handleUsers(w http.ResponseWriter, r *http.Request){
//	switch r.Method {
//	case "GET":
//		users,err:=json.Marshal(users)
//		if err!=nil {
//			w.WriteHeader(http.StatusInternalServerError)
//			fmt.Fprint(w,"{\"message\": \""+err.Error()+"\"}")
//		}else {
//			w.WriteHeader(http.StatusOK)
//			w.Write(users)
//		}
//	default:
//		w.WriteHeader(http.StatusNotFound)
//		fmt.Fprint(w,"{\"message\": \"not found\"}")
//	}
//}

func getUser(c *gin.Context) {
	id := c.Param("id")
	var user User
	found := false
	//类似于数据库的SQL查询
	for _, u := range users {
		if strings.EqualFold(id, strconv.Itoa(u.ID)) {
			user = u
			found = true
			break
		}
	}
	if found {
		c.JSON(200, user)
	} else {
		c.JSON(404, gin.H{
			"message": "用户不存在",
		})
	}
}

func createUser(c *gin.Context) {
	name := c.DefaultPostForm("name", "")
	if name != "" {
		u := User{ID: len(users) + 1, Name: name}
		users = append(users, u)
		c.JSON(http.StatusCreated,u)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "请输入用户名称",
		})
	}
}

func updateUserName(c *gin.Context) {
	id := c.Param("id")
	i := -1
	//类似于数据库的SQL查询
	for index, u := range users {
		if strings.EqualFold(id, strconv.Itoa(u.ID)) {
			i = index
			break
		}
	}
	if i >= 0 {
		users[i].Name = c.DefaultPostForm("name",users[i].Name)
		c.JSON(http.StatusOK, users[i])
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "用户不存在",
		})
	}
}