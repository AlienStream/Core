package requester

import (
	"strings"
	"net/http"
	"encoding/json"
	//"fmt"
)

type SoundCloudoEmbed struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Thumbnail   string `json:"thumbnail_url"`
	Author_Name string `json:"author_name"`
	Author_Url  string `json:"url"`
}


func SoundCloudoEmbedLookup(url string) SoundCloudoEmbed {
	var oEmbed SoundCloudoEmbed
	// setup the request
	var base_url string = "http://soundcloud.com/oembed"
	var params = []string{"format=json","url="+ url}
	var curl_url = base_url +"?"+ strings.Join(params, "&")

	// get the data
	client := &http.Client{}
	req,_ := http.NewRequest("GET",curl_url,nil)
    req.Header.Set("User-Agent","AlienStream Master Server v. 1.0")
    resp,_ := client.Do(req)
    defer resp.Body.Close();
    decoder := json.NewDecoder(resp.Body)
    decoder.Decode(&oEmbed)

    //fmt.Printf("decoding " + curl_url + " \n ")
    return oEmbed
}
