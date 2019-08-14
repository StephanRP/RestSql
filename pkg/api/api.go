package api

import (
	"RestSQL/pkg/config"
	"RestSQL/pkg/connLog"
	"RestSQL/routes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var memberResponse []routes.Member

//Write to use route to determine which call to make
func Handler(route string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		var req routes.MemberRequest

		for v := range vars {
			if vars[v] == vars["id"] {
				i, err := strconv.Atoi(vars["id"])
				if err != nil {
					panic(err.Error())
				}
				req.Id = i
			}
			if vars[v] == vars["name"] {
				req.Name = vars["name"]
			}
			if vars[v] == vars["lob"] {
				req.Lob = vars["lob"]
			}
			if vars[v] == vars["pcp"] {
				req.Pcp = vars["pcp"]
			}
			if vars[v] == vars["update"] {
				req.UpdateField = vars["update"]
			}
			if vars[v] == vars["new"] {
				req.NewValue = vars["new"]
			}
		}
		switch {
		case route == config.MemberSearch:
			//connLog.ConnTimeGet(r)
			MemberSearch(req)
		case route == config.MemberAll:
			AllMembers(req)
		case route == config.AddMember:
			AddMember(req)
		case route == config.DeleteMember:
			DeleteMember(req)
		case route == config.UpdateMember:
			UpdateMember(req)
		}
		json.NewEncoder(w).Encode(&memberResponse)
		memberResponse = memberResponse[:0]
	}
}

func MemberSearch(req routes.MemberRequest) {
	defer connLog.FuncTimeTrack(time.Now(), "Member Search Full Function")

	//Logging connection durations
	//TODO: Impliment concurrently or in the same call rather than right before
	defer connLog.ConnTimeGet(config.MemberSearch + req.Name)
	results, err := config.SqlDb().Query("SELECT * FROM members WHERE NAME LIKE '" + req.Name + "%'")
	defer config.SqlDb().Close()

	if err != nil {
		panic(err.Error())
	}
	//Build response based on results
	buildResponse(results, err)
}

//Pulls and encodes all members
func AllMembers(req routes.MemberRequest) {
	defer connLog.FuncTimeTrack(time.Now(), "Member All Full Function")
	fmt.Println("Endpoint Hit: All Members Endpoint")

	//TODO: Run Debug, gets stuck in infinite loop
	//connLog.ConnTimeGet(config.MemberAll)
	results, err := config.SqlDb().Query("SELECT * FROM members")
	defer config.SqlDb().Close()

	if err != nil {
		panic(err.Error())
	}

	buildResponse(results, err)
}

//Adds a new member with auto ++ id
//TODO: Allow additional information updated to seperate table
func AddMember(req routes.MemberRequest) {
	stmt, err := config.SqlDb().Prepare("INSERT INTO members(name, lob, pcp) VALUES(?, ?, ?)")
	defer config.SqlDb().Close()

	if err != nil {
		panic(err.Error())
	}

	stmt.Exec(req.Name, req.Lob, req.Pcp)
	log.Printf("%s Added Successfully", req.Name)
}

//Delete Member
func DeleteMember(req routes.MemberRequest) {
	config.SqlDb().Exec("DELETE FROM members WHERE member_id = ?", req.Id)
	defer config.SqlDb().Close()
	log.Printf("%s Deleted", req.Name)
}

func UpdateMember(req routes.MemberRequest) {
	//unable to take ? with '' around it
	config.SqlDb().Exec("UPDATE members SET "+req.UpdateField+" = '"+req.NewValue+"' WHERE member_id = ?", req.Id)
	defer config.SqlDb().Close()

	log.Printf("%s updated to %s", req.UpdateField, req.NewValue)
}

//Scan results for each row into object Member
func buildResponse(results *sql.Rows, err error) {
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var memberResults routes.Member
		err = results.Scan(&memberResults.Id, &memberResults.Name, &memberResults.Lob, &memberResults.Pcp)

		memberResponse = append(memberResponse, memberResults)
	}
}

//Check health
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

/*func MemberHippa(req routes.MemberRequest) {
	defer connLog.FuncTimeTrack(time.Now(), "Member Hippa Full Function")

	results, err := config.SqlDb().Query("SELECT members.member_id, members.NAME, memberaddress.AddressLine, memberaddress.City, memberaddress.State, memberaddress.Zip FROM memberaddress INNER JOIN members ON members.member_id=? AND memberaddress.member_id=?;", req.Id, req.Address.Id)
	defer config.SqlDb().Close()

	if err != nil {
		log.Panic(err.Error())
	}

	buildResponse(results, err)
}*/
