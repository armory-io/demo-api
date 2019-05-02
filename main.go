package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/sirupsen/logrus"
)

type BucketStatus struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

func main() {
	bucket := flag.String("bucket", "demo-api-bucket", "application bucket")
	port := flag.Int("port", 3000, "application port")
	flag.Parse()
	// some comment
	os.Setenv("AWS_SDK_LOAD_CONFIG", "true")

	sessionConfig := &aws.Config{
		Region:                        aws.String("us-west-2"),
		CredentialsChainVerboseErrors: aws.Bool(true),
	}

	sess, err := session.NewSession(sessionConfig)
	if err != nil {
		logrus.Fatalf("failed to create aws session: %s", err.Error())
	}
	s3Client := s3.New(sess)

	r := http.NewServeMux()
	r.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		_, err := s3Client.HeadBucket(&s3.HeadBucketInput{
			Bucket: aws.String(*bucket),
		})

		if err != nil {
			json.NewEncoder(w).Encode(BucketStatus{
				Status: fmt.Sprintf("unable to communicate with S3 API for bucket %s", *bucket),
				Error:  err.Error(),
			})
			return
		}

		json.NewEncoder(w).Encode(BucketStatus{
			Status: fmt.Sprintf("able to communicate with bucket %s", *bucket),
		})
	})

	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", *port),
		Handler: r,
	}

	logrus.Infof("server starting on port %d", *port)
	if err := s.ListenAndServe(); err != nil {
		logrus.Fatalf("error stopping server: %s", err.Error())
	}
	logrus.Info("server stopped gracefully")
}
