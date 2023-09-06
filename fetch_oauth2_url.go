package main

func FetchOAuth2URL(registryHost string) (string, error) {
	u := fmt.Sprintf("https://%s/v2/", registryHost)
	res, err := http.Get(u)
	if err != nil {
		return "", err
	}

	return "", nil
}
