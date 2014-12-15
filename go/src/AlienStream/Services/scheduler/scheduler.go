package scheduler

import (
    "net/http"
    "time"
    "AlienStream/Services/requester"
    "AlienStream/Services/updater"
)


    /*
    || Scheduler
    */
    func Init() { 
        //rate limit to 30 requests per minute max
        longinterval := time.NewTicker(time.Minute * 60).C;
        shortinterval := time.NewTicker(time.Minute * 30).C;
        go func() {
            for {
                select {
                case <- longinterval:
                    updater.RefreshContent()
                    break
                case <- shortinterval:
                    requester.RefreshContent()
                    break
              }
            }
        }()
    }

    func Test(w http.ResponseWriter, request *http.Request) {
        //fmt.Print("Beginning Pull-----------")
        //requester.RefreshContent(w)
    }

