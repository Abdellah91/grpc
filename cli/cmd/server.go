/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"net"
	"net/url"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"context"

	"github.com/zemirco/couchdb"

	"os"

	students "github.com/Abdellah91/grpc/cli/pkg/students"

	"github.com/joho/godotenv"
)

type Server struct {
	students.UnimplementedRegistrationServer
}

type studentDocument struct {
	couchdb.Document
	FirstName string
	LastName  string
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
	lastName := request.GetLastName()
	//getEnvVars()
	dbUrl := os.Getenv("DB_URL")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	u, err := url.Parse(dbUrl)
	if err != nil {
		panic(err)
	}
	// create a new client
	client, err := couchdb.NewAuthClient(username, password, u)
	if err != nil {
		panic(err)
	}
	//client, err := getDbClient(dbUrl, username, password)

	db := client.Use("registration")
	doc := &studentDocument{
		FirstName: firstName,
		LastName:  lastName,
	}
	res, err := db.Post(doc)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
	return &students.Decision{StudentId: firstName, Message: lastName + " is registred"}, nil

}

func getEnvVars() {

	err := godotenv.Load("/home/properties.env")
	if err != nil {
		panic(err)
	}
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
