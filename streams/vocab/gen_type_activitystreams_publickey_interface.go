package vocab

// A public key represents a public cryptographical key for a user
type ActivityStreamsPublicKey interface {
	// GetActivityStreamsId returns the "id" property if it exists, and nil
	// otherwise.
	GetActivityStreamsId() ActivityStreamsIdProperty
	// GetActivityStreamsOwner returns the "owner" property if it exists, and
	// nil otherwise.
	GetActivityStreamsOwner() ActivityStreamsOwnerProperty
	// GetActivityStreamsPublicKeyPem returns the "publicKeyPem" property if
	// it exists, and nil otherwise.
	GetActivityStreamsPublicKeyPem() ActivityStreamsPublicKeyPemProperty
	// GetActivityStreamsType returns the "type" property if it exists, and
	// nil otherwise.
	GetActivityStreamsType() ActivityStreamsTypeProperty
	// GetTypeName returns the name of this type.
	GetTypeName() string
	// GetUnknownProperties returns the unknown properties for the PublicKey
	// type. Note that this should not be used by app developers. It is
	// only used to help determine which implementation is LessThan the
	// other. Developers who are creating a different implementation of
	// this type's interface can use this method in their LessThan
	// implementation, but routine ActivityPub applications should not use
	// this to bypass the code generation tool.
	GetUnknownProperties() map[string]interface{}
	// IsExtending returns true if the PublicKey type extends from the other
	// type.
	IsExtending(other Type) bool
	// JSONLDContext returns the JSONLD URIs required in the context string
	// for this type and the specific properties that are set. The value
	// in the map is the alias used to import the type and its properties.
	JSONLDContext() map[string]string
	// LessThan computes if this PublicKey is lesser, with an arbitrary but
	// stable determination.
	LessThan(o ActivityStreamsPublicKey) bool
	// Serialize converts this into an interface representation suitable for
	// marshalling into a text or binary format.
	Serialize() (map[string]interface{}, error)
	// SetActivityStreamsId sets the "id" property.
	SetActivityStreamsId(i ActivityStreamsIdProperty)
	// SetActivityStreamsOwner sets the "owner" property.
	SetActivityStreamsOwner(i ActivityStreamsOwnerProperty)
	// SetActivityStreamsPublicKeyPem sets the "publicKeyPem" property.
	SetActivityStreamsPublicKeyPem(i ActivityStreamsPublicKeyPemProperty)
	// SetActivityStreamsType sets the "type" property.
	SetActivityStreamsType(i ActivityStreamsTypeProperty)
	// VocabularyURI returns the vocabulary's URI as a string.
	VocabularyURI() string
}
