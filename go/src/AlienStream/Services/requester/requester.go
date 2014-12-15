package requester

import (
	"fmt"
    //"time"
    //custom packages
	"AlienStream/Models"
)


type SortParams struct {
    sort string
    time string
}

// Fetch New Content
func RefreshContent() {

    // For each community fetch the tracklist given the sources of that community
    // Communities will not be stored with their children's tracks
    var communities = models.Community{}.All()

    var sorts []SortParams
    sorts = append(sorts, SortParams{"hot","hot"})
    sorts = append(sorts, SortParams{"top","today"})
    sorts = append(sorts, SortParams{"top","week"})
    sorts = append(sorts, SortParams{"top","month"})
    sorts = append(sorts, SortParams{"top","year"})
    sorts = append(sorts, SortParams{"top","all"})
  
    jobs := make(chan models.Community, 100)

    // Allocate 5 Workers
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, sorts)
    }

    // For Each Community
    for _,community := range communities{
        jobs <- community
    }
    close(jobs)

}

func worker(id int, jobs <-chan models.Community, sorts []SortParams) {
    for community := range jobs {
        fmt.Println("worker", id, "processing job", community.Name)
        RefreshCommunity(community, sorts)
    }
}

func RefreshCommunity(community models.Community, sorts []SortParams){
        // fetch all date ranges
        for _,params := range sorts {

            // Fetch the Posts from each of it's sources
            var allPosts []Post
            for _,source := range community.Sources{
                switch(source.Type) {
                    case "reddit":
                        fmt.Print("Fetching " + source.URL+"/"+params.sort+"/", "sort="+params.sort+"&t="+params.time + "\n")
                        subreddit := GetSubreddit(source.URL+"/"+params.sort+"/", "sort="+params.sort+"&t="+params.time)
                        posts := subreddit.Posts(community.Name+"&t="+params.time)
                        allPosts = append(allPosts, posts...)                    
                        break
                    default:
                        fmt.Print("Unknown Source Type")
                }
            }

            // Check for a good pull, bail if bad
            if(len(allPosts) >= 1) {

                // Convert the posts into track objects for storage
                for _,post := range allPosts{
                    var expanded ExtPost = post.Expand()
                    expanded.Upsert()
                }

            } else {
                fmt.Printf("bad pull for " + community.Name + "&t="+ params.time + "\n")
            }
        }

        // Save our Changes
        community.Update()

        fmt.Printf("Finished " + community.Name + " \n")
}




/*


func Tracks(w http.ResponseWriter, request *http.Request) {


    //get our own tracks
    tracks := getTracks(community,sort,time)
    all_tracks = append(all_tracks,tracks...)
    community.Tracks = []string{}
    for _,track := range tracks{
        community.Tracks = append(community.Tracks, fmt.Sprintf("%x",string(track.Id)))
    }
    c.Upsert(bson.M{"name": name},community)

    //return all tracks
    data, _ := json.Marshal(all_tracks);
    fmt.Fprintf(w,"%s",data)
}
*/