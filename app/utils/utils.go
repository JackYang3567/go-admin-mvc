package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"admin-mvc/app/models"	
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	//"io/ioutil"
)

type JsonData struct {
   IsSuccessful  bool          //true: 成功,  false: 失败
   ErrorMessage  string        //错误信息
   Data          interface{}   //数据集
}

type Configuration struct {
	Address     	 string
	ReadTimeout 	 int64
	WriteTimeout 	 int64
	Static       	 string
	DbDriverName 	 string
	DataSourceName   string
}

var Config Configuration
var logger *log.Logger

// Convenience function for printing to stdout


func init() {
	loadConfig()
	file, err := os.OpenFile("log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	logger = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
}

func loadConfig() {
	file, err := os.Open("config/config.json")
	if err != nil {
		log.Fatalln("Cannot open config file", err)
	}
	decoder := json.NewDecoder(file)
	
	Config = Configuration{}
	err = decoder.Decode(&Config)
	
	if err != nil {
		log.Fatalln("Cannot get configuration from file", err)
	}
}
func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}


func P(a ...interface{}) {
	fmt.Println(a)
}
 
func GetParentDirectory(dirctory string) string {
	return substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}
 
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}



// Convenience function to redirect to the error message page
func Error_message(writer http.ResponseWriter, request *http.Request, msg string) {
	url := []string{"/err?msg=", msg}
	http.Redirect(writer, request, strings.Join(url, ""), 302)
}

// Checks if the user is logged in and has a session, if not err is not nil
func Session(writer http.ResponseWriter, request *http.Request) (sess models.Session, err error) {
	cookie, err := request.Cookie("_cookie")
	if err == nil {
		sess = models.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}


// parse HTML templates
// pass in a list of file names, and get a template
func ParseTemplateFiles(filenames ...string) (t *template.Template) {
	var files []string
	t = template.New("layout")
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/%s.html", file))
	}
	t = template.Must(t.ParseFiles(files...))
	return
}



func GenerateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
}

// for logging
/*
 调试输出信息写入文件
*/
func Info(args ...interface{}) {
	logger.SetPrefix("INFO ")
	logger.Println(args...)
}

/*
 错误日志信息写入文件
*/
func Danger(args ...interface{}) {
	logger.SetPrefix("ERROR ")
	logger.Println(args...)
}

/*
 错误日志信息写入文件
*/
func Warning(args ...interface{}) {
	logger.SetPrefix("WARNING ")
	logger.Println(args...)
}

// version
func Version() string {
	return "0.1"
}


//判断某个slice中是否包含某个元素
func IsInArray(val string, array []string) bool {
	for _, v := range array {
		if v == val {
			return true
		}
	}
	return false
}