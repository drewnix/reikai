/*
Copyright © 2022 Andrew Tanner <drewnix>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/drewnix/reikai/cmd/reikai/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	log "k8s.io/klog/v2"
)

var (
	cfgFile string
	// serveOpts server.ServeOptions
	version = "devel"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd *cobra.Command

// rootCmd represents the base command when called without any subcommands
func newRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "reikai",
		Short: "Reikai is a micro-service that creates an API endpoint for controlling the spirit world.",
		PreRun: func(cmd *cobra.Command, args []string) {
			// serveOpts.UserAgent = getUserAgent(version, serveOpts.UserAgentComment)
			// log.Infof("kubeops has been configured with: %#v", serveOpts)
			log.Infof("reikai has been configured")
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return server.Serve()
		},
		Version: "devel",
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd = newRootCmd()
	rootCmd.SetVersionTemplate(version)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".reikai" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".reikai")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
