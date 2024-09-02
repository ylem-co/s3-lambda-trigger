package command

func getBaseUrl(env string) string {
	switch env {
	case "production":
		return "https://api.ylem.co"
	case "test":
		return "https://api-test.ylem.co"
	}

	panic("unknown environment specified")
}
