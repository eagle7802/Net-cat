package server

func (s *Server) CheckUserInChat(name string) bool {
	var boolean bool
	for user, q := range s.Connections {
		if user == name || q != q {
			boolean = true
		}
	}
	return boolean
}

func CheckText(text string) string {
	var temp string

	for i := 0; i < len(text); i++ {
		if text[i] == 27 {
			temp += "^["
		} else {
			temp += string(text[i])
		}
	}
	return temp
}

func CheckName(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == 27 {
			return true
		}
	}
	return false
}

func (s *Server) CheckUser(name string) bool {
	for user, q := range s.Connections {
		if user == name || q != q {
			return false
		}
	}
	return true
}

func isPrintable(message string) bool {
	printableFlag := false
	for _, char := range message {
		if char != ' ' && char != '\t' && char != '\n' && char != '\r' {
			printableFlag = true
			if char < 32 || char > 126 {
				return false
			}
		}
	}
	return printableFlag
}
