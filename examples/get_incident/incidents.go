package main

import (
	"fmt"
	"log"
	"os"

	"github.com/madsaune/snowy/auth"
	"github.com/madsaune/snowy/services/table"
)

func main() {
	username := os.Getenv("SN_USERNAME")
	password := os.Getenv("SN_PASSWORD")

	authorizer := &auth.Authorizer{
		Username:    username,
		Password:    password,
		InstanceURL: "https://embriqtest.service-now.com",
	}

	client := table.NewClient(authorizer)
	q := &table.Query{
		Query: table.ToStringPtr("active=true"),
		Limit: table.ToIntPtr(5),
	}

	// Get multiple incidents
	incidentList, err := client.GetAll("incident", q)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	for _, r := range incidentList.ResponseBody.Result {
		fmt.Printf("[%s] %s\n", r["sys_id"], r["short_description"])
	}

	// Get a single incident by sys_id
	incident, err := client.GetOne("incident", "ef2706c29707d510b2a3b68fe153af7d")
	if err != nil {
		log.Fatalf("%+v", err)
	}

	fmt.Printf("[%s] %s\n", incident.Result["sys_id"], incident.Result["short_description"])

	// Get multiple sc_request
	requestList, err := client.GetAll("sc_request", q)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	for _, r := range requestList.ResponseBody.Result {
		fmt.Printf("[%s] %s\n", r["sys_id"], r["short_description"])
	}
}
