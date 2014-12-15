package requester

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "strings"
    "time"
)


////////////////////////////////
// REDDIT INTERMEDIATE OBJECT //
////////////////////////////////
type RedditRoot struct {
    Kind string     `json:"kind"`
    Data RedditData `json:"data"`
}

type RedditData struct {
    Children []RedditDataChild `json:"children"`
}

type RedditDataChild struct {
    Data redditPost  `json:"data"`
}

type redditPost struct {
    Url            string   `json:"url"`
    Id             string   `json:"id"`
    Title          string   `json:"title"`
    Thumbnail      string   `json:"thumbnail"`
    Submitted_by   string   `json:"author"`
    Submitted_time float64  `json:"created_utc"`
    Upvotes        int      `json:"ups"`
    Downvotes      int      `json:"downs"`
    Num_Comments   int      `json:"num_comments"`
    Permalink      string   `json:"permalink"`
}

type SubredditAbout struct {
    Data SubredditInfo `json:"data"`
    
}

type SubredditInfo struct {
    Title       string  `json:"title"`
    Id          string  `json:"name"`
    Thumbnail   string  `json:"header_img"`
    Description string  `json:"header_title"`
    Info        string  `json:"description_html"`
    Subscribers int     `json:"subscribers"`  
}

////////////////////////////
// PREJSON/Storage FORMAT //
////////////////////////////
type SubredditData struct {
    Title        string       `json:"subreddit"`
    Id           string       `json:"id"`
    Thumbnail    string       `json:"thumbnail"`
    Description  string       `json:"description"`
    Subscribers  int          `json:"subscribers"`
    Last_Update  time.Time    `json:"-"`
    SourcePosts  []redditPost `json:"tracks"`
}


func GetSubreddit(subreddit string, sort string) SubredditData {
    sub:= SubredditData{}
    last_post:=""
    url:=""
    failcount:=0
    start_time := time.Now()    
    client:=&http.Client{}

    for len(sub.SourcePosts)<200 && time.Since(start_time)<time.Minute*2 && failcount<3 {
        if (last_post=="") {
            url=fmt.Sprintf("%s.json?%s&limit=1000",subreddit,sort)
        } else {
            url=fmt.Sprintf("%s.json?%s&limit=1000&after=t3_%s",subreddit,sort,last_post)
        }
        req,_ := http.NewRequest("GET",url,nil)
        req.Header.Set("User-Agent","AlienStream Master Server v. 1.0")
        resp, err := client.Do(req);
        defer resp.Body.Close();
    
        if err !=nil {
            fmt.Print("An Error Occured")
        } else {
            var data RedditRoot
            decoder := json.NewDecoder(resp.Body)
            decoder.UseNumber()
            err = decoder.Decode(&data)

            if(err != nil) {
                panic(err)
            }

            if len(data.Data.Children) == 0 {
                failcount++
            }
            
            for _, v := range data.Data.Children {
                if strings.Contains(v.Data.Url,"soundcloud.com") || strings.Contains(v.Data.Url,"youtube.com") || strings.Contains(v.Data.Url,"youtu.be"){
                    sub.SourcePosts = append(sub.SourcePosts, v.Data)
                }
                if last_post == v.Data.Id {
                    failcount++ //not working for some reason
                }
                last_post = v.Data.Id
            }

        }
    }
    
    infoclient:=&http.Client{}
    infosub := fmt.Sprintf("%s/about.json",subreddit[0:(len(subreddit)-5)]);
    inforeq, _ := http.NewRequest("GET",infosub,nil)
    inforeq.Header.Set("User-Agent","AlienStream Master Server v. 1.0")
    inforesp, err := infoclient.Do(inforeq);

    if err!=nil {
        fmt.Print("An Error Occured")
    }


    defer inforesp.Body.Close();
    data := SubredditAbout{}
    temp, _ := ioutil.ReadAll(inforesp.Body)
    
    err = json.Unmarshal(temp,&data)

    if(err != nil) {
        panic("Requester failed to fetch from source");
    }
     
    sub.Thumbnail = data.Data.Thumbnail
    sub.Id = data.Data.Id
    sub.Description = data.Data.Description
    sub.Subscribers = data.Data.Subscribers
    sub.Title = data.Data.Title
    
    return sub
}

