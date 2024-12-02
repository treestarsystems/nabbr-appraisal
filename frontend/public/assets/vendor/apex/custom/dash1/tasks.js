var options = {
	plotOptions: {
		pie: {
			customScale: 0.8,
			donut: {
				size: "20%",
			},
		},
	},
	chart: {
		width: 204,
		type: "donut",
	},
	labels: ["Pending", "Completed", "New"],
	series: [40, 30, 20],
	legend: { show: false },
	dataLabels: {
		enabled: false,
	},
	stroke: {
		width: 0,
	},

	colors: ["#3c92b1", "#a9bd7a", "#bf7a6a"],
};
var chart = new ApexCharts(document.querySelector("#tasks"), options);
chart.render();
