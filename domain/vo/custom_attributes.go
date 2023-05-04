package vo

type CustomAttributes map[string]any

func NewCustomAttributes(customAttributes map[string]any) CustomAttributes {
	return CustomAttributes(customAttributes)
}

func (vo *CustomAttributes) Validate() error {
	return nil
}

func (vo CustomAttributes) GetAttribute(attributeName string) (attributeValue any, find bool) {
	attributeValue, find = vo[attributeName]
	return attributeValue, find
}
