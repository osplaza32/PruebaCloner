package Entidad
type Version struct {
	Content []byte
	Version time.Time
	Name uuid.UUID
	Status bool
	Deleted time.Time
}