package tmpl

type Middlewares interface {
	GenerateToken(userId int64) (string, error)
}
