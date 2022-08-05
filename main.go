package main

import (
	"github.com/gin-gonic/gin"
	"errors"
	"net/http"
)

type Game struct {
	ID 			string  	`json:"id"`
	Title 		string 		`json:"title"`
	Developer 	string 		`json:"developer"`
	Price 		float64 	`json:"price"`
}

func (g *Game) setTitle(title string)  {
	g.Title = title
}

func (g *Game) setDeveloper(developer string)  {
	g.Developer = developer
}

func (g *Game) setPrice(price float64)  {
	g.Price = price
}

var games = []*Game{
	{ID: "1", Title: "Halo 3", Developer: "Bungie", Price: 56.99},
	{ID: "2", Title: "Call of Duty Moder Warfare", Developer: "Infinity Ward", Price: 56.99},
	{ID: "3", Title: "Gears of War", Developer: "Epic Games", Price: 56.99},
}

func getGames(context *gin.Context){
	context.IndentedJSON(http.StatusOK, games)
}

func getGame(context *gin.Context){
	id := context.Param("id")
	game, err := getGameById(id)
	if err != nil{
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "game not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, game)
}

func addGame(context *gin.Context){
	var newGame Game

	if err := context.BindJSON(&newGame); err != nil{
		return
	}

	games = append(games, &newGame)
	context.IndentedJSON(http.StatusCreated, newGame)
}

func putGame(context *gin.Context){
	id := context.Param("id")
	var newGame Game

	game, err := getGameById(id)

	if err != nil{
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "game not found"})
		return
	}

	if err := context.BindJSON(&newGame); err != nil{
		return
	}

	game.setTitle(newGame.Title)
	game.setDeveloper(newGame.Developer)
	game.setPrice(newGame.Price)
}

func deleteGame(context *gin.Context) {
	context.BindJSON(http.StatusOK)
	return
}


func getGameById(id string) (*Game, error){
	for i, g := range games{
		if g.ID == id{
			return games[i], nil
		}
	}
	return nil, errors.New("game not found")
}

func main() {
	router := gin.Default()
	router.GET("/games", getGames)
	router.GET("/games/:id", getGame)
	router.PUT("/games/:id", putGame)
	router.POST("/games", addGame)
	router.Run("localhost:9090")
}