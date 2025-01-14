package vocab

import "net/url"

// The public key for an ActivityStreams actor
type ActivityStreamsPublicKeyProperty interface {
	// Clear ensures no value of this property is set. Calling
	// IsActivityStreamsPublicKey afterwards will return false.
	Clear()
	// Get returns the value of this property. When IsActivityStreamsPublicKey
	// returns false, Get will return any arbitrary value.
	Get() ActivityStreamsPublicKey
	// GetIRI returns the IRI of this property. When IsIRI returns false,
	// GetIRI will return any arbitrary value.
	GetIRI() *url.URL
	// GetType returns the value in this property as a Type. Returns nil if
	// the value is not an ActivityStreams type, such as an IRI or another
	// value.
	GetType() Type
	// HasAny returns true if the value or IRI is set.
	HasAny() bool
	// IsActivityStreamsPublicKey returns true if this property is set and not
	// an IRI.
	IsActivityStreamsPublicKey() bool
	// IsIRI returns true if this property is an IRI.
	IsIRI() bool
	// JSONLDContext returns the JSONLD URIs required in the context string
	// for this property and the specific values that are set. The value
	// in the map is the alias used to import the property's value or
	// values.
	JSONLDContext() map[string]string
	// KindIndex computes an arbitrary value for indexing this kind of value.
	// This is a leaky API detail only for folks looking to replace the
	// go-fed implementation. Applications should not use this method.
	KindIndex() int
	// LessThan compares two instances of this property with an arbitrary but
	// stable comparison. Applications should not use this because it is
	// only meant to help alternative implementations to go-fed to be able
	// to normalize nonfunctional properties.
	LessThan(o ActivityStreamsPublicKeyProperty) bool
	// Name returns the name of this property: "publicKey".
	Name() string
	// Serialize converts this into an interface representation suitable for
	// marshalling into a text or binary format. Applications should not
	// need this function as most typical use cases serialize types
	// instead of individual properties. It is exposed for alternatives to
	// go-fed implementations to use.
	Serialize() (interface{}, error)
	// Set sets the value of this property. Calling IsActivityStreamsPublicKey
	// afterwards will return true.
	Set(v ActivityStreamsPublicKey)
	// SetIRI sets the value of this property. Calling IsIRI afterwards will
	// return true.
	SetIRI(v *url.URL)
	// SetType attempts to set the property for the arbitrary type. Returns an
	// error if it is not a valid type to set on this property.
	SetType(t Type) error
}
