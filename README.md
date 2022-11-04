### Bidding Task

Take the CSV provided by the teams, and generate a CHIP-0007 compatible json,
calculate the sha256 of the json file and append it to each line in the csv as
a (filename.output.csv)

### main function
- Open the csv file
- Read all lines of the csv data
- Send the data into the GetAllLines function

### Create a struct to handle the csv data format, and chip-0007 Format

### GetAllLines function
- create a variable of an array of the csvFileJson struct
- Range through the data while excluding the header
- Create a variable to represent each rows of the data
- And a variable to store the Chip-0007 format data
- Range through each rows and store each column under the respective header columns and chip-0007 fields
- remove semicolon from the attributes and convert it to string and range over it to get the specific key and value pairs
- Append the attributes to the chip.Attributes field
- Convert the structs of both the csv and chip-0007 to json files
- add the .json extension to the files
- Open the json files and calculate sha256 on the chip-0007 json files
- Append the hash to the csv json files
- Create an output.json file from the array of various rows
- Send the output.json file and the output.csv file to ConvertJSONToCSV function to create our updated csv file

### ConvertJSONToCSV
- Open the json file 
- Decode the json file into an array of struct
- Create the output.csv file
- create a csv writer that writes to the csv file
- Write the header row into the csv file
- Range through the array of struct created and append the rows of the struct to each
rows of an array of csvRow variable
- write(append) the array of csvRow into the csv writer
- return nothing or error if any occurs
