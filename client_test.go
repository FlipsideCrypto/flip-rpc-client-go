package flip

import (
	"fmt"
	"os"
	"testing"
)

func makeCondition() Condition {
	gte := Gte{
		PartitionID: "sorted_set:ad43bf8e-0f0c-4102-be91-52bc84150af2:current_balances:flipside",
		Value:       10000000,
	}
	c := Condition{
		Gte: gte,
	}
	return c
}

func TestClient_GetConditionMembers(t *testing.T) {
	client := getClient(t)

	c, err := client.GetConditionMembers(makeCondition())
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if c == nil {
		t.Fatal("count is nil")
	}

	fmt.Fprintln(os.Stdout, "GetConditionMembers")
	fmt.Fprintln(os.Stdout, *c)
	fmt.Println("")
}

func TestClient_IntersectMembersToCondition(t *testing.T) {
	client := getClient(t)

	intersectMembers := make([]string, 0)
	intersectMembers = append(intersectMembers, "a090b025a1489aa6c9204d7b85ac77d51b814402d5cbdec27335575bb46e4f20")

	c, err := client.IntersectMembersToCondition(intersectMembers, makeCondition())
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if c == nil {
		t.Fatal("count is nil")
	}

	fmt.Fprintln(os.Stdout, "IntersectMembersToCondition")
	fmt.Fprintln(os.Stdout, *c)
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

func TestClient_GetRefreshJob(t *testing.T) {
	client := getClient(t)
	jobID := "0f5f8149-03b7-4809-ae3b-0b319cf062c0"

	c, err := client.GetRefreshJob(jobID)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if c == nil {
		t.Fatal("count is nil")
	}

	fmt.Fprintln(os.Stdout, "GetRefreshJob")
	fmt.Fprintln(os.Stdout, *c)
	fmt.Println("")
}

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
