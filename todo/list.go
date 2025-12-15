package todo

type List struct {
	tasks map[int]Task
}

func NewList() *List {

	return &List{
		tasks: make(map[int]Task),
	}
}

func (list *List) AddTask(task Task) error {

	if _, ok := list.tasks[task.Id]; ok {
		return ErrTaskAlreasyExist
	}

	list.tasks[task.Id] = task

	return nil
}

func (list *List) ListTasks() map[int]Task {

	tmp := make(map[int]Task, len(list.tasks))
	for k, v := range list.tasks {
		tmp[k] = v
	}

	return tmp
}

func (list *List) ListNotCompletedTasks() map[int]Task {
	notCompletedtask := make(map[int]Task)

	for k, v := range list.tasks {
		if !list.tasks[k].IsCompleted {
			notCompletedtask[k] = v
		}
	}
	return notCompletedtask

}
func (list *List) ListCompletedTasks() map[int]Task {
	completedTask := make(map[int]Task)

	for k, v := range list.tasks {
		if list.tasks[k].IsCompleted {
			completedTask[k] = v
		}
	}
	return completedTask

}

func (list *List) CompleteTask(id int) error {
	task, ok := list.tasks[id]
	if !ok {
		return ErrTaskNotFound
	}

	task.Complete()

	list.tasks[id] = task

	return nil

}

func (list *List) DeleteTask(id int) error {

	_, ok := list.tasks[id]

	if !ok {
		return ErrTaskNotFound
	}

	delete(list.tasks, id)

	return nil

}
