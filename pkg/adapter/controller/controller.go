package controller

// Controller holds the controllers for the entire across application
type Controller struct {
	TestUser  interface{ TestUser }
	TestTodo  interface{ TestTodo }
	Auth      interface{ Auth }
	Teammate  interface{ Teammate }
	Workspace interface{ Workspace }
	Color     interface{ Color }
	Icon      interface{ Icon }
}
