document.addEventListener("DOMContentLoaded", function () {
	var calendarEl = document.getElementById("dayGrid");
	var calendar = new FullCalendar.Calendar(calendarEl, {
		headerToolbar: {
			left: "prevYear,prev,next,nextYear today",
			center: "title",
			right: "dayGridMonth,dayGridWeek,dayGridDay",
		},
		initialDate: "2022-10-12",
		navLinks: true, // can click day/week names to navigate views
		editable: true,
		dayMaxEvents: true, // allow "more" link when too many events
		events: [
			{
				title: "All Day Event",
				start: "2022-10-01",
				color: "#64a2b9",
			},
			{
				title: "Long Event",
				start: "2022-10-07",
				end: "2022-10-10",
				color: "#6fb4ce",
			},
			{
				groupId: 999,
				title: "Birthday",
				start: "2022-10-09T16:00:00",
				color: "#5990a5",
			},
			{
				groupId: 999,
				title: "Birthday",
				start: "2022-10-16T16:00:00",
				color: "#7dbcd3",
			},
			{
				title: "Conference",
				start: "2022-10-11",
				end: "2022-10-13",
				color: "#8cc3d8",
			},
			{
				title: "Meeting",
				start: "2022-10-14T10:30:00",
				end: "2022-10-14T12:30:00",
				color: "#9acbdd",
			},
			{
				title: "Lunch",
				start: "2022-10-16T12:00:00",
				color: "#6fb4ce",
			},
			{
				title: "Meeting",
				start: "2022-10-18T14:30:00",
				color: "#a9d2e2",
			},
			{
				title: "Interview",
				start: "2022-10-21T17:30:00",
				color: "#b7dae7",
			},
			{
				title: "Meeting",
				start: "2022-10-22T20:00:00",
				color: "#a9d2e2",
			},
			{
				title: "Birthday",
				start: "2022-10-13T07:00:00",
				color: "#8cc3d8",
			},
			{
				title: "Click for Google",
				url: "http://bootstrap.gallery/",
				start: "2022-10-28",
				color: "#7dbcd3",
			},
			{
				title: "Interview",
				start: "2022-10-20",
				color: "#6fb4ce",
			},
			{
				title: "Product Launch",
				start: "2022-10-29",
				color: "#9acbdd",
			},
			{
				title: "Leave",
				start: "2022-10-25",
				color: "#8cc3d8",
			},
		],
	});

	calendar.render();
});
