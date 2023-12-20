package main

import (
	"context"
	"fmt"
	"idolTime-gorse/client"
)

func testFeed() {
	// Create a client
	gorse := client.NewGorseClient("http://localhost:8088", "api-key")
	//client.NewGorseClient("http://localhost:8088", "")

	// Insert feedback
	resu, err := gorse.InsertFeedback(context.Background(), []client.Feedback{
		{FeedbackType: "like", UserId: "lily", ItemId: "test2", Timestamp: "2023-11-24"},
		{FeedbackType: "like", UserId: "lily", ItemId: "test3", Timestamp: "2023-11-24"},
		{FeedbackType: "like", UserId: "lily", ItemId: "test4", Timestamp: "2023-11-24"},
	})
	if err != nil {
		panic(err)
	}
	items, _ := gorse.GetNeighbors(context.Background(), "test12", 10)
	fmt.Println(resu, err)
	fmt.Println(items, err)
	//key := getMaxValueKey(maps)
	// Get recommendation.
	s, _ := gorse.GetRecommend(context.Background(), "lily", "", 10)
	//scores, _ := gorse.GetUserNeighbors(context.Background(), "lily", 0, 10)
	//gorse.get
	for _, s1 := range s {
		fmt.Println(s1)
	}

}

func main() {
	testFeed()
}
