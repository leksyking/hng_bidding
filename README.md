### Bidding Task

Take the CSV provided by the teams, and generate a CHIP-0007 compatible json,
calculate the sha256 of the json file and append it to each line in the csv as
a (filename.output.csv)

### main function
- Open the csv file
- Read all lines of the csv data
- Send the data into the GetAllLines function

### Create a struct to handle the csv data format

### GetAllLines function
- create a variable of an array of the struct
- Range through the data while excluding the header
- Create a variable to represent each rows of the data
- Range through each rows and store each column under the respective header columns
- Convert the row to json files
- add the .json extension to the files
- Open the json files and calculate sha256 on the json files
- Append the hash to the json files
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
- write the array of csvRow into the csv writer
