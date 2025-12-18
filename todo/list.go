package todo

import "github.com/google/uuid"

type List struct {
	tasks map[uuid.UUID]Task
}

func NewList() *List {

	return &List{
		tasks: make(map[uuid.UUID]Task),
	}
}

func (list *List) AddTask(task Task) error {

	if _, ok := list.tasks[task.Id]; ok {
		return ErrTaskAlreadyExist
	}

	list.tasks[task.Id] = task

	return nil
}

func (list *List) ListTasks() map[uuid.UUID]Task {

	tmp := make(map[uuid.UUID]Task, len(list.tasks))
	for k, v := range list.tasks {
		tmp[k] = v
	}

	return tmp
}

func (list *List) GetTask(id uuid.UUID) (Task, error) {

	foundedTask, ok := list.tasks[id]

	if !ok {
		return Task{}, ErrTaskNotFound
	}
	return foundedTask, nil

}

func (list *List) ListUnCompletedTasks() map[uuid.UUID]Task {
	notCompletedtask := make(map[uuid.UUID]Task)

	for k, v := range list.tasks {
		if !list.tasks[k].IsCompleted {
			notCompletedtask[k] = v
		}
	}
	return notCompletedtask

}
func (list *List) ListCompletedTasks() map[uuid.UUID]Task {
	completedTask := make(map[uuid.UUID]Task)

	for k, v := range list.tasks {
		if list.tasks[k].IsCompleted {
			completedTask[k] = v
		}
	}
	return completedTask

}

func (list *List) CompleteTask(id uuid.UUID) error {
	task, ok := list.tasks[id]
	if !ok {
		return ErrTaskNotFound
	}

	task.Complete()

	list.tasks[id] = task

	return nil

}

func (list *List) DeleteTask(id uuid.UUID) error {

	_, ok := list.tasks[id]

	if !ok {
		return ErrTaskNotFound
	}

	delete(list.tasks, id)

	return nil

}
