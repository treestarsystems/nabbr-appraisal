// Sparkline 1
var options1 = {
	series: [
		{
			name: "Likes",
			data: [10, 20, 30, 50, 40, 30, 60],
		},
	],
	chart: {
		type: "line",
		height: 40,
		sparkline: {
			enabled: true,
		},
	},
	colors: ["#6fb4ce"],
	stroke: {
		curve: "smooth",
		width: 5,
	},
	tooltip: {
		y: {
			formatter: function (val) {
				return val;
			},
		},
	},
	fill: {
		type: "gradient",
		gradient: {
			shade: "dark",
			type: "horizontal",
			shadeIntensity: 0.5,
			gradientToColors: undefined,
			inverseColors: true,
		},
	},
};

var chart1 = new ApexCharts(
	document.querySelector("#sparklineLine1"),
	options1
);
chart1.render();

// Sparkline 2
var options2 = {
	series: [
		{
			name: "Views",
			data: [10, 20, 30, 50, 40, 30, 60],
		},
	],
	chart: {
		type: "line",
		height: 40,
		sparkline: {
			enabled: true,
		},
	},
	colors: ["#6fb4ce"],
	stroke: {
		curve: "smooth",
		width: 5,
	},
	tooltip: {
		y: {
			formatter: function (val) {
				return val;
			},
		},
	},
	fill: {
		type: "gradient",
		gradient: {
			shade: "dark",
			type: "horizontal",
			shadeIntensity: 0.5,
			gradientToColors: undefined,
			inverseColors: true,
		},
	},
};

var chart2 = new ApexCharts(
	document.querySelector("#sparklineLine2"),
	options2
);
chart2.render();

// Sparkline 3
var options3 = {
	series: [
		{
			name: "Users",
			data: [10, 20, 30, 50, 40, 30, 60],
		},
	],
	chart: {
		type: "line",
		height: 40,
		sparkline: {
			enabled: true,
		},
	},
	colors: ["#6fb4ce"],
	stroke: {
		curve: "smooth",
		width: 5,
	},
	tooltip: {
		y: {
			formatter: function (val) {
				return val;
			},
		},
	},
	fill: {
		type: "gradient",
		gradient: {
			shade: "dark",
			type: "horizontal",
			shadeIntensity: 0.5,
			gradientToColors: undefined,
			inverseColors: true,
		},
	},
};

var chart3 = new ApexCharts(
	document.querySelector("#sparklineLine3"),
	options3
);
chart3.render();

// Sparkline 4
var options4 = {
	series: [
		{
			name: "Sales",
			data: [10, 20, 30, 50, 40, 30, 60],
		},
	],
	chart: {
		type: "line",
		height: 40,
		sparkline: {
			enabled: true,
		},
	},
	colors: ["#6fb4ce"],
	stroke: {
		curve: "smooth",
		width: 5,
	},
	tooltip: {
		y: {
			formatter: function (val) {
				return val;
			},
		},
	},
	fill: {
		type: "gradient",
		gradient: {
			shade: "dark",
			type: "horizontal",
			shadeIntensity: 0.5,
			gradientToColors: undefined,
			inverseColors: true,
		},
	},
};

var chart4 = new ApexCharts(
	document.querySelector("#sparklineLine4"),
	options4
);
chart4.render();
