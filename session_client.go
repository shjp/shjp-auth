package auth

// SessionClient provides the top level interface client for user session
type SessionClient interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte) error
}

// SessionClientOptions provides the interface to populate a SessionClient
type SessionClientOptions interface {
	NewClient() (SessionClient, error)
	String() string
}
