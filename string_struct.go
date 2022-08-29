package imap

type StringStructMap map[string]struct{}

func (s StringStructMap)ToSlice() []string {
	slices := make([]string,0,len(s))
	for str,_ := range s {
		slices = append(slices,str)
	}
	return slices
}

func (s StringStructMap)Exist(str string) bool {
	_,exist := s[str]
	return exist
}

func (s StringStructMap)Set(str string){
	s[str] = struct{}{}
}

func (s StringStructMap)Delete(str string){
	delete(s,str)
}

type SortStringStructMap struct {
	keys            []string
	stringStructMap StringStructMap
}

func (s SortStringStructMap)ToSlice() []string {
	return s.keys
}

func (s SortStringStructMap)Set(str string)  {
	if !s.Exist(str) {
		s.keys = append(s.keys,str)
		s.stringStructMap[str] = struct{}{}
	}
}

func (s SortStringStructMap)Exist(str string) bool {
	_,exist := s.stringStructMap[str]
	return exist
}