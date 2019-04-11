
type Atom interface {
	Interpret() Branch // make empty interface
}

type GenericAtom struct {}

func (atom *GenericAtom) Interpret() Branch {

}
