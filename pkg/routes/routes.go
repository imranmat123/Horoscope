package routes

import (
	"github.com/labstack/echo/v4"
	"hrorscope/api"
	"hrorscope/internal/handlers"
	"hrorscope/internal/logic"
	"net/http"
	"strconv"
	"time"
)

func SetupRoutes(e *echo.Echo) {
	//home
	e.GET("/", handlers.HomeHandler)

	e.POST("/ask", func(c echo.Context) error {
		userName := c.FormValue("name")
		DoB := c.FormValue("DateOfBirth")
		currentTime := time.Now()
		month := logic.SetMonth(DoB[3:5])
		day, err := strconv.Atoi(DoB[:2])
		starSign := logic.Starsign(month, day)
		// Convert to string using Format
		currentTimeStr := currentTime.Format("2006-01-02 15:04:05")
		style := `
	- Career changes
	- Transition
	- Catastrophic events
	- Employer-related challenges
	- Business closure
	- Commute stress
	- Panic and fear of change
	- Adaptation
	- Moving on from setbacks
	- Success and achievement
	- Good fortune
	- Long-term positive outcomes
	- Change
	- Opportunity for growth
	- Future possibilities
`

		theme := `
Your response should:
	- limit of 200 words.
	- Provide **inspirational** and **motivational** guidance.
	- Referance everything to the given star-sign.
	- Offer **practical advice** on overcoming challenges.
	- Reflect on **career growth**, **progression**, and **adapting to change**.
	- Focus on **resilience** and **positive outcomes**.
	- Provide **forward-thinking** advice on opportunities and how to adapt to transitions.
	- Do not refer to birthdays, age, or third-person perspectives.
	- Mention **respected historical figures from the past with the same star-sign** **NO MOVIE OR TV STARS**only when relevant and in the context of **resilience** or **overcoming obstacles**. Do not focus on their birthdates.
	- Avoid using the phrase "Happy Birthday" or "celebrating life." Instead, focus on **personal growth**, **career advancement**, and **embracing opportunities**.
`
		finalPrompt := theme + "\n\n" +
			"Name: " + userName + "\n" +
			"Date of Birth: " + DoB + "\n" +
			"Star Sign: " + starSign + "\n" +
			"Todays Date: " + currentTimeStr + "\n" +
			"Use at least 3 of the following themes to shape the response: " + style
		reply, err := askGPT.AskGPT(finalPrompt)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Error talking to GPT: "+err.Error())
		}

		return c.Render(http.StatusOK, "home", map[string]interface{}{
			"Response": reply,
		})
	})
}
