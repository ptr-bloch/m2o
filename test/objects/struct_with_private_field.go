package objects

type StructWithPrivateField struct {
	private int
}

func NewStructWithPrivateField() StructWithPrivateField {
	return StructWithPrivateField{
		private: 100,
	}
}
