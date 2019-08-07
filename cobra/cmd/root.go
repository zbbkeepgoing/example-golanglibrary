/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
  "fmt"
  "os"
  "github.com/spf13/cobra"

  homedir "github.com/mitchellh/go-homedir"
  "github.com/spf13/viper"

)


var cfgFile string

var studentCmd = &cobra.Command{
  Use:   "student",
  Short: "this is a command of student",
  Long:  `a clent of student which you can use this client to query student by some flag`,
  PreRun: func(cmd *cobra.Command, args []string) {
    fmt.Printf("Inside student PreRun with args: %v\n", args)
  },
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Printf("Inside student Run with args: %v\n", args)
  },
  PostRun: func(cmd *cobra.Command, args []string) {
    fmt.Printf("Inside student PostRun with args: %v\n", args)
  },
  PersistentPostRun: func(cmd *cobra.Command, args []string) {
    fmt.Printf("Inside student PersistentPostRun with args: %v\n", args)
  },
}

var teacherCmd = &cobra.Command{
  Use:   "teacher",
  Short: "this is a command of teacher",
  Long:  `a clent of teacher which you can use this client to query teacher by some flag`,
  PreRun: func(cmd *cobra.Command, args []string) {
    fmt.Printf("Inside teacher PreRun with args: %v\n", args)
  },
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Printf("Inside teacher Run with args: %v\n", args)
  },
  PostRun: func(cmd *cobra.Command, args []string) {
    fmt.Printf("Inside teacher PostRun with args: %v\n", args)
  },

}



// rootCmd represents the base command when called without any subcommands
var userCmd = &cobra.Command{
  Use:   "user",
  Short: "this is a command of user",
  Long: `a clent of user which you can use this client to query user by some flag`,
  // Uncomment the following line if your bare application
  // has an action associated with it:
  Args: func(cmd *cobra.Command, args []string) error {
    fmt.Printf("args is : %s \n",args[0])
    return nil
  },
  PersistentPreRun: func(cmd *cobra.Command, args []string) {
    fmt.Printf("Inside user PersistentPreRun with args: %v\n", args)
  },
  PreRun: func(cmd *cobra.Command, args []string) {
    fmt.Printf("Inside user PreRun with args: %v\n", args)
  },
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Printf("Inside user Run with args: %v\n", args)
  },
  PostRun: func(cmd *cobra.Command, args []string) {
    fmt.Printf("Inside user PostRun with args: %v\n", args)
  },
  PersistentPostRun: func(cmd *cobra.Command, args []string) {
    fmt.Printf("Inside user PersistentPostRun with args: %v\n", args)
  },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
  if err := userCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func init() {
  userCmd.PersistentFlags().StringP("UserName","n","xxx","the username of this User,you can select from username.default xxx")
  userCmd.PersistentFlags().StringP("UserId","i","0","the id of this User,you can select from id,default 0")
  userCmd.PersistentFlags().IntP("Sex","s",0,"female is 0,male is 1,defalut 0")
  userCmd.LocalFlags().StringP("Profession","p","student","the profession of user,default student")

  studentCmd.LocalFlags().String("School","PrimarySchool","the school of student,default student")

  userCmd.MarkFlagRequired("UserName")
  userCmd.AddCommand(studentCmd)
  userCmd.AddCommand(teacherCmd)
}


// initConfig reads in config file and ENV variables if set.
func initConfig() {
  if cfgFile != "" {
    // Use config file from the flag.
    viper.SetConfigFile(cfgFile)
  } else {
    // Find home directory.
    home, err := homedir.Dir()
    //fmt.Printf("path of home is : %s \n",home)
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    // Search config in home directory with name ".examplecobra" (without extension).
    viper.AddConfigPath(home)
    viper.SetConfigName(".cobra")
  }

 // viper.AutomaticEnv() // read in environment variables that match

  // If a config file is found, read it in.
  if err := viper.ReadInConfig(); err != nil {
    fmt.Println("Can't read config:", err)
    os.Exit(1)
  }
}

