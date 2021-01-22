package main

import (
	"fmt"

	"github.com/cz-theng/czkit-go/log"
	"github.com/cz-theng/ethspider"
	"github.com/spf13/cobra"
)

var (
	_version    bool
	_configFile string
)

var rootCMD = &cobra.Command{
	Use:   "ethspider",
	Short: "start an ethspider",
	Long:  `start an ethspider`,
	Run:   _main,
}

func init() {
	rootCMD.PersistentFlags().BoolVarP(&_version, "version", "v", false, "print version of ethspider")
	rootCMD.PersistentFlags().StringVarP(&_configFile, "config", "c", "", "config file path for ethspider")
}

func main() {
	rootCMD.Execute()
}

func dumpVersion() {
	fmt.Printf("%s\n", ethspider.Version())
}

var (
	_spider ethspider.Spider
)

func _main(cmd *cobra.Command, args []string) {
	if _version {
		dumpVersion()
		return
	}
	if len(_configFile) > 0 {
		if err := loadConfig(_configFile); err != nil {
			return
		}
	} else {
		cmd.Usage()
		return
	}

	logFile := "ethspider.log"
	logPath := "./"
	if len(config.LogPath) > 0 {
		logPath = config.LogPath
	}
	if len(config.LogFile) > 0 {
		logFile = config.LogFile
	}

	logNameOpt := log.WithLogName(logFile)
	logPathOpt := log.WithLogPath(logPath)
	log.Init(logNameOpt, logPathOpt)

	rpcAddrOpt := ethspider.WithRPCAddr(config.RPCAddr)
	_spider.Init(rpcAddrOpt)
	_spider.Start()
}
