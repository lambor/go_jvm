package heap
type SymRef struct {
	cp *ConstantPool
	className string
	class *Class
}

func newClassRef(cp *ConstantPool,classInfo *classfile.ConstantClassInfo) *ClassRef {
	ref:=&ClassRef{}
	ref.cp=cp
	ref.className=classInfo.Name()
	return ref
}
