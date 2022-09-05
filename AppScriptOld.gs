function myFunction() {
  // barcode = e.queryString.barcode;
   // Open Google Sheet using ID
  var ss = SpreadsheetApp.openById(fileID);
  var sheet = ss.getSheetByName(sheetName);
  // Read all data rows from Google Sheet
  const values = sheet.getRange(2, 1, sheet.getLastRow() - 1, sheet.getLastColumn()).getValues();
  //values.forEach(element => console.log(element));
  
 /* var item, items = [];
  values.forEach(element => {
  //  console.log(element)
    var nameArr = element.split(',');
    item = {};
    item.SupplierName = nameArr[3];
    item.Brand = nameArr[4];
    item.id = idsArray[i];
    items.push(item);
    }); 

  return ContentService.createTextOutput(values[0]).setMimeType(ContentService.MimeType.TEXT);
  */

  // Ref: https://stackoverflow.com/a/21231012
const letterToColumn = letter => {
  let column = 0,
    length = letter.length;
  for (let i = 0; i < length; i++) {
    column += (letter.charCodeAt(i) - 64) * Math.pow(26, length - i - 1);
  }
  return column;
};

const columnLetters = ["N", "X", "S"]; // Column letters you want to retrieve.
const fields = values.map(r => columnLetters.map(e => r[letterToColumn(e) - 1]));

const selected = fields.filter(line => line[0].toString() == barcode);

// Converts data rows in json format
const result = JSON.stringify({ItemCode: selected[0],SupplierName:selected[1],BarcCode:selected[2],});
// const result = JSON.stringify(values.map(([a,b,c,d,e]) => ({SupplierName: d,Brand:e,})));
 // or
//  const result = JSON.stringify({SupplierName: values.d,Brand:values.e,});
// Returns Result
return ContentService.createTextOutput(result).setMimeType(ContentService.MimeType.JSON);

}
