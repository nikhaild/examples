package main

import (
	"io"
	"math/rand"
	"os"

	"github.com/go-echarts/go-echarts/charts"
	"github.com/go-echarts/go-echarts/components"
	"github.com/go-echarts/go-echarts/opts"
)

var (
	itemCntPie = 4
	seasons    = []string{"Spring", "Summer", "Autumn ", "Winter"}
)

func generatePieItems() []opts.PieData {
	items := make([]opts.PieData, 0)
	for i := 0; i < itemCntPie; i++ {
		items = append(items, opts.PieData{Name: seasons[i], Value: rand.Intn(100)})
	}
	return items
}

func pieBase() *charts.Pie {
	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Pie-basic-example",
		}),
	)

	pie.AddSeries("pie", generatePieItems())
	return pie
}

func pieShowLabel() *charts.Pie {
	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Pie-with-label-opts",
		}),
	)

	pie.AddSeries("pie", generatePieItems()).
		SetSeriesOptions(charts.WithLabelOpts(
			opts.Label{
				Show:      true,
				Formatter: "{b}: {c}",
			}),
		)
	return pie
}

func pieRadius() *charts.Pie {
	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Pie-Radius",
		}),
	)

	pie.AddSeries("pie", generatePieItems()).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show:      true,
				Formatter: "{b}: {c}",
			}),
			charts.WithPieChartOpts(opts.PieChart{
				Radius: []string{"40%", "75%"},
			}),
		)
	return pie
}

func pieRoseArea() *charts.Pie {
	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Pie-Rose(Area)",
		}),
	)

	pie.AddSeries("pie", generatePieItems()).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show:      true,
				Formatter: "{b}: {c}",
			}),
			charts.WithPieChartOpts(opts.PieChart{
				Radius:   []string{"40%", "75%"},
				RoseType: "area",
			}),
		)
	return pie
}

func pieRoseRadius() *charts.Pie {
	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Pie-Rose(Radius)",
		}),
	)

	pie.AddSeries("pie", generatePieItems()).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show:      true,
				Formatter: "{b}: {c}",
			}),
			charts.WithPieChartOpts(opts.PieChart{
				Radius:   []string{"30%", "75%"},
				RoseType: "radius",
			}),
		)
	return pie
}

func pieRoseAreaRadius() *charts.Pie {
	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Pie-Rose(Area/Radius)",
		}),
	)

	pie.AddSeries("area", generatePieItems()).
		SetSeriesOptions(
			charts.WithPieChartOpts(opts.PieChart{
				Radius:   []string{"30%", "75%"},
				RoseType: "area",
				Center:   []string{"25%", "50%"},
			}),
		)

	pie.AddSeries("pie", generatePieItems()).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show:      true,
				Formatter: "{b}: {c}",
			}),
			charts.WithPieChartOpts(opts.PieChart{
				Radius:   []string{"30%", "75%"},
				RoseType: "radius",
				Center:   []string{"75%", "50%"},
			}),
		)
	return pie
}

func pieInPie() *charts.Pie {
	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Pie-in-pie",
		}),
	)

	pie.AddSeries("area", generatePieItems(),
		charts.WithLabelOpts(opts.Label{
			Show:      true,
			Formatter: "{b}: {c}",
		}),
		charts.WithPieChartOpts(opts.PieChart{
			Radius:   []string{"50%", "55%"},
			RoseType: "area",
		}),
	)

	pie.AddSeries("radius", generatePieItems(),
		charts.WithPieChartOpts(opts.PieChart{
			Radius:   []string{"0%", "45%"},
			RoseType: "radius",
		}),
	)
	return pie
}

func main() {
	page := components.NewPage()
	page.AddCharts(
		pieBase(),
		pieShowLabel(),
		pieRadius(),
		pieRoseArea(),
		pieRoseRadius(),
		pieRoseAreaRadius(),
		pieInPie(),
	)
	f, err := os.Create("html/pie.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(os.Stdout, f))
}
