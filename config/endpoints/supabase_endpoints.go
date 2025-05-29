package endpoints

const (
	AUTH_BASE   = "/auth/v1"
	TABLES_BASE = "/rest/v1"
)

var Supabase = struct {
	Auth struct {
		SignIn  string
		SignUp  string
		Refresh string
	}
	Tables struct {
		Profiles  string
		Content   string
		Watchlist string
	}
}{
	Auth: struct {
		SignIn  string
		SignUp  string
		Refresh string
	}{
		SignIn:  AUTH_BASE + "/token?grant_type=password",
		SignUp:  AUTH_BASE + "/signup",
		Refresh: AUTH_BASE + "/token?grant_type=refresh_token",
	},
	Tables: struct {
		Profiles  string
		Content   string
		Watchlist string
	}{
		Profiles:  TABLES_BASE + "/profiles",
		Content:   TABLES_BASE + "/content",
		Watchlist: TABLES_BASE + "/watchlist",
	},
}
