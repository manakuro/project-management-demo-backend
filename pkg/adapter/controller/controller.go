package controller

// Controller holds the controllers for the entire across application
type Controller struct {
	Auth               interface{ Auth }
	Color              interface{ Color }
	FavoriteProject    interface{ FavoriteProject }
	FavoriteWorkspace  interface{ FavoriteWorkspace }
	Icon               interface{ Icon }
	Me                 interface{ Me }
	MyTasksTabStatus   interface{ MyTasksTabStatus }
	Project            interface{ Project }
	ProjectBaseColor   interface{ ProjectBaseColor }
	ProjectIcon        interface{ ProjectIcon }
	ProjectLightColor  interface{ ProjectLightColor }
	ProjectTeammate    interface{ ProjectTeammate }
	TaskColumn         interface{ TaskColumn }
	Teammate           interface{ Teammate }
	TeammateTaskColumn interface{ TeammateTaskColumn }
	TestTodo           interface{ TestTodo }
	TestUser           interface{ TestUser }
	Workspace          interface{ Workspace }
	WorkspaceTeammate  interface{ WorkspaceTeammate }
}
