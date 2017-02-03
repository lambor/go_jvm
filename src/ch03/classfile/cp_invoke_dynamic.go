package classfile
type ConstantMethodTypeInfo struct {
	descriptor_index uint16
}

func (self *ConstantMethodTypeInfo) readInfo(reader *ClassReader) {
	self.descriptor_index = reader.readUint16()
}

type ConstantMethodHandleInfo struct {
	reference_kind uint8
	reference_index uint16
}

func (self *ConstantMethodHandleInfo) readInfo(reader *ClassReader) {
	self.reference_kind = reader.readUint8()
	self.reference_index = reader.readUint16()
}

type ConstantInvokeDynamicInfo struct {
	bootstrap_method_attr_index uint16
	name_and_type_index uint16
}

func (self *ConstantInvokeDynamicInfo) readInfo(reader *ClassReader) {
	self.bootstrap_method_attr_index = reader.readUint16()
	self.name_and_type_index = reader.readUint16()
}
