package get
import (
    "io/ioutil"
    "net/http"
)

func GetResponseBody(url string) (response string, err error) {
    resp, err := http.Get(url)
    defer resp.Body.Close()
    if err != nil {
        return "", err
    }

    byteArray, err := ioutil.ReadAll(resp.Body)
    return string(byteArray), err
}
