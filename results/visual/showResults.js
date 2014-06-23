var fs = require('fs');
var d3 = require('d3');
var xmldom = require('xmldom');

var fileName = process.argv[2]

var filePath = "./results/" + fileName + ".csv"

// Construct the graph.
var margin = {top: 20, right: 20, bottom: 30, left: 50},
    width = 960 - margin.left - margin.right,
    height = 500 - margin.top - margin.bottom;

var x = d3.scale.linear()
    .range([0, width]);

var y = d3.scale.linear()
    .range([height, 0]);

var xAxis = d3.svg.axis()
    .scale(x)
    .orient("bottom");

var yAxis = d3.svg.axis()
    .scale(y)
    .orient("left");

var line = d3.svg.line()
    .interpolate("basis")
    .x(function(d) { return x(d[0]); })
    .y(function(d) { return y(d[1]); });

var svg = d3.select("body").append("svg")
    .attr("width", width + margin.left + margin.right)
    .attr("height", height + margin.top + margin.bottom)
  .append("g")
    .attr("transform", "translate(" + margin.left + "," + margin.top + ")");

var dataset = [];
var tokens = [];

fs.readFile(filePath, 'utf8', function (err, data) {
  d3.csv.parseRows(data, function(d) {

    dataset.push([+d[0], +d[1]]);
    tokens.push(d[2]);
  });

  x.domain(d3.extent(dataset, function(d) { return d[0]; }));
  y.domain(d3.extent(dataset, function(d) { return d[1]; }));

  svg.append("path")
      .datum(dataset)
      .attr("class", "line")
      .attr("d", line);

  svg.append("g")
      .attr("class", "x axis")
      .attr("transform", "translate(0," + height + ")")
      .call(xAxis);

  svg.append("g")
      .attr("class", "y axis")
      .call(yAxis)
    .append("text")
      .attr("transform", "rotate(-90)")
      .attr("y", 6)
      .attr("dy", ".71em")
      .style("text-anchor", "end")
      .text("Cohesion");

  var svgGraph = d3.select('svg');
  var svgXML = (new xmldom.XMLSerializer()).serializeToString(svgGraph[0][0]);



  var tokensString = "var tokens = ["
  tokens.forEach(function(d) {
    tokensString = tokensString + "'" + d +"',"
  })
  tokensString = tokensString + "];\n"
  var scriptString = "<script>\n" + tokensString + "\nfunction checkField(val) {\ndocument.getElementById('tokenDiv').innerHTML = tokens[val];\n};\n</script>"
  var tokenQuery = "<input type='number' name='txt' value='0' onchange='checkField(this.value)'>"
  var styleString = "<style>\nbody {\nfont: 10px sans-serif;\n}\npath,\nline {\nfill: none;\nstroke: #000;\n}\n</style>\n"
  var html = "<!DOCTYPE html>\n<html>\n<head>\n<meta charset='utf-8' />\n" + styleString + "\n" + scriptString + "</head>\n<body>" + svgXML + "<div id='tokenDiv'></div>" + tokenQuery + "\n</body>\n</html>"

  // Save the file in .html format.
  var writeName = "./results/visual/graphs/" + fileName + ".html";
  fs.writeFile(writeName, html);

  console.log(fileName + ".html");
})
