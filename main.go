package main

type Counter struct {
    total           int
    increment_qty   int
    history         []int
}

func (c *Counter) increment() {
    c.history = append(c.history, c.total) // instead of calling c.history.append(c.total), go passes c.history as a parameter into the append function
    c.total += c.increment_qty
}

func (c *Counter) setIncrementQty(qty int) { // This function takes one parameter, 'qty' of type int
    c.increment_qty = qty
}

func (c Counter) getTotal() int { // getTotal() has a return type of 'int'
    return c.total
}

func main() {
    num := 5 // ':=' initializes 'num' with a value of 5
    num = 0 // After a variable is already initialized, using ':=' will result in an error

    counter := Counter{ // Here we are creating a new 'Counter' object and initializing some of its fields
        total: num,
        history: []int{}, // We are creating a blank slice here (basically an array). By default this would be nil (null)
        // We are choosing not to initialize increment_qty here so it will fall back to a default value which, for int, is 0
    }

    counter.setIncrementQty(5)
    counter.increment()

    println(counter.getTotal())
}

