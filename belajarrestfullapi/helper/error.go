package helper

func ReturnIfError(data interface{}, err error) (dataReturn interface{}, errReturn error) {
	if err != nil {
		return nil, err
	}
	return data, nil
}

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
