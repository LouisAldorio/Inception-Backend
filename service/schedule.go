package service

import (
	"context"
	"fmt"
	"log"
	"myapp/config"
	"myapp/graph/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetSchedule(username string) []*model.Schedule {
	var schedules []*model.Schedule

	client := config.MongodbConnect()
	collection := client.Database("Inception").Collection("Schedule")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	findOptions := options.FindOptions{}
	cur, err := collection.Find(ctx, bson.D{
		{"involvedUsers", bson.M{"$elemMatch": bson.D{{"$eq", username}}}},
	}, &findOptions)

	for cur.Next(ctx) {
		var result bson.M
		err = cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}

		var days []*string
		for _, v := range result["days"].(primitive.A) {
			day := fmt.Sprintf("%v", v)
			days = append(days, &day)
		}

		var involvedUsers []*string
		for _, v := range result["involvedUsers"].(primitive.A) {
			user := fmt.Sprintf("%v", v)
			involvedUsers = append(involvedUsers, &user)
		}

		schedule := &model.Schedule{
			ID:                    result["_id"].(primitive.ObjectID).Hex(),
			ScheduleName:          fmt.Sprintf("%v", result["scheduleName"]),
			CommodityName:         fmt.Sprintf("%v", result["commodityName"]),
			DealedUnit:            fmt.Sprintf("%v", result["dealedUnit"]),
			StartDate:             fmt.Sprintf("%v", result["startDate"]),
			EndDate:               fmt.Sprintf("%v", result["endDate"]),
			Day:                   days,
			StartTime:             fmt.Sprintf("%v", result["startTime"]),
			EndTime:               fmt.Sprintf("%v", result["endTime"]),
			InvolvedUsersUsername: involvedUsers,
		}
		schedules = append(schedules, schedule)
	}

	return schedules
}

func CreateSchedule(ctx context.Context, input model.NewSchedule) *model.Schedule {
	client := config.MongodbConnect()
	collection := client.Database("Inception").Collection("Schedule")

	cur, err := collection.InsertOne(ctx, bson.D{
		{"scheduleName", input.ScheduleName},
		{"commodityName", input.CommodityName},
		{"dealedUnit", input.DealedUnit},
		{"startDate", input.StartDate},
		{"endDate", input.EndDate},
		{"days", input.Day},
		{"startTime", input.StartTime},
		{"endTime", input.EndTime},
		{"involvedUsers", input.InvolvedUsersUsername},
	})

	if err != nil {
		log.Println(err)
	}

	result := model.Schedule{
		ID:                    cur.InsertedID.(primitive.ObjectID).Hex(),
		ScheduleName:          input.ScheduleName,
		CommodityName:         input.CommodityName,
		DealedUnit:            input.DealedUnit,
		StartDate:             input.StartDate,
		EndDate:               input.EndDate,
		Day:                   input.Day,
		StartTime:             input.StartTime,
		EndTime:               input.EndTime,
		InvolvedUsersUsername: input.InvolvedUsersUsername,
	}
	return &result
}
