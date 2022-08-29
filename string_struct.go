package imap

type StringStruct map[string]struct{}

func (s StringStruct)ToSlice() []string {
	slices := make([]string,0,len(s))
	for str,_ := range s {
		slices = append(slices,str)
	}
	return slices
}

func (s StringStruct)Exist(str string) bool {
	_,exist := s[str]
	return exist
}

func (s StringStruct)Set(str string){
	s[str] = struct{}{}
}

func (s StringStruct)Delete(str string){
	delete(s,str)
}

type SortStringStruct struct {
	keys []string
	stringStruct StringStruct
}

func (s SortStringStruct)ToSlice() []string {
	return s.keys
}

func (s SortStringStruct)Set(str string)  {
	if !s.Exist(str) {
		s.keys = append(s.keys,str)
		s.stringStruct[str] = struct{}{}
	}
}

func (s SortStringStruct)Exist(str string) bool {
	_,exist := s.stringStruct[str]
	return exist
}