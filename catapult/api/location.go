package locationSvc

import (
	"encoding/json"

	"time"

	"github.com/BorisBorshevsky/GolangDemos/catapult"
	"github.com/BorisBorshevsky/GolangDemos/catapult/addons/cache"
	"github.com/BorisBorshevsky/GolangDemos/catapult/addons/cache/cache_provider"
	"github.com/BorisBorshevsky/GolangDemos/catapult/addons/circuit-breaker"
	"github.com/BorisBorshevsky/GolangDemos/catapult/addons/logger"
	"github.com/BorisBorshevsky/GolangDemos/catapult/addons/payload"
	"github.com/BorisBorshevsky/GolangDemos/catapult/addons/rurl"
	"github.com/BorisBorshevsky/GolangDemos/catapult/api/entities"
)

var cl *location

type Location interface {
	Alive(feature ...catapult.ClientFeature) (*entities.Alive, error)
	SuppliersLocation(env string, params map[string]string, features ...catapult.ClientFeature) (*entities.LocationServiceResponse, error)
}

type location struct {
	*catapult.Client
	SuppliersLocationFeatures []catapult.ClientFeature
}

func init() {
	//host := "location-scrum15.gett.io" //todo from env
	client := catapult.NewClient()
	cl = &location{
		Client: client,
	}

	cl.Wrap(rurl.Host("location.gtforge.com"))
	cl.Wrap(logger.Curl())
	cl.Wrap(payload.String("SomeString"))

	client.RequestTimeout = time.Millisecond * 2000

}

type RequestObject struct {
	A string                 `api:"header:a"`
	B string                 `api:"param:a"`
	C map[string]interface{} `api:"payload"`
}

func WrapRequest(req interface{}) func(*catapult.Request) *catapult.Request {
	return func(req *catapult.Request) *catapult.Request {
		req.Wrap(rurl.AddParam("test", "val"))
		return req
	}
}

func (l *location) Alive(requestObject RequestObject, feature ...catapult.ClientFeature) (*entities.Alive, error) {
	req := cl.NewRequest()
	//req.Wrap(WithOptions)

	req.Wrap(rurl.Path("/alive"))
	req.Wrap(cache.Add(cacheProvider.RedisTTLCache(time.Second * 10)))

	req.Wrap(circuitBreaker.Add())
	req.Wrap(cache.AddFallbackCache(cacheProvider.RedisTTLCache(time.Second * 300)))

	req.SetDecoder2(new(entities.Alive))

	req.SetDecoder(func(response *catapult.Response) (interface{}, error) {
		res := new(entities.Alive)
		err := json.NewDecoder(response).Decode(res)
		return res, err
	})

	resp, err := WrapRequest(requestObject)(req).Send()
	if err != nil {
		return nil, err
	}

	decoded, err := resp.Decode()
	if err != nil {
		return nil, err
	}

	return decoded.(*entities.Alive), err
}

func Alive(feature ...catapult.ClientFeature) (*entities.Alive, error) {
	return cl.Alive(RequestObject{}, feature...)
}
