package actions

import "github.com/gobuffalo/buffalo"
import(
	"net/http"
	"net/url"
)
import "io/ioutil"
import "encoding/json"
import "strings"
import "time"

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	return c.Redirect(302, login_url)
}

func AuthHandler(c buffalo.Context) error {
	s := c.Session()

	var myClient = &http.Client{Timeout: 10 * time.Second}
	var dat map[string]interface{}

	for s.Get("access_token")==nil {
					resp, err := myClient.PostForm("https://login.salesforce.com/services/oauth2/token",
				                                url.Values{"grant_type":{"authorization_code"},
																								"code":{c.Param("code")},
																								"client_id":{client_id},
				                                        "client_secret":{client_secret},
				                                        "redirect_uri":{redirect_uri}})
					if err != nil{
						panic(err)
					}
					defer resp.Body.Close()

					err = json.NewDecoder(resp.Body).Decode(&dat)
					if err != nil{
						return err
					}

					if dat["access_token"] == nil {
							return c.Redirect(302, login_url)
					}

					acc := dat["access_token"].(string)

					s.Set("access_token", acc[strings.Index(acc,"!")+1:])
	}

	url1 := "https://na30.salesforce.com/services/data/v41.0/sobjects/Account"
  req1, _ := http.NewRequest("GET", url1, nil)
  req1.Header.Add("Authorization", "Bearer " + s.Get("access_token").(string))
  req1.Header.Add("X-PrettyPrint", "1")

  url2 := "https://na30.salesforce.com/services/data/v41.0/sobjects/Account/0013600000CuqCdAAJ"
  req2, _ := http.NewRequest("GET", url2, nil)
  req2.Header.Add("Authorization", "Bearer " + s.Get("access_token").(string))
  req2.Header.Add("X-PrettyPrint", "1")

  resp1, err1 := myClient.Do(req1)
  if err1!= nil{
    panic(err1)
  }
  defer resp1.Body.Close()

  resp2, err2 := myClient.Do(req2)
  if err2!= nil{
    panic(err2)
  }
  defer resp2.Body.Close()

  body1, _ := ioutil.ReadAll(resp1.Body)
  body2, _ := ioutil.ReadAll(resp2.Body)

  c.Set("recent_items", string(body1))
  c.Set("account_single", string(body2))

  return c.Render(200, r.HTML("map.html"))

}
