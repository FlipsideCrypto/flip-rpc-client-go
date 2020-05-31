# flip-client-go

Go client for accessing Flip RPC Interface

## Usage

### Initialization

```golang
config := Config{APIKey: "<api-key>" BaseURL: "<url-to-the-rpc>"}
client, err := NewClient(config)
```

### Get Segment Members

Evaluate a segment's conditions and retrieve its members.

```golang
gte := segment.Gte{
    PartitionID: "sorted_set:ad43bf8e-0f0c-4102-be91-52bc84150af2:current_balances:flipside",
    Value:       10000000,
}
condition := segment.Condition{
    Gte: gte,
}

result, err := client.GetSegmentMembers(condition)
```

### Intersect Members to Segment

Identify the intersection of an array of members to a segment's evaluated conditions.

```go
intersectMembers := make([]string, 0)
intersectMembers = append(intersectMembers, "a090b025a1489aa6c9204d7b85ac77d51b814402d5cbdec27335575bb46e4f20")

gte := segment.Gte{
    PartitionID: "sorted_set:ad43bf8e-0f0c-4102-be91-52bc84150af2:current_balances:flipside",
    Value:       10000000,
}
condition := segment.Condition{
    Gte: gte,
}

result, err := client.IntersectMembersToSegment(intersectMembers, condition)
```

### Get Member Partitions

For a particular entityID and memberID return the partitions that the member belongs to.

```go
entityID := "ad43bf8e-0f0c-4102-be91-52bc84150af2"
memberID := "a0969f676e0274c34fffb4261b59d3de48de0d5845ed9780ac43045cf954ed81"

result, err := client.GetMemberPartitions(entityID, memberID)
```

### ExecuteDynamicQuery

Execute a query.

```golang
gte := segment.Gte{
    PartitionID: "sorted_set:ad43bf8e-0f0c-4102-be91-52bc84150af2:current_balances:flipside",
    Value:       10000000,
}

segments := make(map[string]segment.Condition)
segments["large_balance_holder"] = segment.Condition{
    Gte: gte,
}

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

debug := false

result, err := client.ExecuteDynamicQuery(query, debug)

```

### Get Datasets

Retreive available datasets and apply optional filters (`entityID`, `ownerID`).

```go
// Get the datasets corresponding to entityID and ownerID
entityID := "ad43bf8e-0f0c-4102-be91-52bc84150af2"
ownderID := ""

result, err := client.GetDatasets(entityID, ownerID)
```

### Refresh

Refresh a metric or dataset using the provided filters.

```go

// Incremental Refresh of Metrics
entityID := "ad43bf8e-0f0c-4102-be91-52bc84150af2"

tags := make([]string, 0)
tags = append(tags, "60m")

ra := RefreshArgs{
    EntityID: entityID,
    Stage:    "INCREMENTAL_REFRESH",
    Tags:     tags,
}

result, err := client.Refresh(entityID, ownerID)

// Incremental Refresh of all Datasets

refreshArgs := RefreshArgs{
    Stage:    "DATASET_BUILD",
}

result, err := client.Refresh(refreshArgs)
```

### Get Refresh Job

Get the status of a refresh job.

```go
// Retrieve the status of this job
jobID := "0f5f8149-03b7-4809-ae3b-0b319cf062c0"

result, err := client.GetRefreshJob(jobID)
```

## Test

```bash
go test
```
