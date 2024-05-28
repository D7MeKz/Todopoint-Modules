package d7mongo

import (
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

type setupOptions struct {
	// env is a file path to authenicate the mongodb.
	env string
}

type SetupOption func(o *setupOptions) error

func WithEnv(env string) SetupOption {
	return func(o *setupOptions) error {
		o.env = env
		return nil
	}
}

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

func (m *setupOptions) loadEnv() {
	err := godotenv.Load(m.env)
	if err != nil {
		panic("Error loading .env file")
	}
}

func (m *setupOptions) getUri() string {
	m.loadEnv()
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		return "localhost:27017"
	}
	return uri
}

func (m *setupOptions) getAuth() options.Credential {
	// load env
	m.loadEnv()

	// extract httpdata
	username := os.Getenv("MONGODB_USERNAME")
	pw := os.Getenv("MONGODB_PASSWORD")

	if username != "" && pw != "" {
		username = "admin"
		pw = "admin"
	}
	return options.Credential{
		Username: username,
		Password: pw,
	}
}
