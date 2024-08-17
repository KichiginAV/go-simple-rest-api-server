package handlers

import "fmt"

type result struct {
	OK  bool   `json:"ok"`
	Msg string `json:"message"`
}

func ResponseErr(s string) result {
	return result{
		OK:  false,
		Msg: s,
	}
}

func ResponseOK(s string) result {
	return result{
		OK:  true,
		Msg: s,
	}
}

func ResponseOKf(s string, args ...any) result {
	return result{
		OK:  true,
		Msg: fmt.Sprintf(s, args...),
	}
}

func ResponseErrf(format string, args ...any) result {
	return result{
		OK:  true,
		Msg: fmt.Sprintf(format, args...),
	}
}
