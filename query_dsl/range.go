package query

type Range struct {
	Gt     int    `json:"gt"`
	Gte    int    `json:"gte"`
	Lt     int    `json:"lt"`
	Lte    int    `json:"lte"`
	Format string `json:"format"`
}
