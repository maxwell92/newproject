package resource

type IResource interface {
	Get() interface{}
	List() interface{}
	Update()
	Create()
	Delete()

	Encode() []byte
	Decode() string
}
