package services

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

func CreateTopic(svc *sns.Client, topicName string) (*sns.CreateTopicOutput, error) {

	input := &sns.CreateTopicInput{
		Name: aws.String(topicName),
	}

	rs, err := svc.CreateTopic(context.TODO(), input)

	if err != nil {
		return nil, err
	}

	return rs, nil
}

func ListTopics(svc *sns.Client) (*sns.ListTopicsOutput, error) {

	input := &sns.ListTopicsInput{}

	rs, err := svc.ListTopics(context.TODO(), input)

	if err != nil {
		return nil, err
	}

	return rs, nil
}

func SubscribeToTopic(svc *sns.Client, topicArn, protocol, endpoont string) error {

	input := &sns.SubscribeInput{
		TopicArn: aws.String(topicArn),
		Protocol: aws.String(protocol),
		Endpoint: aws.String(endpoont),
	}

	_, err := svc.Subscribe(context.TODO(), input)

	if err != nil {
		return err
	}

	return nil
}

func DeleteTopic(svc *sns.Client, topicARN string) error {

	input := &sns.DeleteTopicInput{
		TopicArn: aws.String(topicARN),
	}

	_, err := svc.DeleteTopic(context.TODO(), input)

	if err != nil {
		return err
	}

	return nil
}

func ListSubscriptions(svc *sns.Client, topicARN string) (*sns.ListSubscriptionsByTopicOutput, error) {

	input := &sns.ListSubscriptionsByTopicInput{
		TopicArn: aws.String(topicARN),
	}

	res, err := svc.ListSubscriptionsByTopic(context.TODO(), input)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func Unsubscribe(svc *sns.Client, subscriptionARN string) error {

	input := &sns.UnsubscribeInput{
		SubscriptionArn: aws.String(subscriptionARN),
	}

	_, err := svc.Unsubscribe(context.TODO(), input)

	return err
}

func PublishMessageToTopic(svc *sns.Client, topicARN, message string) (*sns.PublishOutput, error) {

	input := &sns.PublishInput{
		TopicArn: aws.String(topicARN),
		Message:  aws.String(message),
	}

	rs, err := svc.Publish(context.TODO(), input)

	if err != nil {
		return nil, err
	}

	return rs, nil
}
