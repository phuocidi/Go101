import $ from 'jquery';

window.jQuery = $;
global.jQuery = $;

$( () => {
	var socket = undefined;
	$("#chatbox").submit( (e) => {
		e.preventDefault();
		if (!$("#chatbox textarea").val()) return false;
		if (!socket) {
			alert("Error: There is no socket connection.");
			return false;
		}
		socket.send($("#chatbox textarea").val());
		
		$("#chatbox textarea").val("");
		return false;
	});

	if (!window["WebSocket"]) {
		alert("Error: Your browser does not support web sockets.")
	} else {
		socket = new WebSocket("ws://localhost:8080/room");
		socket.onclose = () => {
			alert("Connection has been closed.")
		}
		socket.onmessage = (e) => {
			$("#messages").append($("<li>").text(e.data));
		}
	}
});

