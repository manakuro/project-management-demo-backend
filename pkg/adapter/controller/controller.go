package controller

// Controller holds the controllers for the entire across application.
type Controller struct {
	Auth                    interface{ Auth }
	Color                   interface{ Color }
	FavoriteProject         interface{ FavoriteProject }
	FavoriteWorkspace       interface{ FavoriteWorkspace }
	FileType                interface{ FileType }
	Icon                    interface{ Icon }
	Me                      interface{ Me }
	Project                 interface{ Project }
	ProjectBaseColor        interface{ ProjectBaseColor }
	ProjectIcon             interface{ ProjectIcon }
	ProjectLightColor       interface{ ProjectLightColor }
	ProjectTask             interface{ ProjectTask }
	ProjectTaskColumn       interface{ ProjectTaskColumn }
	ProjectTaskListStatus   interface{ ProjectTaskListStatus }
	ProjectTaskSection      interface{ ProjectTaskSection }
	ProjectTeammate         interface{ ProjectTeammate }
	Tag                     interface{ Tag }
	Task                    interface{ Task }
	TaskCollaborator        interface{ TaskCollaborator }
	TaskColumn              interface{ TaskColumn }
	TaskFeed                interface{ TaskFeed }
	TaskFeedLike            interface{ TaskFeedLike }
	TaskFile                interface{ TaskFile }
	TaskLike                interface{ TaskLike }
	TaskListCompletedStatus interface{ TaskListCompletedStatus }
	TaskListSortStatus      interface{ TaskListSortStatus }
	TaskPriority            interface{ TaskPriority }
	TaskSection             interface{ TaskSection }
	TaskTag                 interface{ TaskTag }
	Teammate                interface{ Teammate }
	TeammateTask            interface{ TeammateTask }
	TeammateTaskColumn      interface{ TeammateTaskColumn }
	TeammateTaskListStatus  interface{ TeammateTaskListStatus }
	TeammateTaskSection     interface{ TeammateTaskSection }
	TeammateTaskTabStatus   interface{ TeammateTaskTabStatus }
	TestTodo                interface{ TestTodo }
	TestUser                interface{ TestUser }
	Workspace               interface{ Workspace }
	WorkspaceTeammate       interface{ WorkspaceTeammate }
}
