package controller

import (
	"fmt"
	"math/rand"
	"os"
	"s3bench/config"
	"s3bench/object"
	"s3bench/report"
	"s3bench/s3"
	"time"

	log "github.com/sirupsen/logrus"
)

var bucketQueue chan string

type task struct {
	t          s3.Operation
	bucketName string
	objectName string
	content    []byte
	record     bool
}

func init() {
	report.SummaryHandler = &report.Summary{}
}

func initWorkers(jobs chan task, results chan bool) {
	for i := 0; i < config.Script.Workers; i++ {
		worker := &report.Worker{
			Name:     fmt.Sprintf("worker-%d", i+1),
			Endpoint: config.Script.S3.Endpoints[i%len(config.Script.S3.Endpoints)],
		}
		report.SummaryHandler.Workers = append(report.SummaryHandler.Workers, worker)
		s3Handler := s3.New(config.Script.S3.Endpoints[0], config.Script.S3.AccessKey, config.Script.S3.SecretKey, config.Script.S3.Region)
		go func() {
			var errCount int = 0
			var totalCount int = 0
			for job := range jobs {
				switch job.t {
				case s3.CreateBucket:
					s3Handler.CreateBucket(&s3.CreateBucketInput{
						BucketName: job.bucketName,
					})
				case s3.PutObject:
					func() {
						if job.record {
							totalCount++
							start := time.Now()
							defer func() {
								worker.Speed = (float64(len(job.content)) / time.Now().Sub(start).Seconds())
								worker.TotalSize += len(job.content)
								worker.Count = totalCount
								worker.ErrorRate = fmt.Sprintf("%d%%", (errCount/totalCount)*100)
							}()
						}
						if _, err := s3Handler.Upload(&s3.UploadInput{
							Bucket:  job.bucketName,
							Key:     job.objectName,
							Content: job.content,
						}); err != nil {
							errCount++
						}
					}()
				case s3.GetObject:
					func() {
						var resp *s3.GetObjectOutput
						var err error
						if job.record {
							totalCount++
							start := time.Now()
							defer func() {
								if resp != nil {
									worker.Speed = (float64(resp.Size) / time.Now().Sub(start).Seconds())
									worker.TotalSize += resp.Size
								}
								worker.Count = totalCount
								worker.ErrorRate = fmt.Sprintf("%d%%", (errCount/totalCount)*100)
							}()
						}
						resp, err = s3Handler.GetObject(&s3.GetObjectInput{
							BucketName: job.bucketName,
							Key:        job.objectName,
						})
						if err != nil {
							errCount++
						}
					}()
				case s3.DeleteObject:
					s3Handler.DeleteObject(&s3.DeleteObjectInput{
						BucketName: job.bucketName,
						Key:        job.objectName,
					})
				case s3.DeleteBucket:
					s3Handler.DeleteBucket(&s3.DeleteBucketInput{
						BucketName: job.bucketName,
					})
				}
				results <- true
			}
		}()
	}
}

// Running ...
func Running() {
	taskQ := make(chan task, 100)
	resultQ := make(chan bool)
	initWorkers(taskQ, resultQ)
	for _, stage := range config.Script.Workstages {
		switch stage.Type {
		case s3.CreateBucket:
			for i := 0; i < stage.Buckets; i++ {
				taskQ <- task{
					t:          s3.CreateBucket,
					bucketName: fmt.Sprintf("%s-%d", stage.BucketPrefix, i+1),
					record:     false,
				}
				log.Infof("Create Bucket: %s", fmt.Sprintf("%s-%d", stage.BucketPrefix, i+1))
			}
			for i := 0; i < stage.Buckets; i++ {
				<-resultQ
			}
		case s3.PutObject:
			oMap := make(map[string]int)
			for _, size := range stage.Sizes {
				oMap[size.Size] = size.Rate
			}
			obj, err := object.New(oMap)
			if err != nil {
				os.Exit(1)
			}
			send := func(bucketName, objectName string, record bool) {
				taskQ <- task{
					t:          s3.PutObject,
					bucketName: bucketName,
					objectName: objectName,
					content:    obj.Get(),
					record:     true,
				}
			}
			if stage.Duration != 0 {
				report.SummaryHandler.Duration = int32(stage.Duration)
				var t int = 0
				timer(stage.Duration, &t)
				go func() {
					for {
						if t == stage.Duration {
							return
						}
						send(
							fmt.Sprintf("%s-%d", stage.BucketPrefix, rand.Intn(stage.Buckets)+1),
							fmt.Sprintf("%s-%d", stage.ObjectPrefix, rand.Intn(stage.Objects)+1),
							true,
						)
						<-resultQ
					}
				}()
				report.Start(&t)
			} else {
				for i := 0; i < stage.Buckets; i++ {
					for j := 0; j < stage.Objects; j++ {
						send(
							fmt.Sprintf("%s-%d", stage.BucketPrefix, i+1),
							fmt.Sprintf("%s-%d", stage.ObjectPrefix, j+1),
							false,
						)
						log.Infof("Upload file %s to bucket: %s", fmt.Sprintf("%s-%d", stage.BucketPrefix, i+1), fmt.Sprintf("%s-%d", stage.ObjectPrefix, j+1))
					}
				}
				for i := 0; i < stage.Buckets*stage.Objects; i++ {
					<-resultQ
				}
			}
		case s3.GetObject:
			if stage.Duration != 0 {
				report.SummaryHandler.Duration = int32(stage.Duration)
				var t int = 0
				timer(stage.Duration, &t)
				go func() {
					for {
						if t == stage.Duration {
							return
						}
						taskQ <- task{
							t:          s3.GetObject,
							bucketName: fmt.Sprintf("%s-%d", stage.BucketPrefix, rand.Intn(stage.Buckets)+1),
							objectName: fmt.Sprintf("%s-%d", stage.ObjectPrefix, rand.Intn(stage.Objects)+1),
							record:     true,
						}
						<-resultQ
					}
				}()
				report.Start(&t)
			}
		case s3.DeleteObject:
			for i := 0; i < stage.Buckets; i++ {
				for j := 0; j < stage.Objects; j++ {
					taskQ <- task{
						t:          s3.DeleteObject,
						bucketName: fmt.Sprintf("%s-%d", stage.BucketPrefix, i+1),
						objectName: fmt.Sprintf("%s-%d", stage.ObjectPrefix, j+1),
						record:     false,
					}
					log.Infof("Delete file %s in bucket: %s", fmt.Sprintf("%s-%d", stage.BucketPrefix, i+1), fmt.Sprintf("%s-%d", stage.ObjectPrefix, j+1))
				}
			}
			for i := 0; i < stage.Buckets*stage.Objects; i++ {
				<-resultQ
			}
		case s3.DeleteBucket:
			for i := 0; i < stage.Buckets; i++ {
				taskQ <- task{
					t:          s3.DeleteBucket,
					bucketName: fmt.Sprintf("%s-%d", stage.BucketPrefix, i+1),
					record:     false,
				}
				log.Infof("Delete bucket: %s", fmt.Sprintf("%s-%d", stage.BucketPrefix, i+1))
			}
			for i := 0; i < stage.Buckets; i++ {
				<-resultQ
			}
		default:
			log.Error("The S3 Operation is not suppoted.")
			os.Exit(1)
		}
	}
}

func timer(d int, t *int) {
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for {
			if *t >= d {
				return
			}
			select {
			case <-ticker.C:
				*t++
			}
		}
	}()
}
