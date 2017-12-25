package catapult

type ClientFeature interface {
	Register(ctx *Ctx)
}
