
// Package slot is generated by gogll. Do not edit. 
package slot

import(
	"bytes"
	"fmt"
	
	"github.com/goccmack/gogll/examples/json/parser/symbols"
)

type Label int

const(
	Array0R0 Label = iota
	Array0R1
	Array0R2
	Array1R0
	Array1R1
	Array1R2
	Array1R3
	GoGLL0R0
	GoGLL0R1
	Member0R0
	Member0R1
	Member0R2
	Member0R3
	Members0R0
	Members0R1
	Members1R0
	Members1R1
	Members1R2
	Members1R3
	Object0R0
	Object0R1
	Object0R2
	Object1R0
	Object1R1
	Object1R2
	Object1R3
	Value0R0
	Value0R1
	Value1R0
	Value1R1
	Value2R0
	Value2R1
	Value3R0
	Value3R1
	Value4R0
	Value4R1
	Value5R0
	Value5R1
	Value6R0
	Value6R1
	Values0R0
	Values0R1
	Values1R0
	Values1R1
	Values1R2
	Values1R3
)

type Slot struct {
	NT      symbols.NT
	Alt     int
	Pos     int
	Symbols symbols.Symbols
	Label 	Label
}

type Index struct {
	NT      symbols.NT
	Alt     int
	Pos     int
}

func GetAlternates(nt symbols.NT) []Label {
	alts, exist := alternates[nt]
	if !exist {
		panic(fmt.Sprintf("Invalid NT %s", nt))
	}
	return alts
}

func GetLabel(nt symbols.NT, alt, pos int) Label {
	l, exist := slotIndex[Index{nt,alt,pos}]
	if exist {
		return l
	}
	panic(fmt.Sprintf("Error: no slot label for NT=%s, alt=%d, pos=%d", nt, alt, pos))
}

func (l Label) EoR() bool {
	return l.Slot().EoR()
}

func (l Label) Head() symbols.NT {
	return l.Slot().NT
}

func (l Label) Index() Index {
	s := l.Slot()
	return Index{s.NT, s.Alt, s.Pos}
}

func (l Label) Alternate() int {
	return l.Slot().Alt
}

func (l Label) Pos() int {
	return l.Slot().Pos
}

func (l Label) Slot() *Slot {
	s, exist := slots[l]
	if !exist {
		panic(fmt.Sprintf("Invalid slot label %d", l))
	}
	return s
}

func (l Label) String() string {
	return l.Slot().String()
}

func (l Label) Symbols() symbols.Symbols {
	return l.Slot().Symbols
}

func (s *Slot) EoR() bool {
	return s.Pos >= len(s.Symbols)
}

func (s *Slot) String() string {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "%s : ", s.NT)
	for i, sym := range s.Symbols {
		if i == s.Pos {
			fmt.Fprintf(buf, "∙")
		}
		fmt.Fprintf(buf, "%s ", sym)
	}
	if s.Pos >= len(s.Symbols) {
		fmt.Fprintf(buf, "∙")
	}
	return buf.String()
}

