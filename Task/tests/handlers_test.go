package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	task "github.com/marioTiara/todolistwebapi/Task"
	mockdb "github.com/marioTiara/todolistwebapi/Task/mock"
	"github.com/marioTiara/todolistwebapi/utils"
	"github.com/stretchr/testify/require"
)

func TestPostTasksHandler(t *testing.T) {
	taskReq := randomAddTaskRequest()
	taskinput := task.Task{
		Title:       taskReq.Title,
		Descryption: taskReq.Descryption,
	}

	taskRepoResponse := task.Task{
		ID:          12,
		Title:       taskReq.Title,
		Descryption: taskReq.Descryption,
	}

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
				repo.EXPECT().Create(gomock.Eq(taskinput)).Times(1).
					Return(taskRepoResponse)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mockdb.NewMockRepository(ctrl)
			tc.buildStubs(repo)

			// server:= ne
			// //Marshal body data to JSON
			// data, err:= json.Marshal(tc.body)
			// require.NoError(t, err)

			// url:="/v1/tasks"
			// request, err:= http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			// require.NoError(t,err)
			// ser

		})
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
