// Morris Bar Chart
Morris.Bar({
	element: "morrisBar",
	data: [
		{ x: "2016 Q1", y: 2, z: 4, a: 2 },
		{ x: "2016 Q2", y: 5, z: 3, a: 1 },
		{ x: "2016 Q3", y: 2, z: 7, a: 4 },
		{ x: "2016 Q4", y: 5, z: 6, a: 3 },
	],
	xkey: "x",
	ykeys: ["y", "z", "a"],
	labels: ["Y", "Z", "A"],
	resize: true,
	hideHover: "auto",
	gridLineColor: "#575e6d",
	barColors: ["#30758e", "#3c92b1", "#63a8c1", "#8abed0", "#b1d3e0", "#c5dee8"],
});
