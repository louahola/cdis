package auth

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
"golang.org/x/oauth2"
	"io/ioutil"
"encoding/json"
)

type AccessToken struct {
	Token  string
	Expiry int64
}

func readHttpBody(response *http.Response) string {

	fmt.Println("Reading body")

	bodyBuffer := make([]byte, 5000)
	var str string

	count, err := response.Body.Read(bodyBuffer)

	for ; count > 0; count, err = response.Body.Read(bodyBuffer) {

	if err != nil {

	}

	str += string(bodyBuffer[:count])
	}

	return str

}

//Converts a code to an Auth_Token
func GetAccessToken(client_id string, code string, secret string, callbackUri string) AccessToken {
	fmt.Println("GetAccessToken")
	//https://graph.facebook.com/oauth/access_token?client_id=YOUR_APP_ID&redirect_uri=YOUR_REDIRECT_URI&client_secret=YOUR_APP_SECRET&code=CODE_GENERATED_BY_FACEBOOK
	response, err := http.Get("https://graph.facebook.com/oauth/access_token?client_id=" +
	client_id + "&redirect_uri=" + callbackUri +
	"&client_secret=" + secret + "&code=" + code + "&fields=email")

	if err == nil {

	auth := readHttpBody(response)
		fmt.Println(auth)

	var token AccessToken

	tokenArr := strings.Split(auth, "&")

	token.Token = strings.Split(tokenArr[0], "=")[1]
	expireInt, err := strconv.Atoi(strings.Split(tokenArr[1], "=")[1])

	if err == nil {
	token.Expiry = int64(expireInt)
	}

	return token
	}

	var token AccessToken

	return token
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// generate loginURL
	fbConfig := &oauth2.Config{
	// ClientId: FBAppID(string), ClientSecret : FBSecret(string)
	// Example - ClientId: "1234567890", ClientSecret: "red2drdff6e2321e51aedcc94e19c76ee"

	ClientID:     "1455598258078537", // change this to yours
	ClientSecret: "82793411ea1962ae7c8ea64345f6cecd",
	RedirectURL:  "http://72.231.185.204:8080/FBLogin", // change this to your webserver adddress
	Scopes:       []string{"email"},
	Endpoint: oauth2.Endpoint{
	AuthURL:  "https://www.facebook.com/dialog/oauth",
	TokenURL: "https://graph.facebook.com/oauth/access_token",
	},
	}

	url := fbConfig.AuthCodeURL("") + "&auth_type=rerequest&fields=email"
	fmt.Println(url)

	// Home page will display a button for login to Facebook

	w.Write([]byte("<html><title>Golang Login Facebook Example</title> <body> <a href='" + url + "'><button>Login with Facebook!</button> </a> </body></html>"))
}

func FBLogin(w http.ResponseWriter, r *http.Request) {
// grab the code fragment

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	code := r.FormValue("code")

	defer r.Body.Close()
	st, _ := ioutil.ReadAll(r.Body)
	fmt.Println("resp:" + string(st))

	ClientId := "1455598258078537" // change this to yours
	ClientSecret := "82793411ea1962ae7c8ea64345f6cecd"
	RedirectURL := "http://72.231.185.204:8080/FBLogin"

	accessToken := GetAccessToken(ClientId, code, ClientSecret, RedirectURL)

	response, err := http.Get("https://graph.facebook.com/me?access_token=" + accessToken.Token + "&fields=name,email")
	if err != nil {
		return
	}
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	var user map[string]interface{}
	err = decoder.Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "ERROR: %s", err.Error())
		return
	}

	id, _ := user["id"]
	name, _ := user["name"]
	email, _ := user["email"]

	w.Write([]byte(fmt.Sprintf("Username %s ID is %s Email is %s<br>", name, id, email)))


//
//	img := "https://graph.facebook.com/" + id.(string) + "/picture?width=180&height=180"
//
//	w.Write([]byte("Photo is located at " + img + "<br>"))
//	// see https://www.socketloop.com/tutorials/golang-download-file-example on how to save FB file to disk
//
//	w.Write([]byte("<img src='" + img + "'>"))
}
