package main

import (
	"fmt"

	"github.com/ruraomsk/rui"
)

const tabsDemoText = `
GridLayout {
	style = demoPage,
	content = [
		TabsLayout { id = tabsLayout, width = 100%, height = 100%, tabs = top, tab-close-button = true,
			content = [
				View { width = 300px, height = 200px, background-color = #FFFF0000, title = "ПК1"},
				View { width = 400px, height = 250px, background-color = #FF00FF00, title = "ПК2"},
				View { width = 100px, height = 400px, background-color = #FF0000FF, title = "ПК3"},
				View { width = 300px, height = 200px, background-color = #FF000000, title = "ПК4"},
				View { width = 300px, height = 200px, background-color = #FFFF0000, title = "ПК5"},
				View { width = 400px, height = 250px, background-color = #FF00FF00, title = "ПК6"},
				View { width = 100px, height = 400px, background-color = #FF0000FF, title = "ПК7"},
				View { width = 300px, height = 200px, background-color = #FF000000, title = "ПК8"},
				View { width = 300px, height = 200px, background-color = #FFFF0000, title = "ПК9"},
				View { width = 400px, height = 250px, background-color = #FF00FF00, title = "ПК10"},
				View { width = 100px, height = 400px, background-color = #FF0000FF, title = "ПК11"},
				View { width = 300px, height = 200px, background-color = #FF000000, title = "ПК12"},
				View { width = 300px, height = 200px, background-color = #FFFF0000, title = "ПК13"},
				View { width = 400px, height = 250px, background-color = #FF00FF00, title = "ПК14"},
				View { width = 100px, height = 400px, background-color = #FF0000FF, title = "ПК15"},
				View { width = 300px, height = 200px, background-color = #FF000000, title = "ПК16"},
				View { width = 300px, height = 200px, background-color = #FFFF0000, title = "ПК17"},
				View { width = 400px, height = 250px, background-color = #FF00FF00, title = "ПК18"},
				View { width = 100px, height = 400px, background-color = #FF0000FF, title = "ПК19"},
				View { width = 300px, height = 200px, background-color = #FF000000, title = "ПК20"},
				View { width = 300px, height = 200px, background-color = #FFFF0000, title = "ПК21"},
			]
		},
		ListLayout {
			style = optionsPanel,
			content = [
				GridLayout {
					style = optionsTable,
					content = [
						TextView { row = 0, text = "Tabs location" },
						DropDownList { row = 0, column = 1, id = tabsTypeList, 
							items = ["top", "bottom", "left", "right", "left list", "right list", "hidden"]
						}
					]
				}
			]
		}
	]
}
`

func createTabsDemo(session rui.Session) rui.View {
	view := rui.CreateViewFromText(session, tabsDemoText)
	if view == nil {
		return nil
	}

	rui.Set(view, "tabsTypeList", rui.DropDownEvent, func(_ rui.DropDownList, number int) {
		rui.Set(view, "tabsLayout", rui.Tabs, number)
	})

	rui.Set(view, "tabsLayout", rui.TabCloseEvent, func(index int) {
		rui.ShowMessage("", fmt.Sprintf(`The close button of the tab "%d" was clicked`, index), session)
	})
	return view
}
