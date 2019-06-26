package teammodels

type TeamRequest struct {
	Name    string
	LogoURL string
}

type Team struct {
	TeamID  string
	Name    string
	LogoURL string
	// Array of league ids to which this team plays in
	Leagues []string
}
