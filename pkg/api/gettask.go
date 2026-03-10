package api

import (
	"net/http"
	"strconv"

	"github.com/GeaRvant/go_finalProject/pkg/db"
)

type TaskDTO struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func getTasksHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	tasks, err := db.Tasks(50)
	if err != nil {
		writeError(w, "Failed to fetch tasks", http.StatusInternalServerError)
		return
	}

	resp := []TaskDTO{}

	for _, t := range tasks {
		resp = append(resp, TaskDTO{
			ID:      strconv.FormatInt(t.ID, 10),
			Date:    t.Date,
			Title:   t.Title,
			Comment: t.Comment,
			Repeat:  t.Repeat,
		})
	}

	writeJson(w, map[string]any{
		"tasks": resp,
	})
}
