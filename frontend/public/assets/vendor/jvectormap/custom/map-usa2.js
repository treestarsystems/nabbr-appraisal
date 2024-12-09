// USA map 1
$(function () {
	var cityAreaData = [
		230.2, 750.9, 440.28, 180.15, 69.35, 280.9, 510.5, 99.6, 135.5,
	];
	$("#us-map3").vectorMap({
		map: "us_aea_en",
		scaleColors: ["#bf7a6a"],
		normalizeFunction: "polynomial",

		zoomOnScroll: false,
		zoomMin: 1,
		hoverColor: true,
		regionStyle: {
			initial: {
				fill: "#bf7a6a",
			},
			hover: {
				"fill-opacity": 0.8,
			},
		},
		markerStyle: {
			initial: {
				fill: "#d2a968",
				stroke: "#FFFFFF",
				r: 5,
			},
		},
		backgroundColor: "transparent",
	});
});
