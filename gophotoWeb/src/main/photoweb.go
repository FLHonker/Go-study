package main

import (
	"net/http"
	"io"
	"log"
	"os"
	"io/ioutil"
	"html/template"
	"path"
	"runtime/debug"
)

const (
	UPLOAD_DIR = "../src/uploads"
	TEMPLATE_DIR = "../src/views"
	ListDir = 0x0001
)

//全局变量，用于存放所有模板内容
var templates = make(map[string]*template.Template)

func init() {
	fileInfoArr, err := ioutil.ReadDir(TEMPLATE_DIR)
	check(err)
	var templateName, templatePath string
	for _, fileInfo := range fileInfoArr {
		templateName = fileInfo.Name()
		if ext := path.Ext(templateName); ext != ".html" {
			continue
		}
		templatePath = TEMPLATE_DIR + "/" + templateName
		log.Println("Loading templates:" + templateName)
		//Must封装，确保了模板不能解析成功时,一定会触发错误处理流程。
		t := template.Must(template.ParseFiles(templatePath))
		//去除后缀“.html”
		str := templateName
		tmpl := string([]byte(str)[:len(str)-5])
		templates[tmpl] = t
	}
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//io.WriteString(w, "<form method=\"POST\" action=\"/upload\" " +
		//	" enctype=\"multipart/form-data\">" +
		//	"Choose an image to upload: <input name=\"image\" type=\"file\" />" +
		//	"<input type\"submit\" values\"Upload\" />" +
		//	"</form>")
		//t, err := template.ParseFiles("upload.html")
		if err := renderHtml(w, "upload", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//t.Execute(w, nil)
	}

	if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		//if err != nil {
		//	http.Error(w, err.Error(), http.StatusInternalServerError)
		//	return
		//}
		check(err)
		filename := h.Filename
		defer f.Close()
		t, err := os.Create(UPLOAD_DIR + "/" + filename)
		//if err != nil {
		//	http.Error(w, err.Error(), http.StatusInternalServerError)
		//	return
		//}
		check(err)
		defer t.Close()
		_, err = io.Copy(t, f)
		check(err)
		http.Redirect(w, r, "/view?id=" + filename, http.StatusFound)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	imagePath := UPLOAD_DIR + "/" + imageId
	if exists := isExist(imagePath); !exists {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}

func isExist(path string) bool {
	_, err :=os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

//在网页上列出该目录下存放的所有文件
func listenHandler(w http.ResponseWriter, r *http.Request) {
	fileInfoArr, err := ioutil.ReadDir(UPLOAD_DIR)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	check(err)
	//
	//var listHtml string
	//for _, fileInfo := range fileInfoArr {
	//	imgid := fileInfo.Name()
	//	listHtml += "<li><a href=\"/view?id=" + imgid + "\">imgid</a></li>"
	//}
	//
	//io.WriteString(w, "<ol>"+listHtml+"</ol>")
	locals := make(map[string]interface{})
	images := []string{}
	for _, fileInfo := range fileInfoArr {
		images = append(images, fileInfo.Name())
	}
	locals["images"] = images
	//t, err := template.ParseFiles("list.html")
	err = renderHtml(w, "list", locals)
	check(err)
	//t.Execute(w, locals)
}

//模板渲染函数
func renderHtml(w http.ResponseWriter, tmpl string, locals map[string]interface{})(err error) {
	err = templates[tmpl].Execute(w, locals)
	return err
}

//错误处理
func check(err error) {
	if err != nil {
		panic(err)
	}
}

//闭包封装,避免运行崩溃
func safeHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err, ok := recover().(error); ok {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				// 或者输出到自定义的50x作物界面
				//w.WriteHeader(http.StatusInternalServerError)
				//renderHtml(w, "error", err)

				// logging
				log.Println("WARN: panic in %v - %v", fn, err)
				log.Println(string(debug.Stack()))
			}
		}()
		fn(w, r)
	}
}

// 静态资源和动态请求的分离,闭包实现
func staticDirHandler(mux *http.ServeMux, prefix string, staticDir string, flags int) {
	mux.HandleFunc(prefix, func(w http.ResponseWriter, r *http.Request) {
		file := staticDir + r.URL.Path[len(prefix)-1:]
		if (flags & ListDir) == 0 {
			if exists := isExist(file); !exists {
				http.NotFound(w, r)
				return
			}
		}
		http.ServeFile(w, r, file)
	})
}

func main() {
	//注册方法
	mux := http.NewServeMux()
	staticDirHandler(mux, "/assets/", "../src/style", 0)
	http.HandleFunc("/upload", safeHandler(uploadHandler))
	http.HandleFunc("/view", safeHandler(viewHandler))
	http.HandleFunc("/", safeHandler(listenHandler))
	err := http.ListenAndServe(":8811", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
}
