package tests

import (
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	task "github.com/marioTiara/todolistwebapi/Task"
	mockdb "github.com/marioTiara/todolistwebapi/Task/mock"
	"github.com/marioTiara/todolistwebapi/utils"
)

func TestPostTasksHandler(c *gin.Context) {
	taskReq := randomAddTaskRequest()

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(repo *mockdb.MockRepository)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"title":       taskReq.Title,
				"description": taskReq.Descryption,
				"file_name":   taskReq.FileName,
			},
			buildStubs: func(repo *mockdb.MockRepository) {
				repo.EXPECT().Create()
			},
		},
	}
}

func randomAddTaskRequest() (request task.AddTaskRequest) {
	request = task.AddTaskRequest{
		Title:       utils.RandomString(20),
		Descryption: utils.RandomString(50),
		FileName:    utils.RandomString(20),
	}

	return request
}

func randomTask() task.Task {
	task := task.Task{
		ID:          12,
		Title:       utils.RandomString(20),
		Descryption: utils.RandomString(150),
	}
	return task
}
