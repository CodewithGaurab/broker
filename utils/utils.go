package utils

type Data struct {
	Head     []byte
	Body     []byte
	Exchange string
	Type     int
}

// Type 1 - Produced. 0 - Consume
