package user
/*
import (
	"fmt"
	"net/http"
	"encoding/json"
    "io/ioutil"
    "sync"
    "time"
    "strings"
    
    //custom packages
    "AlienStream/Models"
    "AlienStream/Services/db"
)


var store = sessions.NewCookieStore(securecookie.GenerateRandomKey(4))

func clear(b []byte) {
    for i := 0; i < len(b); i++ {
        b[i] = 0;
    }
}

func Crypt(password []byte) ([]byte, error) {
    defer clear(password)
    return bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
}



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



func Info(w http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
        Error(request, "Not A POSTable Resource", 405)
        return
    }

	session, _ := store.Get(r, "user-login")
	username := session.Values["username"];
	if username != nil {
		// lookup user info
		fmt.Fprint(w,"{'Response','SUCCESS'}")
	} else {
		// send 403
		fmt.Fprint(w,"{'Response','Not Logged In'}")
	}
}

func Login(w http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
        Error(request, "Not A GETable Resource", 405)
        return
    }

	if r.Method == "POST" {
		session, _ := store.Get(r, "user-login")
		session.Options = &sessions.Options{ 
			Path: "/",
		}
        r.ParseForm()
        u := r.FormValue("username")
        pass := r.FormValue("password")
		
		dbsession := 
		defer dbsession.Close()
		result := User{}
		c := dbsession.DB("users").C("people")
		err = c.Find(bson.M{"name":u}).One(&result)
            	if err != nil {
			panic(err)
		}
		if bcrypt.CompareHashAndPassword([]byte(result.Password),[]byte(pass)) == nil {
            		session.Values["username"] = u;
            		err = session.Save(r, w)
			fmt.Fprint(w,"{'Response','SUCCESS'}")
        	} else {
			fmt.Fprint(w,"{'Response','Login Failed'}")
        	}
	} else {
		fmt.Fprint(w,"{'Response','Invalid Request'}")
	}
}

func Logout(w http.ResponseWriter, request *http.Request) {
	session, _ := store.Get(r, "user-login")
	session.Options.MaxAge = -1	
	session.Save(r,w)
	fmt.Fprint(w,"{'Response','User Logged Out'}")
}

func Save(w http.ResponseWriter, request *http.Request) {
	if r.Method == "POST" {
		session, _ := store.Get(r, "user-login")
		username := session.Values["username"];
        	savedata := r.FormValue("save")
		if username != nil {
			dbsession, _ := mgo.Dial("direct.destruct-o.com")
			defer dbsession.Close()
			c := dbsession.DB("users").C("people")
			c.Update(bson.M{"name":username},bson.M{"$set": bson.M{"save":savedata}});
			fmt.Fprintf(w,"{'Response','Value Saved'}")
		} else {
			fmt.Fprint(w,"{'Response','Not Logged In'}")
		}
	} else {
		fmt.Fprint(w,"{'Response','Invalid Request'}")
	}
}

func Load(w http.ResponseWriter, request *http.Request) {
	session, _ := store.Get(r, "user-login")
	username := session.Values["username"];
	if username != nil {
		dbsession, _ := mgo.Dial("direct.destruct-o.com")
		defer dbsession.Close()
		c := dbsession.DB("users").C("people")
		result := User{}
		c.Find(bson.M{"name": username}).One(&result)
		fmt.Fprintf(w,"{'Response','%s'}",result.Save)
	} else {
		fmt.Fprint(w,"{'Response','Not Logged In'}")
	}

}

func Register(w http.ResponseWriter, request *http.Request) {
	if r.Method == "POST" {
        	u := r.FormValue("username")
        	pass,_ := Crypt([]byte (r.FormValue("password")))
		dbsession, err := mgo.Dial("direct.destruct-o.com")
		if err != nil {
			panic(err)
		}
		defer dbsession.Close()
		c := dbsession.DB("users").C("people")
		err = c.Insert(&User{u, string(pass), "thefakevalue"})
		if err != nil {
			panic(err)
		}
		fmt.Fprint(w,"{'Response','User Registered'}")
	} else {
		fmt.Fprint(w,"{'Response','Invalid Request'}")
	}
}
*/
