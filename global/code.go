package global

const (
	ServerError = iota
	ClientError
)

type JikeApiKey struct {
	LastTime int64
	ApiKey   string
}

const (
	JikeApi = "https://api.jike.xyz"
)

var ApiKeys = [...]JikeApiKey{
	{
		LastTime: 0,
		ApiKey:"13630.491c6ecf9696138af75289baa5503eea.91daf0941c069585ca1388b002c6f588",
	},
	{
		LastTime: 0,
		ApiKey:"13752.31f345e1febd960b0ca0b67d599ccee2.04630e22a221d44073366182d83223fe",
	},
	{
		LastTime: 0,
		ApiKey:"13751.a0356c2e7406681e9157dc87d1a9a068.10befacfd5244590d4b9b93b536ead62",
	},
	{
		LastTime: 0,
		ApiKey:"13750.88665a02304f8c6596bd3d655be5bed3.d3e28a5ce0a4b229c548182c9fd9e50f",
	},
	{
		LastTime: 0,
		ApiKey:"13753.7eba1d61a2fe1f8601efe5dd074f8fed.5eba22bce65fe8e5fc50f9f047874c73",
	},
	{
		LastTime: 0,
		ApiKey:"13754.6f7bb5c0eb3fa3b70033853ea0cf0587.af6764b02a141466882f310303876a2c",
	},
}
