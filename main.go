package main

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"log"
	"time"
)

func main() {
	metric := "CPUUtilization"
	namespace := "AWS/EC2"
	dimensionName := "InstanceId"
	dimensionValue := "<your-instance-id>"
	dimension := cloudwatch.Dimension{
		Name:  &dimensionName,
		Value: &dimensionValue,
	}
	dimensionArray := []*cloudwatch.Dimension{
		&dimension,
	}
	statistics := "Average"
	statisticsArray := []*string{&statistics}

	period := int64(300)
	endTime := time.Now().UTC()
	startTime := endTime.Add(time.Duration(-5) * time.Minute)
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create CloudWatch client
	svc := cloudwatch.New(sess)

	input := cloudwatch.GetMetricStatisticsInput{
		Dimensions:         dimensionArray,
		EndTime:            &endTime,
		ExtendedStatistics: nil,
		MetricName:         &metric,
		Namespace:          &namespace,
		Period:             &period,
		StartTime:          &startTime,
		Statistics:         statisticsArray,
		Unit:               nil,
	}

	metricStatistics, err := svc.GetMetricStatistics(&input)
	if err != nil {
		log.Panic(err)
	}
	if len(metricStatistics.Datapoints) == 0 {
		log.Panic(errors.New("metrics is empty"))
	}
	fmt.Printf("Metrics %+v", metricStatistics)
}
