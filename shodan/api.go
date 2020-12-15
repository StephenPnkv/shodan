
package shodan
import (
  "fmt"
  "encoding/json"
  "net/http"
)

type APIInfo struct{
  QueryCredits  int `json:"query_credits"`
  ScanCredits   int `json:"scan_credits"`
  Telnet        bool `json:"telnet"`
  Plan          string `json:"plan"`
  HTTPS         bool `json:"https"`
  Unlocked      bool `json:"unlocked"`
}

func (s *Client) APIInfo()(*APIInfo, error){
  res, err := http.Get(fmt.Sprintf("%s/api-info?key=%s", BaseURL, s.apiKey))
  if err != nil{
    return nil, err
  }
  defer res.Body.Close()
  var data APIInfo
  if err := json.NewDecoder(res.Body).Decode(&data); err != nil{
    return nil, err
  }

  return &data, nil
}
