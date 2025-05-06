package logic

func SetMonth(a string) string {
	if a == "01" {
		return "January"
	} else if a == "02" {
		return "February"
	} else if a == "03" {
		return "March"
	} else if a == "04" {
		return "April"
	} else if a == "05" {
		return "May"
	} else if a == "06" {
		return "June"
	} else if a == "07" {
		return "July"
	} else if a == "08" {
		return "August"
	} else if a == "09" {
		return "September"
	} else if a == "10" {
		return "October"
	} else if a == "11" {
		return "November"
	} else if a == "12" {
		return "December"
	}
	return ""
}

func Starsign(month string, day int) string {
	//Aries: March 21 – April 19
	if month == "March" && day > 20 && day < 32 || month == "April" && day > 0 && day < 20 {
		return "Aries"
	}
	//Taurus: April 20 – May 20
	if month == "April" && day > 19 && day < 31 || month == "May" && day > 0 && day < 21 {
		return "Taurus"
	}
	//Gemini: May 21 – June 21
	if month == "May" && day > 20 && day < 32 || month == "June" && day > 0 && day < 22 {
		return "Gemini"
	}
	//Cancer: June 22 – July 22
	if month == "June" && day > 21 && day < 31 || month == "July" && day > 0 && day < 23 {
		return "Cancer"
	}
	//Leo: July 23 – August 22
	if month == "July" && day > 22 && day < 32 || month == "August" && day > 0 && day < 23 {
		return "Leo"
	}
	//Virgo: August 23 – September 22
	if month == "August" && day > 22 && day < 32 || month == "September" && day > 0 && day < 23 {
		return "Leo"
	}
	//Libra: September 23 – October 23
	if month == "September" && day > 22 && day < 31 || month == "October" && day > 0 && day < 24 {
		return "Libra"
	}
	//Scorpio: October 24 – November 21
	if month == "October" && day > 23 && day < 32 || month == "November" && day > 0 && day < 22 {
		return "Scorpio"
	}
	//Sagittarius: November 22 – December 21
	if month == "November" && day > 21 && day < 31 || month == "December" && day > 0 && day < 22 {
		return "Sagittarius"
	}
	//Capricorn: December 22 – January 19
	if month == "December" && day > 21 && day < 32 || month == "January" && day > 0 && day < 20 {
		return "Capricorn"
	}
	//Aquarius: January 20 – February 18
	if month == "January" && day > 19 && day < 32 || month == "February" && day > 0 && day < 19 {
		return "Aquarius"
	}
	//Pisces: February 19 – March 20
	if month == "February" && day > 18 && day < 29 || month == "March" && day > 0 && day < 21 {
		return "Pisces"
	}
	return ""
}
