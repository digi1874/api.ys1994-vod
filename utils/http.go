/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-21 21:14:06
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-03-21 21:14:31
 */

package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetHTTPJSON 获取http的json
func GetHTTPJSON(url string, v interface{}) error {
  response, err := http.Get(url)
  if err == nil && response.StatusCode == http.StatusOK {
    err = json.NewDecoder(response.Body).Decode(&v)
  }
  response.Body.Close()

  if err != nil || response.StatusCode != http.StatusOK {
    fmt.Println("ERROR: http.Get", url, response.StatusCode, err)
  }

  return err
}
