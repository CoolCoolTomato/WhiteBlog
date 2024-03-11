package admin

import (
	"WhiteBlog/config"
	"WhiteBlog/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type MonthArticleCount struct {
	Month string
	Count int
}

// 图表数据
func getArticleChart() []MonthArticleCount {
	var months []MonthArticleCount
	formatArticles, _ := models.GetArticles()
	now := time.Now()
	currMonth := int(now.Month())
	currYear := now.Year()
	// 计算需要统计的最早的月份
	earliestMonth := currMonth - 12 // 最近12个月
	earliestYear := currYear
	if earliestMonth <= 0 {
		earliestMonth += 12
		earliestYear -= 1
	}
	// 初始化需要统计的月份数据，从最早需要统计的月份开始
	for year, month := currYear, currMonth; !(year == earliestYear && month == earliestMonth); {
		months = append([]MonthArticleCount{{Month: fmt.Sprintf("%d月", month), Count: 0}}, months...)
		month--
		if month <= 0 {
			month += 12
			year -= 1
		}
	}
	// 统计每个月的文章数量
	for _, article := range formatArticles {
		articleMonth := int(article.CreatedDate.Month())
		articleYear := article.CreatedDate.Year()
		// 在月份切片中找到对应的月份并增加计数
		for i := 0; i < len(months); i++ {
			monthStr := fmt.Sprintf("%d月", articleMonth)
			if months[i].Month == monthStr {
				if articleYear == currYear {
					months[i].Count++
					break
				}
				if articleYear == currYear-1 && articleMonth > earliestMonth {
					months[i].Count++
					break
				}
			}
		}
	}
	return months
}

func Index(c *gin.Context) {
	theConfig := config.GetConfig()
	rootClass, _ := models.GetRootClasses()
	articleCountMap := getArticleChart()
	poets, _ := models.GetPoets()
	classNum := models.GetClassNum()
	articleNum := models.GetArticleNum()
	c.HTML(http.StatusOK, "admin/index.html", gin.H{
		"rootClass":       rootClass,
		"articleCountMap": articleCountMap,
		"poets":           poets,
		"title":           config.GetConfig().User.Username,
		"articleNum":      articleNum,
		"classNum":        classNum,
		"config":          theConfig,
	})
}
