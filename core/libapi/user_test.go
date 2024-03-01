package libapi

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/openaccounting/oa-server/core/model"
	"github.com/openaccounting/oa-server/core/model/db"
	"github.com/openaccounting/oa-server/core/model/types"
	"github.com/openaccounting/oa-server/core/util"
)

func TestMain(m *testing.M) {
	cnf := types.Config{
		WebUrl:          "https://domain.com",
		Address:         "",
		Port:            8080,
		ApiPrefix:       "",
		KeyFile:         "",
		CertFile:        "",
		DatabaseAddress: "192.168.0.176:31878",
		Database:        "openaccounting",
		User:            "openaccounting",
		Password:        "openaccounting",
		MailgunDomain:   "mg.domain.com",
		MailgunKey:      "",
		MailgunEmail:    "noreply@domain.com",
		MailgunSender:   "Sender",
	}

	connectionString := cnf.User + ":" + cnf.Password + "@tcp(" + cnf.DatabaseAddress + ")/" + cnf.Database

	db, err := db.NewDB(connectionString)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to connect to database with: %s", err.Error()))
	}

	bc := &util.StandardBcrypt{}

	model.NewModel(db, bc, cnf)

	os.Exit(m.Run())
}

func TestCreateUser(t *testing.T) {
	u := types.User{
		Id:           "111",
		FirstName:    "Akshay",
		LastName:     "Dua",
		Email:        "dakshay@gmail.com",
		Password:     "babis",
		AgreeToTerms: true,
	}

	if e1 := CreateUser(&u); e1 != nil {
		t.Error("Could not create user", u, "Error is", e1)
	}
}
