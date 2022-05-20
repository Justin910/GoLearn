package main

type StruDBDelegate struct {
	Err error
}

type StruDBRow struct {
	err error
}

func (dd *StruDBDelegate) QueryRow(query string, args ...interface{}) DBRow {
	return &StruDBRow{
		err: dd.Err,
	}
}

func (ddr *StruDBRow) Scan(dest ...interface{}) error {

	if ddr.err != nil {
		return ddr.err
	}

	for i, _ := range dest {
		switch d := dest[i].(type) {
		case *string:
			if d == nil {
				return ddr.err
			}
			*d = string("小明")
		}
	}

	return ddr.err
}
