package main

import (
	"net/http"
	"fmt"
	"strings"
	"html/template"
	"log"
	"net/url"
	"strconv"
	"regexp"
	"time"
	"crypto/md5"
	"io"
	"os"
	"github.com/astaxie/beego/session"
)

var globalSessions *session.Manager
// 然后在 init 函数中初始化
func init() {
	cf := &session.ManagerConfig{}
	globalSessions, _ = session.NewManager("memory", cf)
	//管理session销毁, 启动
	go globalSessions.GC()
}

/*
Test
welcome
 */
func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() 	//解析url传递的参数，对于POST则解析响应包的主题
	// 注意，若谷没有滴啊用ParseForm方法，下面将无法获取表单数据

	// 打印信息输出到服务端
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, " "))
	}
	fmt.Fprintf(w, "Hello FrankLiu!") 	//输出到客户端
}

/*
登录处理
 */
func login(w http.ResponseWriter, r *http.Request) {
	sess, _ := globalSessions.SessionStart(w, r)  // 创建session
	fmt.Println("method:", r.Method)		//获取请求的方法
	if r.Method == "GET" {
		//获取系统时间,生成时间戳
		curtime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(curtime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("../view/login.tpl")
		t.Execute(w, token)
	}else {		//请求的是登录数据，那么执行登录的逻辑判断
		// session 设置
		sess.Set("username", r.Form["username"])
		http.Redirect(w, r, "/", 302)

		//验证表单输入
		if ok := check(w, r); !ok {
			fmt.Println("部分输入格式不正确，请验证输入！")
			//return
		}

		//r.ParseForm()	//默认不解析表单，必须显式调用
		//fmt.Println("usernam:", r.Form["username"])
		//fmt.Println("password:", r.Form["password"])

		// 或者，可以这样写：
		//r.FormValue()会自动调用r.ParseForm(),所以不必提前调用
		fmt.Println("usernam:", r.FormValue("username"))
		fmt.Println("password:", r.FormValue("password"))

		// 对form数据进行操作
		v := url.Values{}
		v.Set("username", "lala")
		v.Add("tele", "1504656556")
		v.Add("friend", "Mike")
		v.Add("friend", "Amy")
		// v.Encode() = "username=lala&tele=1504656556&friend=Mike&friend=Amy"
		fmt.Println("username2:", v.Get("username"))
		fmt.Println("tele:", v.Get("tele"))
		fmt.Println("friend1:", v.Get("friend"))
		fmt.Println("friend2:", v.Get("friend"))

		//预防跨站脚本
		//func HTMLEscape(w io.Writer, b []byte) //把b进行转义之后写到w
		//func HTMLEscapeString(s string) string //转义s之后返回结果字符串
		//func HTMLEscaper(args ...interface{}) string //支持多个参数一起转义,返回结果字符串
		template.HTMLEscape(w, []byte(r.Form.Get("username")))	//输出到客户端

		//使用template.HTML类型
		t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
		err = t.ExecuteTemplate(w, "T", template.HTML("<script>alert('you have been pwned')</script>"))
		if err != nil {
			//....
		}

		// 验证表单token是否合法,防止重复提交
		token := r.Form.Get("token")
		if token != "" {
			//验证token的合法性
			//......

			fmt.Println("token:", token)
			template.HTMLEscape(w, []byte(token))
		} else {
			//不存在token,报错
			fmt.Println("不存在token")
		}
	}
}

/*
上传文件处理
 */
func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)	//获取请求的方法
	if r.Method == "GET" {
		//生成时间戳
		curtime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(curtime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("../view/upload.html")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("../../upload/" + handler.Filename, os.O_WRONLY | os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)	//将文件赋值到服务端文件夹
	}
}

