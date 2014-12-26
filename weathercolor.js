var Blink1 = require('node-blink1'),
	blink1,
	util = require('util');

try {
	blink1 = new Blink1();
} catch(e) {

}
var ranges = [{ 
	temp : 90, 
	rgb : [207, 57, 39],
}, { 
	temp : 80, 
	rgb : [238, 126, 38],
}, { 
	temp : 70, 
	rgb : [253, 249, 41],
}, { 
	temp : 60, 
	rgb : [110, 210, 40],
}, { 
	temp : 50, 
	rgb : [90, 219, 140],
}, { 
	temp : 40, 
	rgb : [68, 184, 219],
}, { 
	temp : 30, 
	rgb : [75, 56, 156],
}, {
	temp : 20, 
	rgb : [148, 69, 188], // hsl(280, 47%, 50%}, )
}, { 
	temp : 10, 
	rgb : [213, 117, 219] // hsl(296, 59%, 66%)
}];

module.exports = function(temp) {
	var upper = ranges.reduce(function(prev, current) {
			return (current.temp < prev.temp && current.temp > temp) 
				? current 
				: prev;
		}),
		lower = ranges.reduceRight(function(prev, current) {
			return (current.temp > prev.temp && current.temp < temp) 
				? current 
				: prev;
		}),
		rgb, pct;

	if (upper.temp === lower.temp) { // weather extremes
		rgb = upper.rgb;
	} else {
		pct = (temp - lower.temp)/(upper.temp-lower.temp);
		rgb = [];
		for (var i = 0; i < 3; i++) {
			rgb[i] = ((upper.rgb[i] - lower.rgb[i]) * pct) + lower.rgb[i];
			rgb[i] = rgb[i].toFixed(2);
		};
	}
	// if (blink1) blink1.fadeToRGB(100, rgb[0], rgb[1], rgb[2])
	return rgb;
}