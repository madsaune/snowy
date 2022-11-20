package table

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	"github.com/madsaune/snowy/auth"
	"github.com/madsaune/snowy/client"
)

type Client struct {
	baseClient *client.Client
}

// NewClient initializes a new client for the ServiceNow Table API
func NewClient(authorizer *auth.Authorizer) *Client {
	return &Client{
		baseClient: client.New(authorizer),
	}
}

// ListResponse contains the response data
// Result is the response body as an interface
type ListResponse struct {
	Count        *int
	Link         *string
	ResponseBody ListResponseBody
}

type ListResponseBody struct {
	Result []map[string]interface{}
}

type SingleResponse struct {
	SingleResponseBody
}

type SingleResponseBody struct {
	Result map[string]interface{}
}

type Query struct {
	// Name-value pairs to use to filter the result set.
	Filter *url.Values

	// Data retrieval operation for reference and choice fields.
	// Based on this value, retrieves the display value and/or the actual value from the database.
	//
	// Valid values:
	//
	// true: Returns the display values for all fields.
	// false: Returns the actual values from the database.
	// all: Returns both actual and display values.
	//
	// Default: false
	DisplayValue interface{}

	// Flag that indicates whether to exclude Table API links for reference fields.
	//
	// Valid values:
	//
	// true: Exclude Table API links for reference fields.
	// false: Include Table API links for reference fields.
	//
	// Default: false
	ExcludeReferenceLink *bool

	// List of field names to include in the response.
	Fields *[]string

	// Maximum number of records to return.
	Limit *int

	// Flag that indicates whether to execute a select count(*) query on the table to return the
	// number of rows in the associated table.
	//
	// Valid values:
	//
	// true: Do not execute a select count(*).
	// false: Execute a select count(*).
	//
	// Default: false
	NoCount *bool

	// Starting record index for which to begin retrieving records.
	// Do not pass a negative number.
	//
	// Default: 0
	Offset *uint64

	// Encoded query used to filter the result set.
	//
	// Syntax: <col_name><operator><value>
	//
	// <col_name>: Name of the table column to filter against.
	// <operator>: Supports the following values:
	//
	//    =: Exactly matches <value>.
	//    !=: Does not match <value>.
	//    ^: Logically AND multiple query statements.
	//    ^OR: Logically OR multiple query statements.
	//    LIKE: <col_name> contains the specified string. Only works for <col_name> fields whose data type is string.
	//    STARTSWITH: <col_name> starts with the specified string. Only works for <col_name> fields whose data type is string.
	//    ENDSWITH: <col_name> ends with the specified string. Only works for <col_name> fields whose data type is string.
	//
	// <value>: Value to match against.
	//
	// All parameters are case-sensitive. Queries can contain more than one entry. Delimit with "^".
	// e.g "caller_id=javascript:gs.getUserID()^active=true"
	//
	// Encoded queries also supports order by functionality. To sort responses based on certain fields,
	// use the ORDERBY and ORDERBYDESC clauses:
	//
	// Syntax:
	//
	//    ORDERBY<col_name>
	//    ORDERBYDESC<col_name>
	//
	// For example: "active=true^ORDERBYnumber^ORDERBYDESCcategory"
	// If part of the query is invalid, such as by specifying an invalid field name, the instance
	// ignores the invalid part. It then returns rows using only the valid portion of the query
	Query *string

	// Name of the category to use for queries.
	QueryCategory *string

	// Flag that indicates whether to restrict the record search to only the domains for which
	// the logged in user is configured.
	//
	// Valid values:
	//
	// false: Exclude the record if it is in a domain that the currently logged in user is not
	// configured to access.
	//
	// true: Include the record even if it is in a domain that the currently logged in user is
	// not configured to access.
	QueryNoDomain *bool

	// Flag that indicates whether to remove the Link header from the response. The Link header
	// provides various URLs to relative pages in the record set which you can use to paginate
	// the returned record set.
	//
	// Valid values:
	//
	// true: Remove the Link header from the response.
	// false: Do not remove the Link header from the response.
	SuppressPaginationHeader *bool

	// UI view for which to render the data. Determines the fields returned in the response.
	//
	// Valid values:
	//
	// desktop
	// mobile
	// both
	//
	// If you also specify the Fields parameter, it takes precedent.
	View *string
}

