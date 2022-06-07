package shared

type Marshaler func(msg [][]byte) (ZeroMessage, error)
type UnMarshaler func(ZeroMessage) ([][]byte, error)
