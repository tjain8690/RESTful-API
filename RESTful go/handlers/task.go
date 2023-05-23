package handlers

import (
    // Import necessary packages

    "your-project-name/models"
    "your-project-name/utils"
)

// ...

// UpdateTask updates an existing task
func UpdateTask(w http.ResponseWriter, r *http.Request) {
    // Parse task ID from URL
    taskID, err := strconv.Atoi(utils.GetURLParam(r, "id"))
    if err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, "Invalid task ID")
        return
    }

    // Read request body
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }

    // Find the task with the given ID
    var updatedTask models.Task
    for i, task := range tasks {
        if task.ID == taskID {
            // Update the task fields
            if err := json.Unmarshal(body, &updatedTask); err != nil {
                utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
                return
            }

            updatedTask.ID = taskID
            tasks[i] = updatedTask

            utils.RespondWithJSON(w, http.StatusOK, updatedTask)
            return
        }
    }

    utils.RespondWithError(w, http.StatusNotFound, "Task not found")
}

// DeleteTask deletes a task by its ID
func DeleteTask(w http.ResponseWriter, r *http.Request) {
    // Parse task ID from URL
    taskID, err := strconv.Atoi(utils.GetURLParam(r, "id"))
    if err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, "Invalid task ID")
        return
    }

    // Find the task with the given ID and remove it from the tasks slice
    for i, task := range tasks {
        if task.ID == taskID {
            tasks = append(tasks[:i], tasks[i+1:]...)
            utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Task deleted successfully"})
            return
        }
    }

    utils.RespondWithError(w, http.StatusNotFound, "Task not found")
}
