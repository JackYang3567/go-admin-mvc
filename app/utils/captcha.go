package utils

import (
	"image/color"
	"image/png"
	"net/http"	
	"github.com/julienschmidt/httprouter"
	 "admin-mvc/app/utils/captcha"
	 qrcode "github.com/skip2/go-qrcode"
	 "fmt"
	 "strings"
)

type CaptchaList []*CaptchaElement

type CaptchaElement struct {
    Id string
}
var CapList CaptchaList
func (cl *CaptchaList) addElement(id string) {
    e := &CaptchaElement{
        Id: id,
    }
    *cl = append(*cl, e)
}

var cap *captcha.Captcha

func init() {
	cap = captcha.New()
	if err := cap.SetFont("public/fonts/Asterix-Blink.ttf"); err != nil {
		panic(err.Error())
	}
	/*
	   //We can load font not only from localfile, but also from any []byte slice
	   	fontContenrs, err := ioutil.ReadFile("comic.ttf")
	   	if err != nil {
	   		panic(err.Error())
	   	}
	   	err = cap.AddFontFromBytes(fontContenrs)
	   	if err != nil {
	   		panic(err.Error())
	   	}
	*/
	cap.SetSize(128, 64)
	cap.SetDisturbance(captcha.MEDIUM)
	cap.SetFrontColor(color.RGBA{255, 255, 255, 255})
	cap.SetBkgColor(color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})
}

func GetCaptcha(w http.ResponseWriter, r *http.Request ,_ httprouter.Params) {
	img, str := cap.Create(6, captcha.ALL)
	png.Encode(w, img)
	//println(str)

	//CapList = CaptchaList{}
	CapList.addElement(strings.ToUpper(str))
	
	//println(list)
	for _, v := range CapList {
		
		println(v.Id)
	}
	
}

func GetCaptchaByQuery(w http.ResponseWriter, r *http.Request ,_ httprouter.Params) {
	str := r.URL.RawQuery
	img := cap.CreateCustom(str)
	png.Encode(w, img)
}

func Qrcode(w http.ResponseWriter, r *http.Request ,_ httprouter.Params){
   
    err := qrcode.WriteFile("https://example.org", qrcode.Medium, 256, "public/qr/qr.png")
	if err != nil {
        fmt.Println("write error")
	}
	
	//var png []byte
   // png, err := qrcode.Encode("https://example.org", qrcode.Medium, 256)
}