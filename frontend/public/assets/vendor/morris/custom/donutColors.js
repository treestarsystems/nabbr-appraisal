// Morris Donut
Morris.Donut({
	element: "donutColors",
	data: [
		{ value: 30, label: "foo" },
		{ value: 15, label: "bar" },
		{ value: 10, label: "baz" },
		{ value: 5, label: "A really really long label" },
	],
	backgroundColor: "#2a3039",
	labelColor: "#95a0b1",
	colors: ["#30758e", "#3c92b1", "#63a8c1", "#8abed0", "#b1d3e0", "#c5dee8"],
	resize: true,
	hideHover: "auto",
	gridLineColor: "#575e6d",
	formatter: function (x) {
		return x + "%";
	},
});