var slots = map[Label]*Slot{ 
	Array0R0: {
		symbols.NT_Array, 0, 0, 
		symbols.Symbols{  
			symbols.T_2, 
			symbols.T_3,
		}, 
		Array0R0, 
	},
	Array0R1: {
		symbols.NT_Array, 0, 1, 
		symbols.Symbols{  
			symbols.T_2, 
			symbols.T_3,
		}, 
		Array0R1, 
	},
	Array0R2: {
		symbols.NT_Array, 0, 2, 
		symbols.Symbols{  
			symbols.T_2, 
			symbols.T_3,
		}, 
		Array0R2, 
	},
	Array1R0: {
		symbols.NT_Array, 1, 0, 
		symbols.Symbols{  
			symbols.T_2, 
			symbols.NT_Values, 
			symbols.T_3,
		}, 
		Array1R0, 
	},
	Array1R1: {
		symbols.NT_Array, 1, 1, 
		symbols.Symbols{  
			symbols.T_2, 
			symbols.NT_Values, 
			symbols.T_3,
		}, 
		Array1R1, 
	},
	Array1R2: {
		symbols.NT_Array, 1, 2, 
		symbols.Symbols{  
			symbols.T_2, 
			symbols.NT_Values, 
			symbols.T_3,
		}, 
		Array1R2, 
	},
	Array1R3: {
		symbols.NT_Array, 1, 3, 
		symbols.Symbols{  
			symbols.T_2, 
			symbols.NT_Values, 
			symbols.T_3,
		}, 
		Array1R3, 
	},
	GoGLL0R0: {
		symbols.NT_GoGLL, 0, 0, 
		symbols.Symbols{  
			symbols.NT_Value,
		}, 
		GoGLL0R0, 
	},
	GoGLL0R1: {
		symbols.NT_GoGLL, 0, 1, 
		symbols.Symbols{  
			symbols.NT_Value,
		}, 
		GoGLL0R1, 
	},
	Member0R0: {
		symbols.NT_Member, 0, 0, 
		symbols.Symbols{  
			symbols.T_7, 
			symbols.T_1, 
			symbols.NT_Value,
		}, 
		Member0R0, 
	},
	Member0R1: {
		symbols.NT_Member, 0, 1, 
		symbols.Symbols{  
			symbols.T_7, 
			symbols.T_1, 
			symbols.NT_Value,
		}, 
		Member0R1, 
	},
	Member0R2: {
		symbols.NT_Member, 0, 2, 
		symbols.Symbols{  
			symbols.T_7, 
			symbols.T_1, 
			symbols.NT_Value,
		}, 
		Member0R2, 
	},
	Member0R3: {
		symbols.NT_Member, 0, 3, 
		symbols.Symbols{  
			symbols.T_7, 
			symbols.T_1, 
			symbols.NT_Value,
		}, 
		Member0R3, 
	},
	Members0R0: {
		symbols.NT_Members, 0, 0, 
		symbols.Symbols{  
			symbols.NT_Member,
		}, 
		Members0R0, 
	},
	Members0R1: {
		symbols.NT_Members, 0, 1, 
		symbols.Symbols{  
			symbols.NT_Member,
		}, 
		Members0R1, 
	},
	Members1R0: {
		symbols.NT_Members, 1, 0, 
		symbols.Symbols{  
			symbols.NT_Member, 
			symbols.T_0, 
			symbols.NT_Members,
		}, 
		Members1R0, 
	},
	Members1R1: {
		symbols.NT_Members, 1, 1, 
		symbols.Symbols{  
			symbols.NT_Member, 
			symbols.T_0, 
			symbols.NT_Members,
		}, 
		Members1R1, 
	},
	Members1R2: {
		symbols.NT_Members, 1, 2, 
		symbols.Symbols{  
			symbols.NT_Member, 
			symbols.T_0, 
			symbols.NT_Members,
		}, 
		Members1R2, 
	},
	Members1R3: {
		symbols.NT_Members, 1, 3, 
		symbols.Symbols{  
			symbols.NT_Member, 
			symbols.T_0, 
			symbols.NT_Members,
		}, 
		Members1R3, 
	},
	Object0R0: {
		symbols.NT_Object, 0, 0, 
		symbols.Symbols{  
			symbols.T_9, 
			symbols.T_10,
		}, 
		Object0R0, 
	},
	Object0R1: {
		symbols.NT_Object, 0, 1, 
		symbols.Symbols{  
			symbols.T_9, 
			symbols.T_10,
		}, 
		Object0R1, 
	},
	Object0R2: {
		symbols.NT_Object, 0, 2, 
		symbols.Symbols{  
			symbols.T_9, 
			symbols.T_10,
		}, 
		Object0R2, 
	},
	Object1R0: {
		symbols.NT_Object, 1, 0, 
		symbols.Symbols{  
			symbols.T_9, 
			symbols.NT_Members, 
			symbols.T_10,
		}, 
		Object1R0, 
	},
	Object1R1: {
		symbols.NT_Object, 1, 1, 
		symbols.Symbols{  
			symbols.T_9, 
			symbols.NT_Members, 
			symbols.T_10,
		}, 
		Object1R1, 
	},
	Object1R2: {
		symbols.NT_Object, 1, 2, 
		symbols.Symbols{  
			symbols.T_9, 
			symbols.NT_Members, 
			symbols.T_10,
		}, 
		Object1R2, 
	},
	Object1R3: {
		symbols.NT_Object, 1, 3, 
		symbols.Symbols{  
			symbols.T_9, 
			symbols.NT_Members, 
			symbols.T_10,
		}, 
		Object1R3, 
	},
	Value0R0: {
		symbols.NT_Value, 0, 0, 
		symbols.Symbols{  
			symbols.T_7,
		}, 
		Value0R0, 
	},
	Value0R1: {
		symbols.NT_Value, 0, 1, 
		symbols.Symbols{  
			symbols.T_7,
		}, 
		Value0R1, 
	},
	Value1R0: {
		symbols.NT_Value, 1, 0, 
		symbols.Symbols{  
			symbols.T_6,
		}, 
		Value1R0, 
	},
	Value1R1: {
		symbols.NT_Value, 1, 1, 
		symbols.Symbols{  
			symbols.T_6,
		}, 
		Value1R1, 
	},
	Value2R0: {
		symbols.NT_Value, 2, 0, 
		symbols.Symbols{  
			symbols.NT_Object,
		}, 
		Value2R0, 
	},
	Value2R1: {
		symbols.NT_Value, 2, 1, 
		symbols.Symbols{  
			symbols.NT_Object,
		}, 
		Value2R1, 
	},
	Value3R0: {
		symbols.NT_Value, 3, 0, 
		symbols.Symbols{  
			symbols.NT_Array,
		}, 
		Value3R0, 
	},
	Value3R1: {
		symbols.NT_Value, 3, 1, 
		symbols.Symbols{  
			symbols.NT_Array,
		}, 
		Value3R1, 
	},
	Value4R0: {
		symbols.NT_Value, 4, 0, 
		symbols.Symbols{  
			symbols.T_8,
		}, 
		Value4R0, 
	},
	Value4R1: {
		symbols.NT_Value, 4, 1, 
		symbols.Symbols{  
			symbols.T_8,
		}, 
		Value4R1, 
	},
	Value5R0: {
		symbols.NT_Value, 5, 0, 
		symbols.Symbols{  
			symbols.T_4,
		}, 
		Value5R0, 
	},
	Value5R1: {
		symbols.NT_Value, 5, 1, 
		symbols.Symbols{  
			symbols.T_4,
		}, 
		Value5R1, 
	},
	Value6R0: {
		symbols.NT_Value, 6, 0, 
		symbols.Symbols{  
			symbols.T_5,
		}, 
		Value6R0, 
	},
	Value6R1: {
		symbols.NT_Value, 6, 1, 
		symbols.Symbols{  
			symbols.T_5,
		}, 
		Value6R1, 
	},
	Values0R0: {
		symbols.NT_Values, 0, 0, 
		symbols.Symbols{  
			symbols.NT_Value,
		}, 
		Values0R0, 
	},
	Values0R1: {
		symbols.NT_Values, 0, 1, 
		symbols.Symbols{  
			symbols.NT_Value,
		}, 
		Values0R1, 
	},
	Values1R0: {
		symbols.NT_Values, 1, 0, 
		symbols.Symbols{  
			symbols.NT_Value, 
			symbols.T_0, 
			symbols.NT_Values,
		}, 
		Values1R0, 
	},
	Values1R1: {
		symbols.NT_Values, 1, 1, 
		symbols.Symbols{  
			symbols.NT_Value, 
			symbols.T_0, 
			symbols.NT_Values,
		}, 
		Values1R1, 
	},
	Values1R2: {
		symbols.NT_Values, 1, 2, 
		symbols.Symbols{  
			symbols.NT_Value, 
			symbols.T_0, 
			symbols.NT_Values,
		}, 
		Values1R2, 
	},
	Values1R3: {
		symbols.NT_Values, 1, 3, 
		symbols.Symbols{  
			symbols.NT_Value, 
			symbols.T_0, 
			symbols.NT_Values,
		}, 
		Values1R3, 
	},
}

