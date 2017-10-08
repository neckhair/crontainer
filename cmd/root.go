// Copyright © 2017 Philippe Hässig <phil@neckhair.ch>

package cmd

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/neckhair/crontainer/crontainer"
)

var cfgFile string

func waitForQuit() {
	var endWaiter sync.WaitGroup
	endWaiter.Add(1)

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

	go func() {
		// Block until the channel receives a signal
		<-signalChannel

		endWaiter.Done()
		log.Println("Byebye, it was a pleasure serving you.")
	}()

	endWaiter.Wait()
}

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "crontainer",
	Short: "Like cron, but for a single user",
	Long: `crontainer runs regular tasks defined in its config file.

It is mainly inteded to be run inside a Docker container and
designed to be run as an unprivileged user.`,

	PreRun: func(cmd *cobra.Command, args []string) {
	},

	Run: func(cmd *cobra.Command, args []string) {
		log.Println("---> Begin scheduling <---")

		engine := crontainer.NewCronEngine()
		engine.Initialize(viper.GetViper())
		engine.Start()
		defer engine.Stop()

		waitForQuit()
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.crontainer.yaml)")

	RootCmd.Flags().String("command", "", "Command to run")
	RootCmd.Flags().String("schedule", "* * * * * *", "Cron like schedule including seconds")

	viper.BindPFlag("command", RootCmd.Flags().Lookup("command"))
	viper.BindPFlag("schedule", RootCmd.Flags().Lookup("schedule"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName(".crontainer") // name of config file (without extension)
		viper.AddConfigPath("$HOME")       // adding home directory as first search path
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
