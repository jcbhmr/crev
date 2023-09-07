package main

func FindOAuth2EndpointURL(registryHost string) (string, error) {
	u := fmt.Sprintf("https://%s/v2/", registryHost)
	res, err := http.Get(u)
	if err != nil {
		return "", err
	}

	return "", nil
}
