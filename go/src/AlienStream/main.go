package main

import (
	//golang packages
	
	"net/http"

	//vendor packages
	"github.com/gorilla/mux"

	//custom packages
	"AlienStream/Controllers/Artist"
	"AlienStream/Controllers/Community"
	//"AlienStream/Controllers/Track"
	//"AlienStream/Controllers/User"
	//"AlienStream/Models"
	"AlienStream/Services/db"
	//"AlienStream/Services/logging"
	//"AlienStream/Services/requester"
	"AlienStream/Services/scheduler"

)


func main() {
	db.Init()
	go scheduler.Init()
    /*		
    || Routes
	*/

    router := mux.NewRouter()
    router.StrictSlash(true)


    	router.HandleFunc("/test/refresh/content",addDefaultHeaders(scheduler.Test))
	//Artist

    	//Aggregation
		router.HandleFunc("/artist/trending/", addDefaultHeaders(artist.Trending))
		router.HandleFunc("/artist/popular/", addDefaultHeaders(artist.Popular))

		//CRUD
		router.HandleFunc("/artist/{name}/", addDefaultHeaders(artist.Info))
		router.HandleFunc("/artist/{name}/create/", addDefaultHeaders(artist.Create))
		router.HandleFunc("/artist/{name}/edit/", addDefaultHeaders(artist.Edit))
		router.HandleFunc("/artist/{name}/flag/", addDefaultHeaders(artist.Flag))

		//Info
		router.HandleFunc("/artist/{name}/tracks/", addDefaultHeaders(artist.Tracks))
		router.HandleFunc("/artist/{name}/communities/", addDefaultHeaders(artist.Communities))


    //Community

		//Aggregation
		router.HandleFunc("/community/trending/", addDefaultHeaders(community.Trending))
		router.HandleFunc("/community/popular/", addDefaultHeaders(community.Popular))

		//Info
		router.HandleFunc("/community/{name}/", addDefaultHeaders(community.Info))
		router.HandleFunc("/community/{name}/tracks/", addDefaultHeaders(community.Tracks))
		//CRUD
		router.HandleFunc("/community/{name}/create/", addDefaultHeaders(community.Create))
		router.HandleFunc("/community/{name}/edit/", addDefaultHeaders(community.Edit))
		router.HandleFunc("/community/{name}/delete/", addDefaultHeaders(community.Delete))

		//Social
		//router.HandleFunc("/community/{name}/favorite/", addDefaultHeaders(community.Favorite))
		//router.HandleFunc("/community/{name}/flag/", addDefaultHeaders(community.Flag))


	//Track
		/*
		//CRUD
		router.HandleFunc("/track/{id}", track.Info)
		router.HandleFunc("/track/{id}/create", track.Create)
		router.HandleFunc("/track/{id}/edit", track.Edit)

		//Aggregation
		router.HandleFunc("/track/trending", track.Trending)
		router.HandleFunc("/track/popular", track.Popular)

		//Social
		router.HandleFunc("/track/{id}/favorite", track.Favorite)
		router.HandleFunc("/track/{id}/flag", track.Flag)
		router.HandleFunc("/track/{id}/upvote", track.Upvote)
		router.HandleFunc("/track/{id}/downvote", track.Downvote)
	
	//User
		
		//CRUD
		router.HandleFunc("/user/", user.Info) 
		router.HandleFunc("/user/login", user.Login)
		router.HandleFunc("/user/logout", user.Logout)
		router.HandleFunc("/user/Register", user.Register)

		//Info
		router.HandleFunc("/user/communities/favorited", user.FavoritedCommunities)
		router.HandleFunc("/user/communities/flagged", user.FlaggedCommunities)
		router.HandleFunc("/user/tracks/upvoted", user.UpvotedTracks)
		router.HandleFunc("/user/tracks/downvoted", user.DownvotedTracks)
		router.HandleFunc("/user/tracks/favorited", user.FavoritedTracks)
		router.HandleFunc("/user/tracks/flagged", user.FlaggedTracks)
		*/
	//Misc

		//Conversion API
		//router.HandleFunc("/request/convert/", request.Convert)

	http.Handle("/", &MyServer{router})
    http.ListenAndServe(":8080", nil);

}
type MyServer struct {
    r *mux.Router
}

func (s *MyServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
    if origin := req.Header.Get("Origin"); origin != "" {
        rw.Header().Set("Access-Control-Allow-Origin", origin)
        rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
        rw.Header().Set("Access-Control-Allow-Headers",
            "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
    }
    // Stop here if its Preflighted OPTIONS request
    if req.Method == "OPTIONS" {
        return
    }
    // Lets Gorilla work
    s.r.ServeHTTP(rw, req)
}


//http://stackoverflow.com/questions/12830095/setting-http-headers-in-golang
func addDefaultHeaders(fn http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, request *http.Request) {
        fn(w, request)
    }
}
