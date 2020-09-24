package main

import (
    "encoding/json"
    "log"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
)

type Campaign struct {
    ID string `json:"id"`
    Name string `json:"name"`
    Status string `json:"status"`
    StartDate string `json:"startDate"`
    Channel string `json:"channel"`
}

var campaigns []Campaign

func getCampaigns(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(campaigns)
}

func getCampaign(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for _, item := range campaigns {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item);
            return
        }
    }
}

func createCampaign(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var newCampaign Campaign
    json.NewDecoder(r.Body).Decode(&newCampaign)
    newCampaign.ID = strconv.Itoa(len(campaigns) + 1)
    campaigns = append(campaigns, newCampaign)
    
    json.NewEncoder(w).Encode(newCampaign)
}

func updateCampaign(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for i, item := range campaigns {
        if item.ID == params["id"] {
            campaigns = append(campaigns[:i], campaigns[i+1:]...)
            var newCampaign Campaign
            json.NewDecoder(r.Body).Decode(&newCampaign)
            newCampaign.ID = params["id"]
            campaigns = append(campaigns, newCampaign)
            json.NewEncoder(w).Encode(newCampaign);
            return
        }
    }
}

func deleteCampaign(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for i, item := range campaigns {
        if item.ID == params["id"] {
            campaigns = append(campaigns[:i], campaigns[i+1:]...)
            break
        }
    }
    json.NewEncoder(w).Encode(campaigns)
}

func main() {

    // dummy data
    campaigns = append(campaigns, 
        Campaign{ID: "1", Name: "VIP Campaign", Status: "Running", StartDate: "09/08/2020", Channel: "Email"},
        Campaign{ID: "2", Name: "Wednesday Madness", Status: "Recurring", StartDate: "09/09/2020", Channel: "Social Media"})

    router := mux.NewRouter()

    router.HandleFunc("/campaigns", getCampaigns).Methods(http.MethodGet)
    router.HandleFunc("/campaign", createCampaign).Methods(http.MethodPost)
    router.HandleFunc("/campaign/{id}", getCampaigns).Methods(http.MethodGet)
    router.HandleFunc("/campaign/{id}", updateCampaign).Methods(http.MethodPut)
    router.HandleFunc("/campaign/{id}", deleteCampaign).Methods(http.MethodDelete)

    log.Fatal(http.ListenAndServe(":8081", router))
}
