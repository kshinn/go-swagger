package validate

// Validatable types implementing this interface allow customizing their validation
// this will be used instead of the reflective valditation based on the spec document.
// the implementations are assumed to have been generated by the swagger tool so they should
// contain all the validations obtained from the spec
type Validatable interface {
	Validate() *Result
}