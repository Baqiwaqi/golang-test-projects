package main

import (
	"encoding/json"
	"errors"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Goal struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
	Done      bool      `json:"done"`
}

var goals []Goal

// getGoals function
func getGoals(c *gin.Context) {
	result, err := checkGoals(goals)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"goals": "No goals found",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"goals": result,
		})

	}
}

func checkGoals(goals []Goal) (*[]Goal, error) {
	if len(goals) > 0 {
		return &goals, nil
	}
	return nil, errors.New("there are no goals stored")
}

// delete goals
func deleteGoal(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := deleteGoalbyID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"goals": "No goal found with id: " + strconv.Itoa(id),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"goals": "Goal deleted",
		})
	}
}

func deleteGoalbyID(id int) (*Goal, error) {
	for i, goal := range goals {
		if goal.ID == id {
			goals = append(goals[:i], goals[i+1:]...)
			return &goal, nil
		}
	}
	return &Goal{}, errors.New("no goal found")
}

// get single goal
func getGoal(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	goal, err := getGoalByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"goal": "Could not find goal",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"goal": goal,
		})

	}
}

func getGoalByID(id int) (*Goal, error) {
	for _, goal := range goals {
		if goal.ID == id {
			return &goal, nil
		}
	}
	return nil, errors.New("Goal not found")
}

// update goal
func updateGoal(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	goal, err := updateGoalById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"goal": "Could not find goal",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"goal": goal,
		})

	}
}

func updateGoalById(id int) (*Goal, error) {
	for _, goal := range goals {
		if goal.ID == id {
			deleteGoalbyID(id)
			goal.Done = true
			goals = append(goals, goal)
			return &goal, nil
		}
	}
	return &Goal{}, errors.New("Goal not found")
}

// create goal
func createGoal(c *gin.Context) {
	var goal Goal
	goal.ID = rand.Intn(100)
	_ = json.NewDecoder(c.Request.Body).Decode(&goal)
	goals = append(goals, goal)
	c.JSON(http.StatusOK, gin.H{
		"goals": goals,
	})
}

func main() {
	router := gin.Default()

	goals = append(goals, Goal{ID: 1, Title: "Read", CreatedAt: time.Now(), Done: false})
	goals = append(goals, Goal{ID: 2, Title: "Write", CreatedAt: time.Now(), Done: false})

	// getGoals fucntion
	router.GET("/goals", getGoals)

	// getGoal fucntion
	router.GET("/goals/:id", getGoal)

	// create Goal
	router.POST("/goals", createGoal)

	// update goals
	router.PUT("/goals/:id", updateGoal)

	// delete goal
	router.DELETE("/goals/:id", deleteGoal)

	router.Run()
}
