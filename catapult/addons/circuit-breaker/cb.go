package circuitBreaker

import (
	"sync"

	"time"

	"github.com/BorisBorshevsky/GolangDemos/catapult"
	"github.com/pkg/errors"
	"github.com/rubyist/circuitbreaker"
)

type ctxKey int

var cbCtxKey = ctxKey(0)

var cbError = errors.New("circuit breaker - red light")
var cbStore = store{
	breakers: make(map[string]*circuit.Breaker),
	Mutex:    &sync.Mutex{},
}

type store struct {
	breakers map[string]*circuit.Breaker
	*sync.Mutex
}

type cbAddon struct {
	key     string
	options *circuit.Options
}

var DefaultOptions = &circuit.Options{
	ShouldTrip: circuit.ConsecutiveTripFunc(3),
	WindowTime: time.Second * 10,
}

func Add(options ...*circuit.Options) *cbAddon {
	if len(options) == 1 {
		return &cbAddon{options: options[0]}
	}

	return &cbAddon{options: DefaultOptions}
}

func (s *store) getOrCreate(key string, options *circuit.Options) *circuit.Breaker {
	s.Lock()
	defer s.Unlock()
	if breaker, ok := s.breakers[key]; ok {
		return breaker
	}

	breaker := circuit.NewBreakerWithOptions(options)
	s.breakers[key] = breaker

	return breaker
}

func (c *cbAddon) WithKey(cacheKey string) *cbAddon {
	c.key = cacheKey
	return c
}

func (c *cbAddon) Register(ctx *catapult.Ctx) {

	ctx.AddJustBefore(func(request *catapult.Request) {
		key := c.extractKey(request)
		ctx.Data[cbCtxKey] = key

		cb := cbStore.getOrCreate(key, c.options)

		if !cb.Ready() {
			ctx.SkipDispatch()
			ctx.SetError(cbError)
		}
	})

	ctx.AddAfter(func(response *catapult.Response) {
		cb := cbStore.getOrCreate(ctx.Data[cbCtxKey].(string), c.options)

		if response.Valid() { //todo when are we valid??
			cb.Success()
		} else {
			cb.Fail()
		}

	})

}

func (c *cbAddon) extractKey(request *catapult.Request) string {
	if c.key != "" {
		return c.key
	}
	//todo per path / host
	return request.Raw().URL.String()
}