// session test - count and show
func count(w http.ResponseWriter, r *http.Request) {
	sess, _ := globalSessions.SessionStart(w, r)
	//createtime:= sess.Get("createtime")
	//if createtime == nil {
	//	sess.Set("createtime", time.Now().Unix())
	//} else if (createtime.(int64) + 360) < time.Now().Unix() {
	//	globalSessions.SessionDestroy(w, r)
	//	sess, _ = globalSessions.SessionStart(w, r)
	//}
	ct := sess.Get("countnum")
	if ct == nil {
		sess.Set("countnum", 1)
	} else {
		sess.Set("countnum", ct.(int) + 1)
	}
	t, _ := template.ParseFiles("../view/count.html")
	w.Header().Set("Conten-Type", "text/html")
	t.Execute(w, sess.Get("countnum"))
	//fmt.Println("counter:", sess.Get("countnum"))
}

// 验证表单输入函数
func check(w http.ResponseWriter, r *http.Request) bool {
	r.ParseForm()

	ok := true
	//验证是否为空
	if len(r.Form["username"][0]) == 0 {
		fmt.Println("用户名为空！")
		ok = false
	} else if len(r.Form["password"][0]) == 0 {
		fmt.Println("密码为空！")
		ok = false
	}

	//判断是否正整数，然后转化成int
	getint, err := strconv.Atoi(r.Form.Get("age"))
	if err != nil {
		fmt.Println("数字转化出错！")
		ok = false
	} else if getint > 100 {
		fmt.Println("年龄太大了！")
		ok = false
	}

	//正则匹配方式
	//对于性能要求很高的用户，赢避免使用正则表达式
	//if m, _ := regexp.MatchString("^[0-9]+$", r.Form.Get("age")); !m {
	//	ok = false

	//正则匹配中文用户名
	if m, _ := regexp.MatchString("^[\\x{4e00}-\\x{9fa5}]+$", r.Form.Get("username")); m {
		fmt.Println("中文用户名input")
	}

	//英文用户名
	if m, _ := regexp.MatchString("^[a-zA-Z]+$", r.Form.Get("username")); m {
		fmt.Println("English uername input.")
	}

	//验证email
	if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, r.Form.Get("email")); !m {
		fmt.Println("email格式不正确！")
		ok = false
	}

	//验证手机号
	if m, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, r.Form.Get("tele")); !m {
		fmt.Println("手机号格式不正确！")
		ok = false
	}

	//验证select下拉菜单输入
	slice := []string{"apple", "pear", "banane"}
	sel := false
	for _, v := range slice {
		if v == r.Form.Get("fruit") {
			sel = true
		}
	}
	if !sel {
		fmt.Println("未选择下拉菜单！")
		ok = false
	}

	//验证单选按钮,和下拉菜单做法一样
	radio := false
	slice2 := []string{"1","2"}
	for _, v := range slice2 {
		if v == r.Form.Get("gender") {
			radio = true
		}
	}
	if !radio {
		fmt.Println("性别选择不正确！")
		ok = false
	}

	//验证复选框
	//和单选不太一样,因为接收到的数据是一个slice
	slice3 := []string{"football", "basketball", "tennis"}
	a := Slice_diff(r.Form["interest"], slice3)	//slice比较函数
	if a != nil {
		ok = false
	}

	//验证身份证id
	// 15位身份证号全部是数字,18位身份证号前17位是数字,最后一位是校验位,可能是数字或字符x
	if m, _ := regexp.MatchString(`^(\d{15})$`, r.Form.Get("idcard")); !m {
		fmt.Println("15位身份证号正确！")
	} else if m, _ := regexp.MatchString(`^(\d{17})([0-9]|x)$`, r.Form.Get("idcard")); !m {
		fmt.Println("18位身份证号正确！")
	} else {
		fmt.Println("身份证号不正确！")
		ok = false
	}

	return ok
}

//Slice比较函数
func Slice_diff([]string, []string) error {
	//slice比较算法
	//....

	return nil
}

func main() {
	http.HandleFunc("/", sayhelloName)	//设置访问路由 /
	http.HandleFunc("/login", login)		//设置访问路由login
	http.HandleFunc("/upload", upload)	//上传文件
	http.HandleFunc("/count", count)
	fmt.Println("Server is running...")
	fmt.Println("Please access \"http://localhost:8811/\".")
	err := http.ListenAndServe(":8811", nil)	//设置监听端口
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
