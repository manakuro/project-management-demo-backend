package controller

// Controller holds the controllers for the entire across application
type Controller struct {
	TestUser interface{ TestUser }
}
