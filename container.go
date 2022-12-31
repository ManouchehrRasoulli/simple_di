package simpledi

// Option
// inorder to configure container you can use provided options
type Option func(c *container) error

type container struct {
	// providers
	// list of providers who are introduced into container
	providers []provider

	// invokers
	// list of invokes who container have to peforms
	invoker *invoker

	// collection
	// list of provided items who we have inside container and we are going to use them
	// to satisfy goven funcitons
	collection []input
}

// New
// atlease you have to provide one option for container
func New(ops ...Option) (*container, error) {
	c := container{
		providers:  nil,
		invoker:    nil,
		collection: make([]input, 0),
	}

	for _, op := range ops {
		err := op(&c)
		if err != nil {
			return nil, err
		}
	}

	if c.invoker == nil {
		return nil, DiNeedInvoke
	}

	return &c, nil
}

func (c *container) Run() error {
	retryCout := 0
	runnedProviders := 0
retry:
	for {
		anyOneRunned := false
		for i := 0; i < len(c.providers); i++ {
			if c.providers[i].isCalled {
				continue
			}

			if c.providers[i].readytogo(c.collection) {
				outputs := c.providers[i].call(c.collection)
				if len(outputs) > 0 {
					for _, o := range outputs { // outputs are gone used as input for others
						i := input{
							name:  o.name,
							typ:   o.typ,
							value: o.value,
						}

						c.collection = append(c.collection, i)
					}
				}
				anyOneRunned = true
				runnedProviders++
			}
		}

		if anyOneRunned && runnedProviders == len(c.providers) {
			break retry
		}

		if retryCout > 1 {
			return DiCycleDetected
		}
		retryCout++
	}

	// run invokes
	if !c.invoker.readytogo(c.collection) {
		return DiInvokeNotSatisfied
	}

	return c.invoker.call(c.collection)
}
