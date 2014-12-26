var Forecast = require('forecast.io'),
	options = {
	  APIKey: process.env.FORECAST_API_KEY
	},
	color = require('onecolor'),
	forecast = new Forecast(options),
	weathercolor = require('./weathercolor.js');

forecast.get(41.248098, -96.036211, function (err, res, data) {
	if (err) throw err;
	var nowTemp = data.currently.temperature,
		todaysHigh = data.daily.data[0].temperatureMax,
		todaysLow = data.daily.data[0].temperatureMin;

	// console.log(new Date())
	// console.log("NOW",nowTemp)
	// console.log("LOW",todaysLow)
	// console.log("HIGH", todaysHigh);

	var rgb = weathercolor(nowTemp),
		formattedRGB = "rgb("+ rgb.join(",") + ")",
		hex = color(formattedRGB).hex();
	console.log(hex);
});