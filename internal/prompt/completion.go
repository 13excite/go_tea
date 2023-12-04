package prompt

import "github.com/charmbracelet/bubbles/list"

var (
	cityMap = map[string]struct{}{
		"berlin":    {},
		"munchen":   {},
		"frankfurt": {},
		"leipzig":   {},
		"longon":    {},
		"paris":     {},
		"liverpool": {},
		"koln":      {},
		"lion":      {},
		"flensburg": {},
		"bordo":     {},
		"erfurt":    {},
		"dresden":   {},
	}
)

func getCityNamesCompletion() []list.Item {
	// we know that there are 13 items in the map
	items := make([]list.Item, 13)

	for cityName := range cityMap {
		items = append(items, item(cityName))
	}
	return items

}
