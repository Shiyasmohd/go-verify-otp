package api

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func evnACCOUNTSID() string {
	println("Loading .env file")
	println(godotenv.Unmarshal(".env"))
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("TWILIO_ACCOUNT_SID")
}
func evnAUTHTOKEN() string {
	println("Loading .env file")
	println(godotenv.Unmarshal(".env"))
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("TWILIO_AUTH_TOKEN")
}
func evnSERVICESSID() string {
	println("Loading .env file")
	println(godotenv.Unmarshal(".env"))
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("TWIIO_SERVICES_SID")
}
