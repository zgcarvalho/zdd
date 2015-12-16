package atom

var LigAtoms = [...]string{"Br", "C.1", "C.2", "C.3", "C.ar", "C.cat", "Cl",
	"F", "I", "N.1", "N.2", "N.4", "N.am", "N.ar", "N.pl3", "O.2", "O.3", "O.co2",
	"P.3", "S.3", "S.o2"}

var ProtAtoms = [...]string{"C.2", "C.3", "Ca", "C.ar", "C.cat", "Cd", "Hg",
	"Mg", "Mn", "N.2", "N.3", "N.4", "Na", "N.am", "N.ar", "Ni", "N.pl3", "O.2",
	"O.3", "O.co2", "P.3", "S", "S.3", "Zn"}

var AtomNames = [...]string{"Al", "Br", "C.1", "C.2", "C.3", "Ca", "C.ar", "C.cat",
	"Cl", "Co.oh", "Cr.oh", "Cr.th", "Cu", "F", "Fe", "I", "K", "Li", "Mg", "Mn", "Mo", "N.1",
	"N.2", "N.3", "N.4", "Na", "N.am", "N.ar", "N.pl3", "O.2", "O.3", "O.co2", "P.3", "S.2",
	"S.3", "Se", "Si", "Sn", "S.o", "S.o2", "Zn"}

var NonPolar = map[string]struct{}{"C.1": {}, "C.2": {}, "C.3": {}, "C.ar": {},
	"P.3": {}, "S.2": {}, "S.3": {}, "Se": {}, "Si": {}, "S.o": {}, "S.o2": {}}

var Polar = map[string]struct{}{"Al": {}, "Br": {}, "Ca": {}, "C.cat": {},
	"Cl": {}, "Co.oh": {}, "Cr.oh": {}, "Cr.th": {}, "Cu": {}, "F": {}, "Fe": {},
	"I": {}, "K": {}, "Li": {}, "Mg": {}, "Mn": {}, "Mo": {}, "N.1": {},
	"N.2": {}, "N.3": {}, "N.4": {}, "Na": {}, "N.am": {}, "N.ar": {}, "N.pl3": {},
	"O.2": {}, "O.3": {}, "O.co2": {}, "Sn": {}, "Zn": {}}

type Atom struct {
	Name  string
	Coord [3]float64
}
