package resource

type IResource interface{
	Get() interface{}
	List() interface{}
	Update()
	Create()
	Delete()
}
