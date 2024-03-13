package storage

type BoxProtocol interface {
	AddBox(box Box) error

	NewDirectory(name string)

	RemoveDirectory(name string)

	Push(...Commit)

	Pull()

	Clone()
}

type ossConn struct {
}

type nativeConn struct {
	Path  string
	Boxes map[string]Box
}

func (n nativeConn) AddBox(box Box) error {
	return nil
}

func (n nativeConn) NewDirectory(name string) {
	//TODO implement me
	panic("implement me")
}

func (n nativeConn) RemoveDirectory(name string) {
	//TODO implement me
	panic("implement me")
}

func (n nativeConn) Push(commit ...Commit) {
	//TODO implement me
	panic("implement me")
}

func (n nativeConn) Pull() {
	//TODO implement me
	panic("implement me")
}

func (n nativeConn) Clone() {
	//TODO implement me
	panic("implement me")
}

func NewNativeConn(path string) BoxProtocol {
	return nativeConn{
		Path:  path,
		Boxes: make(map[string]Box),
	}
}
