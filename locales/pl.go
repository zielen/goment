package locales

import (
	"strconv"
	"strings"
)

// EnLocale is the US English language locale.
var PlLocale = newLocale(
	"pl",
	strings.Split("Niedziela_Poniedziałek_Wtorek_Środa_Czwartek_Piątek_Sobota", "_"),
	strings.Split("Nie_Pon_Wt_Śr_Czw_Pt_Sob", "_"),
	strings.Split("Nd_Pn_Wt_Śr_Cz_Pt_So", "_"),
	strings.Split("Styczeń_Luty_Marzec_Kwiecień_Maj_Czerwiec_Lipiec_Sierpień_Wrzesień_Październik_Listopad_Grudzień", "_"),
	strings.Split("Sty_Lut_Mar_Apr_May_Jun_Jul_Aug_Sep_Oct_Nov_Dec", "_"),
	func(num int, period string) string {
		suffix := "th"
		switch num % 10 {
		case 1:
			if num%100 != 11 {
				suffix = "st"
			}
		case 2:
			if num%100 != 12 {
				suffix = "nd"
			}
		case 3:
			if num%100 != 13 {
				suffix = "rd"
			}
		}
		return strconv.Itoa(num) + suffix
	},
	nil,
	week{Dow: 0, Doy: 6},
	longDateFormats{
		"LTS":  "h:mm:ss A",
		"LT":   "h:mm A",
		"L":    "MM/DD/YYYY",
		"LL":   "MMMM D, YYYY",
		"LLL":  "MMMM D, YYYY h:mm A",
		"LLLL": "dddd, MMMM D, YYYY h:mm A",
	},
	relativeTimeFormats{
		"future": "za %s",
		"past":   "%s temu",
		"s":      "kilka sekund",
		"ss":     "%d sekund",
		"m":      "minuta",
		"mm":     "%d minut",
		"h":      "godzina",
		"hh":     "%d godzin",
		"d":      "dzień",
		"dd":     "%d dni",
		"M":      "miesiąc",
		"MM":     "%d miesięcy",
		"y":      "rok",
		"yy":     "%d lat",
	},
	calendarFunctions{
		"sameDay": func(hours int, day int) string {
			return "[Dziś o] LT"
		},
		"nextDay": func(hours int, day int) string {
			return "[Jutro o] LT"
		},
		"nextWeek": func(hours int, day int) string {
			return "dddd [o] LT"
		},
		"lastDay": func(hours int, day int) string {
			return "[Wczoraj o] LT"
		},
		"lastWeek": func(hours int, day int) string {
			return "[Last] dddd [at] LT"
		},
		"sameElse": func(hours int, day int) string {
			return "L"
		},
	},
	`(?i)(Styczeń|Luty|Marzec|Kwiecień|Maj|Czerwiec|Lipiec|Sierpień|Wrzesień|Październik|Listopad|Grudzień)`,
	`(?i)(Sty|Lut|Mar|Kwi|Maj|Cze|Lip|Sie|Wrz|Paź|Lis|Gru)`,
	`(?i)(Niedziela|Poniedziałej|Wtorek|Środa|Czwartek|Piątek|Sobota)`,
	`(?i)(Nie|Pon|Wto|Śr|Czw|Pt|Sob)`,
	`(?i)(Nd|Pn|Wt|Śr|Cz|Pt|Sb)`,
	`\d{1,2}(th|st|nd|rd)`,
)
