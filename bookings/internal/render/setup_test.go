package render

import (
	"encoding/gob"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/nickfang/bookings/internal/config"
	"github.com/nickfang/bookings/internal/models"
)

var session *scs.SessionManager
var testApp config.AppConfig

func TestMain(m *testing.M) {
	// what am I going to put in the session
	gob.Register(models.Reservation{})

	// change this to true when in production
	testApp.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour // 24 hours in seconds
	session.Cookie.Persist = true     // persist session across browser restarts
	session.Cookie.Secure = false
	session.Cookie.SameSite = http.SameSiteLaxMode
	testApp.Session = session

	app = &testApp

	os.Exit(m.Run())
}
