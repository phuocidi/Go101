package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type record struct {
	Date string
	Open float64
}

func makeRecord(row []string) record {
	open, _ := strconv.ParseFloat(row[1], 64)
	return record{
		Date: row[0],
		Open: open,
	}
}

func main() {
	f, err := os.Open("table.csv")
	defer f.Close()

	if err != nil {
		panic(err)
	}

	rdr := csv.NewReader(f)
	rows, err := rdr.ReadAll()
	if err != nil {
		panic(err)
	}

	fmt.Println(`<!DOCTYPE html>
		<head>
			<script
		  	src="http://code.jquery.com/jquery-1.12.4.min.js"
		  	integrity="sha256-ZosEbRLbNQzLpnKIkEdrPv7lOy9C27hHQ+Xp8a4MxAQ="
		  	crossorigin="anonymous">
		  	</script>

		  	<script src="https://code.highcharts.com/highcharts.js"></script>
			<script src="https://code.highcharts.com/modules/exporting.js"></script>
		</head>
		<body>
			<div id="container"></div>
			<table>
				<thead>
					<tr>
						<th>Date</th>
						<th>Open</th>
					</tr>
				</thead>
				
				<tbody>
		`)

	openValues := []string{}

	for i, row := range rows {
		if i == 0 {
			continue
		}
		record := makeRecord(row)
		fmt.Println(`
			<tr>
				<td>` + record.Date + `</td>
				<td>` + fmt.Sprintf("%.2f", record.Open) + `</td>
			</tr> 
			`)
		openValues = append(openValues, fmt.Sprintf("%.2f", record.Open))
	}

	fmt.Println(`
				</tbody>
			</table>

			<script>
Highcharts.chart('container', {

    title: {
        text: 'Solar Employment Growth by Sector, 2010-2016'
    },

    subtitle: {
        text: 'Source: thesolarfoundation.com'
    },
    xAxis: {
		categories:['Jan','Feb','March','April','May','Jun','Jul','Aug','Sep','Oct','Nov','Dec']
    	},
    yAxis: {
        title: {
            text: 'Indexes'
        }
    },
    legend: {
        layout: 'vertical',
        align: 'right',
        verticalAlign: 'middle'
    },

    plotOptions: {
        series: {
            pointStart: 2012
        }
    },

    series: [{
        name: 'NASDAS',
        data: [

` + strings.Join(openValues, ",") + `
        ],
    }]

});
    

			</script>
		</body>`)

}
