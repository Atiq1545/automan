package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/mateors/mcb"
	"github.com/mateors/models"
	"github.com/mateors/mtool"
)

var db *mcb.DB

// var bucketName string

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

	companySetup(db)

}

func main() {

	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("./assets"))))

	http.HandleFunc("/", home)
	http.HandleFunc("/why", why)
	http.HandleFunc("/services", services)
	http.HandleFunc("/request-service", requestService)
	http.HandleFunc("/buy", buy)
	http.HandleFunc("/sell", sell)
	http.HandleFunc("/promotion", promotion)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/about", about)
	http.HandleFunc("/car-details", carDetails)

	http.HandleFunc("/login", login)
	http.HandleFunc("/forget_password", forgetPassword)
	http.HandleFunc("/changepass", changepass)

	http.HandleFunc("/dashboard", dashboard)

	http.HandleFunc("/test", test)

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
	ptmp.Execute(w, data)

}

func services(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseGlob(websiteTemplateAbsPath)
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = ptmp.ParseFiles("page/service.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	data := struct {
		Title string
	}{
		Title: "Services - AutoMan",
	}
	ptmp.Execute(w, data)

}

func requestService(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseGlob(websiteTemplateAbsPath)
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

	ptmp, err := template.ParseGlob(websiteTemplateAbsPath)
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

	ptmp, err := template.ParseGlob(websiteTemplateAbsPath)
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

	ptmp, err := template.ParseGlob(websiteTemplateAbsPath)
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

	ptmp, err := template.ParseGlob(websiteTemplateAbsPath)
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
	}
	ptmp.Execute(w, data)

}

func about(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseGlob(websiteTemplateAbsPath)
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = ptmp.ParseFiles("page/about.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	data := struct {
		Title string
	}{
		Title: "About - AutoMan",
	}
	ptmp.Execute(w, data)
}

func carDetails(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseGlob(websiteTemplateAbsPath)
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

}

//for test
func test(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("page/test.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	type Account struct {
		FirstName string
		LastName  string
		Email     string
	}

	r.ParseForm()
	fmt.Println("Method:", r.Method)

	if r.Method == "POST" {

		FirstName := r.FormValue("first_name")
		LastName := r.FormValue("last_name")
		Email := r.FormValue("email")

		fmt.Println("first_name:", FirstName)
		fmt.Println("last_name:", LastName)
		fmt.Println("email:", Email)

	}

	// var myData Account

	// form := make(url.Values, 0)

	// p := db.Insert(form, &myData) //pass by reference (&myData)
	// fmt.Println("Status:", p.Status)

	var a models.Account
	docID := "account::1"

	pRes := db.InsertIntoBucket(docID, BucketName, a)
	fmt.Println("Company:", pRes.Status)

	data := struct {
		Title string
	}{
		Title: "Test",
	}

	ptmp.Execute(w, data)

}

func login(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("page/login.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	r.ParseForm()
	fmt.Println("Method:", r.Method)

	if r.Method == "POST" {
		// fmt.Println("Form Received", r.Form)

		// method 1 to catch html form vlue

		username := r.FormValue("username")
		password := r.FormValue("password")
		geolocation := r.FormValue("geolocation")
		ip := r.FormValue("ip")
		platform := r.FormValue("platform")
		screenSize := r.FormValue("screen_size")
		battery := r.FormValue("battery")

		fmt.Println("username:", username)
		fmt.Println("password:", password)
		fmt.Println("geolocation:", geolocation)
		fmt.Println("ip:", ip)
		fmt.Println("platform:", platform)
		fmt.Println("screen_size:", screenSize)
		fmt.Println("Battery:", battery)

		fmt.Println()

		Login(r, w, db)
		//method 2 to catch html form value

		// for key, valA := range r.Form {
		// 	fmt.Println(key, valA)
		// }

	}

	data := struct {
		Title      string
		LoginError string
	}{
		Title:      "Login",
		LoginError: "",
	}
	ptmp.Execute(w, data)
}

func forgetPassword(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("page/forget_password.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	data := struct {
		Title      string
		LoginError string
	}{
		Title:      "Forget Password",
		LoginError: "",
	}
	ptmp.Execute(w, data)
}

func dashboard(w http.ResponseWriter, r *http.Request) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("dashboard panicking with value >>", r)
		}
	}()

	//check if user logged in or not
	isLoggedIn, loginType, logData := CheckIfAlreadyLoggedIn(r, db)
	fmt.Println("/dashboard ....")
	fmt.Println(isLoggedIn, loginType)

	if isLoggedIn == false {
		fmt.Println("No session found, redirecting /dashboard to /login")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	var cid, accountID string

	cid = mtool.GetMapValue(logData, "cid")
	accountID = mtool.GetMapValue(logData, "account_id")
	fmt.Println(cid, accountID)

	ptmp, err := template.ParseGlob(adminTemplateAbsPath)
	if err != nil {
		fmt.Println(err.Error())
	}

	pageName := filepath.Join(workingDirectory, "page", "dashboard.gohtml")
	ptmp, err = ptmp.ParseFiles(pageName)
	if err != nil {
		fmt.Println(err.Error())
	}

	data := struct {
		Title           string
		AccountName     string
		CompanyName     string
		CompanyWebsite  string
		LoginType       string
		Page            string
		WebuserCount    int
		SystemuserCount int
	}{
		Title:           "Dashboard",
		AccountName:     mtool.GetMapValue(logData, "account_name"),
		CompanyName:     mtool.GetMapValue(logData, "company_name"),
		CompanyWebsite:  mtool.GetMapValue(logData, "website"),
		LoginType:       "SUPERADMIN",
		Page:            r.RequestURI,
		WebuserCount:    0,
		SystemuserCount: 5,
	}

	if err := ptmp.Execute(w, data); err != nil {
		log.Fatal(err)
	}

}

func changepass(w http.ResponseWriter, r *http.Request) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("changepass panicking with value >>", r)
		}
	}()

	isLoggedIn, loginType, logData := CheckIfAlreadyLoggedIn(r, db)
	if isLoggedIn == false {
		fmt.Println("No session found, redirecting /dashboard to /login")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	//handling form submission
	if r.Method == http.MethodPost {

		data := make(map[string]string, 0)
		r.ParseForm()

		loginID := mtool.GetMapValue(logData, "login_id")
		fmt.Println(loginID)
		r.Form.Add("login_id", loginID)

		isOk := ChangePassword(r, db)
		var message string
		if isOk == true {
			message = "OK"
			//fmt.Println("Password")
		} else {
			message = "ERROR"
		}

		data["message"] = message
		//response return to client
		js, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Something wrong @ json.Marshal", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, string(js))

	}

	if r.Method == http.MethodGet {

		ptmp, err := template.ParseGlob(adminTemplateAbsPath)
		if err != nil {
			fmt.Println(err.Error())
		}

		pageName := filepath.Join(workingDirectory, "page", "changepass.gohtml")
		ptmp, err = ptmp.ParseFiles(pageName)
		if err != nil {
			fmt.Println(err.Error())
		}

		data := struct {
			Title          string
			LoginType      string
			Page           string
			CompanyName    string
			CompanyWebsite string
		}{
			Title:          "ChangePassword",
			LoginType:      loginType,
			Page:           r.RequestURI,
			CompanyName:    mtool.GetMapValue(logData, "company_name"),
			CompanyWebsite: mtool.GetMapValue(logData, "website"),
		}

		if err := ptmp.Execute(w, data); err != nil {
			log.Fatal(err)
		}
	}

}
