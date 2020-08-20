package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"

	"github.com/mateors/mcb"
)

var db *mcb.DB
var bucketName string

var workingDirectory string
var websiteTemplateAbsPath string
var adminTemplateAbsPath string

// E:\GOLANG\src\mateors\tutorial\template\basic_template\define\templates\admin\*.gohtml
// E:\GOLANG\src\mateors\tutorial\template\basic_template\define\templates\admin\*.
// E:\GOLANG\src\mateors\tutorial\template\basic_template\define\templates\website\*.gohtml

// /var/www/share/html/template/basic_template\define\templates\admin\*.gohtml

func init() {

	workingDirectory, _ = os.Getwd()
	//fmt.Println("workingDirectory:", workingDirectory)
	adminTemplateAbsPath = filepath.Join(workingDirectory, "templates", "admin", "*.gohtml")
	//fmt.Println("adminTemplateAbsPath:", adminTemplateAbsPath)

	websiteTemplateAbsPath = filepath.Join(workingDirectory, "templates", "website", "*.gohtml")
	//fmt.Println("websiteTemplateAbsPath:", websiteTemplateAbsPath)

	db = mcb.Connect("128.199.136.190", "mostain", "Mosta!n2020$")
	res, err := db.Ping()
	if err != nil {
		fmt.Println(res)
		os.Exit(1)
	}
	fmt.Println(res, err)

}

func main() {

	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("./assets"))))

	http.HandleFunc("/", home)
	http.HandleFunc("/why", why)
<<<<<<< HEAD
	http.HandleFunc("/services", services)
	http.HandleFunc("/request-service", requestService)
	http.HandleFunc("/buy", buy)
	http.HandleFunc("/sell", sell)
	http.HandleFunc("/promotion", promotion)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/about", about)
	http.HandleFunc("/car-details", carDetails)
=======
	http.HandleFunc("/login", login)

	http.HandleFunc("/forgot_password", forgetPassword)
>>>>>>> 69e929ac2d95134b70ad8c97103f0bb47d55a299

	http.ListenAndServe(":8080", nil)

}

func home(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseGlob(websiteTemplateAbsPath)
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = ptmp.ParseFiles("page/home.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	data := struct {
		Title string
	}{
		Title: "Homepage",
	}

	ptmp.Execute(w, data)

}

func why(w http.ResponseWriter, r *http.Request) {

	//New("abase.gohtml")
	ptmp, err := template.ParseGlob(websiteTemplateAbsPath)
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = ptmp.ParseFiles("page/why.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	data := struct {
		Title string
	}{
		Title: "Why Automan",
	}

	//smtp.SendMail()
	ptmp.Execute(w, data)

}

<<<<<<< HEAD
func services(w http.ResponseWriter, r *http.Request) {

	//New("abase.gohtml")
	ptmp, err := template.ParseGlob("template/*.gohtml")
=======
func login(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("page/login.gohtml")
>>>>>>> 69e929ac2d95134b70ad8c97103f0bb47d55a299
	if err != nil {
		fmt.Println(err.Error())
	}

<<<<<<< HEAD
	_, err = ptmp.ParseFiles("page/service.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	data := struct {
		Title string
	}{
		Title: "Services - AutoMan",
	}

	//smtp.SendMail()
	ptmp.Execute(w, data)

}

func requestService(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseGlob("template/*.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = ptmp.ParseFiles("page/request-service.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	data := struct {
		Title string
	}{
		Title: "Request Services - AutoMan",
	}

	ptmp.Execute(w, data)
}

func buy(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseGlob("template/*.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = ptmp.ParseFiles("page/buy.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	data := struct {
		Title string
	}{
		Title: "Buy - AutoMan",
	}

	ptmp.Execute(w, data)
}

func sell(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseGlob("template/*.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = ptmp.ParseFiles("page/sell.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	data := struct {
		Title string
	}{
		Title: "Sell AutoMan",
	}

	ptmp.Execute(w, data)
}

func promotion(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseGlob("template/*.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = ptmp.ParseFiles("page/promotion.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	data := struct {
		Title string
	}{
		Title: "Promotion - AutoMan",
	}

	ptmp.Execute(w, data)

}

func contact(w http.ResponseWriter, r *http.Request) {

	//New("abase.gohtml")
	ptmp, err := template.ParseGlob("template/*.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = ptmp.ParseFiles("page/contact.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	data := struct {
		Title string
	}{
		Title: "Contact - AutoMan",
=======
	r.ParseForm()

	fmt.Println("Method:", r.Method)

	if r.Method == "POST" {

		fmt.Println("form received", r.Form)

		//method 1 to catch html form value
		username := r.FormValue("username")
		password := r.FormValue("password")
		geolocation := r.FormValue("geolocation")
		ip := r.FormValue("ip")
		platform := r.FormValue("platform")
		screen_size := r.FormValue("screen_size")

		fmt.Println("username:", username)
		fmt.Println("password:", password)
		fmt.Println("geolocation:", geolocation)
		fmt.Println("ip:", ip)
		fmt.Println("platform:", platform)
		fmt.Println("screen_size:", screen_size)

		fmt.Println()
		fmt.Println()

		//Mthod 2 to catch html form value map[string][]string
		for key, valA := range r.Form {

			fmt.Println(key, valA[0])
		}

	}

	data := struct {
		Title      string
		LoginError string
	}{
		Title:      "Login",
		LoginError: "",
>>>>>>> 69e929ac2d95134b70ad8c97103f0bb47d55a299
	}

	//smtp.SendMail()
	ptmp.Execute(w, data)

}

<<<<<<< HEAD
func about(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseGlob("template/*.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = ptmp.ParseFiles("page/about.gohtml")
=======
func forgetPassword(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("page/forget_password.gohtml")
>>>>>>> 69e929ac2d95134b70ad8c97103f0bb47d55a299
	if err != nil {
		fmt.Println(err.Error())
	}

	data := struct {
<<<<<<< HEAD
		Title string
	}{
		Title: "About - AutoMan",
	}

	ptmp.Execute(w, data)
}

func carDetails(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseGlob("template/*.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = ptmp.ParseFiles("page/car-details.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	data := struct {
		Title string
	}{
		Title: "Car Details - AutoMan",
	}

	ptmp.Execute(w, data)
=======
		Title      string
		LoginError string
	}{
		Title:      "Login",
		LoginError: "",
	}

	//smtp.SendMail()
	ptmp.Execute(w, data)

>>>>>>> 69e929ac2d95134b70ad8c97103f0bb47d55a299
}
