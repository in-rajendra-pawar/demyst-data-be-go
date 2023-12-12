package server

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strings"
	// "test/datastore"
	// "test/entity"
	// "test/helper"
)

var (
	getTaskRe    = regexp.MustCompile(`(?:/tasks\/[^/\s]+)*/?$`)
)

type TaskHandler struct {
	//InmemoryStore *datastore.InMemoryMap
}

func NewTaskHandler() *TaskHandler {
	/*
	datastore := datastore.NewInMemoryDataStore()
	return &TaskHandler{
		InmemoryStore: datastore,
	}
	*/
}

func (h *TaskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodGet && getTaskRe.MatchString(r.URL.Path):
		h.GetTask(w, r)
		return
	default:
		helper.NotFound(w, r)
		return
	}
}

func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
/*
	// swagger:operation GET /tasks/ GetTask
	//
	// Insert documentation
	//
	// ---
	// produces:
	// - application/json
	// responses:
	//   '200':
	//     description: task response
	//     schema:
	//       type: array
	//       items:
	//         "$ref": "#/definitions/Task"

	defer r.Body.Close()

	matches := getTaskRe.FindStringSubmatch(r.URL.Path)

	matchArr := strings.Split(matches[0], "/")

	taskIdIndex := len(matchArr) - 1
	taskId := matchArr[taskIdIndex]

	task, ok := h.InmemoryStore.GetData(taskId)

	if !ok {
		message := []byte(`{
			"success": false,
			"message": "Not found / Invalid Task ID",
		   }`)
		helper.ReturnJsonResponse(w, http.StatusNotFound, message)
		return
	}

	if task == nil {
		message := []byte(`{
			"success": false,
			"message": No record found
		   }`)
		helper.ReturnJsonResponse(w, http.StatusNotFound, message)
		return
	}

	data, err := json.MarshalIndent(task, "", "")

	if err != nil {
		log.Println(err)
		message := []byte(`{
			"success": false,
			"message": error marshalling data
		   }`)
		helper.ReturnJsonResponse(w, http.StatusInternalServerError, message)
		return
	}

	helper.ReturnJsonResponse(w, http.StatusOK, data)
	*/
}
