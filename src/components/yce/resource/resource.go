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

type YceMetadata struct {
	SessionId string
	UserId    int32
	OrgId     int32
	DcIdList  []int32
}

type YceResource struct {
	Metadata *YceMetadata
	Spec     IResource
}
