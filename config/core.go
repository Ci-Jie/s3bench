package config

import "s3bench/s3"

// TestingScript ...
type TestingScript struct {
	Version     string      `yaml:"version"`
	Name        string      `yaml:"name"`
	Description string      `yaml:"description"`
	Workers     int         `yaml:"workers"`
	S3          S3          `yaml:"s3"`
	Workstages  []Workstage `yaml:"workstages"`
}

// Workstage ...
type Workstage struct {
	Type         s3.Operation `yaml:"type"`
	Duration     int          `yaml:"duration"`
	Buckets      int          `yaml:"buckets"`
	BucketPrefix string       `yaml:"bucket_prefix"`
	ObjectPrefix string       `yaml:"object_prefix"`
	Sizes        []Size       `yaml:"sizes"`
	Objects      int          `yaml:"objects"`
}

// S3 ...
type S3 struct {
	Endpoints []string `yaml:"endpoints"`
	AccessKey string   `yaml:"access_key"`
	SecretKey string   `yaml:"secret_key"`
	Region    string   `yaml:"region"`
}

// Size ...
type Size struct {
	Size string `yaml:"size"`
	Rate int    `yaml:"rate"`
}

// Script ...
var Script *TestingScript

func init() {
	Script = &TestingScript{}
}
