package db

import(
    "labix.org/v2/mgo"
)

var master_session *mgo.Session

func Init() {
	var err error
	master_session, err = mgo.Dial("localhost")
    if err != nil {
            panic(err)
    }
}


func Connection() *mgo.Session {
	session := master_session.Copy()
    return session
}