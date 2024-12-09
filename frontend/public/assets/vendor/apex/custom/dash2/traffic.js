var options = {
	chart: {
		height: 380,
		type: "line",
		toolbar: {
			show: false,
		},
	},
	dataLabels: {
		enabled: false,
	},
	stroke: {
		curve: "smooth",
		width: 5,
	},
	series: [
		{
			name: "Visitors",
			data: [10, 40, 15, 40, 35, 96, 69],
		},
	],
	grid: {
		borderColor: "#575e6d",
		strokeDashArray: 5,
		xaxis: {
			lines: {
				show: true,
			},
		},
		yaxis: {
			lines: {
				show: false,
			},
		},
		padding: {
			top: 0,
			right: 20,
			bottom: 0,
			left: 20,
		},
	},
	xaxis: {
		categories: ["Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"],
	},
	yaxis: {
		labels: {
			show: false,
		},
	},
	colors: ["#3c92b1"],
	fill: {
		type: "gradient",
		gradient: {
			shadeIntensity: 1,
			opacityFrom: 0.7,
			opacityTo: 0.9,
			colorStops: [
				{
					offset: 0,
					color: "#3c92b1",
					opacity: 1,
				},
				{
					offset: 20,
					color: "#6fb4ce",
					opacity: 1,
				},
				{
					offset: 60,
					color: "#a9bd7a",
					opacity: 1,
				},
				{
					offset: 100,
					color: "#d2a968",
					opacity: 1,
				},
			],
		},
	},
	markers: {
		size: 0,
		opacity: 0.3,
		colors: ["#3c92b1"],
		strokeColor: "#ffffff",
		strokeWidth: 2,
		hover: {
			size: 7,
		},
	},
	tooltip: {
		y: {
			formatter: function (val) {
				return +val + "k";
			},
		},
	},
};

var chart = new ApexCharts(document.querySelector("#traffic"), options);

chart.render();
