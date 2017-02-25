package routes

import (
	"net/http"
    "github.com/labstack/echo"
)

type App struct {
    Id              int     `json:"id"`
    Name            string  `json:"name"`
    Description     string  `json:"description"`
    Requirements    []string  `json:"requirements"`
    Icon            string  `json:"icon"`
}
        
// GET ~/apps
func getApps() echo.HandlerFunc {
    return func(c echo.Context) (err error) {
        // Consider moving into SQL later
        var apps []*App
        apps = make([]*App, 2)
        apps[0] = &App{
            Id: 1,
            Name: "Memory",
            Icon: "/src/Images/sample.jpg",          
        }
        apps[1] = &App{
            Id: 2,
            Name: "Charades",
            Icon: "/src/Images/sample.jpg",        
        }        
        return c.JSON(http.StatusOK, &apps)
    }
}

// GET ~/app/{id}
func getApp() echo.HandlerFunc {
    return func(c echo.Context) (err error) {
        app := &App{
            Id: 1,
            Name: "Memory",
            Description: "Description",
            Requirements: []string{"req 1", "req 2"},
            Icon: "/src/Images/sample.jpg",            
        }                  
        return c.JSON(http.StatusOK, &app)
    }
}