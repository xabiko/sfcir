package actions

import "github.com/gobuffalo/buffalo"
import(
	"net/http"
	"net/url"
)
import "io/ioutil"
import "encoding/json"
//import "strings"
import "time"

func Getaccess(code string) (string, string) {
		var dat map[string]interface{}
		resp, err := http.PostForm("https://login.salesforce.com/services/oauth2/token",
		                              url.Values{"grant_type":{"authorization_code"},
																					"code":{code},
																					"client_id":{client_id},
		                                      "client_secret":{client_secret},
		                                      "redirect_uri":{redirect_uri}})
		if err != nil{
			panic(err)
		}
		defer resp.Body.Close()

		json.NewDecoder(resp.Body).Decode(&dat)

		if dat["access_token"] == nil {
				return "","" //Code expired or incorrect
		}

		return dat["access_token"].(string), dat["refresh_token"].(string)
}

func Makerequest(url string, acc string, client http.Client) string {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Bearer " + acc)
	req.Header.Add("X-PrettyPrint", "1")

	resp, err := client.Do(req)
	if err!= nil{
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

func AuthHandler(c buffalo.Context) error {
	return c.Redirect(302, login_url)
}

func HomeHandler(c buffalo.Context) error {
	s := c.Session()
	http.DefaultClient.Timeout = 10 * time.Second
	myClient := &http.Client{
		Timeout : 10 * time.Second,
	}

	acc, ref := Getaccess(c.Param("code"))
	if acc=="" {
		return c.Redirect(302, login_url)
	}
	s.Set("access_token", acc)
	s.Set("refresh_token", ref)

	url1 := "https://na30.salesforce.com/services/data/v41.0/sobjects/Account"
  url2 := "https://na30.salesforce.com/services/data/v41.0/sobjects/Account/0013600000CuqCdAAJ"

  c.Set("recent_items", Makerequest(url1, s.Get("access_token").(string), *myClient))
  c.Set("account_single", Makerequest(url2, s.Get("access_token").(string), *myClient))

  return c.Render(200, r.HTML("map.html"))
}
