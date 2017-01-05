// Package react implements a basic reactive computation engine
package react

// id represents the unique ID of a given cell
type ident uint64

const testVersion = 4

// engine is the implementation of our reactive computation engine
type engine struct {
	dependencies map[Cell]map[*derived]bool
	topological  []*derived    // topologically sorted list of derived edges
	inputs       map[Cell]bool // set of input cells
}

// New returns a new Reactor implementation to handle computation on a set of cells
func New() *engine {
	return &engine{
		dependencies: make(map[Cell]map[*derived]bool),
		inputs:       make(map[Cell]bool),
	}
}

// CreateInput simples manufactors a basic input block
func (e *engine) CreateInput(initial int) InputCell {
	i := &input{basic: basic{value: initial, e: e}}
	e.inputs[i] = true
	return i
}

// getDepends returns all the cells which depends on a given modified cell
func (e *engine) getDepends(modified Cell) map[Cell]bool {
	result := make(map[Cell]bool)
	for d := range e.dependencies[modified] {
		result[d] = true
		for k := range e.getDepends(d) {
			result[k] = true
		}
	}
	return result
}

// trigger is responsible for updating all the dependent values of a cell with the provided id
func (e *engine) trigger(changed Cell) {
	depends := e.getDepends(changed)
	for _, d := range e.topological {
		if depends[d] {
			d.trigger()
		}
	}
}

// addDependent adds a new derived cell to the list
func (e *engine) addDependent(product *derived, inputs []Cell) {

	for _, source := range inputs {
		// first we create our map of dependencies if we haven't already
		if _, ok := e.dependencies[source]; !ok {
			e.dependencies[source] = make(map[*derived]bool)
		}
		//insert into the set of connected nodes
		e.dependencies[source][product] = true
	}

	sources := e.derivedDepends(inputs)

	//now we try to find the correct position in the topological sort
	i := 0
	for ; i < len(e.topological); i++ {
		c := e.topological[i]
		if sources[c] {
			delete(sources, c)
			if len(sources) == 0 {
				break
			}
		}
	}

	// if we hit the end of the list, just append
	if i == len(e.topological) {
		e.topological = append(e.topological, product)
	} else {
		e.topological = append(e.topological, nil)
		copy(e.topological[i+2:], e.topological[i+1:])
		e.topological[i+1] = product
	}

}

// derivedDepends returns a set of the derived dependieces from a list of candidate cells
func (e *engine) derivedDepends(cells []Cell) map[Cell]bool {
	result := make(map[Cell]bool)
	for _, c := range cells {
		if e.inputs[c] {
			continue
		}
		result[c] = true
	}
	return result
}

// CreateCompute1 generates a ComputeCell which is a unary function of its inputs
func (e *engine) CreateCompute1(operand Cell, cb func(int) int) ComputeCell {
	d := newDerived(e)
	f := func() {
		d.value = cb(operand.Value())
	}
	d.updateCallback = f
	d.updateCallback()
	e.addDependent(d, []Cell{operand})
	return d
}

// CreateCompute2 generates a ComputeCell which is a binary function of its inputs
func (e *engine) CreateCompute2(op1, op2 Cell, cb func(int, int) int) ComputeCell {
	d := newDerived(e)
	f := func() {
		d.value = cb(op1.Value(), op2.Value())
	}
	d.updateCallback = f
	d.updateCallback()
	e.addDependent(d, []Cell{op1, op2})
	return d
}
