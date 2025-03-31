# **N**atural **L**anguage **T**ables
## *or Natural Language Representations of Tabular Data*
---

[![Open in gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#https://github.com/ryancahildebrandt/nlt)
[![This project contains 0% LLM-generated content](https://brainmade.org/88x31-dark.png)](https://brainmade.org/)

## Purpose

This project provides some simple tooling to transform tabular data into a series of natural language statements. RAG based systems often show difficulty accurately working with data stored in tabular format, and any foundation language model is going to be limited in how much of its training data shows tabular data. By presenting tabular data in natural language format, RAG systems may be able to retrieve and represent tabular data more accurately.

---

## Approach

### Theory
This project was inspired by some professional work searching and retrieving information from tables within larger documents, and draws heavily on [this](https://doi.org/10.1007/978-3-031-15146-0_15) conference paper from Dr. Dirk Schlimm of McGill University. The paper provides a formal overview of tables' structure and features, as well as how they are used to store and display information. Of this overview, the most relevant concepts are:

- Tables are 2 dimensional structures composed of a series of horizontal rows, vertical columns, and individual cells that sit at the intersection of a particular row and column
- Generally, rows and columns cohere semantically, such that a given row or column forms some meaningful unit
- Rows and columns are often accompanied by labels, which denote what if any concept a given column or row corresponds to
- Cells in a table can be described by their position on the x and y axes, such that all cells with the same x coordinate belong to the same column and all cells with the same y coordinate belong to the same row

In addition to the core concepts above, the current project makes use of a few additional characteristics:

- Labels for a given table may be contained in the table itself, in the form of one or more "header" cells at the beginning of the row/column
- In addition to the unique coordinates given by a cell's position on the x and y axes, a cell's position can also be described by a unique combination of labels and/or headers from the row and column to which it belongs

### Data Format
While tabular data is frequently stored in databases or spreadsheets, this program works with simpler and more standardized file formats. Each of these needs slightly different handling to read into a standardized table format, briefly described below:

- C/TSV: Imported more or less as is, as c/tsv files generally correspond 1:1 to their tabular format without additional transformation 
- HTML: Read using nested <tr> and <td> tags, and strips all <tbody>, <thead>, and <th> tags for simplicity
- Markdown: Converted to html and treated as above
- JSONL: Interpreted as a series of table rows, with each column name represented in the key for each key value pair
- JSON: While json as a format is flexilbe enough to allow for a range of different tabular data representations, but for the purposes of this project I've implemented parsers for the following 2:
	- Array of arrays:
	```json
	[
		["x1y1", "x2y1", ...],
		["x1y2", "x2y2, ...],
		...
	]
	```
	- Array of objects:
	```json
	[
		{"x1":"y1", "x2":"y1", ...},
		{"x1":"y2", "x1":"y2", ...},
		...
	]
	```
Once the content of each file is read, I use the [gota](https://pkg.go.dev/github.com/go-gota/gota/dataframe) dataframe struct to get the raw data into a standardized tabular format

### Parsing
Once in a dataframe, the data gets read into a custom TableData struct storing the dimensions of the dataframe, all row headers, all column headers, and the individual cells. 
Each cell stores its own x and y coordinates within the table, the headers for the cell's row and column, and the value of the cell.
"Headers" as mentioned here are defined as the 1st n values for a given row/column, concatenated together with a provided delimiter.
Joining cells to create headers allows for a more nuanced and specific representation of data in the table when formatting into natural language statements.
Headers in this sense are not necessarily applicable for every table, and in these cases the number of cells composing the header n can be given as 0.

### Formatting
Once decomposed into its constintuent rows, columns, and cells, the data can then be recomposed into natural language statements.
This recomposisition is handled by any number of TableFormatters, which combine the information contained in a TableData struct and user provided string fields contained in a FormatFields struct.

The FormatFields struct accepts the following fields, some combination of which will be used by any given TableFormatter.
I've tried to name them in a way that suggests their role where possible:
- delim: Delimter between headers for each row/column, for when there are multiple cells constituting the header
	- Large **with** thin crust...
	- Family meal deal **,** no discount...
- link: Clause linking x/y/val labels to the rest of the sentence or another label. Can be used simialrly to preamble in some cases
	- **When** topping is pepperoni...
	- Price **for** double cheese...
- eq: Statement of equality between labels and values
	- Price **is** $12...
	- Number of slices **will be** 6...
- pre: Preamble clause setting the context for the relationship between labels and values. Can be used simialrly to link in some cases
	- **All possible values for** size are...
	- **All available** toppings are...
- val_label: Semantic/category label for cell values
	- **Price** is $12...
	- **Discount** totals %20...
- x_label: Semantic/category label for a given row
	- When **country** is South Korea...
	- In cases where customers ask about **pizza chain** Sbarro...
- y_label: Semantic/category label for a given column
	- When **topping** = peppers...
	- If customer ordered **size** large... 

And there are a range of prebuilt TableFormatters, in addition to the CustomFormatter which accepts a custom formatting string to apply. For the CustomFormatter, you'll need to create a short script using the functions and types here, since it requires that you pass in specific objects rather than a simple string field.
*Some notes:*
- For format strings:
	- square brackets [] indicate an array of values
	- angle brackets <> indicate values taken from the table data
	- parentheses () indicate values taken from user provided fields
- The names for different formatters are admittedly lengthy but should give an idea of when you should use each one
	- Named/Unnamed designates whether a row/column/value has a semantic/category label that describes it
	- Coord indicates that
		  1) each cell will be described by its unique combination of row and column headers and/or label
		  2) both the row and column have headers and/or labels to use
	- NamedRow or NamedCol formatters are appropriate when only the column or row has a header and/or label
	- Val formatters enumerate all of the values in a given row or column, without the corresponding column/row head identifying the value
	- KeyVal formatters formatters enumerate all of the values in a given row or column, with the corresponding column/row head identifying the value
	- Traling numbers after a formatter name indicate a variation on the formatter, basically a different phrasing of the same core information

I've included a quick lookup table to help identify which formatters may be useful given the features your table does or doesn't have. In a lot of cases, since most of the formatters use the exact same fields, the choice of formatter may come down to which output phrasing you prefer. ✓ means the formatter expects that field, O means it's optional, though all fields can effectively be omitted by feeding in a blank string

|                           | x_head | y_head | cell_val | delim | link | eq | pre | val_label | x_label | y_label |
|---------------------------|--------|--------|----------|-------|------|----|-----|-----------|---------|---------|
| UnnamedCoordFormatter1    | ✓      | ✓      | ✓        | O     | ✓    | ✓  |     |           |         |         |
| UnnamedCoordFormatter2    | ✓      | ✓      | ✓        | O     | ✓    | ✓  |     |           |         |         |
| NamedCoordFormatter1      | ✓      | ✓      | ✓        | O     | ✓    | ✓  |     | ✓         | ✓       | ✓       |
| NamedCoordFormatter2      | ✓      | ✓      | ✓        | O     | ✓    | ✓  |     | ✓         | ✓       | ✓       |
| NamedRowFormatter         | ✓      | ✓      | ✓        | O     | ✓    | ✓  |     |           | ✓       |         |
| NamedColFormatter         | ✓      | ✓      | ✓        | O     | ✓    | ✓  |     |           |         | ✓       |
| UnnamedRowKeyValFormatter | ✓      | ✓      | ✓        | O     | ✓    | ✓  |     |           |         |         |
| UnnamedColKeyValFormatter | ✓      | ✓      | ✓        | O     | ✓    | ✓  |     |           |         |         |
| NamedRowKeyValFormatter   | ✓      | ✓      | ✓        | O     | ✓    | ✓  |     |           | ✓       |         |
| NamedColKeyValFormatter   | ✓      | ✓      | ✓        | O     | ✓    | ✓  |     |           |         | ✓       |
| RowValFormatter           | ✓      |        | ✓        | O     | ✓    |    | ✓   |           |         |         |
| ColValFormatter           |        | ✓      | ✓        | O     | ✓    |    | ✓   |           |         |         |

And a breakdown of the structure of each formatter with an example output:
- UnnamedCoordFormatter1
	- Format string: (val_label) (link) <x_head> and <y_head> (eq) <value>
	- Example: price for extra pepperoni and no cheese is $12.00
- UnnamedCoordFormatter2
	- Format string: (link) <x_head> and (y_label), (val_label) (eq) <value>
	- Example: For Extra pepperoni and no cheese, price will be $12.00
- NamedCoordFormatter1
	- Format string: (val_label) (link) (x_label) (eq) <x_head> and (y_label) (eq) <y_head> (eq) <value>
	- Example: Price when size is medium and crust is thin is $15
- NamedCoordFormatter2
	- Format string: (link) (x_label) (eq) <x_head> and (y_label) (eq) <y_head>, (val_label) (eq) <value>
	- Example: When size = medium and crust = thin, price = $15
- NamedRowFormatter
	- Format string: (link) (x_label) (eq) <x_head>, <y_head> (eq) <value>
	- Example: If topping is meat, vegan is false
- NamedColFormatter
	- Format string: (link) (y_label) (eq) <y_head>, <x_head> (eq) <value>
	- Example: If crust is gluten free, price increases by $3
- UnnamedRowKeyValFormatter
	- Format string: (link) (x_head),  [(y_head) (eq) <value>]
	- Example: For daily specials, [Monday is none, Tuesday is taco pizza, Wednesday is wing pizza]
- UnnamedColKeyValFormatter
	- Format string: (link) (y_head), [(x_head) (eq) <value>]
	- Example: For sides, [Wings are $5, Mozz sticks are $7, Cheese curds are $6]
- NamedRowKeyValFormatter
	- Format string: (link) (x_label) (eq) <x_head>, [(y_head) (eq) <value>]
	- Example: When country = South Korea, [Dominos is #1, Pizza Alvolo is #2, PizzaHut is #3]
- NamedColKeyValFormatter
	- Format string: (link) (y_label) (eq) <y_head>, [(x_head) (eq) <value>]
	- Example: In the case that chain is Sbarro, [locations is 600, year founded is 1956, hq is Columbus, Ohio]
- RowValFormatter
	- Format string: (pre) <x_head> (link) [<value>]
	- Example: All possible topping are [sausage, mushroom, olives]
- ColValFormatter
	- Format string: (pre) <y_head> (link) [<value>]
	- Example: Size can be one of [small, medium, large]

---

## Usage
All inputs including input and output files, parsers and formatters, and fields used in the formatters are set via config.json or a separate config file. The basic config.json looks like this:

```json
{
	"infile": "data/test1.csv",
	"outfile": "outputs/output.txt",
	"formatter": "UnnamedCoordFormatter2",
	"row_headers": 2,
	"col_headers": 1,
	"parser": "CSV",
	"delim": "delim",
	"link": "link",
	"eq": "eq",
	"pre": "pre",
	"val_label": "val_label",
	"x_label": "x_label",
	"y_label": "y_label"
}
```

Because inputs are handled via config.json, there are only 3 possible usages of nlt on the command line, not counting the -h flag:

```shell
# run with default options
nlt

# run with pointer to config file
nlt -c foo/bar.json

# run using lastrun.json, which stores the last run of nlt
nlt -l

```
As indicated above, lastrun.json is populated with a copy of the config from the last successful nlt run. This happens automatically on every run, so if you have a specific run you want to save the config for, make sure to make a copy of lastrun.json before trying another config.

---

## Outputs

- [nlt](./nlt) executable

---

## References

Schlimm, D. (2022). Tables as Powerful Representational Tools. In: Giardino, V., Linker, S., Burns, R., Bellucci, F., Boucheix, JM., Viana, P. (eds) Diagrammatic Representation and Inference. Diagrams 2022. Lecture Notes in Computer Science(), vol 13462. Springer, Cham. https://doi.org/10.1007/978-3-031-15146-0_15
