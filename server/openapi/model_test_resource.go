/*
 * TraceTest
 *
 * OpenAPI definition for TraceTest endpoint and resources
 *
 * API version: 0.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// TestResource - Represents a test structured into the Resources format.
type TestResource struct {

	// Represents the type of this resource. It should always be set as 'Test'.
	Type string `json:"type,omitempty"`

	Spec Test `json:"spec,omitempty"`
}

// AssertTestResourceRequired checks if the required fields are not zero-ed
func AssertTestResourceRequired(obj TestResource) error {
	if err := AssertTestRequired(obj.Spec); err != nil {
		return err
	}
	return nil
}

// AssertRecurseTestResourceRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of TestResource (e.g. [][]TestResource), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseTestResourceRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aTestResource, ok := obj.(TestResource)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertTestResourceRequired(aTestResource)
	})
}
