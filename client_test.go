package flip

import (
	"fmt"
	"os"
	"testing"

	"github.com/FlipsideCrypto/flip-rpc-client-go/dynamicquery"
	"github.com/FlipsideCrypto/flip-rpc-client-go/segment"
)

func makeCondition() segment.Condition {
	gte := segment.Gte{
		PartitionID: "sorted_set:ad43bf8e-0f0c-4102-be91-52bc84150af2:current_balances:flipside",
		Value:       100000000,
	}
	c := segment.Condition{
		Gte: gte,
	}
	return c
}

func makeQuery(condition segment.Condition) dynamicquery.Query {
	segments := make(map[string]segment.Condition)
	segments["large_balance_holder"] = condition

	aggregates := make([]dynamicquery.Aggregate, 0)
	aggregates = append(aggregates, dynamicquery.Aggregate{
		Field:             "event_amount",
		Label:             "total_amount",
		DecimalAdjustment: 16,
		Operation:         "sum",
	})

	groupBy := make([]dynamicquery.GroupBy, 0)
	groupBy = append(groupBy, dynamicquery.GroupBy{
		Field:      "block_timestamp",
		Timebucket: "1 day",
		Label:      "metric_date",
	})

	inSegment := dynamicquery.InSegment{
		Field: "event_to",
		Value: "large_balance_holder",
	}
	filter := dynamicquery.Filter{
		InSegment: inSegment,
	}

	query := dynamicquery.Query{
		Table:      "udm_events_aion",
		Schema:     "source",
		Filter:     filter,
		GroupBy:    groupBy,
		Aggregates: aggregates,
		Segments:   segments,
	}
	return query
}

func TestClient_GetSegmentMembers(t *testing.T) {
	client := getClient(t)

	c, err := client.GetSegmentMembers(makeCondition())
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if c == nil {
		t.Fatal("count is nil")
	}

	fmt.Fprintln(os.Stdout, "GetSegmentMembers")
	fmt.Fprintln(os.Stdout, *c)
	fmt.Println("")
}

func TestClient_IntersectMembersToSegment(t *testing.T) {
	client := getClient(t)

	intersectMembers := make([]string, 0)
	intersectMembers = append(intersectMembers, "a090b025a1489aa6c9204d7b85ac77d51b814402d5cbdec27335575bb46e4f20")

	c, err := client.IntersectMembersToSegment(intersectMembers, makeCondition())
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if c == nil {
		t.Fatal("count is nil")
	}

	fmt.Fprintln(os.Stdout, "IntersectMembersToSegment")
	fmt.Fprintln(os.Stdout, *c)
	fmt.Println("")
}

func TestClient_ExecuteDynamicQuery(t *testing.T) {
	client := getClient(t)

	resp, err := client.ExecuteDynamicQuery(makeQuery(makeCondition()), false)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if resp == nil {
		t.Fatal("resp is nil")
	}

	fmt.Fprintln(os.Stdout, "ExecuteDynamicQuery")
	fmt.Fprintln(os.Stdout, *resp)
	fmt.Println("")
}

func TestClient_ExecuteDynamicQueryWithError(t *testing.T) {
	client := getClient(t)
	query := makeQuery(makeCondition())
	query.Table = "foobar"
	resp, err := client.ExecuteDynamicQuery(query, false)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if resp == nil {
		t.Fatal("resp is nil")
	}
	if resp.Error == "" {
		t.Fatal("error should not be nil")
	}

	fmt.Fprintln(os.Stdout, "ExecuteDynamicQuery")
	fmt.Fprintln(os.Stdout, *resp)
	fmt.Println("")
}

func TestClient_GetMemberPartitions(t *testing.T) {
	client := getClient(t)
	entityID := "ad43bf8e-0f0c-4102-be91-52bc84150af2"
	memberID := "a0969f676e0274c34fffb4261b59d3de48de0d5845ed9780ac43045cf954ed81"

	c, err := client.GetMemberPartitions(entityID, memberID)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if c == nil {
		t.Fatal("count is nil")
	}

	fmt.Fprintln(os.Stdout, "GetMemberPartitions")
	fmt.Fprintln(os.Stdout, *c)
	fmt.Println("")
}

func TestClient_GetDatasets(t *testing.T) {
	client := getClient(t)
	entityID := "ad43bf8e-0f0c-4102-be91-52bc84150af2"

	c, err := client.GetDatasets(entityID, "")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if c == nil {
		t.Fatal("count is nil")
	}

	fmt.Fprintln(os.Stdout, "GetDatasets")
	fmt.Fprintln(os.Stdout, *c)
	fmt.Println("")
}

// func TestClient_Refresh(t *testing.T) {
// 	client := getClient(t)
// 	entityID := "ad43bf8e-0f0c-4102-be91-52bc84150af2"

// 	tags := make([]string, 0)
// 	tags = append(tags, "60m")

// 	ra := RefreshArgs{
// 		EntityID: entityID,
// 		Stage:    "INCREMENTAL_REFRESH",
// 		Tags:     tags,
// 	}
// 	c, err := client.Refresh(ra)
// 	if err != nil {
// 		t.Fatalf("Unexpected error: %v", err)
// 	}
// 	if c == nil {
// 		t.Fatal("count is nil")
// 	}

// 	fmt.Fprintln(os.Stdout, "Refresh")
// 	fmt.Fprintln(os.Stdout, *c)
// 	fmt.Println("")
// }

// func TestClient_GetRefreshJob(t *testing.T) {
// 	client := getClient(t)
// 	jobID := "0f5f8149-03b7-4809-ae3b-0b319cf062c0"

// 	c, err := client.GetRefreshJob(jobID)
// 	if err != nil {
// 		t.Fatalf("Unexpected error: %v", err)
// 	}
// 	if c == nil {
// 		t.Fatal("count is nil")
// 	}

// 	fmt.Fprintln(os.Stdout, "GetRefreshJob")
// 	fmt.Fprintln(os.Stdout, *c)
// 	fmt.Println("")
// }

func getClient(t *testing.T) Client {
	apiKey := os.Getenv("FLIP_API_KEY")
	if apiKey == "" {
		panic("Missing environment variable: `FLIP_API_KEY`")
	}

	baseURL := os.Getenv("FLIP_BASE_URL")
	if baseURL == "" {
		panic("Missing environment variable: `FLIP_BASE_URL`")
	}

	config := Config{
		APIKey:  apiKey,
		BaseURL: baseURL,
	}
	client, err := NewClient(config)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	return client
}
