package main

import "fmt"

func renderIndexPage() string {
	return `
<html>
	<body>
		<h1>Weather App</h1>
		<div>
			<form action="/weather.html" method="GET">
				<label for="addr">Address:</label>
				<input name="addr" />
			</form>
		</div>
	</body>
</html>`
}

func main() {
	fmt.Println(renderIndexPage())
}