func (ro *Query) String() string {
	var queryString []string
	queryString = append(queryString, ro.Filter.Encode())

	if reflect.TypeOf(ro.DisplayValue).String() == "string" {
		queryString = append(queryString, "sysparm_display_value="+ro.DisplayValue.(string))
	} else if reflect.TypeOf(ro.DisplayValue).String() == "bool" {
		queryString = append(queryString, "sysparm_display_value="+strconv.FormatBool(ro.DisplayValue.(bool)))
	}

	queryString = append(queryString, "sysparm_exclude_reference_link="+strconv.FormatBool(*ro.ExcludeReferenceLink))
	queryString = append(queryString, "sysparm_fields="+strings.Join(*ro.Fields, ","))
	queryString = append(queryString, "sysparm_limit="+strconv.Itoa(*ro.Limit))
	queryString = append(queryString, "sysparm_no_count="+strconv.FormatBool(*ro.NoCount))
	queryString = append(queryString, "sysparm_offset="+strconv.FormatUint(*ro.Offset, 10))
	queryString = append(queryString, "sysparm_query="+url.QueryEscape(*ro.Query))
	queryString = append(queryString, "sysparm_query_category="+*ro.QueryCategory)
	queryString = append(queryString, "sysparm_query_no_domain="+strconv.FormatBool(*ro.QueryNoDomain))
	queryString = append(queryString, "sysparm_suppress_pagination_header="+strconv.FormatBool(*ro.SuppressPaginationHeader))
	queryString = append(queryString, "sysparm_view="+*ro.View)

	return strings.Join(queryString, "&")
}

func (ro *Query) Values() *url.Values {

	// TODO: Add params from ro.Filter to params
	// params := ro.Filter

	params := url.Values{}

	t := reflect.TypeOf(ro.DisplayValue)

	if t != nil {
		switch t.String() {
		case "string":
			params.Add("sysparm_display_value", ro.DisplayValue.(string))
		case "bool":
			params.Add("sysparm_display_value", strconv.FormatBool(ro.DisplayValue.(bool)))
		}
	}

	if ro.ExcludeReferenceLink != nil {
		params.Add("sysparm_exclude_reference_link", strconv.FormatBool(*ro.ExcludeReferenceLink))
	}

	if ro.Fields != nil {
		params.Add("sysparm_fields", strings.Join(*ro.Fields, ","))
	}

	if ro.Limit != nil {
		params.Add("sysparm_limit", strconv.Itoa(*ro.Limit))
	}

	if ro.NoCount != nil {

		params.Add("sysparm_no_count", strconv.FormatBool(*ro.NoCount))
	}

	if ro.Offset != nil {
		params.Add("sysparm_offset", strconv.FormatUint(*ro.Offset, 10))
	}

	if ro.Query != nil {
		params.Add("sysparm_query", url.QueryEscape(*ro.Query))
	}

	if ro.QueryCategory != nil {
		params.Add("sysparm_query_category", *ro.QueryCategory)
	}

	if ro.QueryNoDomain != nil {
		params.Add("sysparm_query_no_domain", strconv.FormatBool(*ro.QueryNoDomain))
	}

	if ro.SuppressPaginationHeader != nil {
		params.Add("sysparm_suppress_pagination_header", strconv.FormatBool(*ro.SuppressPaginationHeader))
	}

	if ro.View != nil {
		params.Add("sysparm_view", *ro.View)
	}

	return &params
}

func (c *Client) GetOne(tableName, sysID string) (*SingleResponse, error) {
	urlPath := fmt.Sprintf("api/now/table/%s/%s", tableName, sysID)
	resp, err := c.baseClient.Get(context.Background(), urlPath, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var res SingleResponseBody
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, fmt.Errorf("[error][table/get]: failed to decode json response: %v", err)
	}

	gr := &SingleResponse{res}

	return gr, nil
}

func (c *Client) GetAll(tableName string, query *Query) (*ListResponse, error) {
	var values *url.Values
	if query != nil {
		values = query.Values()
	}

	resp, err := c.baseClient.Get(context.Background(), "/api/now/table/"+tableName, values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	totalCountStr := resp.Header.Get("X-Total-Count")
	var count int
	if totalCountStr != "" {
		count, err = strconv.Atoi(totalCountStr)
		if err != nil {
			return nil, err
		}
	}

	link := resp.Header.Get("Link")

	var res ListResponseBody
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, fmt.Errorf("[error][table/get]: failed to decode json response: %v", err)
	}

	gr := &ListResponse{
		Link:         &link,
		Count:        &count,
		ResponseBody: res,
	}

	return gr, nil
}

//func ToBoolPtr(b bool) *bool {
//	return &b
//}

func ToStringPtr(s string) *string {
	return &s
}

func ToIntPtr(i int) *int {
	return &i
}