var slotIndex = map[Index]Label { 
	Index{ symbols.NT_Array,0,0 }: Array0R0,
	Index{ symbols.NT_Array,0,1 }: Array0R1,
	Index{ symbols.NT_Array,0,2 }: Array0R2,
	Index{ symbols.NT_Array,1,0 }: Array1R0,
	Index{ symbols.NT_Array,1,1 }: Array1R1,
	Index{ symbols.NT_Array,1,2 }: Array1R2,
	Index{ symbols.NT_Array,1,3 }: Array1R3,
	Index{ symbols.NT_GoGLL,0,0 }: GoGLL0R0,
	Index{ symbols.NT_GoGLL,0,1 }: GoGLL0R1,
	Index{ symbols.NT_Member,0,0 }: Member0R0,
	Index{ symbols.NT_Member,0,1 }: Member0R1,
	Index{ symbols.NT_Member,0,2 }: Member0R2,
	Index{ symbols.NT_Member,0,3 }: Member0R3,
	Index{ symbols.NT_Members,0,0 }: Members0R0,
	Index{ symbols.NT_Members,0,1 }: Members0R1,
	Index{ symbols.NT_Members,1,0 }: Members1R0,
	Index{ symbols.NT_Members,1,1 }: Members1R1,
	Index{ symbols.NT_Members,1,2 }: Members1R2,
	Index{ symbols.NT_Members,1,3 }: Members1R3,
	Index{ symbols.NT_Object,0,0 }: Object0R0,
	Index{ symbols.NT_Object,0,1 }: Object0R1,
	Index{ symbols.NT_Object,0,2 }: Object0R2,
	Index{ symbols.NT_Object,1,0 }: Object1R0,
	Index{ symbols.NT_Object,1,1 }: Object1R1,
	Index{ symbols.NT_Object,1,2 }: Object1R2,
	Index{ symbols.NT_Object,1,3 }: Object1R3,
	Index{ symbols.NT_Value,0,0 }: Value0R0,
	Index{ symbols.NT_Value,0,1 }: Value0R1,
	Index{ symbols.NT_Value,1,0 }: Value1R0,
	Index{ symbols.NT_Value,1,1 }: Value1R1,
	Index{ symbols.NT_Value,2,0 }: Value2R0,
	Index{ symbols.NT_Value,2,1 }: Value2R1,
	Index{ symbols.NT_Value,3,0 }: Value3R0,
	Index{ symbols.NT_Value,3,1 }: Value3R1,
	Index{ symbols.NT_Value,4,0 }: Value4R0,
	Index{ symbols.NT_Value,4,1 }: Value4R1,
	Index{ symbols.NT_Value,5,0 }: Value5R0,
	Index{ symbols.NT_Value,5,1 }: Value5R1,
	Index{ symbols.NT_Value,6,0 }: Value6R0,
	Index{ symbols.NT_Value,6,1 }: Value6R1,
	Index{ symbols.NT_Values,0,0 }: Values0R0,
	Index{ symbols.NT_Values,0,1 }: Values0R1,
	Index{ symbols.NT_Values,1,0 }: Values1R0,
	Index{ symbols.NT_Values,1,1 }: Values1R1,
	Index{ symbols.NT_Values,1,2 }: Values1R2,
	Index{ symbols.NT_Values,1,3 }: Values1R3,
}

var alternates = map[symbols.NT][]Label{ 
	symbols.NT_GoGLL:[]Label{ GoGLL0R0 },
	symbols.NT_Array:[]Label{ Array0R0,Array1R0 },
	symbols.NT_Values:[]Label{ Values0R0,Values1R0 },
	symbols.NT_Value:[]Label{ Value0R0,Value1R0,Value2R0,Value3R0,Value4R0,Value5R0,Value6R0 },
	symbols.NT_Object:[]Label{ Object0R0,Object1R0 },
	symbols.NT_Members:[]Label{ Members0R0,Members1R0 },
	symbols.NT_Member:[]Label{ Member0R0 },
}
