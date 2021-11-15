package cmd

import (
	"io/ioutil"
	"os"
	"s3bench/config"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var rootCmd = &cobra.Command{
	Use:   "S3bench",
	Short: "",
	Long:  "",
}

var log = logrus.New()
var filePath string

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(newStartCmd())
	rootCmd.PersistentFlags().StringVarP(&filePath, "file", "f", "", "Provide a testing file")
}

// Execute ...
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}

func initConfig() {
	log.Out = os.Stdout
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	if err := yaml.Unmarshal(content, config.Script); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
