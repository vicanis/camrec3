package api

import (
	"html/template"
	"io"
)

func renderPlayer(src string, w io.Writer) error {
	return template.Must(template.New("index").Parse(`
		<!DOCTYPE html>
		<html>
			<head>
				<title>Event video</title>
				<style>
					body {
						margin: 0;
						overflow: hidden;
					}
					video {
						width: 100vw;
						height: 100vh;
						padding: 1em;
						box-sizing: border-box;
					}
				</style>
				<script>
					window.addEventListener('load', () => {
						const video = document.getElementById("video");
						video.addEventListener('click', () => video.play());
					});
				</script>
			</head>
			<body>
				<video id="video" src="{{.}}" />
			</body>
		</html>
	`)).Execute(w, src)
}
