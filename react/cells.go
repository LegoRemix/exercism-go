// Package react implements a simple reactive computation engine
package react

// basic represents a simple value with no other properties
type basic struct {
	value int     // the value of this given cell
	e     *engine // the pointer back to our main engine
}

// Value simply returns the set value for this basic cell
func (b basic) Value() int {
	return b.value

}

// input represents a basic cell which has a settable value
type input struct {
	basic
}

// SetValue simply returns the set value for this basic cell
func (i *input) SetValue(update int) {
	if i.value == update {
		return
	}

	i.value = update
	i.e.trigger(i)
}

// derived represents a cell whose value is dependent on a other inputs
type derived struct {
	basic                                       // we also are a basic cell
	updateCallback func()                       // the callback we use to update our basic value
	callbacks      map[CallbackHandle]func(int) // a map of all the callbacks registered to this
	handleGen      ident                        // the next id for our callback handle
}

// RemoveCallback removes a callback from a derived cell
func (d *derived) RemoveCallback(handle CallbackHandle) {
	delete(d.callbacks, handle)
}

// nextHandle gets the next id for a callback for this derived struct
func (d *derived) nextHandle() CallbackHandle {
	id := d.handleGen
	d.handleGen++
	return id
}

// AddCallback inserts an additional callback to be triggered when the value of derived changes
func (d *derived) AddCallback(cb func(int)) CallbackHandle {
	handle := d.nextHandle()
	d.callbacks[handle] = cb
	return handle
}

// trigger is run when a derived value is triggered, we update our value, and propogate any changes
func (d *derived) trigger() {
	prev := d.value
	d.updateCallback()
	curr := d.value
	//only if the value actually changed do we run our update callbacks
	if curr != prev {
		for _, cb := range d.callbacks {
			cb(curr)
		}
	}
}

// newDerived returns an initialized pointer to a derived cell
func newDerived(e *engine) *derived {
	d := &derived{
		callbacks: make(map[CallbackHandle]func(int)),
	}
	d.e = e
	return d
}
