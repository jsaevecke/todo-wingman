package task

type Option func(*task)

type task struct {
	status      Status
	title       string
	description string
	tags        []string
}

func (task task) FilterValue() string {
	return task.title
}

func (task task) Title() string {
	return task.title
}

func (task task) Description() string {
	return task.description
}

func New(title string, options ...Option) task {
	task := task{
		title: title,
	}

	for _, option := range options {
		option(&task)
	}

	return task
}

func WithDescription(description string) Option {
	return func(task *task) {
		task.description = description
	}
}

func WithTags(tags []string) Option {
	return func(task *task) {
		task.tags = tags
	}
}

func WithStatus(status Status) Option {
	return func(task *task) {
		task.status = status
	}
}
