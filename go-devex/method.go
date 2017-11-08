import "fmt"

type Secret struct {
	key string
}

func (s Secret) PrintKey() string {
	return fmt.Sprintf("%s", s.Key)
}
