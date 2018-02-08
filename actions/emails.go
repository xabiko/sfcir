package actions

import "github.com/gobuffalo/buffalo"
import(
	"net/http"
	// "net/url"
)
import "io/ioutil"
// import "encoding/json"
// import "strings"
import "time"

func EmailHandler(c buffalo.Context) error {
	s := c.Session()
  var myClient = &http.Client{Timeout: 10 * time.Second}

	url1 := "https://na30.salesforce.com/services/data/v41.0/query/?q=SELECT+Household_Email__c+from+Account"
  req1, _ := http.NewRequest("GET", url1, nil)
  req1.Header.Add("Authorization", "Bearer " + s.Get("access_token").(string))
  req1.Header.Add("X-PrettyPrint", "1")

  resp1, err := myClient.Do(req1)
  if err!= nil{
    panic(err)
  }
  defer resp1.Body.Close()

  body1, _ := ioutil.ReadAll(resp1.Body)

  c.Set("emails", string(body1))

  return c.Render(200, r.HTML("emails.html"))

}
