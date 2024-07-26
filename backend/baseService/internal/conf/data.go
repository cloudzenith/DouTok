package conf

type Data struct {
	Database struct {
		Source string
	}
	Redis struct {
		Source   string
		Password string
	}
}
