package controller

// Controller holds the controllers for the entire across application
type Controller struct {
	Auth            interface{ Auth }
	Color           interface{ Color }
	Icon            interface{ Icon }
	Project         interface{ Project }
	ProjectTeammate interface{ ProjectTeammate }
	Teammate        interface{ Teammate }
	TestTodo        interface{ TestTodo }
	TestUser        interface{ TestUser }
	Workspace       interface{ Workspace }
}
