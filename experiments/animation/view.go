package main

var loader= []string{"|o---------|",
					 "|-o--------|" ,
					 "|--o-------|" ,
					 "|---o------|" ,
					 "|----o-----|" ,
					 "|-----o----|" ,
					 "|------o---|" ,
					 "|-------o--|" ,
					 "|--------o-|" ,
					 "|---------o|" ,
				 	}

func (m model) View() string {
	return loader[m.counter]
	return "?"
}
