package main

import (
	"context"
	"fmt"

	"github.com/Asad2730/SNSExample/services"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

func main() {

	topicName := "MyTopic"
	subscriberEmail := "your-email@example.com"
	message := "Hello, SNS subscribers!"

	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		panic(err)
	}

	svc := sns.NewFromConfig(cfg)

	_, err = services.CreateTopic(svc, topicName)
	if err != nil {
		panic(err)
	}

	rs, err := services.ListTopics(svc)

	if err != nil {
		fmt.Println("Error listing topics:", err)
		return
	}

	topicArn := ""

	for _, topic := range rs.Topics {
		fmt.Println("Topic ARNS:", *topic.TopicArn)
		topicArn = *topic.TopicArn
	}

	services.SubscribeToTopic(svc, topicArn, "email", subscriberEmail)

	services.DeleteTopic(svc, topicArn)

	services.ListSubscriptions(svc, topicArn)

	services.Unsubscribe(svc, topicArn)

	services.PublishMessageToTopic(svc, topicArn, message)

}
