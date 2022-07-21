package attendance_transactions

var CauseCodes = map[string]string{
	"ARB": "Timlön",
	"BE2": "Beredskapstid 2",
	"BER": "Beredskapstid (BE1 is used in PAXml)",
	"BE1": "Beredskapstid (BE1 is used in PAXml)",
	"FLX": "Flextid +/-",
	"HLG": "Helglön",
	"JO2": "Jourtid 2 (JR2 is used in PAXml)",
	"JR2": "Jourtid 2 (JR2 is used in PAXml)",
	"JOR": "Jourtid (JR1 is used in PAXml)",
	"JR1": "Jourtid (JR1 is used in PAXml)",
	"MER": "Mertid",
	"OB1": "OB-ersättning 1",
	"OB2": "OB-ersättning 2",
	"OB3": "OB-ersättning 3",
	"OB4": "OB-ersättning 4",
	"OB5": "OB-ersättning 5",
	"OK0": "Extratid \u2013 Komptid (NV9 is used in PAXml)",
	"NV9": "Extratid – Komptid (NV9 is used in PAXml)",
	"OK1": "Övertid 1 \u2013 Komptid (ÖK1 is used in PAXml)",
	"ÖK1": "Övertid 1 – Komptid (ÖK1 is used in PAXml)",
	"OK2": "Övertid 2 \u2013 Komptid (ÖK2 is used in PAXml)",
	"ÖK2": "Övertid 2 – Komptid (ÖK2 is used in PAXml)",
	"OK3": "Övertid 3 \u2013 Komptid (ÖK3 is used in PAXml)",
	"ÖK3": "Övertid 3 – Komptid (ÖK3 is used in PAXml)",
	"OK4": "Övertid 4 \u2013 Komptid (ÖK4 is used in PAXml)",
	"ÖK4": "Övertid 4 – Komptid (ÖK4 is used in PAXml)",
	"OK5": "Övertid 5 \u2013 Komptid (ÖK5 is used in PAXml)",
	"ÖK5": "Övertid 5 – Komptid (ÖK5 is used in PAXml)",
	"OT1": "Övertid 1 \u2013 Betalning (ÖT1 is used in PAXml)",
	"ÖT1": "Övertid 1 – Betalning (ÖT1 is used in PAXml)",
	"OT2": "Övertid 2 \u2013 Betalning (ÖT2 is used in PAXml)",
	"ÖT2": "Övertid 2 – Betalning (ÖT2 is used in PAXml)",
	"OT3": "Övertid 3 \u2013 Betalning (ÖT3 is used in PAXml)",
	"ÖT3": "Övertid 3 – Betalning (ÖT3 is used in PAXml)",
	"OT4": "Övertid 4 \u2013 Betalning (ÖT4 is used in PAXml)",
	"ÖT4": "Övertid 4 – Betalning (ÖT4 is used in PAXml)",
	"OT5": "Övertid 5 \u2013 Betalning (ÖT5 is used in PAXml)",
	"ÖT5": "Övertid 5 – Betalning (ÖT5 is used in PAXml)",
	"RES": "Restid (RE1 is used in PAXml)",
	"RE1": "Restid (RE1 is used in PAXml)",
	"TID": "Arbetstid",
}
