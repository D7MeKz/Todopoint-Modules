package d7mysql

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"modules/v2/d7mysql/ent"
	"os"
)

type EntClient struct {
	options *setupOptions
	Client  *EntClient
}

type setupOptions struct {
	// env is a file path to authenicate the mongodb.
	env string
}

type SetupOption func(o *setupOptions) error

//func WithEnv(env string) SetupOption {
//	return func(o *setupOptions) error {
//		o.env = env
//		return nil
//	}
//}

func mergeSetupOptions(opts ...SetupOption) *setupOptions {
	var o setupOptions
	for _, opt := range opts {
		err := opt(&o)
		if err != nil {
			return nil
		}
	}
	return &o
}

func (m *setupOptions) getDsn() string {
	//err := godotenv.Load(m.env)
	//if err != nil {
	//	panic("Error loading .env file")
	//}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))
	return dsn
}

func NewEntClient(opt ...SetupOption) (*ent.Client, error) {
	opts := mergeSetupOptions(opt...)
	dsn := opts.getDsn()

	client, err := ent.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	// Run auto migration tool
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client, err

}
