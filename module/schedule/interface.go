package schedule

type Model interface {
	Update(id int, schedule string) error
}
