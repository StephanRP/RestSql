package api

//Main Stuff:
//myRouter.HandleFunc("/filterMembers/{id}", api.ReturnFilteredMembers).Methods("POST")
//myRouter.HandleFunc("/updateMember/{id}/{update}/{new}", api.UpdateMember).Methods("PUT")

//Search Members based on Id or name
/*func MemberSearch(w http.ResponseWriter, r *http.Request) {
	defer conLog.FuncTimeTrack(time.Now(), "Member Search Full Function")
	vars := mux.Vars(r)
	key := vars["id"]
	//conLog.ConnectionLog(key)

	var searchField string

	for _, chars := range key {
		if !unicode.IsLetter(chars) {
			searchField = "member_id"
		} else {
			searchField = "NAME"
		}
	}

	//Logging connection durations
	//TODO: Impliment concurrently or in the same call rather than right before
	conLog.ConTimeGet(config.MemberSearch + key)
	results, err := config.SqlDb().Query("SELECT * FROM members WHERE " + searchField + " LIKE '" + key + "%'")
	defer config.SqlDb().Close()

	if err != nil {
		panic(err.Error())
	}

	//Append results to member array then clear to empty slice
	for results.Next() {
		var memberResults restSql.Member
		err = results.Scan(&memberResults.Id, &memberResults.Name, &memberResults.Lob, &memberResults.Pcp)

		memberResponse = append(memberResponse, memberResults)
	}
	json.NewEncoder(w).Encode(&memberResponse)

	memberResponse = memberResponse[:0]

}

//Returns members filtered by string
func ReturnFilteredMembers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	results, err := config.SqlDb().Query("SELECT * FROM members WHERE NAME LIKE '" + key + "%'" + "OR Lob LIKE '" + key + "%' OR Pcp LIKE '" + key + "%'")
	defer config.SqlDb().Close()

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var memberResults restSql.Member
		err = results.Scan(&memberResults.Id, &memberResults.Name, &memberResults.Lob, &memberResults.Pcp)

		memberResponse = append(memberResponse, memberResults)
	}
	json.NewEncoder(w).Encode(&memberResponse)
	memberResponse = memberResponse[:0]
}

//Returns single member based on id
func ReturnSingleMember(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	i, err := strconv.Atoi(key)

	if err != nil {
		panic(err.Error())
	}
	//Execute the single member query
	var singleMember restSql.Member
	err = config.SqlDb().QueryRow("SELECT * FROM members where member_id = ?", i).Scan(&singleMember.Id,
		&singleMember.Name, &singleMember.Lob, &singleMember.Pcp)

	defer config.SqlDb().Close()

	json.NewEncoder(w).Encode(&singleMember)
}

//Pulls and encodes all members
func AllMembers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: All Members Endpoint")

	results, err := config.SqlDb().Query("SELECT * FROM members")
	defer config.SqlDb().Close()

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		//Scan results for each row into object Member
		var memberResults restSql.Member
		err = results.Scan(&memberResults.Id, &memberResults.Name, &memberResults.Lob, &memberResults.Pcp)

		memberResponse = append(memberResponse, memberResults)
	}
	json.NewEncoder(w).Encode(&memberResponse)
	memberResponse = memberResponse[:0]
}

//Adds a new member with auto ++ id
func AddMember(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Name := vars["name"]
	Lob := vars["lob"]
	Pcp := vars["pcp"]

	stmt, err := config.SqlDb().Prepare("INSERT INTO members(name, lob, pcp) VALUES(?, ?, ?)")
	defer config.SqlDb().Close()

	if err != nil {
		panic(err.Error())
	}

	stmt.Exec(Name, Lob, Pcp)

	log.Printf("%s Added Successfully", Name)

	json.NewEncoder(w).Encode(Name + " Added Successfully")
}

//Deletes single member based on id
func DeleteMember(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	i, err := strconv.Atoi(key)

	config.SqlDb().Exec("DELETE FROM members WHERE member_id = ?", i)
	defer config.SqlDb().Close()

	if err != nil {
		panic(err.Error())
	}
}

//Updates a single field at a time
func UpdateMember(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	i, err := strconv.Atoi(key)
	updateField := vars["update"]
	newValue := vars["new"]

	//unable to take ? with '' around it
	config.SqlDb().Exec("UPDATE members SET "+updateField+" = '"+newValue+"' WHERE member_id = ?", i)
	defer config.SqlDb().Close()

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%s updated to %s", updateField, newValue)

	var singleMember restSql.Member
	err = config.SqlDb().QueryRow("SELECT * FROM members where member_id = ?", i).Scan(&singleMember.Id,
		&singleMember.Name, &singleMember.Lob, &singleMember.Pcp)

	json.NewEncoder(w).Encode(&singleMember)
}

//Filler
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint hit")
}*/
