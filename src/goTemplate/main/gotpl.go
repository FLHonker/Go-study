package main

import (
	"html/template"
	"os"
	"fmt"
	"strings"
)

type Friend struct {
	Fname	string
}

type Person struct {
	UserName 	string
	Emails 		[]string
	Friends		[]*Friend
}

//模板函数
func EmailDealWith(args ...interface{}) string {
	ok := false
	var s string
	if len(args) == 1 {
		s, ok = args[0].(string)
	}
	if !ok {
		s = fmt.Sprint(args...)
	}
	//find the @ symbol
	substrs := strings.Split(s, "@")
	if len(substrs) != 2 {
		return s
	}
	//replace the @ bt "at"
	return (substrs[0] + " at " + substrs[1])
}

func friends() {
	f1 := Friend{Fname: "changzhan"}
	f2 := Friend{Fname: "ghost"}
	t := template.New("fieldname example")
	//关联模板函数
	t = t.Funcs(template.FuncMap{"emailDeal": EmailDealWith})
	t, _ = t.Parse(`Hello,{{.UserName}}!
				{{range .Emails}}
				an email: {{.|emailDeal}}
				{{end}}
				{{with .Friends}}
				{{range .}}
					my friend name is {{.Fname}}
				{{end}}
				{{end}}
				`)
	p := Person{UserName: "FrankLiu", Emails: []string{"frankliu624@gmail.com", "frankliu@163.com"}, Friends: []*Friend{&f1, &f2}}
	t.Execute(os.Stdout, p)
}

func ifelse() {
	tEmpty := template.New("template test")
	tEmpty = template.Must(tEmpty.Parse(" 空 pipeline if demo: {{if ``}} 不会输出 . {{end}}\n"))
	tEmpty.Execute(os.Stdout, nil)
	tWithValue := template.New("template test")
	tWithValue = template.Must(tWithValue.Parse(" 不为空的 pipeline if demo: {{if `anything`}} 我有内容"))
	tWithValue.Execute(os.Stdout, nil)
	tIfElse := template.New("template test")
	tIfElse = template.Must(tIfElse.Parse("if-else demo: {{if `anything`}} if 部分 {{else}} else 部分 "))
	tIfElse.Execute(os.Stdout, nil)
}

func mustOpt() {
	tOk := template.New("first")
	template.Must(tOk.Parse(" some static text /* and a comment */"))
	fmt.Println("The first one parsed OK.")
	template.Must(template.New("second").Parse("some static text {{ .Name }}"))
	fmt.Println("The second one parsed OK.")
	fmt.Println("The next one ought to fail.")
	tErr := template.New("check parse error with Must")
	template.Must(tErr.Parse(" some static text {{ .Name }"))	// 故意缺少一个"}"
}

func nestingTpl() {
	s1, _ := template.ParseFiles("tpl/header.html", "tpl/content.html", "tpl/footer.html")
	s1.ExecuteTemplate(os.Stdout, "header", nil)
	fmt.Println()
	s1.ExecuteTemplate(os.Stdout, "content", nil)
	fmt.Println()
	s1.ExecuteTemplate(os.Stdout, "footer.html", nil)
	fmt.Println()
	s1.Execute(os.Stdout, nil)
}

func main() {
	friends()
	//ifelse()
	mustOpt()
	nestingTpl()	//没有任何输出
	//因为在默认的情况下没有默认的子模板,所以不会输出任何的东西。
}