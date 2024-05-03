package awsconfig

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/spf13/viper"
)

type AppConfig struct {
	AccessKeyID     string `mapstructure:"AWS_ACCESS_KEY"`
	SecretAccessKey string `mapstructure:"AWS_SECRET_KEY"`
}

var (
	appConfig AppConfig
	cfg       aws.Config
	svc       *dynamodb.Client
)

func init() {
	viper.SetEnvPrefix("MY_APP") // Set prefix for environment variables
	viper.AutomaticEnv()         // Automatically map environment variables

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file")
	} else {
		fmt.Println("Error reading config file:", err)
	}

	if err := viper.Unmarshal(&appConfig); err != nil {
		log.Fatalf("unable to decode config: %v", err)
	}

	var err error
	cfg, err = config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("ap-southeast-1"),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     appConfig.AccessKeyID,
				SecretAccessKey: appConfig.SecretAccessKey,
			},
		}))

	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	svc = dynamodb.NewFromConfig(cfg)
}

func GetDynamoDBClient() *dynamodb.Client {
	return svc
}
