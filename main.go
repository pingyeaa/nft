package main

import (
	"fmt"
	"log"
	"nft/discord"
	"nft/opensea"
	"time"
)

func main() {

	for {
		webhook := GetConfig("discord", "webhook")
		name := GetConfig("opensea", "name")
		slug := GetConfig("opensea", "slug")
		template := fmt.Sprintf("templates/%s.html", slug)
		templateOut := fmt.Sprintf("o-%s.html", slug)
		imageOut := fmt.Sprintf("%s.png", slug)
		var vars = map[string]string{}

		stats, err := opensea.GetStats(slug)
		if err != nil {
			log.Fatalln(err.Error())
		}

		vars["name"] = name
		vars["time"] = fmt.Sprintf(`%s`, time.Now().Format("2006-01-02 15:04:05"))

		vars["total_volume"] = fmt.Sprintf("%.2fE", stats.Stats.TotalVolume)
		rate := InsertAndCalcRate(slug, "total_volume", stats.Stats.TotalVolume)
		vars["total_volume_rate"] = SetFloatColor(rate, "较上次", "")

		vars["floor_price"] = fmt.Sprintf("%.2fE", stats.Stats.FloorPrice)
		rate = InsertAndCalcRate(slug, "floor_price", stats.Stats.FloorPrice)
		vars["floor_price_rate"] = SetFloatColor(rate, "", "")

		vars["holders"] = fmt.Sprintf("%.f", stats.Stats.NumOwners)
		rate = InsertAndCalcRate(slug, "holders", stats.Stats.NumOwners)
		vars["holders_rate"] = SetFloatColor(rate, "", "")

		vars["avg_price"] = fmt.Sprintf("%.2fE", stats.Stats.AveragePrice)
		rate = InsertAndCalcRate(slug, "avg_price", stats.Stats.AveragePrice)
		vars["avg_price_rate"] = SetFloatColor(rate, "", "")

		vars["total_sales"] = fmt.Sprintf("%.f", stats.Stats.TotalSales)
		rate = InsertAndCalcRate(slug, "total_sales", stats.Stats.TotalSales)
		vars["total_sales_rate"] = SetFloatColor(rate, "", "")

		vars["count"] = fmt.Sprintf("%.f", stats.Stats.Count)

		vars["market_cap"] = fmt.Sprintf("%.2fE", stats.Stats.MarketCap)
		rate = InsertAndCalcRate(slug, "market_cap", stats.Stats.MarketCap)
		vars["market_cap_rate"] = SetFloatColor(rate, "", "")

		vars["one_day_volume"] = fmt.Sprintf("%.2fE", stats.Stats.OneDayVolume)
		rate = fmt.Sprintf("%.2fE", stats.Stats.OneDayChange)
		vars["one_day_volume_rate"] = SetFloatColor(rate, "", "")

		vars["one_day_sales"] = fmt.Sprintf("%.f", stats.Stats.OneDaySales)
		rate = InsertAndCalcRate(slug, "one_day_sales", stats.Stats.OneDaySales)
		vars["one_day_sales_rate"] = SetFloatColor(rate, "", "")

		vars["one_day_avg"] = fmt.Sprintf("%.2fE", stats.Stats.OneDayAveragePrice)
		rate = InsertAndCalcRate(slug, "one_day_avg", stats.Stats.OneDayAveragePrice)
		vars["one_day_avg_rate"] = SetFloatColor(rate, "", "")

		vars["thirty_day_volume"] = fmt.Sprintf("%.2fE", stats.Stats.ThirtyDayVolume)
		rate = fmt.Sprintf("%.2fE", stats.Stats.ThirtyDayChange)
		vars["thirty_day_volume_rate"] = SetFloatColor(rate, "", "")

		vars["thirty_day_sales"] = fmt.Sprintf("%.f", stats.Stats.ThirtyDaySales)
		rate = InsertAndCalcRate(slug, "thirty_day_sales", stats.Stats.ThirtyDaySales)
		vars["thirty_day_sales_rate"] = SetFloatColor(rate, "", "")

		vars["thirty_day_avg"] = fmt.Sprintf("%.2fE", stats.Stats.ThirtyDayAveragePrice)
		rate = InsertAndCalcRate(slug, "thirty_day_avg", stats.Stats.ThirtyDayAveragePrice)
		vars["thirty_day_avg_rate"] = SetFloatColor(rate, "", "")

		vars["seven_day_volume"] = fmt.Sprintf("%.2fE", stats.Stats.SevenDayVolume)
		rate = fmt.Sprintf("%.2fE", stats.Stats.SevenDayChange)
		vars["seven_day_volume_rate"] = SetFloatColor(rate, "", "")

		vars["seven_day_sales"] = fmt.Sprintf("%.f", stats.Stats.SevenDaySales)
		rate = InsertAndCalcRate(slug, "seven_day_sales", stats.Stats.SevenDaySales)
		vars["seven_day_sales_rate"] = SetFloatColor(rate, "", "")

		vars["seven_day_avg"] = fmt.Sprintf("%.2fE", stats.Stats.SevenDayAveragePrice)
		rate = InsertAndCalcRate(slug, "seven_day_avg", stats.Stats.SevenDayAveragePrice)
		vars["seven_day_avg_rate"] = SetFloatColor(rate, "", "")

		ReplaceVar(template, vars, templateOut)
		Html2Image(templateOut, imageOut)
		err = discord.PushFile(webhook, imageOut)
		if err != nil {
			log.Println(err.Error())
		}

		time.Sleep(time.Hour * 12)
	}
}
