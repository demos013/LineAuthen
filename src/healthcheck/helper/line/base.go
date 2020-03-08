package line

import (
	"encoding/json"
	"errors"
	"fmt"
	"healthcheck/config"
	"healthcheck/helper/encoding"
	"healthcheck/helper/file"
	"healthcheck/helper/util"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

// LoginResponse -
type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	IDToken      string `json:"id_token"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
}

func Authorization(state string) (err error) {
	config := config.New()
	url := fmt.Sprintf("%s?response_type=%s&client_id=%s&redirect_uri=%s&state=%s&scope=%s", config.LineURL(), "code", config.LineChannel(), config.AppBaseURL(), state, "openid%20profile")
	err = openbrowser(url)
	return
}

func GetBearer(code string) (LoginResponse, error) {
	config := config.New()
	var res LoginResponse

	requestBody := url.Values{}
	requestBody.Set("grant_type", "authorization_code")
	requestBody.Set("client_id", config.LineChannel())
	requestBody.Set("client_secret", config.LineSecret())
	requestBody.Set("code", code)
	requestBody.Set("redirect_uri", config.AppBaseURL())

	client := &http.Client{}
	r, _ := http.NewRequest("POST", config.LineLoginURL(), strings.NewReader(requestBody.Encode())) // URL-encoded payload

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(requestBody.Encode())))

	resp, _ := client.Do(r)

	if resp.StatusCode != 200 {
		return res, errors.New("Bad Request")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func openbrowser(url string) (err error) {

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
	return
}

func SendReport(input util.WebsiteCheckStatusOutput) (int, error) {
	config := config.New()
	requestBody := url.Values{}
	requestBody.Set("total_websites", strconv.Itoa(input.TotalWebsite))
	requestBody.Set("success", strconv.Itoa(input.Success))
	requestBody.Set("failure", strconv.Itoa(input.Failure))
	requestBody.Set("total_time", strconv.FormatInt(input.TotalTime, 10))

	client := &http.Client{}
	r, _ := http.NewRequest("POST", config.LineReportURL(), strings.NewReader(requestBody.Encode()))

	txt, _ := file.ReadFile("public/bearer/line_access_token.txt")
	bearer, _ := encoding.Decode(txt)
	r.Header.Add("Authorization", bearer.(string))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(requestBody.Encode())))

	resp, err := client.Do(r)

	if err != nil {
		return 500, err
	}

	return resp.StatusCode, nil
}
