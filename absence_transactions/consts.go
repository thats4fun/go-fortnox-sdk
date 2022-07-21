package absence_transactions

const (
	ASK = "ASK"
	FPE = "FPE"
	FRA = "FRA"
	HAV = "HAV"
	KOV = "KOM"
	MIL = "MIL"
	NAR = "NAR"
	OS1 = "OS1"
	OS2 = "OS2"
	OS3 = "OS3"
	OS4 = "OS4"
	OS5 = "OS5"
	PAP = "PAP"
	PEM = "PEM"
	PER = "PER"
	SEM = "SEM"
	SJK = "SJK"
	SMB = "SMB"
	SVE = "SVE"
	TJL = "TJL"
	UTB = "UTB"
	VAB = "VAB"
)

var Codes = [...]string{
	"ASK",
	"FPE",
	"FRA",
	"HAV",
	"KOM",
	"MIL",
	"NAR",
	"OS1",
	"OS2",
	"OS3",
	"OS4",
	"OS5",
	"PAP",
	"PEM",
	"PER",
	"SEM",
	"SJK",
	"SMB",
	"SVE",
	"TJL",
	"UTB",
	"VAB",
}

var codeSet = map[string]struct{}{
	"ASK": {},
	"FPE": {},
	"FRA": {},
	"HAV": {},
	"KOM": {},
	"MIL": {},
	"NAR": {},
	"OS1": {},
	"OS2": {},
	"OS3": {},
	"OS4": {},
	"OS5": {},
	"PAP": {},
	"PEM": {},
	"PER": {},
	"SEM": {},
	"SJK": {},
	"SMB": {},
	"SVE": {},
	"TJL": {},
	"UTB": {},
	"VAB": {},
}

func IsCodeInCodeSet(code string) bool {
	_, ok := codeSet[code]
	return ok
}

// CauseCodes is usable "CauseCodes"
var CauseCodes = map[string]string{
	"ASK": "Arbetsskada",
	"ATF": "Arbetstidsförkortning",
	"FPE": "Föräldraledig",
	"FRA": "Frånvaro, övrigt (FR1 is used in PAXml)",
	"FR1": "Frånvaro, övrigt (FR1 is used in PAXml)",
	"HAV": "Graviditetspenning",
	"KOM": "Kompledig",
	"MIL": "Militärtjänst (max 60 dagar)",
	"NAR": "Närståendevård (NÄR is used in PAXml)",
	"NÄR": "Närståendevård (NÄR is used in PAXml)",
	"OS1": "Sjuk-OB 1",
	"OS2": "Sjuk-OB 2",
	"OS3": "Sjuk-OB 3",
	"OS4": "Sjuk-OB 4",
	"OS5": "Sjuk-OB 5",
	"PAP": "Pappadagar",
	"PEM": "Permission",
	"PER": "Permitterad",
	"SEM": "Semester",
	"SJK": "Sjukfrånvaro",
	"SMB": "Smittbärare",
	"SVE": "Svenska för invandrare",
	"TJL": "Tjänstledig",
	"UTB": "Facklig utbildning (FAC is used in PAXml)",
	"FAC": "Facklig utbildning (FAC is used in PAXml)",
	"VAB": "Vård av barn",
}
