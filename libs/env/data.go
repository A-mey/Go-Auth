package env

func getData(path string) map[string]interface{} {
	client := CreateClient()
	secret, err := client.Logical().Read(path)
	if err != nil {
		return nil
	}

	if secret == nil {
		return nil
	}

	return secret.Data
}
