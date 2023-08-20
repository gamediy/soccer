package soccer

const (
	PlayTypeHandicap = "Handicap"
)

type Play interface {
	WonCheck(openResult OpenResult, calcRule string) (err error) //结算
	
}

type OpenResult struct {
	Result     string
	BoutStatus int
}
