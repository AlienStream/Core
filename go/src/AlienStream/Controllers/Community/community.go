package community

import (
	"fmt"
	"net/http"
    "encoding/json"
    "github.com/gorilla/mux"

    //custom packages
    "AlienStream/Models"
)

func Info(w http.ResponseWriter, request *http.Request) {

    //extract vars from the request
    request.ParseForm()

    name := mux.Vars(request)["name"]
    community := models.Community{}.ByName(name)
    data, _ := json.Marshal(community);

    fmt.Fprintf(w,"%s",data)
}



func Tracks(w http.ResponseWriter, request *http.Request) {

    request.ParseForm()


    name := mux.Vars(request)["name"]
    // TODO make sort work
    
    time := request.FormValue("t")

    if(time == "") {
        time = "hot"
    }

    community := models.Community{}.ByName(name)

    data, _ := json.Marshal(community.Tracks(time));
    fmt.Fprintf(w,"%s",data)
}


// CRUD
func Create(w http.ResponseWriter, request *http.Request) {
	request.ParseForm()
    var community models.Community

	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&community)
	if (err != nil) {
		panic(err)
	}
    	//set up the community with the data passed in our URI
    	if(community.Create()) {
        	fmt.Fprintf(w,"SUCCESS: %s Created",community.Name)
    	} else {
        	fmt.Fprintf(w,"FAILED: Duplicate to %s FOUND",community.Name)
    	}   
}

func Edit(w http.ResponseWriter, request *http.Request) {
    // TODO
    fmt.Fprintf(w,"here's what's trending\n")
}

func Delete(w http.ResponseWriter, request *http.Request) {

    //extract vars from the request
    request.ParseForm()
    var name string = mux.Vars(request)["name"]
    community := models.Community{}.ByName(name)

    if(community.Delete()) {
        fmt.Fprintf(w,"SUCCESS: %s Deleted",name)
    } else {
        fmt.Fprintf(w,"FAILED: %s Not FOUND",name)
    }  
}

// Aggregation
func Trending(w http.ResponseWriter, request *http.Request) {

    //TODO
    var communities = models.Community{}.All()

    var compact_communities []models.CompactCommunity

    for _,community := range communities{
        compact_communities = append(compact_communities, community.Compact())
    }

    output, _ := json.Marshal(compact_communities)
    fmt.Fprintf(w,"%s",output)
}

func Popular(w http.ResponseWriter, request *http.Request) {

    //TODO
    var communities = models.Community{}.All()

    var compact_communities []models.CompactCommunity

    for _,community := range communities{
        compact_communities = append(compact_communities, community.Compact())
    }

    output, _ := json.Marshal(compact_communities)
    fmt.Fprintf(w,"%s",output)
}

// Social 
func Favorite(w http.ResponseWriter, request *http.Request) {

    //TODO
    var communities = models.Community{}.All()

    var compact_communities []models.CompactCommunity

    for _,community := range communities{
        compact_communities = append(compact_communities, community.Compact())
    }

    output, _ := json.Marshal(compact_communities)
    fmt.Fprintf(w,"%s",output)
}
