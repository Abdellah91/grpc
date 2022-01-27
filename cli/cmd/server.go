/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"context"

	students "github.com/Abdellah91/grpc/cli/pkg/students"
)

type Server struct {
	students.UnimplementedRegistrationServer
}

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server called")
		listener, err := net.Listen("tcp", ":4041")
		if err != nil {
			panic(err)
		}

		srv := grpc.NewServer()
		students.RegisterRegistrationServer(srv, &Server{})
		reflection.Register(srv)
		log.Printf("GRPC server listening on %v", listener.Addr())

		if e := srv.Serve(listener); e != nil {
			panic(e)
		}
	},
}

func (s *Server) AddStudent(ctx context.Context, request *students.Student) (*students.Decision, error) {

	firstName := request.GetFirstName()
	return &students.Decision{StudentId: "1", Message: firstName + " is registred with id " + "1"}, nil

}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
