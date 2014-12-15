package requester

import (
    "strings"
    "AlienStream/Services/db"
    "labix.org/v2/mgo/bson"
)
// A post is the raw data associated with a posting on a valid source
type Post struct {
    Dislikes       int     `json:"dislikes"`
    Num_Comments   int     `json:"num_comments"`
    Permalink      string  `json:"permalink"`
    Source         string  `json:"source"`
    Submitted_by   string  `json:"author"`
    Submitted_time float64 `json:"created_utc"`
    Thumbnail      string  `json:"thumbnail"`
    Title          string  `json:"title"`
    Likes          int     `json:"likes"`
    Url            string  `json:"url"`
}

// A Extended Post is the Data we'll predict on, it's been fully expanded
type ExtPost struct {
    Content_flag          int      // user flagged as non­music (response)
    Content_sysflag       int     // system flagged as non­music (qualitative prediction)
    Description           string   // The description from the Source
    Dislikes              int    
    Duplicate_flag        int      // user flagged as duplicate (response)
    Duplicate_sysflag     int     // system flagged as duplicate (qualitative prediction)
    Likes                 int 
    Num_Comments          int    
    Permalink             string   // direct link to Source
    Provider              string   // where the track is from
    Processed             bool     // If the data has been evaluated yet
    Provider_Title        string   // The title from the Provider
    Provider_Description  string   // The description from the Privder
    Provider_Thumbnail    string   // The description from the Privder
    Provider_Channel      string   // The content owner
    Source                string   
    Submitted_by          string   // who submitted it
    Submitted_time        float64  // used for clustering duplicates
    Thumbnail             string   // not used for prediction
    Title                 string   // the title from the Source
    Url                   string 
}

// Transform Functions
func (subdata SubredditData) Posts(source string) []Post {
    var posts []Post
    for _,sourcepost := range subdata.SourcePosts{
        posts = append(posts,sourcepost.toPost(source))
    }
    return posts   
}

// Reddit Transform
func (sourcepost redditPost) toPost(source string) Post {
    var post Post = Post{
        Dislikes: sourcepost.Downvotes,      
        Num_Comments: sourcepost.Num_Comments,
        Permalink: "http://reddit.com"+sourcepost.Permalink,
        Source: source,
        Submitted_by: sourcepost.Submitted_by,
        Submitted_time: sourcepost.Submitted_time,
        Thumbnail: sourcepost.Thumbnail,
        Title: sourcepost.Title,
        Likes: sourcepost.Upvotes,
        Url: sourcepost.Url,
    }
    return post
}

func (ext ExtPost) Lookup() ExtPost {
    switch(ext.Provider) {
        case "Youtube":
            var video YoutubeVideo   = YoutubeVideoLookup(ext.Url);
            ext.Provider_Title       = video.Snippet.Title
            ext.Provider_Description = video.Snippet.Description
            ext.Provider_Thumbnail   = video.Snippet.Thumbnail.MaxRes.URL
            ext.Provider_Channel     = video.Snippet.Channel
            break
        case "Soundcloud":
            var track SoundCloudoEmbed = SoundCloudoEmbedLookup(ext.Url);
            ext.Provider_Title       = track.Title
            ext.Provider_Description = track.Description
            ext.Provider_Thumbnail   = track.Thumbnail
            ext.Provider_Channel     = track.Author_Name
            break
        default: 
            break
    }

    return ext
}


// Post To Expanded 
func (post Post) Expand() ExtPost {
    var ext ExtPost
    var provider string

    if(strings.Contains(post.Url,"soundcloud.com") ) {
        provider = "Soundcloud"
    }

    if(strings.Contains(post.Url,"youtu.be") ) {
        var v_id string = strings.TrimPrefix(post.Url, "http://youtu.be/")
        v_id = strings.TrimPrefix(v_id, "https://youtu.be/")
        v_id = strings.TrimPrefix(v_id, "youtu.be/")
        post.Url = "http://www.youtube.com/watch?v=" + v_id
    }

    if(strings.Contains(post.Url,"youtube.com") ) {
        provider = "Youtube"
    }

    ext.Content_flag = 0
    ext.Content_sysflag = -1
    ext.Dislikes = post.Dislikes
    ext.Duplicate_sysflag = -1
    ext.Duplicate_flag = 0
    ext.Likes = post.Likes
    ext.Num_Comments = post.Num_Comments
    ext.Permalink = post.Permalink
    ext.Source = post.Source
    ext.Processed = false
    ext.Provider = provider
    ext.Submitted_by = post.Submitted_by
    ext.Submitted_time = post.Submitted_time
    ext.Thumbnail = post.Thumbnail
    ext.Title = post.Title
    ext.Url = post.Url

    ext = ext.Lookup()

    return ext
}

// store in DB
func (post ExtPost) Upsert() bool {
    session := db.Connection()
    defer session.Close()
    collection := session.DB("alien").C("post")

    _,err := collection.Upsert(bson.M{"url": post.Url,"source":post.Source}, &post)
    if err != nil {
        panic(err)
    } 

    return true
}

